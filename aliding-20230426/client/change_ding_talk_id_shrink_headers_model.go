// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDingTalkIdShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ChangeDingTalkIdShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ChangeDingTalkIdShrinkHeaders
	GetAccountContextShrink() *string
}

type ChangeDingTalkIdShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ChangeDingTalkIdShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ChangeDingTalkIdShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ChangeDingTalkIdShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ChangeDingTalkIdShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ChangeDingTalkIdShrinkHeaders) SetCommonHeaders(v map[string]*string) *ChangeDingTalkIdShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChangeDingTalkIdShrinkHeaders) SetAccountContextShrink(v string) *ChangeDingTalkIdShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ChangeDingTalkIdShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
