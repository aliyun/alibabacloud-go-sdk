// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEvidenceRequestsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeAiEngineEvidenceRequestsResponseBodyItems) *DescribeAiEngineEvidenceRequestsResponseBody
	GetItems() []*DescribeAiEngineEvidenceRequestsResponseBodyItems
	SetMaxResults(v int32) *DescribeAiEngineEvidenceRequestsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAiEngineEvidenceRequestsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeAiEngineEvidenceRequestsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAiEngineEvidenceRequestsResponseBody
	GetTotalCount() *int32
}

type DescribeAiEngineEvidenceRequestsResponseBody struct {
	// The list of evidence requests.
	Items []*DescribeAiEngineEvidenceRequestsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries to return on each page in a paging query. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. If a next page exists, this field returns a value.
	//
	// > If this parameter returns a value, a next page exists. You can use the returned **NextToken*	- as a request parameter to retrieve the next page. Repeat this process until no value is returned, which indicates that all data has been retrieved.
	//
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 01A1CEE2-*****9-019A-8B44-2*****C2C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of evidence requests that meet the specified conditions.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAiEngineEvidenceRequestsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEvidenceRequestsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEvidenceRequestsResponseBody) GetItems() []*DescribeAiEngineEvidenceRequestsResponseBodyItems {
	return s.Items
}

func (s *DescribeAiEngineEvidenceRequestsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAiEngineEvidenceRequestsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAiEngineEvidenceRequestsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAiEngineEvidenceRequestsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAiEngineEvidenceRequestsResponseBody) SetItems(v []*DescribeAiEngineEvidenceRequestsResponseBodyItems) *DescribeAiEngineEvidenceRequestsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBody) SetMaxResults(v int32) *DescribeAiEngineEvidenceRequestsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBody) SetNextToken(v string) *DescribeAiEngineEvidenceRequestsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBody) SetRequestId(v string) *DescribeAiEngineEvidenceRequestsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBody) SetTotalCount(v int32) *DescribeAiEngineEvidenceRequestsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBody) Validate() error {
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

type DescribeAiEngineEvidenceRequestsResponseBodyItems struct {
	// The list of attack source groups for the request.
	AttackSources []*DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSources `json:"AttackSources,omitempty" xml:"AttackSources,omitempty" type:"Repeated"`
	// The evidence ID.
	//
	// example:
	//
	// EVD-20241025-001
	EvidenceId *string `json:"EvidenceId,omitempty" xml:"EvidenceId,omitempty"`
	// The host of the request.
	//
	// example:
	//
	// api.example.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The HTTP response status code.
	//
	// example:
	//
	// 500
	HttpStatus *int32 `json:"HttpStatus,omitempty" xml:"HttpStatus,omitempty"`
	// The evidence impact tag. Valid values:
	//
	// - **block**: Indicates that the request is blocked.
	//
	// - **challenge**: Indicates that the request is challenged.
	//
	// - **monitor**: Indicates that the request is monitored.
	//
	// - **bypass**: Indicates a suspected bypass.
	//
	// - **attempt**: Indicates reconnaissance or probing.
	//
	// example:
	//
	// bypass
	ImpactType *string `json:"ImpactType,omitempty" xml:"ImpactType,omitempty"`
	// The HTTP method.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The ID of the associated node. This corresponds to the NodeId in the node list.
	//
	// example:
	//
	// node-109-04
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The masked path of the request.
	//
	// example:
	//
	// /api/orders/query
	RequestPath *string `json:"RequestPath,omitempty" xml:"RequestPath,omitempty"`
	// The time when the original request was made. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1792973400000
	RequestTime *int64 `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	// The protected object to which the request belongs.
	//
	// example:
	//
	// www.example.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// Specifies whether the request is a suspected bypass. Valid values:
	//
	// - **1**: Indicates a suspected bypass.
	//
	// - **0**: Indicates that the request is not a bypass.
	//
	// example:
	//
	// 1
	SuspectedBypass *int32 `json:"SuspectedBypass,omitempty" xml:"SuspectedBypass,omitempty"`
	// The trace ID of the original request.
	//
	// example:
	//
	// trace-20241025-109-04
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeAiEngineEvidenceRequestsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEvidenceRequestsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) GetAttackSources() []*DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSources {
	return s.AttackSources
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) GetEvidenceId() *string {
	return s.EvidenceId
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) GetHost() *string {
	return s.Host
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) GetHttpStatus() *int32 {
	return s.HttpStatus
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) GetImpactType() *string {
	return s.ImpactType
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) GetMethod() *string {
	return s.Method
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) GetRequestPath() *string {
	return s.RequestPath
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) GetRequestTime() *int64 {
	return s.RequestTime
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) GetResource() *string {
	return s.Resource
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) GetSuspectedBypass() *int32 {
	return s.SuspectedBypass
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) SetAttackSources(v []*DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSources) *DescribeAiEngineEvidenceRequestsResponseBodyItems {
	s.AttackSources = v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) SetEvidenceId(v string) *DescribeAiEngineEvidenceRequestsResponseBodyItems {
	s.EvidenceId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) SetHost(v string) *DescribeAiEngineEvidenceRequestsResponseBodyItems {
	s.Host = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) SetHttpStatus(v int32) *DescribeAiEngineEvidenceRequestsResponseBodyItems {
	s.HttpStatus = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) SetImpactType(v string) *DescribeAiEngineEvidenceRequestsResponseBodyItems {
	s.ImpactType = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) SetMethod(v string) *DescribeAiEngineEvidenceRequestsResponseBodyItems {
	s.Method = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) SetNodeId(v string) *DescribeAiEngineEvidenceRequestsResponseBodyItems {
	s.NodeId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) SetRequestPath(v string) *DescribeAiEngineEvidenceRequestsResponseBodyItems {
	s.RequestPath = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) SetRequestTime(v int64) *DescribeAiEngineEvidenceRequestsResponseBodyItems {
	s.RequestTime = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) SetResource(v string) *DescribeAiEngineEvidenceRequestsResponseBodyItems {
	s.Resource = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) SetSuspectedBypass(v int32) *DescribeAiEngineEvidenceRequestsResponseBodyItems {
	s.SuspectedBypass = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) SetTraceId(v string) *DescribeAiEngineEvidenceRequestsResponseBodyItems {
	s.TraceId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItems) Validate() error {
	if s.AttackSources != nil {
		for _, item := range s.AttackSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSources struct {
	// The list of sources of this type.
	Sources []*DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSourcesSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The attack source type.
	//
	// example:
	//
	// ip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSources) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSources) GetSources() []*DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSourcesSources {
	return s.Sources
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSources) GetType() *string {
	return s.Type
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSources) SetSources(v []*DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSourcesSources) *DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSources {
	s.Sources = v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSources) SetType(v string) *DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSources {
	s.Type = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSources) Validate() error {
	if s.Sources != nil {
		for _, item := range s.Sources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSourcesSources struct {
	// The display attribute of the source, such as the region code.
	//
	// example:
	//
	// US-CA
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The source value, such as an IP address.
	//
	// example:
	//
	// 2XX.0.113.1X
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSourcesSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSourcesSources) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSourcesSources) GetAttribute() *string {
	return s.Attribute
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSourcesSources) GetValue() *string {
	return s.Value
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSourcesSources) SetAttribute(v string) *DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSourcesSources {
	s.Attribute = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSourcesSources) SetValue(v string) *DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSourcesSources {
	s.Value = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponseBodyItemsAttackSourcesSources) Validate() error {
	return dara.Validate(s)
}
