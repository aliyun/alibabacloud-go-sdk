// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileDownloadInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetFileDownloadInfoHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetFileDownloadInfoHeadersAccountContext) *GetFileDownloadInfoHeaders
	GetAccountContext() *GetFileDownloadInfoHeadersAccountContext
}

type GetFileDownloadInfoHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetFileDownloadInfoHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetFileDownloadInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetFileDownloadInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetFileDownloadInfoHeaders) GetAccountContext() *GetFileDownloadInfoHeadersAccountContext {
	return s.AccountContext
}

func (s *GetFileDownloadInfoHeaders) SetCommonHeaders(v map[string]*string) *GetFileDownloadInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFileDownloadInfoHeaders) SetAccountContext(v *GetFileDownloadInfoHeadersAccountContext) *GetFileDownloadInfoHeaders {
	s.AccountContext = v
	return s
}

func (s *GetFileDownloadInfoHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileDownloadInfoHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetFileDownloadInfoHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetFileDownloadInfoHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetFileDownloadInfoHeadersAccountContext) SetAccountId(v string) *GetFileDownloadInfoHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetFileDownloadInfoHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
