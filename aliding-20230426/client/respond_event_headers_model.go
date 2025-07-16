// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRespondEventHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RespondEventHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *RespondEventHeadersAccountContext) *RespondEventHeaders
	GetAccountContext() *RespondEventHeadersAccountContext
}

type RespondEventHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *RespondEventHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s RespondEventHeaders) String() string {
	return dara.Prettify(s)
}

func (s RespondEventHeaders) GoString() string {
	return s.String()
}

func (s *RespondEventHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RespondEventHeaders) GetAccountContext() *RespondEventHeadersAccountContext {
	return s.AccountContext
}

func (s *RespondEventHeaders) SetCommonHeaders(v map[string]*string) *RespondEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RespondEventHeaders) SetAccountContext(v *RespondEventHeadersAccountContext) *RespondEventHeaders {
	s.AccountContext = v
	return s
}

func (s *RespondEventHeaders) Validate() error {
	return dara.Validate(s)
}

type RespondEventHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s RespondEventHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s RespondEventHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *RespondEventHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *RespondEventHeadersAccountContext) SetAccountId(v string) *RespondEventHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *RespondEventHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
