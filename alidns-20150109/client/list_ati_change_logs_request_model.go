// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAtiChangeLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListAtiChangeLogsRequest
	GetAgentId() *string
	SetClientToken(v string) *ListAtiChangeLogsRequest
	GetClientToken() *string
	SetEndTimestamp(v int64) *ListAtiChangeLogsRequest
	GetEndTimestamp() *int64
	SetOperationType(v string) *ListAtiChangeLogsRequest
	GetOperationType() *string
	SetOperatorAccount(v string) *ListAtiChangeLogsRequest
	GetOperatorAccount() *string
	SetPageNumber(v int32) *ListAtiChangeLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAtiChangeLogsRequest
	GetPageSize() *int32
	SetStartTimestamp(v int64) *ListAtiChangeLogsRequest
	GetStartTimestamp() *int64
	SetTimeRange(v string) *ListAtiChangeLogsRequest
	GetTimeRange() *string
}

type ListAtiChangeLogsRequest struct {
	// The agent ID that is assigned by CNNIC after real-name verification. The AgentID serves as the unique identifier that binds the agent to the real-name registered contact.
	//
	// example:
	//
	// Justin@underarmour
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// Ensures the idempotency of the request. Generate a unique parameter value from your client to ensure that the value is unique across different requests. The ClientToken value supports only ASCII characters and cannot exceed 64 characters in length.
	//
	// - If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may vary for each API request.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The end time of the query (timestamp).
	//
	// example:
	//
	// 1474335170000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The operation type of the Operation logs log record, such as modifying an agent.
	//
	// example:
	//
	// 2074753647748672512
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The UID of the operator.
	//
	// example:
	//
	// 1646808646232999
	OperatorAccount *string `json:"OperatorAccount,omitempty" xml:"OperatorAccount,omitempty"`
	// The current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. Maximum value: 100. Default value: 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time of the query (timestamp).
	//
	// example:
	//
	// 1474335170000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// Ignored.
	//
	// example:
	//
	// 忽略
	TimeRange *string `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s ListAtiChangeLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAtiChangeLogsRequest) GoString() string {
	return s.String()
}

func (s *ListAtiChangeLogsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAtiChangeLogsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListAtiChangeLogsRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *ListAtiChangeLogsRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *ListAtiChangeLogsRequest) GetOperatorAccount() *string {
	return s.OperatorAccount
}

func (s *ListAtiChangeLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAtiChangeLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAtiChangeLogsRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *ListAtiChangeLogsRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *ListAtiChangeLogsRequest) SetAgentId(v string) *ListAtiChangeLogsRequest {
	s.AgentId = &v
	return s
}

func (s *ListAtiChangeLogsRequest) SetClientToken(v string) *ListAtiChangeLogsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListAtiChangeLogsRequest) SetEndTimestamp(v int64) *ListAtiChangeLogsRequest {
	s.EndTimestamp = &v
	return s
}

func (s *ListAtiChangeLogsRequest) SetOperationType(v string) *ListAtiChangeLogsRequest {
	s.OperationType = &v
	return s
}

func (s *ListAtiChangeLogsRequest) SetOperatorAccount(v string) *ListAtiChangeLogsRequest {
	s.OperatorAccount = &v
	return s
}

func (s *ListAtiChangeLogsRequest) SetPageNumber(v int32) *ListAtiChangeLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAtiChangeLogsRequest) SetPageSize(v int32) *ListAtiChangeLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAtiChangeLogsRequest) SetStartTimestamp(v int64) *ListAtiChangeLogsRequest {
	s.StartTimestamp = &v
	return s
}

func (s *ListAtiChangeLogsRequest) SetTimeRange(v string) *ListAtiChangeLogsRequest {
	s.TimeRange = &v
	return s
}

func (s *ListAtiChangeLogsRequest) Validate() error {
	return dara.Validate(s)
}
