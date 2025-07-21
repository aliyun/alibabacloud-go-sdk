// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExchangesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListExchangesResponseBodyData) *ListExchangesResponseBody
	GetData() *ListExchangesResponseBodyData
	SetRequestId(v string) *ListExchangesResponseBody
	GetRequestId() *string
}

type ListExchangesResponseBody struct {
	// The returned data.
	Data *ListExchangesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// FEBA5E0C-50D0-4FA6-A794-4901E5465***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListExchangesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExchangesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExchangesResponseBody) GetData() *ListExchangesResponseBodyData {
	return s.Data
}

func (s *ListExchangesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExchangesResponseBody) SetData(v *ListExchangesResponseBodyData) *ListExchangesResponseBody {
	s.Data = v
	return s
}

func (s *ListExchangesResponseBody) SetRequestId(v string) *ListExchangesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExchangesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListExchangesResponseBodyData struct {
	// The exchanges.
	Exchanges []*ListExchangesResponseBodyDataExchanges `json:"Exchanges,omitempty" xml:"Exchanges,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page.``
	//
	// 	- If the value of this parameter is empty, the next query is not required and the token used to start the next query is unavailable.``
	//
	// 	- If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.``
	//
	// example:
	//
	// AAAANDQBYW1xcC1jbi03cHAybXdiY3AwMGEBdmhvc3QBAXNkZndhYWJhATE2NDkzMTM4OTU5NDIB4o3z1pPwWzk4aYuiRffi8R6-****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListExchangesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListExchangesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExchangesResponseBodyData) GetExchanges() []*ListExchangesResponseBodyDataExchanges {
	return s.Exchanges
}

func (s *ListExchangesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExchangesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExchangesResponseBodyData) SetExchanges(v []*ListExchangesResponseBodyDataExchanges) *ListExchangesResponseBodyData {
	s.Exchanges = v
	return s
}

func (s *ListExchangesResponseBodyData) SetMaxResults(v int32) *ListExchangesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListExchangesResponseBodyData) SetNextToken(v string) *ListExchangesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListExchangesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListExchangesResponseBodyDataExchanges struct {
	// The attributes. This parameter is unavailable in the current version.
	//
	// example:
	//
	// test
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// Indicates whether the exchange was automatically deleted.
	//
	// example:
	//
	// false
	AutoDeleteState *bool `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	// The timestamp that indicates when the exchange was created. Unit: milliseconds.
	//
	// example:
	//
	// 1580886216000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The exchange type.
	//
	// example:
	//
	// DIRECT
	ExchangeType *string `json:"ExchangeType,omitempty" xml:"ExchangeType,omitempty"`
	// The exchange name.
	//
	// example:
	//
	// amq.direct
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The vhost name.
	//
	// example:
	//
	// test
	VHostName *string `json:"VHostName,omitempty" xml:"VHostName,omitempty"`
}

func (s ListExchangesResponseBodyDataExchanges) String() string {
	return dara.Prettify(s)
}

func (s ListExchangesResponseBodyDataExchanges) GoString() string {
	return s.String()
}

func (s *ListExchangesResponseBodyDataExchanges) GetAttributes() map[string]interface{} {
	return s.Attributes
}

func (s *ListExchangesResponseBodyDataExchanges) GetAutoDeleteState() *bool {
	return s.AutoDeleteState
}

func (s *ListExchangesResponseBodyDataExchanges) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListExchangesResponseBodyDataExchanges) GetExchangeType() *string {
	return s.ExchangeType
}

func (s *ListExchangesResponseBodyDataExchanges) GetName() *string {
	return s.Name
}

func (s *ListExchangesResponseBodyDataExchanges) GetVHostName() *string {
	return s.VHostName
}

func (s *ListExchangesResponseBodyDataExchanges) SetAttributes(v map[string]interface{}) *ListExchangesResponseBodyDataExchanges {
	s.Attributes = v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetAutoDeleteState(v bool) *ListExchangesResponseBodyDataExchanges {
	s.AutoDeleteState = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetCreateTime(v int64) *ListExchangesResponseBodyDataExchanges {
	s.CreateTime = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetExchangeType(v string) *ListExchangesResponseBodyDataExchanges {
	s.ExchangeType = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetName(v string) *ListExchangesResponseBodyDataExchanges {
	s.Name = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetVHostName(v string) *ListExchangesResponseBodyDataExchanges {
	s.VHostName = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) Validate() error {
	return dara.Validate(s)
}
