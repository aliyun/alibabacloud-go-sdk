// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDialogDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *GetDialogDetailRequest
	GetSessionId() *string
}

type GetDialogDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1906623923815534xxx
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s GetDialogDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDialogDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDialogDetailRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetDialogDetailRequest) SetSessionId(v string) *GetDialogDetailRequest {
	s.SessionId = &v
	return s
}

func (s *GetDialogDetailRequest) Validate() error {
	return dara.Validate(s)
}
