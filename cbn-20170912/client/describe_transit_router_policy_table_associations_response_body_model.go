// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouterPolicyTableAssociationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssociations(v []*DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations) *DescribeTransitRouterPolicyTableAssociationsResponseBody
	GetAssociations() []*DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations
	SetMaxResults(v int32) *DescribeTransitRouterPolicyTableAssociationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeTransitRouterPolicyTableAssociationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeTransitRouterPolicyTableAssociationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeTransitRouterPolicyTableAssociationsResponseBody
	GetTotalCount() *int32
}

type DescribeTransitRouterPolicyTableAssociationsResponseBody struct {
	// The associate relationships between the policy routing table and network connectivity instances.
	Associations []*DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations `json:"Associations,omitempty" xml:"Associations,omitempty" type:"Repeated"`
	// The number of entries returned per page in a paginated query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query in a paginated query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 13526224-5780-4426-8BDF-BC8B08700F22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTransitRouterPolicyTableAssociationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyTableAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBody) GetAssociations() []*DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations {
	return s.Associations
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBody) SetAssociations(v []*DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations) *DescribeTransitRouterPolicyTableAssociationsResponseBody {
	s.Associations = v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBody) SetMaxResults(v int32) *DescribeTransitRouterPolicyTableAssociationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBody) SetNextToken(v string) *DescribeTransitRouterPolicyTableAssociationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBody) SetRequestId(v string) *DescribeTransitRouterPolicyTableAssociationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBody) SetTotalCount(v int32) *DescribeTransitRouterPolicyTableAssociationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBody) Validate() error {
	if s.Associations != nil {
		for _, item := range s.Associations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations struct {
	// The ID of the association between the policy route table and the network instance connection.
	//
	// example:
	//
	// pta-vfan3mbxoespx7xvyt
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The ID of the network instance connection.
	//
	// example:
	//
	// tr-attach-123qsrads
	AttachmentId *string `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	// The instance ID of the policy routing table.
	//
	// example:
	//
	// ptbl-123adfzcx
	PolicyTableId *string `json:"PolicyTableId,omitempty" xml:"PolicyTableId,omitempty"`
	// The status of the association. Valid values:
	//
	// - **Active**: Active.
	//
	// - **Associating**: Associating.
	//
	// - **Dissociating**: Dissociating.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations) GetAssociationId() *string {
	return s.AssociationId
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations) GetPolicyTableId() *string {
	return s.PolicyTableId
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations) GetStatus() *string {
	return s.Status
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations) SetAssociationId(v string) *DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations {
	s.AssociationId = &v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations) SetAttachmentId(v string) *DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations {
	s.AttachmentId = &v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations) SetPolicyTableId(v string) *DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations {
	s.PolicyTableId = &v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations) SetStatus(v string) *DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations {
	s.Status = &v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponseBodyAssociations) Validate() error {
	return dara.Validate(s)
}
