// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseChatInstanceSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionIds(v []*string) *CloseChatInstanceSessionsRequest
	GetSessionIds() []*string
}

type CloseChatInstanceSessionsRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// ["8C9F2D4E-7A6B-4F1C-9E53-0B2C8D3F9A4B"]
	SessionIds []*string `json:"sessionIds,omitempty" xml:"sessionIds,omitempty" type:"Repeated"`
}

func (s CloseChatInstanceSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseChatInstanceSessionsRequest) GoString() string {
	return s.String()
}

func (s *CloseChatInstanceSessionsRequest) GetSessionIds() []*string {
	return s.SessionIds
}

func (s *CloseChatInstanceSessionsRequest) SetSessionIds(v []*string) *CloseChatInstanceSessionsRequest {
	s.SessionIds = v
	return s
}

func (s *CloseChatInstanceSessionsRequest) Validate() error {
	return dara.Validate(s)
}
