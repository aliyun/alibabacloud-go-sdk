// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudMigrationPrecheckResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeCloudMigrationPrecheckResultResponseBodyItems) *DescribeCloudMigrationPrecheckResultResponseBody
	GetItems() []*DescribeCloudMigrationPrecheckResultResponseBodyItems
	SetPageNumber(v int64) *DescribeCloudMigrationPrecheckResultResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCloudMigrationPrecheckResultResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeCloudMigrationPrecheckResultResponseBody
	GetRequestId() *string
	SetTotalSize(v int32) *DescribeCloudMigrationPrecheckResultResponseBody
	GetTotalSize() *int32
}

type DescribeCloudMigrationPrecheckResultResponseBody struct {
	Items      []*DescribeCloudMigrationPrecheckResultResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int64                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalSize  *int32                                                   `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeCloudMigrationPrecheckResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudMigrationPrecheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudMigrationPrecheckResultResponseBody) GetItems() []*DescribeCloudMigrationPrecheckResultResponseBodyItems {
	return s.Items
}

func (s *DescribeCloudMigrationPrecheckResultResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCloudMigrationPrecheckResultResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCloudMigrationPrecheckResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudMigrationPrecheckResultResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *DescribeCloudMigrationPrecheckResultResponseBody) SetItems(v []*DescribeCloudMigrationPrecheckResultResponseBodyItems) *DescribeCloudMigrationPrecheckResultResponseBody {
	s.Items = v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBody) SetPageNumber(v int64) *DescribeCloudMigrationPrecheckResultResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBody) SetPageSize(v int64) *DescribeCloudMigrationPrecheckResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBody) SetRequestId(v string) *DescribeCloudMigrationPrecheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBody) SetTotalSize(v int32) *DescribeCloudMigrationPrecheckResultResponseBody {
	s.TotalSize = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudMigrationPrecheckResultResponseBodyItems struct {
	Detail             *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	GmtCreated         *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified        *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	SourceAccount      *string `json:"SourceAccount,omitempty" xml:"SourceAccount,omitempty"`
	SourceCategory     *string `json:"SourceCategory,omitempty" xml:"SourceCategory,omitempty"`
	SourceIpAddress    *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	SourcePassword     *string `json:"SourcePassword,omitempty" xml:"SourcePassword,omitempty"`
	SourcePort         *int64  `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	TargetEip          *string `json:"TargetEip,omitempty" xml:"TargetEip,omitempty"`
	TargetInstanceName *string `json:"TargetInstanceName,omitempty" xml:"TargetInstanceName,omitempty"`
	TaskId             *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName           *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeCloudMigrationPrecheckResultResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudMigrationPrecheckResultResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) GetDetail() *string {
	return s.Detail
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) GetSourceAccount() *string {
	return s.SourceAccount
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) GetSourceCategory() *string {
	return s.SourceCategory
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) GetSourceIpAddress() *string {
	return s.SourceIpAddress
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) GetSourcePassword() *string {
	return s.SourcePassword
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) GetSourcePort() *int64 {
	return s.SourcePort
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) GetTargetEip() *string {
	return s.TargetEip
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) GetTargetInstanceName() *string {
	return s.TargetInstanceName
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) SetDetail(v string) *DescribeCloudMigrationPrecheckResultResponseBodyItems {
	s.Detail = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) SetGmtCreated(v string) *DescribeCloudMigrationPrecheckResultResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) SetGmtModified(v string) *DescribeCloudMigrationPrecheckResultResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) SetSourceAccount(v string) *DescribeCloudMigrationPrecheckResultResponseBodyItems {
	s.SourceAccount = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) SetSourceCategory(v string) *DescribeCloudMigrationPrecheckResultResponseBodyItems {
	s.SourceCategory = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) SetSourceIpAddress(v string) *DescribeCloudMigrationPrecheckResultResponseBodyItems {
	s.SourceIpAddress = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) SetSourcePassword(v string) *DescribeCloudMigrationPrecheckResultResponseBodyItems {
	s.SourcePassword = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) SetSourcePort(v int64) *DescribeCloudMigrationPrecheckResultResponseBodyItems {
	s.SourcePort = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) SetTargetEip(v string) *DescribeCloudMigrationPrecheckResultResponseBodyItems {
	s.TargetEip = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) SetTargetInstanceName(v string) *DescribeCloudMigrationPrecheckResultResponseBodyItems {
	s.TargetInstanceName = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) SetTaskId(v int64) *DescribeCloudMigrationPrecheckResultResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) SetTaskName(v string) *DescribeCloudMigrationPrecheckResultResponseBodyItems {
	s.TaskName = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
