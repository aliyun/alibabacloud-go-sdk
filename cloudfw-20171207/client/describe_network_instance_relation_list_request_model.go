// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInstanceRelationListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectType(v string) *DescribeNetworkInstanceRelationListRequest
	GetConnectType() *string
	SetFirewallConfigureStatus(v string) *DescribeNetworkInstanceRelationListRequest
	GetFirewallConfigureStatus() *string
	SetLang(v string) *DescribeNetworkInstanceRelationListRequest
	GetLang() *string
}

type DescribeNetworkInstanceRelationListRequest struct {
	// example:
	//
	// cen
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// example:
	//
	// notconfigured
	FirewallConfigureStatus *string `json:"FirewallConfigureStatus,omitempty" xml:"FirewallConfigureStatus,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeNetworkInstanceRelationListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInstanceRelationListRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInstanceRelationListRequest) GetConnectType() *string {
	return s.ConnectType
}

func (s *DescribeNetworkInstanceRelationListRequest) GetFirewallConfigureStatus() *string {
	return s.FirewallConfigureStatus
}

func (s *DescribeNetworkInstanceRelationListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNetworkInstanceRelationListRequest) SetConnectType(v string) *DescribeNetworkInstanceRelationListRequest {
	s.ConnectType = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListRequest) SetFirewallConfigureStatus(v string) *DescribeNetworkInstanceRelationListRequest {
	s.FirewallConfigureStatus = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListRequest) SetLang(v string) *DescribeNetworkInstanceRelationListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListRequest) Validate() error {
	return dara.Validate(s)
}
