// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourceDataDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDataSourceDataDownloadUrlRequest
	GetLang() *string
	SetDataSourceId(v int64) *DescribeDataSourceDataDownloadUrlRequest
	GetDataSourceId() *int64
	SetRegId(v string) *DescribeDataSourceDataDownloadUrlRequest
	GetRegId() *string
}

type DescribeDataSourceDataDownloadUrlRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 69
	DataSourceId *int64 `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// The region code.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeDataSourceDataDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceDataDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceDataDownloadUrlRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDataSourceDataDownloadUrlRequest) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *DescribeDataSourceDataDownloadUrlRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeDataSourceDataDownloadUrlRequest) SetLang(v string) *DescribeDataSourceDataDownloadUrlRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataSourceDataDownloadUrlRequest) SetDataSourceId(v int64) *DescribeDataSourceDataDownloadUrlRequest {
	s.DataSourceId = &v
	return s
}

func (s *DescribeDataSourceDataDownloadUrlRequest) SetRegId(v string) *DescribeDataSourceDataDownloadUrlRequest {
	s.RegId = &v
	return s
}

func (s *DescribeDataSourceDataDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
