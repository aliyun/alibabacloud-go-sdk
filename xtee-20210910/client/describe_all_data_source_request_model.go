// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAllDataSourceRequest
	GetLang() *string
	SetRegId(v string) *DescribeAllDataSourceRequest
	GetRegId() *string
}

type DescribeAllDataSourceRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
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
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeAllDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAllDataSourceRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAllDataSourceRequest) SetLang(v string) *DescribeAllDataSourceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetRegId(v string) *DescribeAllDataSourceRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAllDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
