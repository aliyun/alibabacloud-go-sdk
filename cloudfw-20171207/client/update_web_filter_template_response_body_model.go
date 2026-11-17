// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebFilterTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWebFilterTemplateResponseBody
	GetRequestId() *string
}

type UpdateWebFilterTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWebFilterTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebFilterTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWebFilterTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWebFilterTemplateResponseBody) SetRequestId(v string) *UpdateWebFilterTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWebFilterTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
