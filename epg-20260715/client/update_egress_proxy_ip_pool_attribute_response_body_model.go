// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEgressProxyIpPoolAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEgressProxyIpPoolAttributeResponseBody
	GetRequestId() *string
}

type UpdateEgressProxyIpPoolAttributeResponseBody struct {
	// example:
	//
	// 019F3A67-2D1B-537B-B08D-B2DEB69E8E8B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEgressProxyIpPoolAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyIpPoolAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyIpPoolAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEgressProxyIpPoolAttributeResponseBody) SetRequestId(v string) *UpdateEgressProxyIpPoolAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEgressProxyIpPoolAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
