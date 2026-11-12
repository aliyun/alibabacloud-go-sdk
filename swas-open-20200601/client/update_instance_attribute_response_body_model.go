// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateInstanceAttributeResponseBody
	GetRequestId() *string
}

type UpdateInstanceAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceAttributeResponseBody) SetRequestId(v string) *UpdateInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
