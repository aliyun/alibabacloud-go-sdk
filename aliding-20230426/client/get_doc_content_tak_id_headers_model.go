// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocContentTakIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDocContentTakIdHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetDocContentTakIdHeadersAccountContext) *GetDocContentTakIdHeaders
	GetAccountContext() *GetDocContentTakIdHeadersAccountContext
}

type GetDocContentTakIdHeaders struct {
	CommonHeaders  map[string]*string                       `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetDocContentTakIdHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetDocContentTakIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDocContentTakIdHeaders) GoString() string {
	return s.String()
}

func (s *GetDocContentTakIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDocContentTakIdHeaders) GetAccountContext() *GetDocContentTakIdHeadersAccountContext {
	return s.AccountContext
}

func (s *GetDocContentTakIdHeaders) SetCommonHeaders(v map[string]*string) *GetDocContentTakIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDocContentTakIdHeaders) SetAccountContext(v *GetDocContentTakIdHeadersAccountContext) *GetDocContentTakIdHeaders {
	s.AccountContext = v
	return s
}

func (s *GetDocContentTakIdHeaders) Validate() error {
	return dara.Validate(s)
}

type GetDocContentTakIdHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetDocContentTakIdHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetDocContentTakIdHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetDocContentTakIdHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetDocContentTakIdHeadersAccountContext) SetAccountId(v string) *GetDocContentTakIdHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetDocContentTakIdHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
