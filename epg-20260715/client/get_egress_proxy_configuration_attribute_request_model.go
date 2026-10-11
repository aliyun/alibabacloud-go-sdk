// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEgressProxyConfigurationAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyConfigurationId(v string) *GetEgressProxyConfigurationAttributeRequest
	GetEgressProxyConfigurationId() *string
}

type GetEgressProxyConfigurationAttributeRequest struct {
	// The ID of the egress proxy configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// epc-hn302f6a9a284c4f****
	EgressProxyConfigurationId *string `json:"EgressProxyConfigurationId,omitempty" xml:"EgressProxyConfigurationId,omitempty"`
}

func (s GetEgressProxyConfigurationAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyConfigurationAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetEgressProxyConfigurationAttributeRequest) GetEgressProxyConfigurationId() *string {
	return s.EgressProxyConfigurationId
}

func (s *GetEgressProxyConfigurationAttributeRequest) SetEgressProxyConfigurationId(v string) *GetEgressProxyConfigurationAttributeRequest {
	s.EgressProxyConfigurationId = &v
	return s
}

func (s *GetEgressProxyConfigurationAttributeRequest) Validate() error {
	return dara.Validate(s)
}
