// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *DescribeUserConfigsRequest
	GetConfig() *string
	SetOwnerId(v int64) *DescribeUserConfigsRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeUserConfigsRequest
	GetSecurityToken() *string
}

type DescribeUserConfigsRequest struct {
	// The feature whose configurations you want to query. You can specify only one feature in each request. Valid values: oss, green_manager, waf, cc_rule, ddos_dispatch, edge_safe, blocked_regions, http_acl_policy, bot_manager, and ip_reputation.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss
	Config        *string `json:"Config,omitempty" xml:"Config,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeUserConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserConfigsRequest) GetConfig() *string {
	return s.Config
}

func (s *DescribeUserConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUserConfigsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeUserConfigsRequest) SetConfig(v string) *DescribeUserConfigsRequest {
	s.Config = &v
	return s
}

func (s *DescribeUserConfigsRequest) SetOwnerId(v int64) *DescribeUserConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserConfigsRequest) SetSecurityToken(v string) *DescribeUserConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeUserConfigsRequest) Validate() error {
	return dara.Validate(s)
}
