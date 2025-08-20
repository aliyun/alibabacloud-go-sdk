// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSparkAppAttemptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListSparkAppAttemptsRequest
	GetAppId() *string
	SetDBClusterId(v string) *ListSparkAppAttemptsRequest
	GetDBClusterId() *string
	SetPageNumber(v int64) *ListSparkAppAttemptsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSparkAppAttemptsRequest
	GetPageSize() *int64
}

type ListSparkAppAttemptsRequest struct {
	// The ID of the Spark application.
	//
	// > You can call the [ListSparkApps](https://help.aliyun.com/document_detail/455888.html) operation to query all application IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// s202204132018hzprec1ac****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// example:
	//
	// amv-uf6o6m8p6x***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The page number. The value must be an integer that is greater than 0. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **10*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSparkAppAttemptsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSparkAppAttemptsRequest) GoString() string {
	return s.String()
}

func (s *ListSparkAppAttemptsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListSparkAppAttemptsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListSparkAppAttemptsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSparkAppAttemptsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSparkAppAttemptsRequest) SetAppId(v string) *ListSparkAppAttemptsRequest {
	s.AppId = &v
	return s
}

func (s *ListSparkAppAttemptsRequest) SetDBClusterId(v string) *ListSparkAppAttemptsRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListSparkAppAttemptsRequest) SetPageNumber(v int64) *ListSparkAppAttemptsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSparkAppAttemptsRequest) SetPageSize(v int64) *ListSparkAppAttemptsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSparkAppAttemptsRequest) Validate() error {
	return dara.Validate(s)
}
