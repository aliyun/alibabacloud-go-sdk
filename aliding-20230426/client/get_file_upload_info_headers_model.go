// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileUploadInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetFileUploadInfoHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetFileUploadInfoHeadersAccountContext) *GetFileUploadInfoHeaders
	GetAccountContext() *GetFileUploadInfoHeadersAccountContext
}

type GetFileUploadInfoHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetFileUploadInfoHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetFileUploadInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetFileUploadInfoHeaders) GetAccountContext() *GetFileUploadInfoHeadersAccountContext {
	return s.AccountContext
}

func (s *GetFileUploadInfoHeaders) SetCommonHeaders(v map[string]*string) *GetFileUploadInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFileUploadInfoHeaders) SetAccountContext(v *GetFileUploadInfoHeadersAccountContext) *GetFileUploadInfoHeaders {
	s.AccountContext = v
	return s
}

func (s *GetFileUploadInfoHeaders) Validate() error {
	return dara.Validate(s)
}

type GetFileUploadInfoHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetFileUploadInfoHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadInfoHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetFileUploadInfoHeadersAccountContext) SetAccountId(v string) *GetFileUploadInfoHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetFileUploadInfoHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
