// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormDataByIDHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetFormDataByIDHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetFormDataByIDHeadersAccountContext) *GetFormDataByIDHeaders
	GetAccountContext() *GetFormDataByIDHeadersAccountContext
}

type GetFormDataByIDHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetFormDataByIDHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetFormDataByIDHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetFormDataByIDHeaders) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetFormDataByIDHeaders) GetAccountContext() *GetFormDataByIDHeadersAccountContext {
	return s.AccountContext
}

func (s *GetFormDataByIDHeaders) SetCommonHeaders(v map[string]*string) *GetFormDataByIDHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFormDataByIDHeaders) SetAccountContext(v *GetFormDataByIDHeadersAccountContext) *GetFormDataByIDHeaders {
	s.AccountContext = v
	return s
}

func (s *GetFormDataByIDHeaders) Validate() error {
	return dara.Validate(s)
}

type GetFormDataByIDHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetFormDataByIDHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetFormDataByIDHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetFormDataByIDHeadersAccountContext) SetAccountId(v string) *GetFormDataByIDHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetFormDataByIDHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
