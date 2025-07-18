// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamingDataServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListStreamingDataServicesRequest
	GetDBInstanceId() *string
	SetPageNumber(v int32) *ListStreamingDataServicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListStreamingDataServicesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListStreamingDataServicesRequest
	GetRegionId() *string
}

type ListStreamingDataServicesRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp10g78o9807yv9h3
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Page number, greater than 0 and not exceeding the maximum value of Integer, default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of records per page, with the following values:
	//
	// - 30 (default)
	//
	// - 50
	//
	// - 100
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListStreamingDataServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStreamingDataServicesRequest) GoString() string {
	return s.String()
}

func (s *ListStreamingDataServicesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListStreamingDataServicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListStreamingDataServicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListStreamingDataServicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStreamingDataServicesRequest) SetDBInstanceId(v string) *ListStreamingDataServicesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListStreamingDataServicesRequest) SetPageNumber(v int32) *ListStreamingDataServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStreamingDataServicesRequest) SetPageSize(v int32) *ListStreamingDataServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListStreamingDataServicesRequest) SetRegionId(v string) *ListStreamingDataServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListStreamingDataServicesRequest) Validate() error {
	return dara.Validate(s)
}
