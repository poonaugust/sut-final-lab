package entity_test

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"github.com/poonaugust/sut-final-lab/backend/entity"
)

func TestEmployeeValidate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("Is complete", func(t *testing.T) {
		employee := entity.Employees{
			Name:   "Poon",
			Salary: 50000,
		}

		ok, err := govalidator.ValidateStruct(employee)

		g.Expect(ok).To(gomega.BeTrue())
		g.Expect(err).To(gomega.BeNil())

	})
}
