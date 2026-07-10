// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeLangfuseUsersRequest
	GetDBInstanceId() *string
	SetEmail(v string) *DescribeLangfuseUsersRequest
	GetEmail() *string
	SetName(v string) *DescribeLangfuseUsersRequest
	GetName() *string
	SetPageNumber(v int64) *DescribeLangfuseUsersRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeLangfuseUsersRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeLangfuseUsersRequest
	GetRegionId() *string
}

type DescribeLangfuseUsersRequest struct {
	// The Langfuse instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lfs-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// john@company.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The username.
	//
	// example:
	//
	// john
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLangfuseUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseUsersRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeLangfuseUsersRequest) GetEmail() *string {
	return s.Email
}

func (s *DescribeLangfuseUsersRequest) GetName() *string {
	return s.Name
}

func (s *DescribeLangfuseUsersRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLangfuseUsersRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLangfuseUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLangfuseUsersRequest) SetDBInstanceId(v string) *DescribeLangfuseUsersRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLangfuseUsersRequest) SetEmail(v string) *DescribeLangfuseUsersRequest {
	s.Email = &v
	return s
}

func (s *DescribeLangfuseUsersRequest) SetName(v string) *DescribeLangfuseUsersRequest {
	s.Name = &v
	return s
}

func (s *DescribeLangfuseUsersRequest) SetPageNumber(v int64) *DescribeLangfuseUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLangfuseUsersRequest) SetPageSize(v int64) *DescribeLangfuseUsersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLangfuseUsersRequest) SetRegionId(v string) *DescribeLangfuseUsersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLangfuseUsersRequest) Validate() error {
	return dara.Validate(s)
}
