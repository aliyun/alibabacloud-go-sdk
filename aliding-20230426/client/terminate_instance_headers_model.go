// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateInstanceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TerminateInstanceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *TerminateInstanceHeadersAccountContext) *TerminateInstanceHeaders
	GetAccountContext() *TerminateInstanceHeadersAccountContext
}

type TerminateInstanceHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *TerminateInstanceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s TerminateInstanceHeaders) String() string {
	return dara.Prettify(s)
}

func (s TerminateInstanceHeaders) GoString() string {
	return s.String()
}

func (s *TerminateInstanceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TerminateInstanceHeaders) GetAccountContext() *TerminateInstanceHeadersAccountContext {
	return s.AccountContext
}

func (s *TerminateInstanceHeaders) SetCommonHeaders(v map[string]*string) *TerminateInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TerminateInstanceHeaders) SetAccountContext(v *TerminateInstanceHeadersAccountContext) *TerminateInstanceHeaders {
	s.AccountContext = v
	return s
}

func (s *TerminateInstanceHeaders) Validate() error {
	return dara.Validate(s)
}

type TerminateInstanceHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s TerminateInstanceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s TerminateInstanceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *TerminateInstanceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *TerminateInstanceHeadersAccountContext) SetAccountId(v string) *TerminateInstanceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *TerminateInstanceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
