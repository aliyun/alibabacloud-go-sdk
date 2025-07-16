// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeptNoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDeptNoHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetDeptNoHeadersAccountContext) *GetDeptNoHeaders
	GetAccountContext() *GetDeptNoHeadersAccountContext
}

type GetDeptNoHeaders struct {
	CommonHeaders  map[string]*string              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetDeptNoHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetDeptNoHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDeptNoHeaders) GoString() string {
	return s.String()
}

func (s *GetDeptNoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDeptNoHeaders) GetAccountContext() *GetDeptNoHeadersAccountContext {
	return s.AccountContext
}

func (s *GetDeptNoHeaders) SetCommonHeaders(v map[string]*string) *GetDeptNoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeptNoHeaders) SetAccountContext(v *GetDeptNoHeadersAccountContext) *GetDeptNoHeaders {
	s.AccountContext = v
	return s
}

func (s *GetDeptNoHeaders) Validate() error {
	return dara.Validate(s)
}

type GetDeptNoHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetDeptNoHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetDeptNoHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetDeptNoHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetDeptNoHeadersAccountContext) SetAccountId(v string) *GetDeptNoHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetDeptNoHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
