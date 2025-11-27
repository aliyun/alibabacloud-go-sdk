// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInvocationResultsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *DescribeRCInvocationResultsShrinkRequest
	GetCommandId() *string
	SetContentEncoding(v string) *DescribeRCInvocationResultsShrinkRequest
	GetContentEncoding() *string
	SetIncludeHistory(v bool) *DescribeRCInvocationResultsShrinkRequest
	GetIncludeHistory() *bool
	SetInstanceId(v string) *DescribeRCInvocationResultsShrinkRequest
	GetInstanceId() *string
	SetInvokeId(v string) *DescribeRCInvocationResultsShrinkRequest
	GetInvokeId() *string
	SetInvokeRecordStatus(v string) *DescribeRCInvocationResultsShrinkRequest
	GetInvokeRecordStatus() *string
	SetMaxResults(v int32) *DescribeRCInvocationResultsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeRCInvocationResultsShrinkRequest
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeRCInvocationResultsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRCInvocationResultsShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeRCInvocationResultsShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeRCInvocationResultsShrinkRequest
	GetResourceGroupId() *string
	SetTagShrink(v string) *DescribeRCInvocationResultsShrinkRequest
	GetTagShrink() *string
}

type DescribeRCInvocationResultsShrinkRequest struct {
	// example:
	//
	// c-7d2a745b412b4601b2d47f6a768d****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// example:
	//
	// Base64
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// example:
	//
	// false
	IncludeHistory *bool `json:"IncludeHistory,omitempty" xml:"IncludeHistory,omitempty"`
	// example:
	//
	// rc-i322y2t562oh7o******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// t-7d2a745b412b4601b2d47f6a768d****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// example:
	//
	// Running
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// None
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// None
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmx****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TagShrink       *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DescribeRCInvocationResultsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInvocationResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCInvocationResultsShrinkRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *DescribeRCInvocationResultsShrinkRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *DescribeRCInvocationResultsShrinkRequest) GetIncludeHistory() *bool {
	return s.IncludeHistory
}

func (s *DescribeRCInvocationResultsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCInvocationResultsShrinkRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeRCInvocationResultsShrinkRequest) GetInvokeRecordStatus() *string {
	return s.InvokeRecordStatus
}

func (s *DescribeRCInvocationResultsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRCInvocationResultsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRCInvocationResultsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRCInvocationResultsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRCInvocationResultsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCInvocationResultsShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRCInvocationResultsShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *DescribeRCInvocationResultsShrinkRequest) SetCommandId(v string) *DescribeRCInvocationResultsShrinkRequest {
	s.CommandId = &v
	return s
}

func (s *DescribeRCInvocationResultsShrinkRequest) SetContentEncoding(v string) *DescribeRCInvocationResultsShrinkRequest {
	s.ContentEncoding = &v
	return s
}

func (s *DescribeRCInvocationResultsShrinkRequest) SetIncludeHistory(v bool) *DescribeRCInvocationResultsShrinkRequest {
	s.IncludeHistory = &v
	return s
}

func (s *DescribeRCInvocationResultsShrinkRequest) SetInstanceId(v string) *DescribeRCInvocationResultsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCInvocationResultsShrinkRequest) SetInvokeId(v string) *DescribeRCInvocationResultsShrinkRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeRCInvocationResultsShrinkRequest) SetInvokeRecordStatus(v string) *DescribeRCInvocationResultsShrinkRequest {
	s.InvokeRecordStatus = &v
	return s
}

func (s *DescribeRCInvocationResultsShrinkRequest) SetMaxResults(v int32) *DescribeRCInvocationResultsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeRCInvocationResultsShrinkRequest) SetNextToken(v string) *DescribeRCInvocationResultsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeRCInvocationResultsShrinkRequest) SetPageNumber(v int32) *DescribeRCInvocationResultsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCInvocationResultsShrinkRequest) SetPageSize(v int32) *DescribeRCInvocationResultsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRCInvocationResultsShrinkRequest) SetRegionId(v string) *DescribeRCInvocationResultsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCInvocationResultsShrinkRequest) SetResourceGroupId(v string) *DescribeRCInvocationResultsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRCInvocationResultsShrinkRequest) SetTagShrink(v string) *DescribeRCInvocationResultsShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *DescribeRCInvocationResultsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
