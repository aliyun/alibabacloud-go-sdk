// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEgressProxyIpPoolAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyIpPoolId(v string) *GetEgressProxyIpPoolAttributeRequest
	GetEgressProxyIpPoolId() *string
}

type GetEgressProxyIpPoolAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eppool-gw80d96a96ca0b00****
	EgressProxyIpPoolId *string `json:"EgressProxyIpPoolId,omitempty" xml:"EgressProxyIpPoolId,omitempty"`
}

func (s GetEgressProxyIpPoolAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyIpPoolAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetEgressProxyIpPoolAttributeRequest) GetEgressProxyIpPoolId() *string {
	return s.EgressProxyIpPoolId
}

func (s *GetEgressProxyIpPoolAttributeRequest) SetEgressProxyIpPoolId(v string) *GetEgressProxyIpPoolAttributeRequest {
	s.EgressProxyIpPoolId = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeRequest) Validate() error {
	return dara.Validate(s)
}
