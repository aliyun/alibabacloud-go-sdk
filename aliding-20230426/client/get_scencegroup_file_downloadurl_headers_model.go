// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScencegroupFileDownloadurlHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetScencegroupFileDownloadurlHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetScencegroupFileDownloadurlHeadersAccountContext) *GetScencegroupFileDownloadurlHeaders
	GetAccountContext() *GetScencegroupFileDownloadurlHeadersAccountContext
}

type GetScencegroupFileDownloadurlHeaders struct {
	CommonHeaders  map[string]*string                                  `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetScencegroupFileDownloadurlHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetScencegroupFileDownloadurlHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetScencegroupFileDownloadurlHeaders) GoString() string {
	return s.String()
}

func (s *GetScencegroupFileDownloadurlHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetScencegroupFileDownloadurlHeaders) GetAccountContext() *GetScencegroupFileDownloadurlHeadersAccountContext {
	return s.AccountContext
}

func (s *GetScencegroupFileDownloadurlHeaders) SetCommonHeaders(v map[string]*string) *GetScencegroupFileDownloadurlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetScencegroupFileDownloadurlHeaders) SetAccountContext(v *GetScencegroupFileDownloadurlHeadersAccountContext) *GetScencegroupFileDownloadurlHeaders {
	s.AccountContext = v
	return s
}

func (s *GetScencegroupFileDownloadurlHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScencegroupFileDownloadurlHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetScencegroupFileDownloadurlHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetScencegroupFileDownloadurlHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetScencegroupFileDownloadurlHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetScencegroupFileDownloadurlHeadersAccountContext) SetAccountId(v string) *GetScencegroupFileDownloadurlHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetScencegroupFileDownloadurlHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
