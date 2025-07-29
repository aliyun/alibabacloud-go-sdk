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
	// The permission mode of the agent. Valid values:
	//
	// admin: the admin mode, which provides full permissions. restricted: the restricted mode, which provides partial permissions. Default value: admin.
	//
	// example:
	//
	// admin
	AgentMode *string `json:"AgentMode,omitempty" xml:"AgentMode,omitempty"`
	// Specifies whether to obtain the credentials that are used to access the cluster over the internal network.
	//
	// 	- `true`: obtains the credentials that are used to access the cluster over the internal network.
	//
	// 	- `false`: obtains the credentials that are used to access the cluster over the Internet.
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
