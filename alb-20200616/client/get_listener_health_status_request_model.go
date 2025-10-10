// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListenerHealthStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeRule(v bool) *GetListenerHealthStatusRequest
	GetIncludeRule() *bool
	SetListenerId(v string) *GetListenerHealthStatusRequest
	GetListenerId() *string
	SetMaxResults(v int64) *GetListenerHealthStatusRequest
	GetMaxResults() *int64
	SetNextToken(v string) *GetListenerHealthStatusRequest
	GetNextToken() *string
}

type GetListenerHealthStatusRequest struct {
	// Specifies whether to return the health check results of forwarding rules. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	IncludeRule *bool `json:"IncludeRule,omitempty" xml:"IncludeRule,omitempty"`
	// The listener ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **30**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// 	- If this is your first query or no next queries are to be sent, ignore this parameter.
	//
	// 	- If a next query is to be sent, set the value to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetListenerHealthStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetListenerHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusRequest) GetIncludeRule() *bool {
	return s.IncludeRule
}

func (s *GetListenerHealthStatusRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *GetListenerHealthStatusRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *GetListenerHealthStatusRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetListenerHealthStatusRequest) SetIncludeRule(v bool) *GetListenerHealthStatusRequest {
	s.IncludeRule = &v
	return s
}

func (s *GetListenerHealthStatusRequest) SetListenerId(v string) *GetListenerHealthStatusRequest {
	s.ListenerId = &v
	return s
}

func (s *GetListenerHealthStatusRequest) SetMaxResults(v int64) *GetListenerHealthStatusRequest {
	s.MaxResults = &v
	return s
}

func (s *GetListenerHealthStatusRequest) SetNextToken(v string) *GetListenerHealthStatusRequest {
	s.NextToken = &v
	return s
}

func (s *GetListenerHealthStatusRequest) Validate() error {
	return dara.Validate(s)
}
