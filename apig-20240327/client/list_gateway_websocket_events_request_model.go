// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayWebsocketEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionId(v string) *ListGatewayWebsocketEventsRequest
	GetConnectionId() *string
	SetEndTime(v int64) *ListGatewayWebsocketEventsRequest
	GetEndTime() *int64
	SetMaxResults(v int32) *ListGatewayWebsocketEventsRequest
	GetMaxResults() *int32
	SetModel(v string) *ListGatewayWebsocketEventsRequest
	GetModel() *string
	SetNextToken(v string) *ListGatewayWebsocketEventsRequest
	GetNextToken() *string
	SetProvider(v string) *ListGatewayWebsocketEventsRequest
	GetProvider() *string
	SetSource(v string) *ListGatewayWebsocketEventsRequest
	GetSource() *string
	SetStartTime(v int64) *ListGatewayWebsocketEventsRequest
	GetStartTime() *int64
	SetTurnId(v string) *ListGatewayWebsocketEventsRequest
	GetTurnId() *string
}

type ListGatewayWebsocketEventsRequest struct {
	// The connection identifier.
	//
	// example:
	//
	// connection-example
	ConnectionId *string `json:"connectionId,omitempty" xml:"connectionId,omitempty"`
	// The end time. The value is a UNIX timestamp in seconds since the UTC epoch. Example: 1793919600.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1793919600
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The maximum number of results per page for cursor-based pagination.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The model.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// opaque-token-from-previous-page
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The provider.
	//
	// example:
	//
	// dashscope
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// The log source.
	//
	// example:
	//
	// access.log
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The start time. The value is a UNIX timestamp in seconds since the UTC epoch. Example: 1793916000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1793916000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The turn identifier.
	//
	// example:
	//
	// turn-example
	TurnId *string `json:"turnId,omitempty" xml:"turnId,omitempty"`
}

func (s ListGatewayWebsocketEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayWebsocketEventsRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayWebsocketEventsRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *ListGatewayWebsocketEventsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListGatewayWebsocketEventsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGatewayWebsocketEventsRequest) GetModel() *string {
	return s.Model
}

func (s *ListGatewayWebsocketEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGatewayWebsocketEventsRequest) GetProvider() *string {
	return s.Provider
}

func (s *ListGatewayWebsocketEventsRequest) GetSource() *string {
	return s.Source
}

func (s *ListGatewayWebsocketEventsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListGatewayWebsocketEventsRequest) GetTurnId() *string {
	return s.TurnId
}

func (s *ListGatewayWebsocketEventsRequest) SetConnectionId(v string) *ListGatewayWebsocketEventsRequest {
	s.ConnectionId = &v
	return s
}

func (s *ListGatewayWebsocketEventsRequest) SetEndTime(v int64) *ListGatewayWebsocketEventsRequest {
	s.EndTime = &v
	return s
}

func (s *ListGatewayWebsocketEventsRequest) SetMaxResults(v int32) *ListGatewayWebsocketEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGatewayWebsocketEventsRequest) SetModel(v string) *ListGatewayWebsocketEventsRequest {
	s.Model = &v
	return s
}

func (s *ListGatewayWebsocketEventsRequest) SetNextToken(v string) *ListGatewayWebsocketEventsRequest {
	s.NextToken = &v
	return s
}

func (s *ListGatewayWebsocketEventsRequest) SetProvider(v string) *ListGatewayWebsocketEventsRequest {
	s.Provider = &v
	return s
}

func (s *ListGatewayWebsocketEventsRequest) SetSource(v string) *ListGatewayWebsocketEventsRequest {
	s.Source = &v
	return s
}

func (s *ListGatewayWebsocketEventsRequest) SetStartTime(v int64) *ListGatewayWebsocketEventsRequest {
	s.StartTime = &v
	return s
}

func (s *ListGatewayWebsocketEventsRequest) SetTurnId(v string) *ListGatewayWebsocketEventsRequest {
	s.TurnId = &v
	return s
}

func (s *ListGatewayWebsocketEventsRequest) Validate() error {
	return dara.Validate(s)
}
