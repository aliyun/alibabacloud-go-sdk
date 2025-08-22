// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModuleAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateModuleAttributeResponseBody
	GetRequestId() *string
}

type UpdateModuleAttributeResponseBody struct {
	// example:
	//
	// CA05185E-6B90-5941-98D4-7212688AECC8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateModuleAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModuleAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateModuleAttributeResponseBody) SetRequestId(v string) *UpdateModuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModuleAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
