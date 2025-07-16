// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeEventHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UnsubscribeEventHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UnsubscribeEventHeadersAccountContext) *UnsubscribeEventHeaders
	GetAccountContext() *UnsubscribeEventHeadersAccountContext
}

type UnsubscribeEventHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UnsubscribeEventHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UnsubscribeEventHeaders) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeEventHeaders) GoString() string {
	return s.String()
}

func (s *UnsubscribeEventHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UnsubscribeEventHeaders) GetAccountContext() *UnsubscribeEventHeadersAccountContext {
	return s.AccountContext
}

func (s *UnsubscribeEventHeaders) SetCommonHeaders(v map[string]*string) *UnsubscribeEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnsubscribeEventHeaders) SetAccountContext(v *UnsubscribeEventHeadersAccountContext) *UnsubscribeEventHeaders {
	s.AccountContext = v
	return s
}

func (s *UnsubscribeEventHeaders) Validate() error {
	return dara.Validate(s)
}

type UnsubscribeEventHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UnsubscribeEventHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeEventHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UnsubscribeEventHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UnsubscribeEventHeadersAccountContext) SetAccountId(v string) *UnsubscribeEventHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UnsubscribeEventHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
