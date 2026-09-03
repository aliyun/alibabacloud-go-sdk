// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *GetSessionContentRequest
	GetSessionId() *string
}

type GetSessionContentRequest struct {
	// The session ID.
	//
	// example:
	//
	// UUID
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s GetSessionContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSessionContentRequest) GoString() string {
	return s.String()
}

func (s *GetSessionContentRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetSessionContentRequest) SetSessionId(v string) *GetSessionContentRequest {
	s.SessionId = &v
	return s
}

func (s *GetSessionContentRequest) Validate() error {
	return dara.Validate(s)
}
