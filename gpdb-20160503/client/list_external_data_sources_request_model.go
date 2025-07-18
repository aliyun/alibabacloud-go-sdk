// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalDataSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListExternalDataSourcesRequest
	GetDBInstanceId() *string
	SetPageNumber(v int32) *ListExternalDataSourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListExternalDataSourcesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListExternalDataSourcesRequest
	GetRegionId() *string
}

type ListExternalDataSourcesRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: 30.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListExternalDataSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExternalDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListExternalDataSourcesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListExternalDataSourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExternalDataSourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExternalDataSourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListExternalDataSourcesRequest) SetDBInstanceId(v string) *ListExternalDataSourcesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListExternalDataSourcesRequest) SetPageNumber(v int32) *ListExternalDataSourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExternalDataSourcesRequest) SetPageSize(v int32) *ListExternalDataSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListExternalDataSourcesRequest) SetRegionId(v string) *ListExternalDataSourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListExternalDataSourcesRequest) Validate() error {
	return dara.Validate(s)
}
