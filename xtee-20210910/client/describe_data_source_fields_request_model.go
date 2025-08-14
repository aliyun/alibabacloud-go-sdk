// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourceFieldsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDataSourceFieldsRequest
	GetLang() *string
	SetDataSourceCode(v string) *DescribeDataSourceFieldsRequest
	GetDataSourceCode() *string
	SetRegId(v string) *DescribeDataSourceFieldsRequest
	GetRegId() *string
}

type DescribeDataSourceFieldsRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Data source code
	//
	// This parameter is required.
	//
	// example:
	//
	// date_source
	DataSourceCode *string `json:"dataSourceCode,omitempty" xml:"dataSourceCode,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeDataSourceFieldsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceFieldsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceFieldsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDataSourceFieldsRequest) GetDataSourceCode() *string {
	return s.DataSourceCode
}

func (s *DescribeDataSourceFieldsRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeDataSourceFieldsRequest) SetLang(v string) *DescribeDataSourceFieldsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataSourceFieldsRequest) SetDataSourceCode(v string) *DescribeDataSourceFieldsRequest {
	s.DataSourceCode = &v
	return s
}

func (s *DescribeDataSourceFieldsRequest) SetRegId(v string) *DescribeDataSourceFieldsRequest {
	s.RegId = &v
	return s
}

func (s *DescribeDataSourceFieldsRequest) Validate() error {
	return dara.Validate(s)
}
