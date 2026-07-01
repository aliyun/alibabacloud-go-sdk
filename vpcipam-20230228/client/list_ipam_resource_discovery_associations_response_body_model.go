// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamResourceDiscoveryAssociationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListIpamResourceDiscoveryAssociationsResponseBody
	GetCount() *int32
	SetIpamResourceDiscoveryAssociations(v []*ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) *ListIpamResourceDiscoveryAssociationsResponseBody
	GetIpamResourceDiscoveryAssociations() []*ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations
	SetMaxResults(v int32) *ListIpamResourceDiscoveryAssociationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListIpamResourceDiscoveryAssociationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIpamResourceDiscoveryAssociationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListIpamResourceDiscoveryAssociationsResponseBody
	GetTotalCount() *int64
}

type ListIpamResourceDiscoveryAssociationsResponseBody struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// A list of association details.
	IpamResourceDiscoveryAssociations []*ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations `json:"IpamResourceDiscoveryAssociations,omitempty" xml:"IpamResourceDiscoveryAssociations,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to use to retrieve the next page of results. This value is empty when there are no more results to return.
	//
	// - If **NextToken*	- is empty, no subsequent query is needed.
	//
	// - If **NextToken*	- is not empty, its value is the token to start the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F28A239E-F88D-500E-ADE7-FA5E8CA3A170
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that match the query criteria.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamResourceDiscoveryAssociationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIpamResourceDiscoveryAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) GetIpamResourceDiscoveryAssociations() []*ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations {
	return s.IpamResourceDiscoveryAssociations
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) SetCount(v int32) *ListIpamResourceDiscoveryAssociationsResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) SetIpamResourceDiscoveryAssociations(v []*ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) *ListIpamResourceDiscoveryAssociationsResponseBody {
	s.IpamResourceDiscoveryAssociations = v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) SetMaxResults(v int32) *ListIpamResourceDiscoveryAssociationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) SetNextToken(v string) *ListIpamResourceDiscoveryAssociationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) SetRequestId(v string) *ListIpamResourceDiscoveryAssociationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) SetTotalCount(v int64) *ListIpamResourceDiscoveryAssociationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBody) Validate() error {
	if s.IpamResourceDiscoveryAssociations != nil {
		for _, item := range s.IpamResourceDiscoveryAssociations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations struct {
	// The ID of the IPAM instance.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The ID of the resource discovery instance.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the resource discovery instance.
	//
	// example:
	//
	// 1210123456******
	IpamResourceDiscoveryOwnerId *string `json:"IpamResourceDiscoveryOwnerId,omitempty" xml:"IpamResourceDiscoveryOwnerId,omitempty"`
	// The status of the resource discovery instance. Valid values:
	//
	// - **Creating**
	//
	// - **Created**
	//
	// - **Modifying**
	//
	// - **Deleting**
	//
	// - **Deleted**
	//
	// example:
	//
	// Created
	IpamResourceDiscoveryStatus *string `json:"IpamResourceDiscoveryStatus,omitempty" xml:"IpamResourceDiscoveryStatus,omitempty"`
	// The type of the resource discovery. Valid values:
	//
	// - **system**: a default resource discovery that is automatically created by the system.
	//
	// - **custom**: a custom resource discovery that is created by a user.
	//
	// example:
	//
	// custom
	IpamResourceDiscoveryType *string `json:"IpamResourceDiscoveryType,omitempty" xml:"IpamResourceDiscoveryType,omitempty"`
	// The association status. Valid values:
	//
	// - **Created**
	//
	// - **Deleted**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) String() string {
	return dara.Prettify(s)
}

func (s ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) GetIpamId() *string {
	return s.IpamId
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) GetIpamResourceDiscoveryId() *string {
	return s.IpamResourceDiscoveryId
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) GetIpamResourceDiscoveryOwnerId() *string {
	return s.IpamResourceDiscoveryOwnerId
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) GetIpamResourceDiscoveryStatus() *string {
	return s.IpamResourceDiscoveryStatus
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) GetIpamResourceDiscoveryType() *string {
	return s.IpamResourceDiscoveryType
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) GetStatus() *string {
	return s.Status
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) SetIpamId(v string) *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations {
	s.IpamId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) SetIpamResourceDiscoveryId(v string) *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) SetIpamResourceDiscoveryOwnerId(v string) *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations {
	s.IpamResourceDiscoveryOwnerId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) SetIpamResourceDiscoveryStatus(v string) *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations {
	s.IpamResourceDiscoveryStatus = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) SetIpamResourceDiscoveryType(v string) *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations {
	s.IpamResourceDiscoveryType = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) SetStatus(v string) *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations {
	s.Status = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponseBodyIpamResourceDiscoveryAssociations) Validate() error {
	return dara.Validate(s)
}
