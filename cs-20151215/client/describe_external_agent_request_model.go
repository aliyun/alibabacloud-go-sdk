// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExternalAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentMode(v string) *DescribeExternalAgentRequest
	GetAgentMode() *string
	SetPrivateIpAddress(v string) *DescribeExternalAgentRequest
	GetPrivateIpAddress() *string
}

type DescribeExternalAgentRequest struct {
	// The agent permission mode. Valid values:
	//
	// - admin: Administrator mode. Grants all permissions.
	//
	// - restricted: Restricted mode. Grants only partial permissions.
	//
	// Default value: admin.
	//
	// example:
	//
	// admin
	AgentMode *string `json:"AgentMode,omitempty" xml:"AgentMode,omitempty"`
	// Specifies whether to obtain internal network access credentials. Valid values:
	//
	// - `true`: Obtains only internal network connection credentials.
	//
	// - `false`: Obtains only public network connection credentials.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s DescribeExternalAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExternalAgentRequest) GoString() string {
	return s.String()
}

func (s *DescribeExternalAgentRequest) GetAgentMode() *string {
	return s.AgentMode
}

func (s *DescribeExternalAgentRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeExternalAgentRequest) SetAgentMode(v string) *DescribeExternalAgentRequest {
	s.AgentMode = &v
	return s
}

func (s *DescribeExternalAgentRequest) SetPrivateIpAddress(v string) *DescribeExternalAgentRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeExternalAgentRequest) Validate() error {
	return dara.Validate(s)
}
