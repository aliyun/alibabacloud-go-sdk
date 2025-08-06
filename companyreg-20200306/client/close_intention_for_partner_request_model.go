// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseIntentionForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *CloseIntentionForPartnerRequest
	GetBizType() *string
	SetIntentionBizId(v string) *CloseIntentionForPartnerRequest
	GetIntentionBizId() *string
	SetNote(v string) *CloseIntentionForPartnerRequest
	GetNote() *string
}

type CloseIntentionForPartnerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// esp.beian_assist
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// I20211105230733000001
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	// This parameter is required.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s CloseIntentionForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseIntentionForPartnerRequest) GoString() string {
	return s.String()
}

func (s *CloseIntentionForPartnerRequest) GetBizType() *string {
	return s.BizType
}

func (s *CloseIntentionForPartnerRequest) GetIntentionBizId() *string {
	return s.IntentionBizId
}

func (s *CloseIntentionForPartnerRequest) GetNote() *string {
	return s.Note
}

func (s *CloseIntentionForPartnerRequest) SetBizType(v string) *CloseIntentionForPartnerRequest {
	s.BizType = &v
	return s
}

func (s *CloseIntentionForPartnerRequest) SetIntentionBizId(v string) *CloseIntentionForPartnerRequest {
	s.IntentionBizId = &v
	return s
}

func (s *CloseIntentionForPartnerRequest) SetNote(v string) *CloseIntentionForPartnerRequest {
	s.Note = &v
	return s
}

func (s *CloseIntentionForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
