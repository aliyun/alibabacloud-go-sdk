// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListCloudAccountsRequestFilter) *ListCloudAccountsRequest
	GetFilter() []*ListCloudAccountsRequestFilter
	SetInstanceId(v string) *ListCloudAccountsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListCloudAccountsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCloudAccountsRequest
	GetNextToken() *string
}

type ListCloudAccountsRequest struct {
	Filter []*ListCloudAccountsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 分页查询时每页行数。默认值为20，最大值为100。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 查询凭证（Token），取值为上一次API调用返回的NextToken参数值。
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListCloudAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListCloudAccountsRequest) GetFilter() []*ListCloudAccountsRequestFilter {
	return s.Filter
}

func (s *ListCloudAccountsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloudAccountsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCloudAccountsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCloudAccountsRequest) SetFilter(v []*ListCloudAccountsRequestFilter) *ListCloudAccountsRequest {
	s.Filter = v
	return s
}

func (s *ListCloudAccountsRequest) SetInstanceId(v string) *ListCloudAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCloudAccountsRequest) SetMaxResults(v int32) *ListCloudAccountsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCloudAccountsRequest) SetNextToken(v string) *ListCloudAccountsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCloudAccountsRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloudAccountsRequestFilter struct {
	// example:
	//
	// CloudAccountId
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListCloudAccountsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListCloudAccountsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListCloudAccountsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListCloudAccountsRequestFilter) SetName(v string) *ListCloudAccountsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListCloudAccountsRequestFilter) SetValue(v []*string) *ListCloudAccountsRequestFilter {
	s.Value = v
	return s
}

func (s *ListCloudAccountsRequestFilter) Validate() error {
	return dara.Validate(s)
}
