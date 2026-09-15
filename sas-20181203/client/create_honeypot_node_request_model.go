// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoneypotNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowHoneypotAccessInternet(v bool) *CreateHoneypotNodeRequest
	GetAllowHoneypotAccessInternet() *bool
	SetAvailableProbeNum(v int32) *CreateHoneypotNodeRequest
	GetAvailableProbeNum() *int32
	SetNodeName(v string) *CreateHoneypotNodeRequest
	GetNodeName() *string
	SetSecurityGroupProbeIpList(v []*string) *CreateHoneypotNodeRequest
	GetSecurityGroupProbeIpList() []*string
}

type CreateHoneypotNodeRequest struct {
	// Specifies whether to allow the honeypot to access the Internet. Valid values:
	//
	// - **true**: Allowed.
	//
	// - **false**: Not allowed.
	//
	// example:
	//
	// true
	AllowHoneypotAccessInternet *bool `json:"AllowHoneypotAccessInternet,omitempty" xml:"AllowHoneypotAccessInternet,omitempty"`
	// The number of available probes. This parameter is required. If this parameter is not specified, the API returns InvalidParam (400). The minimum value is 20. If the value is less than 20, the API returns InvalidProbeNum (400).
	//
	// example:
	//
	// 20
	AvailableProbeNum *int32 `json:"AvailableProbeNum,omitempty" xml:"AvailableProbeNum,omitempty"`
	// The name of the management node.
	//
	// This parameter is required.
	//
	// example:
	//
	// manageNode
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The list of allowed CIDR blocks. This parameter is required. At least one allowed CIDR block must be specified (such as 0.0.0.0/0). If this parameter is not specified, the API returns InvalidParam (400).
	SecurityGroupProbeIpList []*string `json:"SecurityGroupProbeIpList,omitempty" xml:"SecurityGroupProbeIpList,omitempty" type:"Repeated"`
}

func (s CreateHoneypotNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateHoneypotNodeRequest) GetAllowHoneypotAccessInternet() *bool {
	return s.AllowHoneypotAccessInternet
}

func (s *CreateHoneypotNodeRequest) GetAvailableProbeNum() *int32 {
	return s.AvailableProbeNum
}

func (s *CreateHoneypotNodeRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *CreateHoneypotNodeRequest) GetSecurityGroupProbeIpList() []*string {
	return s.SecurityGroupProbeIpList
}

func (s *CreateHoneypotNodeRequest) SetAllowHoneypotAccessInternet(v bool) *CreateHoneypotNodeRequest {
	s.AllowHoneypotAccessInternet = &v
	return s
}

func (s *CreateHoneypotNodeRequest) SetAvailableProbeNum(v int32) *CreateHoneypotNodeRequest {
	s.AvailableProbeNum = &v
	return s
}

func (s *CreateHoneypotNodeRequest) SetNodeName(v string) *CreateHoneypotNodeRequest {
	s.NodeName = &v
	return s
}

func (s *CreateHoneypotNodeRequest) SetSecurityGroupProbeIpList(v []*string) *CreateHoneypotNodeRequest {
	s.SecurityGroupProbeIpList = v
	return s
}

func (s *CreateHoneypotNodeRequest) Validate() error {
	return dara.Validate(s)
}
