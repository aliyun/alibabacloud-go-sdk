// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInvocationResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *DescribeRCInvocationResultsRequest
	GetCommandId() *string
	SetContentEncoding(v string) *DescribeRCInvocationResultsRequest
	GetContentEncoding() *string
	SetIncludeHistory(v bool) *DescribeRCInvocationResultsRequest
	GetIncludeHistory() *bool
	SetInstanceId(v string) *DescribeRCInvocationResultsRequest
	GetInstanceId() *string
	SetInvokeId(v string) *DescribeRCInvocationResultsRequest
	GetInvokeId() *string
	SetInvokeRecordStatus(v string) *DescribeRCInvocationResultsRequest
	GetInvokeRecordStatus() *string
	SetMaxResults(v int32) *DescribeRCInvocationResultsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeRCInvocationResultsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeRCInvocationResultsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRCInvocationResultsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeRCInvocationResultsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeRCInvocationResultsRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeRCInvocationResultsRequestTag) *DescribeRCInvocationResultsRequest
	GetTag() []*DescribeRCInvocationResultsRequestTag
}

type DescribeRCInvocationResultsRequest struct {
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
	ResourceGroupId *string                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*DescribeRCInvocationResultsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeRCInvocationResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInvocationResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCInvocationResultsRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *DescribeRCInvocationResultsRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *DescribeRCInvocationResultsRequest) GetIncludeHistory() *bool {
	return s.IncludeHistory
}

func (s *DescribeRCInvocationResultsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCInvocationResultsRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeRCInvocationResultsRequest) GetInvokeRecordStatus() *string {
	return s.InvokeRecordStatus
}

func (s *DescribeRCInvocationResultsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRCInvocationResultsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRCInvocationResultsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRCInvocationResultsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRCInvocationResultsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCInvocationResultsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRCInvocationResultsRequest) GetTag() []*DescribeRCInvocationResultsRequestTag {
	return s.Tag
}

func (s *DescribeRCInvocationResultsRequest) SetCommandId(v string) *DescribeRCInvocationResultsRequest {
	s.CommandId = &v
	return s
}

func (s *DescribeRCInvocationResultsRequest) SetContentEncoding(v string) *DescribeRCInvocationResultsRequest {
	s.ContentEncoding = &v
	return s
}

func (s *DescribeRCInvocationResultsRequest) SetIncludeHistory(v bool) *DescribeRCInvocationResultsRequest {
	s.IncludeHistory = &v
	return s
}

func (s *DescribeRCInvocationResultsRequest) SetInstanceId(v string) *DescribeRCInvocationResultsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCInvocationResultsRequest) SetInvokeId(v string) *DescribeRCInvocationResultsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeRCInvocationResultsRequest) SetInvokeRecordStatus(v string) *DescribeRCInvocationResultsRequest {
	s.InvokeRecordStatus = &v
	return s
}

func (s *DescribeRCInvocationResultsRequest) SetMaxResults(v int32) *DescribeRCInvocationResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeRCInvocationResultsRequest) SetNextToken(v string) *DescribeRCInvocationResultsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeRCInvocationResultsRequest) SetPageNumber(v int32) *DescribeRCInvocationResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCInvocationResultsRequest) SetPageSize(v int32) *DescribeRCInvocationResultsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRCInvocationResultsRequest) SetRegionId(v string) *DescribeRCInvocationResultsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCInvocationResultsRequest) SetResourceGroupId(v string) *DescribeRCInvocationResultsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRCInvocationResultsRequest) SetTag(v []*DescribeRCInvocationResultsRequestTag) *DescribeRCInvocationResultsRequest {
	s.Tag = v
	return s
}

func (s *DescribeRCInvocationResultsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCInvocationResultsRequestTag struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRCInvocationResultsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInvocationResultsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeRCInvocationResultsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeRCInvocationResultsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeRCInvocationResultsRequestTag) SetKey(v string) *DescribeRCInvocationResultsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeRCInvocationResultsRequestTag) SetValue(v string) *DescribeRCInvocationResultsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeRCInvocationResultsRequestTag) Validate() error {
	return dara.Validate(s)
}
