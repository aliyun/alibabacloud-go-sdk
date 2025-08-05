// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefaultIPSConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDefaultIPSConfigRequest
	GetLang() *string
}

type DescribeDefaultIPSConfigRequest struct {
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default)
	//
	// 	- **en**
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeDefaultIPSConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefaultIPSConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefaultIPSConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDefaultIPSConfigRequest) SetLang(v string) *DescribeDefaultIPSConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDefaultIPSConfigRequest) Validate() error {
	return dara.Validate(s)
}
