// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayWebsocketEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGatewayWebsocketEventsResponseBody
	GetCode() *string
	SetData(v *ListGatewayWebsocketEventsResponseBodyData) *ListGatewayWebsocketEventsResponseBody
	GetData() *ListGatewayWebsocketEventsResponseBodyData
	SetMessage(v string) *ListGatewayWebsocketEventsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewayWebsocketEventsResponseBody
	GetRequestId() *string
}

type ListGatewayWebsocketEventsResponseBody struct {
	// The business response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The query result.
	//
	// example:
	//
	// {}
	Data *ListGatewayWebsocketEventsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// The specified gateway does not exist.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D0A6A1A0-8793-4C10-AB1A-2C03E770A912
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGatewayWebsocketEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayWebsocketEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayWebsocketEventsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGatewayWebsocketEventsResponseBody) GetData() *ListGatewayWebsocketEventsResponseBodyData {
	return s.Data
}

func (s *ListGatewayWebsocketEventsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewayWebsocketEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayWebsocketEventsResponseBody) SetCode(v string) *ListGatewayWebsocketEventsResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayWebsocketEventsResponseBody) SetData(v *ListGatewayWebsocketEventsResponseBodyData) *ListGatewayWebsocketEventsResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayWebsocketEventsResponseBody) SetMessage(v string) *ListGatewayWebsocketEventsResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayWebsocketEventsResponseBody) SetRequestId(v string) *ListGatewayWebsocketEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayWebsocketEventsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayWebsocketEventsResponseBodyData struct {
	// The data basis.
	//
	// example:
	//
	// observed
	Basis *string `json:"basis,omitempty" xml:"basis,omitempty"`
	// Indicates whether more pages are available.
	//
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The observability events.
	//
	// example:
	//
	// []
	Items []map[string]interface{} `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The actual number of results per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The next pagination token.
	//
	// example:
	//
	// opaque-token-from-previous-page
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListGatewayWebsocketEventsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayWebsocketEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayWebsocketEventsResponseBodyData) GetBasis() *string {
	return s.Basis
}

func (s *ListGatewayWebsocketEventsResponseBodyData) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListGatewayWebsocketEventsResponseBodyData) GetItems() []map[string]interface{} {
	return s.Items
}

func (s *ListGatewayWebsocketEventsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGatewayWebsocketEventsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGatewayWebsocketEventsResponseBodyData) SetBasis(v string) *ListGatewayWebsocketEventsResponseBodyData {
	s.Basis = &v
	return s
}

func (s *ListGatewayWebsocketEventsResponseBodyData) SetHasMore(v bool) *ListGatewayWebsocketEventsResponseBodyData {
	s.HasMore = &v
	return s
}

func (s *ListGatewayWebsocketEventsResponseBodyData) SetItems(v []map[string]interface{}) *ListGatewayWebsocketEventsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListGatewayWebsocketEventsResponseBodyData) SetMaxResults(v int32) *ListGatewayWebsocketEventsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListGatewayWebsocketEventsResponseBodyData) SetNextToken(v string) *ListGatewayWebsocketEventsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListGatewayWebsocketEventsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
