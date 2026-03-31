// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTemplateAttributesResponseBody
	GetRequestId() *string
}

type UpdateTemplateAttributesResponseBody struct {
	// example:
	//
	// A5320F1D-92D9-44BB-A416-5FC525ED6D57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTemplateAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTemplateAttributesResponseBody) SetRequestId(v string) *UpdateTemplateAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTemplateAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}
