// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentIMChannelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelType(v string) *ListAgentIMChannelsRequest
	GetChannelType() *string
	SetMaxResults(v int32) *ListAgentIMChannelsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgentIMChannelsRequest
	GetNextToken() *string
	SetStatus(v string) *ListAgentIMChannelsRequest
	GetStatus() *string
}

type ListAgentIMChannelsRequest struct {
	// The IM channel type. Valid values:
	//
	// - DINGTALK: DingTalk.
	//
	// - FEISHU: Lark.
	//
	// - WECOM: WeCom.
	//
	// example:
	//
	// DINGTALK
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// The maximum number of entries to return per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. You do not need to specify this parameter for the first request. For subsequent requests, use the nextToken value returned in the previous response.
	//
	// example:
	//
	// next-token-1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The IM channel status. Valid values:
	//
	// - CREATING: being created.
	//
	// - READY: ready.
	//
	// - UPDATING: being updated.
	//
	// - FAILED: failed.
	//
	// - DELETING: being deleted.
	//
	// - DELETE_FAILED: deletion failed.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListAgentIMChannelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentIMChannelsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentIMChannelsRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *ListAgentIMChannelsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentIMChannelsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentIMChannelsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAgentIMChannelsRequest) SetChannelType(v string) *ListAgentIMChannelsRequest {
	s.ChannelType = &v
	return s
}

func (s *ListAgentIMChannelsRequest) SetMaxResults(v int32) *ListAgentIMChannelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAgentIMChannelsRequest) SetNextToken(v string) *ListAgentIMChannelsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAgentIMChannelsRequest) SetStatus(v string) *ListAgentIMChannelsRequest {
	s.Status = &v
	return s
}

func (s *ListAgentIMChannelsRequest) Validate() error {
	return dara.Validate(s)
}
