package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example Tests", func() {
	tests := [][][]int{
		{{4, 3, 9, 7, 2, 1}, {2, 9, 3, 49, 4, 1}},
		{{100, 101, 5, 5, 1, 1}, {10, 10201, 25, 25, 1, 1}},
		{{1, 2, 3, 4, 5, 6}, {1, 4, 9, 2, 25, 36}},
	}
	It("To square(root) or not to square(root)", func() {
		for _, t := range tests {
			Expect(SquareOrSquareRoot(t[0])).To(Equal(t[1]))
		}
	})
})
