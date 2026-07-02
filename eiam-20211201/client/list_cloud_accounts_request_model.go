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
	// The filter conditions.
	Filter []*ListCloudAccountsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of records per page.
	//
	// - Default value: 20.
	//
	// - Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the starting position of the next page.
	//
	// - If you do not specify this parameter, the query starts from the first page.
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
	// The name of the filter field. Valid values:
	//
	// - CloudAccountId: the cloud account ID.
	//
	// - CloudAccountExternalId: the external unique identifier of the cloud account.
	//
	// - CloudAccountVendorType: the cloud account type.
	//
	// example:
	//
	// CloudAccountId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The values of the filter field.
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
