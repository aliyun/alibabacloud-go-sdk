// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecuritySuggestionListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListConfigRulesRequestShrink(v string) *GetSecuritySuggestionListShrinkRequest
	GetListConfigRulesRequestShrink() *string
}

type GetSecuritySuggestionListShrinkRequest struct {
	// The request parameters.
	ListConfigRulesRequestShrink *string `json:"ListConfigRulesRequest,omitempty" xml:"ListConfigRulesRequest,omitempty"`
}

func (s GetSecuritySuggestionListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySuggestionListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSecuritySuggestionListShrinkRequest) GetListConfigRulesRequestShrink() *string {
	return s.ListConfigRulesRequestShrink
}

func (s *GetSecuritySuggestionListShrinkRequest) SetListConfigRulesRequestShrink(v string) *GetSecuritySuggestionListShrinkRequest {
	s.ListConfigRulesRequestShrink = &v
	return s
}

func (s *GetSecuritySuggestionListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
