// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIStaffPreviewUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *GetAIStaffPreviewUrlRequest
	GetConversationId() *string
	SetRestart(v bool) *GetAIStaffPreviewUrlRequest
	GetRestart() *bool
}

type GetAIStaffPreviewUrlRequest struct {
	// example:
	//
	// 81bc5a34-1d8d-4ef7-a208-7401c51b054b
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s GetAIStaffPreviewUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAIStaffPreviewUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAIStaffPreviewUrlRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetAIStaffPreviewUrlRequest) GetRestart() *bool {
	return s.Restart
}

func (s *GetAIStaffPreviewUrlRequest) SetConversationId(v string) *GetAIStaffPreviewUrlRequest {
	s.ConversationId = &v
	return s
}

func (s *GetAIStaffPreviewUrlRequest) SetRestart(v bool) *GetAIStaffPreviewUrlRequest {
	s.Restart = &v
	return s
}

func (s *GetAIStaffPreviewUrlRequest) Validate() error {
	return dara.Validate(s)
}
