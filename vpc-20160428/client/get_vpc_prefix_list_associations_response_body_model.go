// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcPrefixListAssociationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *GetVpcPrefixListAssociationsResponseBody
	GetCount() *int64
	SetNextToken(v string) *GetVpcPrefixListAssociationsResponseBody
	GetNextToken() *string
	SetPrefixListAssociation(v []*GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) *GetVpcPrefixListAssociationsResponseBody
	GetPrefixListAssociation() []*GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation
	SetRequestId(v string) *GetVpcPrefixListAssociationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetVpcPrefixListAssociationsResponseBody
	GetTotalCount() *int64
}

type GetVpcPrefixListAssociationsResponseBody struct {
	// The number of entries.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is used to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the network instances that are associated with the prefix list.
	PrefixListAssociation []*GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation `json:"PrefixListAssociation,omitempty" xml:"PrefixListAssociation,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetVpcPrefixListAssociationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVpcPrefixListAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpcPrefixListAssociationsResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *GetVpcPrefixListAssociationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetVpcPrefixListAssociationsResponseBody) GetPrefixListAssociation() []*GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation {
	return s.PrefixListAssociation
}

func (s *GetVpcPrefixListAssociationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVpcPrefixListAssociationsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetVpcPrefixListAssociationsResponseBody) SetCount(v int64) *GetVpcPrefixListAssociationsResponseBody {
	s.Count = &v
	return s
}

func (s *GetVpcPrefixListAssociationsResponseBody) SetNextToken(v string) *GetVpcPrefixListAssociationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetVpcPrefixListAssociationsResponseBody) SetPrefixListAssociation(v []*GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) *GetVpcPrefixListAssociationsResponseBody {
	s.PrefixListAssociation = v
	return s
}

func (s *GetVpcPrefixListAssociationsResponseBody) SetRequestId(v string) *GetVpcPrefixListAssociationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVpcPrefixListAssociationsResponseBody) SetTotalCount(v int64) *GetVpcPrefixListAssociationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetVpcPrefixListAssociationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation struct {
	// List of CIDR addresses where the prefix list is effective in the currently associated resources.
	//
	// example:
	//
	// 192.168.0.0/16
	CidrList *string `json:"CidrList,omitempty" xml:"CidrList,omitempty"`
	// The ID of the Alibaba Cloud account to which the prefix list belongs.
	//
	// example:
	//
	// 153460731706****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The prefix list ID.
	//
	// example:
	//
	// pl-0b7hwu67****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The reason why the association failed.
	//
	// example:
	//
	// failed
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The region ID of the prefix list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the associated resource.
	//
	// example:
	//
	// vtb-bp1drpcfz9srr393h****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the associated resource. Valid values:
	//
	// 	- **vpcRouteTable**: virtual private cloud (VPC) route table.
	//
	// 	- **trRouteTable**: route table of a transit router.
	//
	// example:
	//
	// vpcRouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource associated with the prefix list belongs.
	//
	// example:
	//
	// 132193271328****
	ResourceUid *string `json:"ResourceUid,omitempty" xml:"ResourceUid,omitempty"`
	// The status of the prefix list. Valid values:
	//
	// 	- **Created**
	//
	// 	- **ModifyFailed**
	//
	// 	- **Creating**
	//
	// 	- **Modifying**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) String() string {
	return dara.Prettify(s)
}

func (s GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) GoString() string {
	return s.String()
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) GetCidrList() *string {
	return s.CidrList
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) GetReason() *string {
	return s.Reason
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) GetResourceUid() *string {
	return s.ResourceUid
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) GetStatus() *string {
	return s.Status
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) SetCidrList(v string) *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation {
	s.CidrList = &v
	return s
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) SetOwnerId(v string) *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation {
	s.OwnerId = &v
	return s
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) SetPrefixListId(v string) *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation {
	s.PrefixListId = &v
	return s
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) SetReason(v string) *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation {
	s.Reason = &v
	return s
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) SetRegionId(v string) *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation {
	s.RegionId = &v
	return s
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) SetResourceId(v string) *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation {
	s.ResourceId = &v
	return s
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) SetResourceType(v string) *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation {
	s.ResourceType = &v
	return s
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) SetResourceUid(v string) *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation {
	s.ResourceUid = &v
	return s
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) SetStatus(v string) *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation {
	s.Status = &v
	return s
}

func (s *GetVpcPrefixListAssociationsResponseBodyPrefixListAssociation) Validate() error {
	return dara.Validate(s)
}
