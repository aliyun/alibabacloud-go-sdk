// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatAclPageStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeNatAclPageStatusRequest
	GetLang() *string
}

type DescribeNatAclPageStatusRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeNatAclPageStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatAclPageStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatAclPageStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNatAclPageStatusRequest) SetLang(v string) *DescribeNatAclPageStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNatAclPageStatusRequest) Validate() error {
	return dara.Validate(s)
}
