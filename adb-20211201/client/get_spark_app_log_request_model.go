// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetSparkAppLogRequest
	GetAppId() *string
	SetDBClusterId(v string) *GetSparkAppLogRequest
	GetDBClusterId() *string
	SetLogLength(v int64) *GetSparkAppLogRequest
	GetLogLength() *int64
	SetPageNumber(v int32) *GetSparkAppLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetSparkAppLogRequest
	GetPageSize() *int32
}

type GetSparkAppLogRequest struct {
	// The Spark application ID.
	//
	// > You can call the [ListSparkApps](https://help.aliyun.com/document_detail/612475.html) operation to query the Spark application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s202206061441hz22a35ab000****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// example:
	//
	// amv-2ze6fl8ud7t***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The number of log entries to return. Valid values: 1 to 500. Default value: 300.
	//
	// example:
	//
	// 20
	LogLength *int64 `json:"LogLength,omitempty" xml:"LogLength,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 500
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetSparkAppLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppLogRequest) GoString() string {
	return s.String()
}

func (s *GetSparkAppLogRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetSparkAppLogRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetSparkAppLogRequest) GetLogLength() *int64 {
	return s.LogLength
}

func (s *GetSparkAppLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetSparkAppLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSparkAppLogRequest) SetAppId(v string) *GetSparkAppLogRequest {
	s.AppId = &v
	return s
}

func (s *GetSparkAppLogRequest) SetDBClusterId(v string) *GetSparkAppLogRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkAppLogRequest) SetLogLength(v int64) *GetSparkAppLogRequest {
	s.LogLength = &v
	return s
}

func (s *GetSparkAppLogRequest) SetPageNumber(v int32) *GetSparkAppLogRequest {
	s.PageNumber = &v
	return s
}

func (s *GetSparkAppLogRequest) SetPageSize(v int32) *GetSparkAppLogRequest {
	s.PageSize = &v
	return s
}

func (s *GetSparkAppLogRequest) Validate() error {
	return dara.Validate(s)
}
