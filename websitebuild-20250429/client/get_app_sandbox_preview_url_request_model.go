// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSandboxPreviewUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *GetAppSandboxPreviewUrlRequest
	GetConversationId() *string
	SetRestart(v bool) *GetAppSandboxPreviewUrlRequest
	GetRestart() *bool
}

type GetAppSandboxPreviewUrlRequest struct {
	// example:
	//
	// 593fe1a2-d0b4-4fde-a2b0-78ad6a438d41
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s GetAppSandboxPreviewUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppSandboxPreviewUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAppSandboxPreviewUrlRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetAppSandboxPreviewUrlRequest) GetRestart() *bool {
	return s.Restart
}

func (s *GetAppSandboxPreviewUrlRequest) SetConversationId(v string) *GetAppSandboxPreviewUrlRequest {
	s.ConversationId = &v
	return s
}

func (s *GetAppSandboxPreviewUrlRequest) SetRestart(v bool) *GetAppSandboxPreviewUrlRequest {
	s.Restart = &v
	return s
}

func (s *GetAppSandboxPreviewUrlRequest) Validate() error {
	return dara.Validate(s)
}
