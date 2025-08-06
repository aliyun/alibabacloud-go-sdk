// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIntentionNoteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *SubmitIntentionNoteRequest
	GetBizType() *string
	SetIntentionBizId(v string) *SubmitIntentionNoteRequest
	GetIntentionBizId() *string
	SetNote(v string) *SubmitIntentionNoteRequest
	GetNote() *string
}

type SubmitIntentionNoteRequest struct {
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// I20210927144823000001
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	// This parameter is required.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s SubmitIntentionNoteRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitIntentionNoteRequest) GoString() string {
	return s.String()
}

func (s *SubmitIntentionNoteRequest) GetBizType() *string {
	return s.BizType
}

func (s *SubmitIntentionNoteRequest) GetIntentionBizId() *string {
	return s.IntentionBizId
}

func (s *SubmitIntentionNoteRequest) GetNote() *string {
	return s.Note
}

func (s *SubmitIntentionNoteRequest) SetBizType(v string) *SubmitIntentionNoteRequest {
	s.BizType = &v
	return s
}

func (s *SubmitIntentionNoteRequest) SetIntentionBizId(v string) *SubmitIntentionNoteRequest {
	s.IntentionBizId = &v
	return s
}

func (s *SubmitIntentionNoteRequest) SetNote(v string) *SubmitIntentionNoteRequest {
	s.Note = &v
	return s
}

func (s *SubmitIntentionNoteRequest) Validate() error {
	return dara.Validate(s)
}
