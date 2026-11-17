// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationTemplateId(v string) *CreateApplicationTemplateResponseBody
	GetApplicationTemplateId() *string
	SetRequestId(v string) *CreateApplicationTemplateResponseBody
	GetRequestId() *string
}

type CreateApplicationTemplateResponseBody struct {
	// The ID of the newly created application template.
	//
	// example:
	//
	// ati-4f0691c56b8d4f06
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationTemplateResponseBody) GetApplicationTemplateId() *string {
	return s.ApplicationTemplateId
}

func (s *CreateApplicationTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationTemplateResponseBody) SetApplicationTemplateId(v string) *CreateApplicationTemplateResponseBody {
	s.ApplicationTemplateId = &v
	return s
}

func (s *CreateApplicationTemplateResponseBody) SetRequestId(v string) *CreateApplicationTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
