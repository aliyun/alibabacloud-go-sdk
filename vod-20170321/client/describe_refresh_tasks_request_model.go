// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRefreshTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeRefreshTasksRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeRefreshTasksRequest
	GetEndTime() *string
	SetObjectPath(v string) *DescribeRefreshTasksRequest
	GetObjectPath() *string
	SetObjectType(v string) *DescribeRefreshTasksRequest
	GetObjectType() *string
	SetOwnerAccount(v string) *DescribeRefreshTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *DescribeRefreshTasksRequest
	GetOwnerId() *string
	SetPageNumber(v int32) *DescribeRefreshTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRefreshTasksRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeRefreshTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *DescribeRefreshTasksRequest
	GetResourceOwnerId() *string
	SetStartTime(v string) *DescribeRefreshTasksRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeRefreshTasksRequest
	GetStatus() *string
	SetTaskId(v string) *DescribeRefreshTasksRequest
	GetTaskId() *string
}

type DescribeRefreshTasksRequest struct {
	DomainName           *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ObjectPath           *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType           *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeRefreshTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTasksRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeRefreshTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRefreshTasksRequest) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *DescribeRefreshTasksRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeRefreshTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRefreshTasksRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeRefreshTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRefreshTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRefreshTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRefreshTasksRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *DescribeRefreshTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRefreshTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeRefreshTasksRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeRefreshTasksRequest) SetDomainName(v string) *DescribeRefreshTasksRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetEndTime(v string) *DescribeRefreshTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetObjectPath(v string) *DescribeRefreshTasksRequest {
	s.ObjectPath = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetObjectType(v string) *DescribeRefreshTasksRequest {
	s.ObjectType = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetOwnerAccount(v string) *DescribeRefreshTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetOwnerId(v string) *DescribeRefreshTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetPageNumber(v int32) *DescribeRefreshTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetPageSize(v int32) *DescribeRefreshTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetResourceOwnerAccount(v string) *DescribeRefreshTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetResourceOwnerId(v string) *DescribeRefreshTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetStartTime(v string) *DescribeRefreshTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetStatus(v string) *DescribeRefreshTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetTaskId(v string) *DescribeRefreshTasksRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeRefreshTasksRequest) Validate() error {
	return dara.Validate(s)
}
