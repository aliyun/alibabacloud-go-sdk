// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitMultipartFileUploadHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InitMultipartFileUploadHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *InitMultipartFileUploadHeadersAccountContext) *InitMultipartFileUploadHeaders
	GetAccountContext() *InitMultipartFileUploadHeadersAccountContext
}

type InitMultipartFileUploadHeaders struct {
	CommonHeaders  map[string]*string                            `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *InitMultipartFileUploadHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s InitMultipartFileUploadHeaders) String() string {
	return dara.Prettify(s)
}

func (s InitMultipartFileUploadHeaders) GoString() string {
	return s.String()
}

func (s *InitMultipartFileUploadHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InitMultipartFileUploadHeaders) GetAccountContext() *InitMultipartFileUploadHeadersAccountContext {
	return s.AccountContext
}

func (s *InitMultipartFileUploadHeaders) SetCommonHeaders(v map[string]*string) *InitMultipartFileUploadHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InitMultipartFileUploadHeaders) SetAccountContext(v *InitMultipartFileUploadHeadersAccountContext) *InitMultipartFileUploadHeaders {
	s.AccountContext = v
	return s
}

func (s *InitMultipartFileUploadHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitMultipartFileUploadHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s InitMultipartFileUploadHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s InitMultipartFileUploadHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *InitMultipartFileUploadHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *InitMultipartFileUploadHeadersAccountContext) SetAccountId(v string) *InitMultipartFileUploadHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *InitMultipartFileUploadHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
