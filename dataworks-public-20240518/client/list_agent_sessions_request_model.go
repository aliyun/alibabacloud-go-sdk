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
	// The ID passed by the requester, which is returned as is.
	//
	// example:
	//
	// 676303114031776
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC version. Fixed to 2.0.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The business parameters.
	Params *ListAgentSessionsRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
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
	// The exact filter for the **agent name**. You must specify at least one of this parameter and SessionSourceList.
	//
	// example:
	//
	// chat_cli_chatbi
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// Specifies whether to return only archived sessions. Valid values:
	//
	// - true: Only archived sessions are returned.
	//
	// - false or not specified: Only unarchived sessions are returned.
	Archived *bool `json:"Archived,omitempty" xml:"Archived,omitempty"`
	// The number of entries per page. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page. Set this parameter to 1 for the first page.
	//
	// example:
	//
	// 10
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The exact match for the session ID.
	//
	// example:
	//
	// sess_0f12abc34
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The filter for the **session source list**. You must specify at least one of this parameter and AgentName.
	SessionSourceList []*string `json:"SessionSourceList,omitempty" xml:"SessionSourceList,omitempty" type:"Repeated"`
	// The fuzzy filter for the session title.
	//
	// example:
	//
	// Sales in the last 30 days
	SessionTitle *string `json:"SessionTitle,omitempty" xml:"SessionTitle,omitempty"`
	// The filter for the session tag code list.
	TagList []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
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

func (s *ListAgentSessionsRequestParams) GetArchived() *bool {
	return s.Archived
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

func (s *ListAgentSessionsRequestParams) SetArchived(v bool) *ListAgentSessionsRequestParams {
	s.Archived = &v
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
