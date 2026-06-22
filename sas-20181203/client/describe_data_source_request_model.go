// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigType(v string) *DescribeDataSourceRequest
	GetConfigType() *string
	SetLang(v string) *DescribeDataSourceRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeDataSourceRequest
	GetSourceIp() *string
}

type DescribeDataSourceRequest struct {
	// The type of data source to retrieve. This parameter is currently fixed to DingTalk.
	//
	// example:
	//
	// DingTalk
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 39.155.*.*
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *DescribeDataSourceRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDataSourceRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDataSourceRequest) SetConfigType(v string) *DescribeDataSourceRequest {
	s.ConfigType = &v
	return s
}

func (s *DescribeDataSourceRequest) SetLang(v string) *DescribeDataSourceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataSourceRequest) SetSourceIp(v string) *DescribeDataSourceRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
