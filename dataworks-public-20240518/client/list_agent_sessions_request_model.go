// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ListAgentSessionsRequest
	GetId() *string
	SetJsonrpc(v string) *ListAgentSessionsRequest
	GetJsonrpc() *string
	SetParams(v *ListAgentSessionsRequestParams) *ListAgentSessionsRequest
	GetParams() *ListAgentSessionsRequestParams
}

type ListAgentSessionsRequest struct {
	// example:
	//
	// 676303114031776
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string                         `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Params  *ListAgentSessionsRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
}

func (s ListAgentSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentSessionsRequest) GetId() *string {
	return s.Id
}

func (s *ListAgentSessionsRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *ListAgentSessionsRequest) GetParams() *ListAgentSessionsRequestParams {
	return s.Params
}

func (s *ListAgentSessionsRequest) SetId(v string) *ListAgentSessionsRequest {
	s.Id = &v
	return s
}

func (s *ListAgentSessionsRequest) SetJsonrpc(v string) *ListAgentSessionsRequest {
	s.Jsonrpc = &v
	return s
}

func (s *ListAgentSessionsRequest) SetParams(v *ListAgentSessionsRequestParams) *ListAgentSessionsRequest {
	s.Params = v
	return s
}

func (s *ListAgentSessionsRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentSessionsRequestParams struct {
	// example:
	//
	// chat_cli_chatbi
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 10
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// sess_0f12abc34
	SessionId         *string   `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SessionSourceList []*string `json:"SessionSourceList,omitempty" xml:"SessionSourceList,omitempty" type:"Repeated"`
	SessionTitle      *string   `json:"SessionTitle,omitempty" xml:"SessionTitle,omitempty"`
	TagList           []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
}

func (s ListAgentSessionsRequestParams) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionsRequestParams) GoString() string {
	return s.String()
}

func (s *ListAgentSessionsRequestParams) GetAgentName() *string {
	return s.AgentName
}

func (s *ListAgentSessionsRequestParams) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentSessionsRequestParams) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentSessionsRequestParams) GetSessionId() *string {
	return s.SessionId
}

func (s *ListAgentSessionsRequestParams) GetSessionSourceList() []*string {
	return s.SessionSourceList
}

func (s *ListAgentSessionsRequestParams) GetSessionTitle() *string {
	return s.SessionTitle
}

func (s *ListAgentSessionsRequestParams) GetTagList() []*string {
	return s.TagList
}

func (s *ListAgentSessionsRequestParams) SetAgentName(v string) *ListAgentSessionsRequestParams {
	s.AgentName = &v
	return s
}

func (s *ListAgentSessionsRequestParams) SetMaxResults(v int32) *ListAgentSessionsRequestParams {
	s.MaxResults = &v
	return s
}

func (s *ListAgentSessionsRequestParams) SetNextToken(v string) *ListAgentSessionsRequestParams {
	s.NextToken = &v
	return s
}

func (s *ListAgentSessionsRequestParams) SetSessionId(v string) *ListAgentSessionsRequestParams {
	s.SessionId = &v
	return s
}

func (s *ListAgentSessionsRequestParams) SetSessionSourceList(v []*string) *ListAgentSessionsRequestParams {
	s.SessionSourceList = v
	return s
}

func (s *ListAgentSessionsRequestParams) SetSessionTitle(v string) *ListAgentSessionsRequestParams {
	s.SessionTitle = &v
	return s
}

func (s *ListAgentSessionsRequestParams) SetTagList(v []*string) *ListAgentSessionsRequestParams {
	s.TagList = v
	return s
}

func (s *ListAgentSessionsRequestParams) Validate() error {
	return dara.Validate(s)
}
