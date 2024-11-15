package bank

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

var _ = Describe("Account", func() {
    Describe("Deposit", func() {
        It("should add amount to balance", func() {
            account := Account{}
            err := account.Deposit(100)
            Expect(err).To(BeNil())
            Expect(account.GetBalance()).To(Equal(100.0))
        })

        It("should return an error if the deposit amount is zero or negative", func() {
            account := Account{}
            err := account.Deposit(-50)
            Expect(err).NotTo(BeNil())
            Expect(err.Error()).To(Equal("deposit amount must be greater than zero"))
        })
    })

    Describe("Withdraw", func() {
        It("should deduct amount from balance", func() {
            account := Account{Balance: 100}
            err := account.Withdraw(50)
            Expect(err).To(BeNil())
            Expect(account.GetBalance()).To(Equal(50.0))
        })

        It("should return an error if the withdrawal amount is zero, negative, or greater than the balance", func() {
            account := Account{Balance: 50}
            err := account.Withdraw(-10)
            Expect(err).NotTo(BeNil())
            Expect(err.Error()).To(Equal("withdrawal amount must be greater than zero"))
            
            err = account.Withdraw(100)
            Expect(err).NotTo(BeNil())
            Expect(err.Error()).To(Equal("insufficient funds"))
        })
    })
})
