// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryChatInstanceSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionIds(v []*string) *QueryChatInstanceSessionsRequest
	GetSessionIds() []*string
}

type QueryChatInstanceSessionsRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// ["8C9F2D4E-7A6B-4F1C-9E53-0B2C8D3F9A4B"]
	SessionIds []*string `json:"sessionIds,omitempty" xml:"sessionIds,omitempty" type:"Repeated"`
}

func (s QueryChatInstanceSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryChatInstanceSessionsRequest) GoString() string {
	return s.String()
}

func (s *QueryChatInstanceSessionsRequest) GetSessionIds() []*string {
	return s.SessionIds
}

func (s *QueryChatInstanceSessionsRequest) SetSessionIds(v []*string) *QueryChatInstanceSessionsRequest {
	s.SessionIds = v
	return s
}

func (s *QueryChatInstanceSessionsRequest) Validate() error {
	return dara.Validate(s)
}
