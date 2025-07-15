// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDhcpOptionsSetAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDhcpOptionsSetAttributeResponseBody
	GetRequestId() *string
}

type UpdateDhcpOptionsSetAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDhcpOptionsSetAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDhcpOptionsSetAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDhcpOptionsSetAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDhcpOptionsSetAttributeResponseBody) SetRequestId(v string) *UpdateDhcpOptionsSetAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
