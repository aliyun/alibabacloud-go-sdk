// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeEventHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SubscribeEventHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SubscribeEventHeadersAccountContext) *SubscribeEventHeaders
	GetAccountContext() *SubscribeEventHeadersAccountContext
}

type SubscribeEventHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SubscribeEventHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SubscribeEventHeaders) String() string {
	return dara.Prettify(s)
}

func (s SubscribeEventHeaders) GoString() string {
	return s.String()
}

func (s *SubscribeEventHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SubscribeEventHeaders) GetAccountContext() *SubscribeEventHeadersAccountContext {
	return s.AccountContext
}

func (s *SubscribeEventHeaders) SetCommonHeaders(v map[string]*string) *SubscribeEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubscribeEventHeaders) SetAccountContext(v *SubscribeEventHeadersAccountContext) *SubscribeEventHeaders {
	s.AccountContext = v
	return s
}

func (s *SubscribeEventHeaders) Validate() error {
	return dara.Validate(s)
}

type SubscribeEventHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SubscribeEventHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SubscribeEventHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SubscribeEventHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SubscribeEventHeadersAccountContext) SetAccountId(v string) *SubscribeEventHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SubscribeEventHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
