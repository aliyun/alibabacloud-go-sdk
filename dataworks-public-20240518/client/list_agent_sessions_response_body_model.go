// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v *ListAgentSessionsResponseBodyJsonRpcResponse) *ListAgentSessionsResponseBody
	GetJsonRpcResponse() *ListAgentSessionsResponseBodyJsonRpcResponse
	SetRequestId(v string) *ListAgentSessionsResponseBody
	GetRequestId() *string
}

type ListAgentSessionsResponseBody struct {
	JsonRpcResponse *ListAgentSessionsResponseBodyJsonRpcResponse `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 50C5A9F7-B5BD-58B2-9EB8-ADFFA9E6A56F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAgentSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentSessionsResponseBody) GetJsonRpcResponse() *ListAgentSessionsResponseBodyJsonRpcResponse {
	return s.JsonRpcResponse
}

func (s *ListAgentSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentSessionsResponseBody) SetJsonRpcResponse(v *ListAgentSessionsResponseBodyJsonRpcResponse) *ListAgentSessionsResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *ListAgentSessionsResponseBody) SetRequestId(v string) *ListAgentSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentSessionsResponseBody) Validate() error {
	if s.JsonRpcResponse != nil {
		if err := s.JsonRpcResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentSessionsResponseBodyJsonRpcResponse struct {
	// example:
	//
	// 29d9a29c-a284-48c1-9eaa-4a42c7c616d5
	Id      *string                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	Jsonrpc *string                                             `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Result  *ListAgentSessionsResponseBodyJsonRpcResponseResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListAgentSessionsResponseBodyJsonRpcResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionsResponseBodyJsonRpcResponse) GoString() string {
	return s.String()
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponse) GetId() *string {
	return s.Id
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponse) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponse) GetResult() *ListAgentSessionsResponseBodyJsonRpcResponseResult {
	return s.Result
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponse) SetId(v string) *ListAgentSessionsResponseBodyJsonRpcResponse {
	s.Id = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponse) SetJsonrpc(v string) *ListAgentSessionsResponseBodyJsonRpcResponse {
	s.Jsonrpc = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponse) SetResult(v *ListAgentSessionsResponseBodyJsonRpcResponseResult) *ListAgentSessionsResponseBodyJsonRpcResponse {
	s.Result = v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponse) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentSessionsResponseBodyJsonRpcResponseResult struct {
	AgentSessions []*ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions `json:"AgentSessions,omitempty" xml:"AgentSessions,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgentSessionsResponseBodyJsonRpcResponseResult) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionsResponseBodyJsonRpcResponseResult) GoString() string {
	return s.String()
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResult) GetAgentSessions() []*ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions {
	return s.AgentSessions
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResult) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResult) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResult) SetAgentSessions(v []*ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions) *ListAgentSessionsResponseBodyJsonRpcResponseResult {
	s.AgentSessions = v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResult) SetMaxResults(v int32) *ListAgentSessionsResponseBodyJsonRpcResponseResult {
	s.MaxResults = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResult) SetNextToken(v string) *ListAgentSessionsResponseBodyJsonRpcResponseResult {
	s.NextToken = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResult) SetTotalCount(v int32) *ListAgentSessionsResponseBodyJsonRpcResponseResult {
	s.TotalCount = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResult) Validate() error {
	if s.AgentSessions != nil {
		for _, item := range s.AgentSessions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions struct {
	Meta *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMeta `json:"Meta,omitempty" xml:"Meta,omitempty" type:"Struct"`
	// example:
	//
	// 123456789
	SessionCreatedAt   *int64  `json:"SessionCreatedAt,omitempty" xml:"SessionCreatedAt,omitempty"`
	SessionDescription *string `json:"SessionDescription,omitempty" xml:"SessionDescription,omitempty"`
	// example:
	//
	// sess_0f12abc34
	SessionId    *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SessionTitle *string `json:"SessionTitle,omitempty" xml:"SessionTitle,omitempty"`
	// example:
	//
	// 123456789
	SessionUpdatedAt *int64 `json:"SessionUpdatedAt,omitempty" xml:"SessionUpdatedAt,omitempty"`
}

func (s ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions) GoString() string {
	return s.String()
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions) GetMeta() *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMeta {
	return s.Meta
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions) GetSessionCreatedAt() *int64 {
	return s.SessionCreatedAt
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions) GetSessionDescription() *string {
	return s.SessionDescription
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions) GetSessionId() *string {
	return s.SessionId
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions) GetSessionTitle() *string {
	return s.SessionTitle
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions) GetSessionUpdatedAt() *int64 {
	return s.SessionUpdatedAt
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions) SetMeta(v *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMeta) *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions {
	s.Meta = v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions) SetSessionCreatedAt(v int64) *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions {
	s.SessionCreatedAt = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions) SetSessionDescription(v string) *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions {
	s.SessionDescription = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions) SetSessionId(v string) *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions {
	s.SessionId = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions) SetSessionTitle(v string) *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions {
	s.SessionTitle = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions) SetSessionUpdatedAt(v int64) *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions {
	s.SessionUpdatedAt = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessions) Validate() error {
	if s.Meta != nil {
		if err := s.Meta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMeta struct {
	// example:
	//
	// openapi_sdk
	SessionSource  *string                                                                              `json:"SessionSource,omitempty" xml:"SessionSource,omitempty"`
	SessionStatus  *string                                                                              `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	SessionTagList []*ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMetaSessionTagList `json:"SessionTagList,omitempty" xml:"SessionTagList,omitempty" type:"Repeated"`
}

func (s ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMeta) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMeta) GoString() string {
	return s.String()
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMeta) GetSessionSource() *string {
	return s.SessionSource
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMeta) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMeta) GetSessionTagList() []*ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMetaSessionTagList {
	return s.SessionTagList
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMeta) SetSessionSource(v string) *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMeta {
	s.SessionSource = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMeta) SetSessionStatus(v string) *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMeta {
	s.SessionStatus = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMeta) SetSessionTagList(v []*ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMetaSessionTagList) *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMeta {
	s.SessionTagList = v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMeta) Validate() error {
	if s.SessionTagList != nil {
		for _, item := range s.SessionTagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMetaSessionTagList struct {
	// example:
	//
	// user_123456
	SessionTagCode *string `json:"SessionTagCode,omitempty" xml:"SessionTagCode,omitempty"`
}

func (s ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMetaSessionTagList) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMetaSessionTagList) GoString() string {
	return s.String()
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMetaSessionTagList) GetSessionTagCode() *string {
	return s.SessionTagCode
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMetaSessionTagList) SetSessionTagCode(v string) *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMetaSessionTagList {
	s.SessionTagCode = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResultAgentSessionsMetaSessionTagList) Validate() error {
	return dara.Validate(s)
}
