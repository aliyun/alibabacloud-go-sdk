// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcPublishedRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListVpcPublishedRouteEntriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVpcPublishedRouteEntriesResponseBody
	GetRequestId() *string
	SetRouteEntries(v []*ListVpcPublishedRouteEntriesResponseBodyRouteEntries) *ListVpcPublishedRouteEntriesResponseBody
	GetRouteEntries() []*ListVpcPublishedRouteEntriesResponseBodyRouteEntries
}

type ListVpcPublishedRouteEntriesResponseBody struct {
	// Indicates whether there is a token for the next query. Values:
	//
	// - If **NextToken*	- is empty, it means there is no next query.
	//
	// - If **NextToken*	- has a return value, this value indicates the token for the start of the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 1D0971B2-A35A-42C1-A44C-E91360C36C0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of route entry publishing status information.
	RouteEntries []*ListVpcPublishedRouteEntriesResponseBodyRouteEntries `json:"RouteEntries,omitempty" xml:"RouteEntries,omitempty" type:"Repeated"`
}

func (s ListVpcPublishedRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcPublishedRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcPublishedRouteEntriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcPublishedRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcPublishedRouteEntriesResponseBody) GetRouteEntries() []*ListVpcPublishedRouteEntriesResponseBodyRouteEntries {
	return s.RouteEntries
}

func (s *ListVpcPublishedRouteEntriesResponseBody) SetNextToken(v string) *ListVpcPublishedRouteEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesResponseBody) SetRequestId(v string) *ListVpcPublishedRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesResponseBody) SetRouteEntries(v []*ListVpcPublishedRouteEntriesResponseBodyRouteEntries) *ListVpcPublishedRouteEntriesResponseBody {
	s.RouteEntries = v
	return s
}

func (s *ListVpcPublishedRouteEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVpcPublishedRouteEntriesResponseBodyRouteEntries struct {
	// The destination CIDR block of the route entry.
	//
	// example:
	//
	// 10.0.0.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ID of the route entry.
	//
	// example:
	//
	// rte-bp1mnnr2al0naomnpv****
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
	// List of route entry publishing status information in the publishing targets.
	RoutePublishTargets []*ListVpcPublishedRouteEntriesResponseBodyRouteEntriesRoutePublishTargets `json:"RoutePublishTargets,omitempty" xml:"RoutePublishTargets,omitempty" type:"Repeated"`
	// The ID of the route table.
	//
	// example:
	//
	// vtb-2ze3jgygk9bmsj23s****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s ListVpcPublishedRouteEntriesResponseBodyRouteEntries) String() string {
	return dara.Prettify(s)
}

func (s ListVpcPublishedRouteEntriesResponseBodyRouteEntries) GoString() string {
	return s.String()
}

func (s *ListVpcPublishedRouteEntriesResponseBodyRouteEntries) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ListVpcPublishedRouteEntriesResponseBodyRouteEntries) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *ListVpcPublishedRouteEntriesResponseBodyRouteEntries) GetRoutePublishTargets() []*ListVpcPublishedRouteEntriesResponseBodyRouteEntriesRoutePublishTargets {
	return s.RoutePublishTargets
}

func (s *ListVpcPublishedRouteEntriesResponseBodyRouteEntries) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *ListVpcPublishedRouteEntriesResponseBodyRouteEntries) SetDestinationCidrBlock(v string) *ListVpcPublishedRouteEntriesResponseBodyRouteEntries {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesResponseBodyRouteEntries) SetRouteEntryId(v string) *ListVpcPublishedRouteEntriesResponseBodyRouteEntries {
	s.RouteEntryId = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesResponseBodyRouteEntries) SetRoutePublishTargets(v []*ListVpcPublishedRouteEntriesResponseBodyRouteEntriesRoutePublishTargets) *ListVpcPublishedRouteEntriesResponseBodyRouteEntries {
	s.RoutePublishTargets = v
	return s
}

func (s *ListVpcPublishedRouteEntriesResponseBodyRouteEntries) SetRouteTableId(v string) *ListVpcPublishedRouteEntriesResponseBodyRouteEntries {
	s.RouteTableId = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesResponseBodyRouteEntries) Validate() error {
	return dara.Validate(s)
}

type ListVpcPublishedRouteEntriesResponseBodyRouteEntriesRoutePublishTargets struct {
	// The publishing status of the route entry in the publishing target.
	//
	// example:
	//
	// Published
	PublishStatus *string `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	// The ID of the route publishing target instance.
	//
	// example:
	//
	// ecr-xvuqdfma6x57ei****
	PublishTargetInstanceId *string `json:"PublishTargetInstanceId,omitempty" xml:"PublishTargetInstanceId,omitempty"`
	// The type of the route publishing target.
	//
	// example:
	//
	// ECR
	PublishTargetType *string `json:"PublishTargetType,omitempty" xml:"PublishTargetType,omitempty"`
}

func (s ListVpcPublishedRouteEntriesResponseBodyRouteEntriesRoutePublishTargets) String() string {
	return dara.Prettify(s)
}

func (s ListVpcPublishedRouteEntriesResponseBodyRouteEntriesRoutePublishTargets) GoString() string {
	return s.String()
}

func (s *ListVpcPublishedRouteEntriesResponseBodyRouteEntriesRoutePublishTargets) GetPublishStatus() *string {
	return s.PublishStatus
}

func (s *ListVpcPublishedRouteEntriesResponseBodyRouteEntriesRoutePublishTargets) GetPublishTargetInstanceId() *string {
	return s.PublishTargetInstanceId
}

func (s *ListVpcPublishedRouteEntriesResponseBodyRouteEntriesRoutePublishTargets) GetPublishTargetType() *string {
	return s.PublishTargetType
}

func (s *ListVpcPublishedRouteEntriesResponseBodyRouteEntriesRoutePublishTargets) SetPublishStatus(v string) *ListVpcPublishedRouteEntriesResponseBodyRouteEntriesRoutePublishTargets {
	s.PublishStatus = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesResponseBodyRouteEntriesRoutePublishTargets) SetPublishTargetInstanceId(v string) *ListVpcPublishedRouteEntriesResponseBodyRouteEntriesRoutePublishTargets {
	s.PublishTargetInstanceId = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesResponseBodyRouteEntriesRoutePublishTargets) SetPublishTargetType(v string) *ListVpcPublishedRouteEntriesResponseBodyRouteEntriesRoutePublishTargets {
	s.PublishTargetType = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesResponseBodyRouteEntriesRoutePublishTargets) Validate() error {
	return dara.Validate(s)
}
