// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceGroupAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateResourceGroupAttributeResponseBody
	GetRequestId() *string
}

type UpdateResourceGroupAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 91C7CAB5-3B2E-4FB6-893C-0162C0******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceGroupAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceGroupAttributeResponseBody) SetRequestId(v string) *UpdateResourceGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceGroupAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
