// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlidingAssistantShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateAlidingAssistantShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateAlidingAssistantShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateAlidingAssistantShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateAlidingAssistantShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateAlidingAssistantShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateAlidingAssistantShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateAlidingAssistantShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateAlidingAssistantShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateAlidingAssistantShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAlidingAssistantShrinkHeaders) SetAccountContextShrink(v string) *CreateAlidingAssistantShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateAlidingAssistantShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
