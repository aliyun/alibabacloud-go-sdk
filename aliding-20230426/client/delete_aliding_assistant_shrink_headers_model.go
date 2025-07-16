// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlidingAssistantShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteAlidingAssistantShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteAlidingAssistantShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteAlidingAssistantShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteAlidingAssistantShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlidingAssistantShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteAlidingAssistantShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteAlidingAssistantShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteAlidingAssistantShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteAlidingAssistantShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteAlidingAssistantShrinkHeaders) SetAccountContextShrink(v string) *DeleteAlidingAssistantShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteAlidingAssistantShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
