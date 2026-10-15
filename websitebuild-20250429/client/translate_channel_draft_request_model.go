// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateChannelDraftRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDraftId(v string) *TranslateChannelDraftRequest
	GetDraftId() *string
}

type TranslateChannelDraftRequest struct {
	// The ID of the channel draft.
	//
	// This parameter is required.
	//
	// example:
	//
	// draft-789
	DraftId *string `json:"DraftId,omitempty" xml:"DraftId,omitempty"`
}

func (s TranslateChannelDraftRequest) String() string {
	return dara.Prettify(s)
}

func (s TranslateChannelDraftRequest) GoString() string {
	return s.String()
}

func (s *TranslateChannelDraftRequest) GetDraftId() *string {
	return s.DraftId
}

func (s *TranslateChannelDraftRequest) SetDraftId(v string) *TranslateChannelDraftRequest {
	s.DraftId = &v
	return s
}

func (s *TranslateChannelDraftRequest) Validate() error {
	return dara.Validate(s)
}
