// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppRequirementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *GetAppRequirementRequest
	GetConversationId() *string
}

type GetAppRequirementRequest struct {
	// Session ID
	//
	// example:
	//
	// 81bc5a34-1d8d-4ef7-a208-7401c51b054b
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
}

func (s GetAppRequirementRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppRequirementRequest) GoString() string {
	return s.String()
}

func (s *GetAppRequirementRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetAppRequirementRequest) SetConversationId(v string) *GetAppRequirementRequest {
	s.ConversationId = &v
	return s
}

func (s *GetAppRequirementRequest) Validate() error {
	return dara.Validate(s)
}
