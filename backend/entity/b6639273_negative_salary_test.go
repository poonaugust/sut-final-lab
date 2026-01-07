package entity_test

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"github.com/poonaugust/sut-final-lab/backend/entity"
)

func TestSalaryValidate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("Salary is invalid", func(t *testing.T) {
		employee := entity.Employees{
			Name:   "Poon",
			Salary: 250000,
			EmployeeCode: "HR-1234",
		}

		ok, err := govalidator.ValidateStruct(employee)

		g.Expect(ok).To(gomega.BeFalse())
		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Salary must be between 15000 and 200000"))

	})
}
