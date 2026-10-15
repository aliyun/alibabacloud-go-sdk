// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRewriteChannelTitleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDraftId(v string) *RewriteChannelTitleRequest
	GetDraftId() *string
	SetUserRequirement(v string) *RewriteChannelTitleRequest
	GetUserRequirement() *string
}

type RewriteChannelTitleRequest struct {
	// The channel draft ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7c80359e-835d-4871-a971-8400d008828b
	DraftId *string `json:"DraftId,omitempty" xml:"DraftId,omitempty"`
	// The user personalization requirement in free text. This parameter is optional.
	//
	// example:
	//
	// xxx
	UserRequirement *string `json:"UserRequirement,omitempty" xml:"UserRequirement,omitempty"`
}

func (s RewriteChannelTitleRequest) String() string {
	return dara.Prettify(s)
}

func (s RewriteChannelTitleRequest) GoString() string {
	return s.String()
}

func (s *RewriteChannelTitleRequest) GetDraftId() *string {
	return s.DraftId
}

func (s *RewriteChannelTitleRequest) GetUserRequirement() *string {
	return s.UserRequirement
}

func (s *RewriteChannelTitleRequest) SetDraftId(v string) *RewriteChannelTitleRequest {
	s.DraftId = &v
	return s
}

func (s *RewriteChannelTitleRequest) SetUserRequirement(v string) *RewriteChannelTitleRequest {
	s.UserRequirement = &v
	return s
}

func (s *RewriteChannelTitleRequest) Validate() error {
	return dara.Validate(s)
}
