// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListModelServicesRequest
	GetDBInstanceId() *string
	SetPageNumber(v int32) *ListModelServicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelServicesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListModelServicesRequest
	GetRegionId() *string
}

type ListModelServicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListModelServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelServicesRequest) GoString() string {
	return s.String()
}

func (s *ListModelServicesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListModelServicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelServicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelServicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListModelServicesRequest) SetDBInstanceId(v string) *ListModelServicesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListModelServicesRequest) SetPageNumber(v int32) *ListModelServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelServicesRequest) SetPageSize(v int32) *ListModelServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelServicesRequest) SetRegionId(v string) *ListModelServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListModelServicesRequest) Validate() error {
	return dara.Validate(s)
}
