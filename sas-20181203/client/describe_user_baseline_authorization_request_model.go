// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBaselineAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeUserBaselineAuthorizationRequest
	GetLang() *string
	SetResourceOwnerId(v int64) *DescribeUserBaselineAuthorizationRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeUserBaselineAuthorizationRequest
	GetSourceIp() *string
}

type DescribeUserBaselineAuthorizationRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// 资产所有者ID。
	//
	// example:
	//
	// 1519712934213764
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the asset owner.
	//
	// example:
	//
	// 1.2.3.4
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeUserBaselineAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBaselineAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserBaselineAuthorizationRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeUserBaselineAuthorizationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeUserBaselineAuthorizationRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeUserBaselineAuthorizationRequest) SetLang(v string) *DescribeUserBaselineAuthorizationRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUserBaselineAuthorizationRequest) SetResourceOwnerId(v int64) *DescribeUserBaselineAuthorizationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUserBaselineAuthorizationRequest) SetSourceIp(v string) *DescribeUserBaselineAuthorizationRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUserBaselineAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
