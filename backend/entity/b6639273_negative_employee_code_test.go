package entity_test

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"github.com/poonaugust/sut-final-lab/backend/entity"
)

func TestEmployeeCodeValidate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("EmployeeCode is invalid", func(t *testing.T) {
		employee := entity.Employees{
			Name:   "Poon",
			Salary: 100000,
			EmployeeCode: "HR-12345",
		}

		ok, err := govalidator.ValidateStruct(employee)

		g.Expect(ok).To(gomega.BeFalse())
		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("EmployeeCode mustbe 2 uppercase English letters (A-Z) followed by '-' and 4 digits (0-9)"))

	})
}
