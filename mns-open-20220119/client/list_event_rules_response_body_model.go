// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ListEventRulesResponseBody
	GetCode() *int64
	SetData(v *ListEventRulesResponseBodyData) *ListEventRulesResponseBody
	GetData() *ListEventRulesResponseBodyData
	SetMessage(v string) *ListEventRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEventRulesResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListEventRulesResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ListEventRulesResponseBody
	GetSuccess() *bool
}

type ListEventRulesResponseBody struct {
	// example:
	//
	// 200
	Code *int64                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListEventRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 06273500-249F-5863-121D-74D51123E62C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEventRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventRulesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListEventRulesResponseBody) GetData() *ListEventRulesResponseBodyData {
	return s.Data
}

func (s *ListEventRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEventRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEventRulesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListEventRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEventRulesResponseBody) SetCode(v int64) *ListEventRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListEventRulesResponseBody) SetData(v *ListEventRulesResponseBodyData) *ListEventRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListEventRulesResponseBody) SetMessage(v string) *ListEventRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEventRulesResponseBody) SetRequestId(v string) *ListEventRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventRulesResponseBody) SetStatus(v string) *ListEventRulesResponseBody {
	s.Status = &v
	return s
}

func (s *ListEventRulesResponseBody) SetSuccess(v bool) *ListEventRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListEventRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEventRulesResponseBodyData struct {
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// a88f58d504b8b4c4e0b5e8707e68181f
	NextToken *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageData  []*ListEventRulesResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 50
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3
	Pages *int64 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 130
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEventRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEventRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEventRulesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEventRulesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEventRulesResponseBodyData) GetPageData() []*ListEventRulesResponseBodyDataPageData {
	return s.PageData
}

func (s *ListEventRulesResponseBodyData) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListEventRulesResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListEventRulesResponseBodyData) GetPages() *int64 {
	return s.Pages
}

func (s *ListEventRulesResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *ListEventRulesResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListEventRulesResponseBodyData) SetMaxResults(v int32) *ListEventRulesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListEventRulesResponseBodyData) SetNextToken(v string) *ListEventRulesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListEventRulesResponseBodyData) SetPageData(v []*ListEventRulesResponseBodyDataPageData) *ListEventRulesResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListEventRulesResponseBodyData) SetPageNum(v int64) *ListEventRulesResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListEventRulesResponseBodyData) SetPageSize(v int64) *ListEventRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListEventRulesResponseBodyData) SetPages(v int64) *ListEventRulesResponseBodyData {
	s.Pages = &v
	return s
}

func (s *ListEventRulesResponseBodyData) SetSize(v int64) *ListEventRulesResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListEventRulesResponseBodyData) SetTotal(v int64) *ListEventRulesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListEventRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListEventRulesResponseBodyDataPageData struct {
	// example:
	//
	// --
	DeliveryMode *string                                         `json:"DeliveryMode,omitempty" xml:"DeliveryMode,omitempty"`
	Endpoint     *ListEventRulesResponseBodyDataPageDataEndpoint `json:"Endpoint,omitempty" xml:"Endpoint,omitempty" type:"Struct"`
	EventTypes   []*string                                       `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	MatchRules   [][]*EventMatchRule                             `json:"MatchRules,omitempty" xml:"MatchRules,omitempty" type:"Repeated"`
	// example:
	//
	// rule-xsXDW
	RuleName      *string                                                `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Subscriptions []*ListEventRulesResponseBodyDataPageDataSubscriptions `json:"Subscriptions,omitempty" xml:"Subscriptions,omitempty" type:"Repeated"`
	// example:
	//
	// demo-topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s ListEventRulesResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s ListEventRulesResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListEventRulesResponseBodyDataPageData) GetDeliveryMode() *string {
	return s.DeliveryMode
}

func (s *ListEventRulesResponseBodyDataPageData) GetEndpoint() *ListEventRulesResponseBodyDataPageDataEndpoint {
	return s.Endpoint
}

func (s *ListEventRulesResponseBodyDataPageData) GetEventTypes() []*string {
	return s.EventTypes
}

func (s *ListEventRulesResponseBodyDataPageData) GetMatchRules() [][]*EventMatchRule {
	return s.MatchRules
}

func (s *ListEventRulesResponseBodyDataPageData) GetRuleName() *string {
	return s.RuleName
}

func (s *ListEventRulesResponseBodyDataPageData) GetSubscriptions() []*ListEventRulesResponseBodyDataPageDataSubscriptions {
	return s.Subscriptions
}

func (s *ListEventRulesResponseBodyDataPageData) GetTopicName() *string {
	return s.TopicName
}

func (s *ListEventRulesResponseBodyDataPageData) SetDeliveryMode(v string) *ListEventRulesResponseBodyDataPageData {
	s.DeliveryMode = &v
	return s
}

func (s *ListEventRulesResponseBodyDataPageData) SetEndpoint(v *ListEventRulesResponseBodyDataPageDataEndpoint) *ListEventRulesResponseBodyDataPageData {
	s.Endpoint = v
	return s
}

func (s *ListEventRulesResponseBodyDataPageData) SetEventTypes(v []*string) *ListEventRulesResponseBodyDataPageData {
	s.EventTypes = v
	return s
}

func (s *ListEventRulesResponseBodyDataPageData) SetMatchRules(v [][]*EventMatchRule) *ListEventRulesResponseBodyDataPageData {
	s.MatchRules = v
	return s
}

func (s *ListEventRulesResponseBodyDataPageData) SetRuleName(v string) *ListEventRulesResponseBodyDataPageData {
	s.RuleName = &v
	return s
}

func (s *ListEventRulesResponseBodyDataPageData) SetSubscriptions(v []*ListEventRulesResponseBodyDataPageDataSubscriptions) *ListEventRulesResponseBodyDataPageData {
	s.Subscriptions = v
	return s
}

func (s *ListEventRulesResponseBodyDataPageData) SetTopicName(v string) *ListEventRulesResponseBodyDataPageData {
	s.TopicName = &v
	return s
}

func (s *ListEventRulesResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}

type ListEventRulesResponseBodyDataPageDataEndpoint struct {
	// example:
	//
	// topic
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// example:
	//
	// test-topic
	EndpointValue *string `json:"EndpointValue,omitempty" xml:"EndpointValue,omitempty"`
}

func (s ListEventRulesResponseBodyDataPageDataEndpoint) String() string {
	return dara.Prettify(s)
}

func (s ListEventRulesResponseBodyDataPageDataEndpoint) GoString() string {
	return s.String()
}

func (s *ListEventRulesResponseBodyDataPageDataEndpoint) GetEndpointType() *string {
	return s.EndpointType
}

func (s *ListEventRulesResponseBodyDataPageDataEndpoint) GetEndpointValue() *string {
	return s.EndpointValue
}

func (s *ListEventRulesResponseBodyDataPageDataEndpoint) SetEndpointType(v string) *ListEventRulesResponseBodyDataPageDataEndpoint {
	s.EndpointType = &v
	return s
}

func (s *ListEventRulesResponseBodyDataPageDataEndpoint) SetEndpointValue(v string) *ListEventRulesResponseBodyDataPageDataEndpoint {
	s.EndpointValue = &v
	return s
}

func (s *ListEventRulesResponseBodyDataPageDataEndpoint) Validate() error {
	return dara.Validate(s)
}

type ListEventRulesResponseBodyDataPageDataSubscriptions struct {
	// example:
	//
	// queue
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// example:
	//
	// wyx-vp-complete-queue
	EndpointValue *string `json:"EndpointValue,omitempty" xml:"EndpointValue,omitempty"`
}

func (s ListEventRulesResponseBodyDataPageDataSubscriptions) String() string {
	return dara.Prettify(s)
}

func (s ListEventRulesResponseBodyDataPageDataSubscriptions) GoString() string {
	return s.String()
}

func (s *ListEventRulesResponseBodyDataPageDataSubscriptions) GetEndpointType() *string {
	return s.EndpointType
}

func (s *ListEventRulesResponseBodyDataPageDataSubscriptions) GetEndpointValue() *string {
	return s.EndpointValue
}

func (s *ListEventRulesResponseBodyDataPageDataSubscriptions) SetEndpointType(v string) *ListEventRulesResponseBodyDataPageDataSubscriptions {
	s.EndpointType = &v
	return s
}

func (s *ListEventRulesResponseBodyDataPageDataSubscriptions) SetEndpointValue(v string) *ListEventRulesResponseBodyDataPageDataSubscriptions {
	s.EndpointValue = &v
	return s
}

func (s *ListEventRulesResponseBodyDataPageDataSubscriptions) Validate() error {
	return dara.Validate(s)
}
