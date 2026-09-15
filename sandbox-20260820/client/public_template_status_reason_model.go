// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicTemplateStatusReason interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *PublicTemplateStatusReason
	GetMessage() *string
	SetStep(v string) *PublicTemplateStatusReason
	GetStep() *string
}

type PublicTemplateStatusReason struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Step    *string `json:"step,omitempty" xml:"step,omitempty"`
}

func (s PublicTemplateStatusReason) String() string {
	return dara.Prettify(s)
}

func (s PublicTemplateStatusReason) GoString() string {
	return s.String()
}

func (s *PublicTemplateStatusReason) GetMessage() *string {
	return s.Message
}

func (s *PublicTemplateStatusReason) GetStep() *string {
	return s.Step
}

func (s *PublicTemplateStatusReason) SetMessage(v string) *PublicTemplateStatusReason {
	s.Message = &v
	return s
}

func (s *PublicTemplateStatusReason) SetStep(v string) *PublicTemplateStatusReason {
	s.Step = &v
	return s
}

func (s *PublicTemplateStatusReason) Validate() error {
	return dara.Validate(s)
}
