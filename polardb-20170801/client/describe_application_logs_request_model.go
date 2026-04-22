// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationLogsRequest
	GetApplicationId() *string
	SetComponentName(v string) *DescribeApplicationLogsRequest
	GetComponentName() *string
	SetContainerName(v string) *DescribeApplicationLogsRequest
	GetContainerName() *string
	SetEndTime(v string) *DescribeApplicationLogsRequest
	GetEndTime() *string
	SetKeyword(v string) *DescribeApplicationLogsRequest
	GetKeyword() *string
	SetLevel(v string) *DescribeApplicationLogsRequest
	GetLevel() *string
	SetOwnerAccount(v string) *DescribeApplicationLogsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeApplicationLogsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeApplicationLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApplicationLogsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeApplicationLogsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeApplicationLogsRequest
	GetResourceOwnerAccount() *string
	SetStartTime(v string) *DescribeApplicationLogsRequest
	GetStartTime() *string
	SetType(v string) *DescribeApplicationLogsRequest
	GetType() *string
}

type DescribeApplicationLogsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2026-03-25T02:11:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// Config
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// WARN
	Level        *string `json:"Level,omitempty" xml:"Level,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 3
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2026-03-25T01:57:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// gateway
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApplicationLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationLogsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationLogsRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *DescribeApplicationLogsRequest) GetContainerName() *string {
	return s.ContainerName
}

func (s *DescribeApplicationLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApplicationLogsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeApplicationLogsRequest) GetLevel() *string {
	return s.Level
}

func (s *DescribeApplicationLogsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeApplicationLogsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeApplicationLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApplicationLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApplicationLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationLogsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeApplicationLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApplicationLogsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeApplicationLogsRequest) SetApplicationId(v string) *DescribeApplicationLogsRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationLogsRequest) SetComponentName(v string) *DescribeApplicationLogsRequest {
	s.ComponentName = &v
	return s
}

func (s *DescribeApplicationLogsRequest) SetContainerName(v string) *DescribeApplicationLogsRequest {
	s.ContainerName = &v
	return s
}

func (s *DescribeApplicationLogsRequest) SetEndTime(v string) *DescribeApplicationLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApplicationLogsRequest) SetKeyword(v string) *DescribeApplicationLogsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeApplicationLogsRequest) SetLevel(v string) *DescribeApplicationLogsRequest {
	s.Level = &v
	return s
}

func (s *DescribeApplicationLogsRequest) SetOwnerAccount(v string) *DescribeApplicationLogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeApplicationLogsRequest) SetOwnerId(v int64) *DescribeApplicationLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeApplicationLogsRequest) SetPageNumber(v int32) *DescribeApplicationLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApplicationLogsRequest) SetPageSize(v int32) *DescribeApplicationLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationLogsRequest) SetRegionId(v string) *DescribeApplicationLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationLogsRequest) SetResourceOwnerAccount(v string) *DescribeApplicationLogsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeApplicationLogsRequest) SetStartTime(v string) *DescribeApplicationLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApplicationLogsRequest) SetType(v string) *DescribeApplicationLogsRequest {
	s.Type = &v
	return s
}

func (s *DescribeApplicationLogsRequest) Validate() error {
	return dara.Validate(s)
}
