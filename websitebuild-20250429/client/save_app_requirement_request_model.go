// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAppRequirementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *SaveAppRequirementRequest
	GetConversationId() *string
	SetPrd(v string) *SaveAppRequirementRequest
	GetPrd() *string
}

type SaveAppRequirementRequest struct {
	// Session ID
	//
	// example:
	//
	// 5b7105a2-2999-430b-ba23-ba09149d5434
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// Requirement document content
	//
	// example:
	//
	// prd
	Prd *string `json:"Prd,omitempty" xml:"Prd,omitempty"`
}

func (s SaveAppRequirementRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveAppRequirementRequest) GoString() string {
	return s.String()
}

func (s *SaveAppRequirementRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *SaveAppRequirementRequest) GetPrd() *string {
	return s.Prd
}

func (s *SaveAppRequirementRequest) SetConversationId(v string) *SaveAppRequirementRequest {
	s.ConversationId = &v
	return s
}

func (s *SaveAppRequirementRequest) SetPrd(v string) *SaveAppRequirementRequest {
	s.Prd = &v
	return s
}

func (s *SaveAppRequirementRequest) Validate() error {
	return dara.Validate(s)
}
