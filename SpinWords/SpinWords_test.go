// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

package spinwords_test

import (
	. "codewarrior/kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Sample", func() {
	It("should test that the solution returns the correct value for single word inputs", func() {
		Expect(SpinWords("Welcome")).To(Equal("emocleW"))
		Expect(SpinWords("to")).To(Equal("to"))
		Expect(SpinWords("CodeWars")).To(Equal("sraWedoC"))
	})
	It("should test that the solution returns the correct value for multiple word outputs", func() {
		Expect(SpinWords("Hey fellow warriors")).To(Equal("Hey wollef sroirraw"))
		Expect(SpinWords("Burgers are my favorite fruit")).To(Equal("sregruB are my etirovaf tiurf"))
		Expect(SpinWords("Pizza is the best vegetable")).To(Equal("azziP is the best elbategev"))
	})
})
