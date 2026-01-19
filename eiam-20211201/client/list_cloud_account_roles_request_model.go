// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAccountRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountId(v string) *ListCloudAccountRolesRequest
	GetCloudAccountId() *string
	SetFilter(v []*ListCloudAccountRolesRequestFilter) *ListCloudAccountRolesRequest
	GetFilter() []*ListCloudAccountRolesRequestFilter
	SetInstanceId(v string) *ListCloudAccountRolesRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListCloudAccountRolesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCloudAccountRolesRequest
	GetNextToken() *string
}

type ListCloudAccountRolesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string                               `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	Filter         []*ListCloudAccountRolesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
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

func (s ListCloudAccountRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountRolesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudAccountRolesRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *ListCloudAccountRolesRequest) GetFilter() []*ListCloudAccountRolesRequestFilter {
	return s.Filter
}

func (s *ListCloudAccountRolesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloudAccountRolesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCloudAccountRolesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCloudAccountRolesRequest) SetCloudAccountId(v string) *ListCloudAccountRolesRequest {
	s.CloudAccountId = &v
	return s
}

func (s *ListCloudAccountRolesRequest) SetFilter(v []*ListCloudAccountRolesRequestFilter) *ListCloudAccountRolesRequest {
	s.Filter = v
	return s
}

func (s *ListCloudAccountRolesRequest) SetInstanceId(v string) *ListCloudAccountRolesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCloudAccountRolesRequest) SetMaxResults(v int32) *ListCloudAccountRolesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCloudAccountRolesRequest) SetNextToken(v string) *ListCloudAccountRolesRequest {
	s.NextToken = &v
	return s
}

func (s *ListCloudAccountRolesRequest) Validate() error {
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

type ListCloudAccountRolesRequestFilter struct {
	// example:
	//
	// CloudAccountRoleId
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListCloudAccountRolesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountRolesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListCloudAccountRolesRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListCloudAccountRolesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListCloudAccountRolesRequestFilter) SetName(v string) *ListCloudAccountRolesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListCloudAccountRolesRequestFilter) SetValue(v []*string) *ListCloudAccountRolesRequestFilter {
	s.Value = v
	return s
}

func (s *ListCloudAccountRolesRequestFilter) Validate() error {
	return dara.Validate(s)
}
