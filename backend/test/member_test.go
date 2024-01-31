package test

import (
	"testing"

	"github.com/Popopond/TestbeforeExam/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)
 func TestMember (t *testing.T){
	g := NewGomegaWithT(t)
	t.Run(`Correcct`, func(t *testing.T){
		member := entity.Member{
			ID : 1,
			Name : "piw",
			PhoneNumber: "0930359255",
			Email:"piwpiw@gmail.com" ,
		}
		ok, err := govalidator.ValidateStruct(member)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
	t.Run(`Email is invalid`, func(t *testing.T){
		member := entity.Member{
			ID : 1,
			Name : "piw",
			PhoneNumber: "0930359255",
			Email:"piwpiw@gmail" ,
		}
		ok, err := govalidator.ValidateStruct(member)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email"))
	})
 }