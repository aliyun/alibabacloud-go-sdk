// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowModify(v bool) *DescribeInstanceConfigsRequest
	GetAllowModify() *bool
	SetConfigKey(v string) *DescribeInstanceConfigsRequest
	GetConfigKey() *string
	SetConfigType(v string) *DescribeInstanceConfigsRequest
	GetConfigType() *string
	SetDescription(v string) *DescribeInstanceConfigsRequest
	GetDescription() *string
	SetInstanceId(v string) *DescribeInstanceConfigsRequest
	GetInstanceId() *string
	SetNeedTotal(v bool) *DescribeInstanceConfigsRequest
	GetNeedTotal() *bool
	SetNodeGroupId(v string) *DescribeInstanceConfigsRequest
	GetNodeGroupId() *string
	SetPageNumber(v int32) *DescribeInstanceConfigsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceConfigsRequest
	GetPageSize() *int32
}

type DescribeInstanceConfigsRequest struct {
	// example:
	//
	// true
	AllowModify *bool `json:"AllowModify,omitempty" xml:"AllowModify,omitempty"`
	// example:
	//
	// enable_udf
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// example:
	//
	// FE
	ConfigType  *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	NeedTotal *bool `json:"NeedTotal,omitempty" xml:"NeedTotal,omitempty"`
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeInstanceConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceConfigsRequest) GetAllowModify() *bool {
	return s.AllowModify
}

func (s *DescribeInstanceConfigsRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *DescribeInstanceConfigsRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *DescribeInstanceConfigsRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceConfigsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceConfigsRequest) GetNeedTotal() *bool {
	return s.NeedTotal
}

func (s *DescribeInstanceConfigsRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DescribeInstanceConfigsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceConfigsRequest) SetAllowModify(v bool) *DescribeInstanceConfigsRequest {
	s.AllowModify = &v
	return s
}

func (s *DescribeInstanceConfigsRequest) SetConfigKey(v string) *DescribeInstanceConfigsRequest {
	s.ConfigKey = &v
	return s
}

func (s *DescribeInstanceConfigsRequest) SetConfigType(v string) *DescribeInstanceConfigsRequest {
	s.ConfigType = &v
	return s
}

func (s *DescribeInstanceConfigsRequest) SetDescription(v string) *DescribeInstanceConfigsRequest {
	s.Description = &v
	return s
}

func (s *DescribeInstanceConfigsRequest) SetInstanceId(v string) *DescribeInstanceConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceConfigsRequest) SetNeedTotal(v bool) *DescribeInstanceConfigsRequest {
	s.NeedTotal = &v
	return s
}

func (s *DescribeInstanceConfigsRequest) SetNodeGroupId(v string) *DescribeInstanceConfigsRequest {
	s.NodeGroupId = &v
	return s
}

func (s *DescribeInstanceConfigsRequest) SetPageNumber(v int32) *DescribeInstanceConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceConfigsRequest) SetPageSize(v int32) *DescribeInstanceConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceConfigsRequest) Validate() error {
	return dara.Validate(s)
}
