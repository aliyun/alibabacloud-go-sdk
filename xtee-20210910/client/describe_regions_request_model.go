// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRegionsRequest
	GetLang() *string
	SetRegId(v string) *DescribeRegionsRequest
	GetRegId() *string
}

type DescribeRegionsRequest struct {
	// Set the language type for requests and responses. The default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRegionsRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRegionsRequest) SetLang(v string) *DescribeRegionsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegId(v string) *DescribeRegionsRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
