// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateApplicationTemplateResponseBody
	GetRequestId() *string
}

type UpdateApplicationTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationTemplateResponseBody) SetRequestId(v string) *UpdateApplicationTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
