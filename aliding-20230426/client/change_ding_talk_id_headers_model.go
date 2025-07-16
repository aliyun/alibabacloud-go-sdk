// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDingTalkIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ChangeDingTalkIdHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ChangeDingTalkIdHeadersAccountContext) *ChangeDingTalkIdHeaders
	GetAccountContext() *ChangeDingTalkIdHeadersAccountContext
}

type ChangeDingTalkIdHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ChangeDingTalkIdHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ChangeDingTalkIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s ChangeDingTalkIdHeaders) GoString() string {
	return s.String()
}

func (s *ChangeDingTalkIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ChangeDingTalkIdHeaders) GetAccountContext() *ChangeDingTalkIdHeadersAccountContext {
	return s.AccountContext
}

func (s *ChangeDingTalkIdHeaders) SetCommonHeaders(v map[string]*string) *ChangeDingTalkIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChangeDingTalkIdHeaders) SetAccountContext(v *ChangeDingTalkIdHeadersAccountContext) *ChangeDingTalkIdHeaders {
	s.AccountContext = v
	return s
}

func (s *ChangeDingTalkIdHeaders) Validate() error {
	return dara.Validate(s)
}

type ChangeDingTalkIdHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ChangeDingTalkIdHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ChangeDingTalkIdHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ChangeDingTalkIdHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ChangeDingTalkIdHeadersAccountContext) SetAccountId(v string) *ChangeDingTalkIdHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ChangeDingTalkIdHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
