// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlidingAssistantShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateAlidingAssistantShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateAlidingAssistantShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateAlidingAssistantShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateAlidingAssistantShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlidingAssistantShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateAlidingAssistantShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateAlidingAssistantShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateAlidingAssistantShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateAlidingAssistantShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateAlidingAssistantShrinkHeaders) SetAccountContextShrink(v string) *UpdateAlidingAssistantShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateAlidingAssistantShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
