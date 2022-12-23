// Copyright © 2022 The Homeport Team
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package dtr_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/homeport/duct-tape-resource/internal/dtr"
)

var _ = Describe("Out", func() {
	Context("valid configuration", func() {
		It("should return the given version", func() {
			version := Version{"ref": "foobar"}
			result, err := Out(feed(Config{
				Source: Source{
					Out: Custom{Run: "true"},
				},
				Version: version,
			}))

			Expect(err).NotTo(HaveOccurred())
			Expect(result.Version).To(Equal(version))
		})

		It("should return the given version and metadata if output was available", func() {
			version := Version{"ref": "foobar"}
			result, err := Out(feed(Config{
				Source: Source{
					Out: Custom{Run: "echo foo bar && echo bar foo"},
				},
				Version: version,
			}))

			Expect(err).NotTo(HaveOccurred())
			Expect(result.Version).To(Equal(version))
			Expect(result.Metadata).To(Equal([]Metadata{
				{Name: "foo", Value: "bar"},
				{Name: "bar", Value: "foo"},
			}))
		})
	})
})
