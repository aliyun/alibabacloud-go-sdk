// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMeCorpSubmissionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetMeCorpSubmissionHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetMeCorpSubmissionHeadersAccountContext) *GetMeCorpSubmissionHeaders
	GetAccountContext() *GetMeCorpSubmissionHeadersAccountContext
}

type GetMeCorpSubmissionHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetMeCorpSubmissionHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetMeCorpSubmissionHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetMeCorpSubmissionHeaders) GoString() string {
	return s.String()
}

func (s *GetMeCorpSubmissionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetMeCorpSubmissionHeaders) GetAccountContext() *GetMeCorpSubmissionHeadersAccountContext {
	return s.AccountContext
}

func (s *GetMeCorpSubmissionHeaders) SetCommonHeaders(v map[string]*string) *GetMeCorpSubmissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMeCorpSubmissionHeaders) SetAccountContext(v *GetMeCorpSubmissionHeadersAccountContext) *GetMeCorpSubmissionHeaders {
	s.AccountContext = v
	return s
}

func (s *GetMeCorpSubmissionHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMeCorpSubmissionHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetMeCorpSubmissionHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetMeCorpSubmissionHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetMeCorpSubmissionHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetMeCorpSubmissionHeadersAccountContext) SetAccountId(v string) *GetMeCorpSubmissionHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetMeCorpSubmissionHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
