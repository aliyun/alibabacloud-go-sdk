// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultipartFileUploadInfosHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetMultipartFileUploadInfosHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetMultipartFileUploadInfosHeadersAccountContext) *GetMultipartFileUploadInfosHeaders
	GetAccountContext() *GetMultipartFileUploadInfosHeadersAccountContext
}

type GetMultipartFileUploadInfosHeaders struct {
	CommonHeaders  map[string]*string                                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetMultipartFileUploadInfosHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetMultipartFileUploadInfosHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetMultipartFileUploadInfosHeaders) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetMultipartFileUploadInfosHeaders) GetAccountContext() *GetMultipartFileUploadInfosHeadersAccountContext {
	return s.AccountContext
}

func (s *GetMultipartFileUploadInfosHeaders) SetCommonHeaders(v map[string]*string) *GetMultipartFileUploadInfosHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMultipartFileUploadInfosHeaders) SetAccountContext(v *GetMultipartFileUploadInfosHeadersAccountContext) *GetMultipartFileUploadInfosHeaders {
	s.AccountContext = v
	return s
}

func (s *GetMultipartFileUploadInfosHeaders) Validate() error {
	return dara.Validate(s)
}

type GetMultipartFileUploadInfosHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetMultipartFileUploadInfosHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetMultipartFileUploadInfosHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetMultipartFileUploadInfosHeadersAccountContext) SetAccountId(v string) *GetMultipartFileUploadInfosHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetMultipartFileUploadInfosHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
