// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendPopupHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SendPopupHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SendPopupHeadersAccountContext) *SendPopupHeaders
	GetAccountContext() *SendPopupHeadersAccountContext
}

type SendPopupHeaders struct {
	CommonHeaders  map[string]*string              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SendPopupHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SendPopupHeaders) String() string {
	return dara.Prettify(s)
}

func (s SendPopupHeaders) GoString() string {
	return s.String()
}

func (s *SendPopupHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SendPopupHeaders) GetAccountContext() *SendPopupHeadersAccountContext {
	return s.AccountContext
}

func (s *SendPopupHeaders) SetCommonHeaders(v map[string]*string) *SendPopupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendPopupHeaders) SetAccountContext(v *SendPopupHeadersAccountContext) *SendPopupHeaders {
	s.AccountContext = v
	return s
}

func (s *SendPopupHeaders) Validate() error {
	return dara.Validate(s)
}

type SendPopupHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SendPopupHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SendPopupHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SendPopupHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SendPopupHeadersAccountContext) SetAccountId(v string) *SendPopupHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SendPopupHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
