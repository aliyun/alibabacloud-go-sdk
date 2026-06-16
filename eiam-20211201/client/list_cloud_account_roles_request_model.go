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
	// The ID of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The filter conditions.
	Filter []*ListCloudAccountRolesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of records to return on each page.
	//
	// - The default value is 20.
	//
	// - The maximum value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that specifies the start of the next page of results.
	//
	// - If this parameter is not specified, the query starts from the first page.
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
	// The name of the filter field. Valid values:
	//
	// - CloudAccountRoleId: The ID of the cloud role.
	//
	// - CloudAccountRoleName: The name of the cloud role.
	//
	// - CloudAccountRoleExternalId: The external ID of the cloud role.
	//
	// - CloudAccountRoleUsageType: The usage type of the cloud role.
	//
	// example:
	//
	// CloudAccountRoleId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of values for the filter field.
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
