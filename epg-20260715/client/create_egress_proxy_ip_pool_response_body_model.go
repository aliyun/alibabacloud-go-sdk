// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEgressProxyIpPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyIpPoolId(v string) *CreateEgressProxyIpPoolResponseBody
	GetEgressProxyIpPoolId() *string
	SetRequestId(v string) *CreateEgressProxyIpPoolResponseBody
	GetRequestId() *string
}

type CreateEgressProxyIpPoolResponseBody struct {
	// example:
	//
	// eppool-gw80d96a96ca0b00****
	EgressProxyIpPoolId *string `json:"EgressProxyIpPoolId,omitempty" xml:"EgressProxyIpPoolId,omitempty"`
	// example:
	//
	// 01A05AB7-6C14-5A12-94E7-F20F9F69AE3D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEgressProxyIpPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyIpPoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyIpPoolResponseBody) GetEgressProxyIpPoolId() *string {
	return s.EgressProxyIpPoolId
}

func (s *CreateEgressProxyIpPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEgressProxyIpPoolResponseBody) SetEgressProxyIpPoolId(v string) *CreateEgressProxyIpPoolResponseBody {
	s.EgressProxyIpPoolId = &v
	return s
}

func (s *CreateEgressProxyIpPoolResponseBody) SetRequestId(v string) *CreateEgressProxyIpPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEgressProxyIpPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
