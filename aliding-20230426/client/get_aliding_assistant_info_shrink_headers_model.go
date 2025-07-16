// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlidingAssistantInfoShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAlidingAssistantInfoShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetAlidingAssistantInfoShrinkHeaders
	GetAccountContextShrink() *string
}

type GetAlidingAssistantInfoShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetAlidingAssistantInfoShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAlidingAssistantInfoShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetAlidingAssistantInfoShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAlidingAssistantInfoShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetAlidingAssistantInfoShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetAlidingAssistantInfoShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAlidingAssistantInfoShrinkHeaders) SetAccountContextShrink(v string) *GetAlidingAssistantInfoShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetAlidingAssistantInfoShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
