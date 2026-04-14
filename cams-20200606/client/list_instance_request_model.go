// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelType(v string) *ListInstanceRequest
	GetChannelType() *string
	SetFilterStr(v string) *ListInstanceRequest
	GetFilterStr() *string
	SetInstanceId(v string) *ListInstanceRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ListInstanceRequest
	GetInstanceName() *string
	SetPageIndex(v int64) *ListInstanceRequest
	GetPageIndex() *int64
	SetPageSize(v int64) *ListInstanceRequest
	GetPageSize() *int64
	SetResourceGroupId(v string) *ListInstanceRequest
	GetResourceGroupId() *string
	SetSubmitTime(v string) *ListInstanceRequest
	GetSubmitTime() *string
}

type ListInstanceRequest struct {
	// example:
	//
	// VIBER
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// example:
	//
	// aa
	FilterStr *string `json:"FilterStr,omitempty" xml:"FilterStr,omitempty"`
	// example:
	//
	// r-uf6wd7pkyjwxvlxfhk
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// viber_ins
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 92
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 87
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 11
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 2023-12-12 00:00:00
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s ListInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *ListInstanceRequest) GetFilterStr() *string {
	return s.FilterStr
}

func (s *ListInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstanceRequest) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *ListInstanceRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstanceRequest) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *ListInstanceRequest) SetChannelType(v string) *ListInstanceRequest {
	s.ChannelType = &v
	return s
}

func (s *ListInstanceRequest) SetFilterStr(v string) *ListInstanceRequest {
	s.FilterStr = &v
	return s
}

func (s *ListInstanceRequest) SetInstanceId(v string) *ListInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceRequest) SetInstanceName(v string) *ListInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstanceRequest) SetPageIndex(v int64) *ListInstanceRequest {
	s.PageIndex = &v
	return s
}

func (s *ListInstanceRequest) SetPageSize(v int64) *ListInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceRequest) SetResourceGroupId(v string) *ListInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstanceRequest) SetSubmitTime(v string) *ListInstanceRequest {
	s.SubmitTime = &v
	return s
}

func (s *ListInstanceRequest) Validate() error {
	return dara.Validate(s)
}
