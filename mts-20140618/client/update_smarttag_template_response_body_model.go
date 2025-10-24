// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmarttagTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSmarttagTemplateResponseBody
	GetRequestId() *string
}

type UpdateSmarttagTemplateResponseBody struct {
	// example:
	//
	// 5210DBB0-E327-4D45-ADBC-0B83C8796E26
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSmarttagTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmarttagTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmarttagTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSmarttagTemplateResponseBody) SetRequestId(v string) *UpdateSmarttagTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmarttagTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
