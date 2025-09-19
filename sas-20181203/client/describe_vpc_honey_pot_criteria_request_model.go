// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcHoneyPotCriteriaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVpcHoneyPotCriteriaRequest
	GetLang() *string
}

type DescribeVpcHoneyPotCriteriaRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeVpcHoneyPotCriteriaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcHoneyPotCriteriaRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotCriteriaRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcHoneyPotCriteriaRequest) SetLang(v string) *DescribeVpcHoneyPotCriteriaRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcHoneyPotCriteriaRequest) Validate() error {
	return dara.Validate(s)
}
