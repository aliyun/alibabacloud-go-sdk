// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEventNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeAiEngineEventNodesResponseBodyItems) *DescribeAiEngineEventNodesResponseBody
	GetItems() []*DescribeAiEngineEventNodesResponseBodyItems
	SetMaxResults(v int32) *DescribeAiEngineEventNodesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAiEngineEventNodesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeAiEngineEventNodesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeAiEngineEventNodesResponseBody
	GetTotalCount() *int64
}

type DescribeAiEngineEventNodesResponseBody struct {
	// The list of nodes.
	Items []*DescribeAiEngineEventNodesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries to return on each page in a paging query. Valid values: 1 to 20. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used to query the next page. If a next page exists, this field returns a value.
	//
	// > If this parameter returns a value, a next page exists. You can use the returned **NextToken*	- as a request parameter to obtain the data on the next page until no value is returned, which indicates that all data is obtained.
	//
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// request-demo
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of nodes for the event.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAiEngineEventNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventNodesResponseBody) GetItems() []*DescribeAiEngineEventNodesResponseBodyItems {
	return s.Items
}

func (s *DescribeAiEngineEventNodesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAiEngineEventNodesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAiEngineEventNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAiEngineEventNodesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAiEngineEventNodesResponseBody) SetItems(v []*DescribeAiEngineEventNodesResponseBodyItems) *DescribeAiEngineEventNodesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAiEngineEventNodesResponseBody) SetMaxResults(v int32) *DescribeAiEngineEventNodesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeAiEngineEventNodesResponseBody) SetNextToken(v string) *DescribeAiEngineEventNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAiEngineEventNodesResponseBody) SetRequestId(v string) *DescribeAiEngineEventNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAiEngineEventNodesResponseBody) SetTotalCount(v int64) *DescribeAiEngineEventNodesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAiEngineEventNodesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAiEngineEventNodesResponseBodyItems struct {
	// The details of the node in the specified language. The details are in Markdown format.
	//
	// example:
	//
	// - **Characteristic**: single quote closure
	//
	// - **Risk**: high
	NodeDetail *string `json:"NodeDetail,omitempty" xml:"NodeDetail,omitempty"`
	// The node ID.
	//
	// example:
	//
	// node-109-04
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node in the specified language.
	//
	// example:
	//
	// SQL Injection
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The summary of the node in the specified language.
	//
	// example:
	//
	// The request matches SQL injection characteristics.
	NodeSummary *string `json:"NodeSummary,omitempty" xml:"NodeSummary,omitempty"`
	// The node tag. Valid values:
	//
	// - **bypass**: Suspected bypass.
	//
	// - **protect**: Mitigated.
	//
	// - **attempt**: Reconnaissance and probing.
	//
	// example:
	//
	// bypass
	NodeTag *string `json:"NodeTag,omitempty" xml:"NodeTag,omitempty"`
	// The time when the node occurred. The time is in the UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1792973400000
	OccurredAt *int64 `json:"OccurredAt,omitempty" xml:"OccurredAt,omitempty"`
}

func (s DescribeAiEngineEventNodesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventNodesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventNodesResponseBodyItems) GetNodeDetail() *string {
	return s.NodeDetail
}

func (s *DescribeAiEngineEventNodesResponseBodyItems) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeAiEngineEventNodesResponseBodyItems) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeAiEngineEventNodesResponseBodyItems) GetNodeSummary() *string {
	return s.NodeSummary
}

func (s *DescribeAiEngineEventNodesResponseBodyItems) GetNodeTag() *string {
	return s.NodeTag
}

func (s *DescribeAiEngineEventNodesResponseBodyItems) GetOccurredAt() *int64 {
	return s.OccurredAt
}

func (s *DescribeAiEngineEventNodesResponseBodyItems) SetNodeDetail(v string) *DescribeAiEngineEventNodesResponseBodyItems {
	s.NodeDetail = &v
	return s
}

func (s *DescribeAiEngineEventNodesResponseBodyItems) SetNodeId(v string) *DescribeAiEngineEventNodesResponseBodyItems {
	s.NodeId = &v
	return s
}

func (s *DescribeAiEngineEventNodesResponseBodyItems) SetNodeName(v string) *DescribeAiEngineEventNodesResponseBodyItems {
	s.NodeName = &v
	return s
}

func (s *DescribeAiEngineEventNodesResponseBodyItems) SetNodeSummary(v string) *DescribeAiEngineEventNodesResponseBodyItems {
	s.NodeSummary = &v
	return s
}

func (s *DescribeAiEngineEventNodesResponseBodyItems) SetNodeTag(v string) *DescribeAiEngineEventNodesResponseBodyItems {
	s.NodeTag = &v
	return s
}

func (s *DescribeAiEngineEventNodesResponseBodyItems) SetOccurredAt(v int64) *DescribeAiEngineEventNodesResponseBodyItems {
	s.OccurredAt = &v
	return s
}

func (s *DescribeAiEngineEventNodesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
