// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouterPolicyTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeTransitRouterPolicyTablesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeTransitRouterPolicyTablesResponseBody
	GetNextToken() *string
	SetPolicyTables(v []*DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) *DescribeTransitRouterPolicyTablesResponseBody
	GetPolicyTables() []*DescribeTransitRouterPolicyTablesResponseBodyPolicyTables
	SetRequestId(v string) *DescribeTransitRouterPolicyTablesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeTransitRouterPolicyTablesResponseBody
	GetTotalCount() *int32
}

type DescribeTransitRouterPolicyTablesResponseBody struct {
	// The number of entries returned on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If NextToken is empty, no subsequent query is required.
	//
	// - If NextToken is returned, the value indicates the token for the next query.
	//
	// example:
	//
	// dd20****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The policy route tables.
	PolicyTables []*DescribeTransitRouterPolicyTablesResponseBodyPolicyTables `json:"PolicyTables,omitempty" xml:"PolicyTables,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2BFA6822-240E-4E27-B4C8-AA400EF7474D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTransitRouterPolicyTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyTablesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeTransitRouterPolicyTablesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTransitRouterPolicyTablesResponseBody) GetPolicyTables() []*DescribeTransitRouterPolicyTablesResponseBodyPolicyTables {
	return s.PolicyTables
}

func (s *DescribeTransitRouterPolicyTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTransitRouterPolicyTablesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTransitRouterPolicyTablesResponseBody) SetMaxResults(v int32) *DescribeTransitRouterPolicyTablesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponseBody) SetNextToken(v string) *DescribeTransitRouterPolicyTablesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponseBody) SetPolicyTables(v []*DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) *DescribeTransitRouterPolicyTablesResponseBody {
	s.PolicyTables = v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponseBody) SetRequestId(v string) *DescribeTransitRouterPolicyTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponseBody) SetTotalCount(v int32) *DescribeTransitRouterPolicyTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponseBody) Validate() error {
	if s.PolicyTables != nil {
		for _, item := range s.PolicyTables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTransitRouterPolicyTablesResponseBodyPolicyTables struct {
	// The network instance attachments associated with the policy route table.
	Attachments []*string `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	// The time when the policy route table was created.
	//
	// The time is displayed in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format. The time is displayed in UTC.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2019-10-22T07:44Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the policy route table.
	//
	// example:
	//
	// descriptiontest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the policy route table.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The policy routing table instance ID.
	//
	// example:
	//
	// ptbl-44m0p68spvlrqq****
	PolicyTableId *string `json:"PolicyTableId,omitempty" xml:"PolicyTableId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the policy route table. Valid values:
	//
	// - Creating: The policy route table is being created.
	//
	// - Active: The policy route table is running.
	//
	// - Deleting: The policy route table is being deleted.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags added to the policy route table.
	Tags []*DescribeTransitRouterPolicyTablesResponseBodyPolicyTablesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the Enterprise Edition transit router.
	//
	// example:
	//
	// tr-uf654ttymmljlvh2x****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) GetAttachments() []*string {
	return s.Attachments
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) GetDescription() *string {
	return s.Description
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) GetName() *string {
	return s.Name
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) GetPolicyTableId() *string {
	return s.PolicyTableId
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) GetStatus() *string {
	return s.Status
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) GetTags() []*DescribeTransitRouterPolicyTablesResponseBodyPolicyTablesTags {
	return s.Tags
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) SetAttachments(v []*string) *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables {
	s.Attachments = v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) SetCreationTime(v string) *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables {
	s.CreationTime = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) SetDescription(v string) *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables {
	s.Description = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) SetName(v string) *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables {
	s.Name = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) SetPolicyTableId(v string) *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables {
	s.PolicyTableId = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) SetRegionId(v string) *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables {
	s.RegionId = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) SetStatus(v string) *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables {
	s.Status = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) SetTags(v []*DescribeTransitRouterPolicyTablesResponseBodyPolicyTablesTags) *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables {
	s.Tags = v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) SetTransitRouterId(v string) *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTables) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTransitRouterPolicyTablesResponseBodyPolicyTablesTags struct {
	// The tag key.
	//
	// example:
	//
	// tagtest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// tagtest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTransitRouterPolicyTablesResponseBodyPolicyTablesTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyTablesResponseBodyPolicyTablesTags) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTablesTags) GetKey() *string {
	return s.Key
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTablesTags) GetValue() *string {
	return s.Value
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTablesTags) SetKey(v string) *DescribeTransitRouterPolicyTablesResponseBodyPolicyTablesTags {
	s.Key = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTablesTags) SetValue(v string) *DescribeTransitRouterPolicyTablesResponseBodyPolicyTablesTags {
	s.Value = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponseBodyPolicyTablesTags) Validate() error {
	return dara.Validate(s)
}
