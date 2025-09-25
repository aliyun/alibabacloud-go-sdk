// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDialogLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetDialogLogRequest
	GetId() *string
	SetSessionId(v string) *GetDialogLogRequest
	GetSessionId() *string
}

type GetDialogLogRequest struct {
	// example:
	//
	// 175600129454077743fb03ac54955a4be72ec08f9c216
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1758010668S001w4paq82azm
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s GetDialogLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDialogLogRequest) GoString() string {
	return s.String()
}

func (s *GetDialogLogRequest) GetId() *string {
	return s.Id
}

func (s *GetDialogLogRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetDialogLogRequest) SetId(v string) *GetDialogLogRequest {
	s.Id = &v
	return s
}

func (s *GetDialogLogRequest) SetSessionId(v string) *GetDialogLogRequest {
	s.SessionId = &v
	return s
}

func (s *GetDialogLogRequest) Validate() error {
	return dara.Validate(s)
}
