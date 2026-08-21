// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeRegionsRequest
	GetAcceptLanguage() *string
	SetAuthorizedUserId(v int64) *DescribeRegionsRequest
	GetAuthorizedUserId() *int64
	SetLang(v string) *DescribeRegionsRequest
	GetLang() *string
	SetScene(v string) *DescribeRegionsRequest
	GetScene() *string
	SetUserClientIp(v string) *DescribeRegionsRequest
	GetUserClientIp() *string
	SetVpcType(v string) *DescribeRegionsRequest
	GetVpcType() *string
}

type DescribeRegionsRequest struct {
	// The supported language. Valid values:
	//
	// - zh-CN: Chinese
	//
	// - en-US: English
	//
	// Default value: en-US.
	//
	// > This parameter has a higher priority than Lang.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the Alibaba Cloud account to which the authorized resources belong.
	//
	// example:
	//
	// 141339776561****
	AuthorizedUserId *int64 `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// Default value: **en**.
	//
	// > This parameter has a lower priority than AcceptLanguage.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The scenario. Valid values:
	//
	// - AUTH: built-in authoritative server
	//
	// - FWD: forwarding management
	//
	// - RA: traffic analysis
	//
	// example:
	//
	// AUTH
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The IP address of the user.
	//
	// example:
	//
	// 192.0.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The type of the VPC. Valid values:
	//
	// - STANDARD: standard VPC.
	//
	// - EDS: Elastic Desktop Service (EDS) VPC.
	//
	// example:
	//
	// STANDARD
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeRegionsRequest) GetAuthorizedUserId() *int64 {
	return s.AuthorizedUserId
}

func (s *DescribeRegionsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRegionsRequest) GetScene() *string {
	return s.Scene
}

func (s *DescribeRegionsRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DescribeRegionsRequest) GetVpcType() *string {
	return s.VpcType
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetAuthorizedUserId(v int64) *DescribeRegionsRequest {
	s.AuthorizedUserId = &v
	return s
}

func (s *DescribeRegionsRequest) SetLang(v string) *DescribeRegionsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRegionsRequest) SetScene(v string) *DescribeRegionsRequest {
	s.Scene = &v
	return s
}

func (s *DescribeRegionsRequest) SetUserClientIp(v string) *DescribeRegionsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeRegionsRequest) SetVpcType(v string) *DescribeRegionsRequest {
	s.VpcType = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
