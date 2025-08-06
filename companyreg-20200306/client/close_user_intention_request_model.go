// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseUserIntentionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *CloseUserIntentionRequest
	GetBizType() *string
	SetIntentionBizId(v string) *CloseUserIntentionRequest
	GetIntentionBizId() *string
	SetNote(v string) *CloseUserIntentionRequest
	GetNote() *string
}

type CloseUserIntentionRequest struct {
	// example:
	//
	// esp.bookkeeping
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// I20201027162033000001
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	// This parameter is required.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s CloseUserIntentionRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseUserIntentionRequest) GoString() string {
	return s.String()
}

func (s *CloseUserIntentionRequest) GetBizType() *string {
	return s.BizType
}

func (s *CloseUserIntentionRequest) GetIntentionBizId() *string {
	return s.IntentionBizId
}

func (s *CloseUserIntentionRequest) GetNote() *string {
	return s.Note
}

func (s *CloseUserIntentionRequest) SetBizType(v string) *CloseUserIntentionRequest {
	s.BizType = &v
	return s
}

func (s *CloseUserIntentionRequest) SetIntentionBizId(v string) *CloseUserIntentionRequest {
	s.IntentionBizId = &v
	return s
}

func (s *CloseUserIntentionRequest) SetNote(v string) *CloseUserIntentionRequest {
	s.Note = &v
	return s
}

func (s *CloseUserIntentionRequest) Validate() error {
	return dara.Validate(s)
}
