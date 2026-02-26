// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialIds(v []*string) *ListCredentialsRequest
	GetCredentialIds() []*string
	SetFilter(v []*ListCredentialsRequestFilter) *ListCredentialsRequest
	GetFilter() []*ListCredentialsRequestFilter
	SetInstanceId(v string) *ListCredentialsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListCredentialsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCredentialsRequest
	GetNextToken() *string
	SetStatuses(v []*string) *ListCredentialsRequest
	GetStatuses() []*string
}

type ListCredentialsRequest struct {
	CredentialIds []*string                       `json:"CredentialIds,omitempty" xml:"CredentialIds,omitempty" type:"Repeated"`
	Filter        []*ListCredentialsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
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
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Statuses  []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s ListCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsRequest) GoString() string {
	return s.String()
}

func (s *ListCredentialsRequest) GetCredentialIds() []*string {
	return s.CredentialIds
}

func (s *ListCredentialsRequest) GetFilter() []*ListCredentialsRequestFilter {
	return s.Filter
}

func (s *ListCredentialsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCredentialsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCredentialsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCredentialsRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListCredentialsRequest) SetCredentialIds(v []*string) *ListCredentialsRequest {
	s.CredentialIds = v
	return s
}

func (s *ListCredentialsRequest) SetFilter(v []*ListCredentialsRequestFilter) *ListCredentialsRequest {
	s.Filter = v
	return s
}

func (s *ListCredentialsRequest) SetInstanceId(v string) *ListCredentialsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCredentialsRequest) SetMaxResults(v int32) *ListCredentialsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCredentialsRequest) SetNextToken(v string) *ListCredentialsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCredentialsRequest) SetStatuses(v []*string) *ListCredentialsRequest {
	s.Statuses = v
	return s
}

func (s *ListCredentialsRequest) Validate() error {
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

type ListCredentialsRequestFilter struct {
	// example:
	//
	// CredentialIdentifier
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListCredentialsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListCredentialsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListCredentialsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListCredentialsRequestFilter) SetName(v string) *ListCredentialsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListCredentialsRequestFilter) SetValue(v []*string) *ListCredentialsRequestFilter {
	s.Value = v
	return s
}

func (s *ListCredentialsRequestFilter) Validate() error {
	return dara.Validate(s)
}
