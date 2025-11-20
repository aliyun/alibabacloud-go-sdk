// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotifyMeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetNotifyMeHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetNotifyMeHeadersAccountContext) *GetNotifyMeHeaders
	GetAccountContext() *GetNotifyMeHeadersAccountContext
}

type GetNotifyMeHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetNotifyMeHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetNotifyMeHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetNotifyMeHeaders) GoString() string {
	return s.String()
}

func (s *GetNotifyMeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetNotifyMeHeaders) GetAccountContext() *GetNotifyMeHeadersAccountContext {
	return s.AccountContext
}

func (s *GetNotifyMeHeaders) SetCommonHeaders(v map[string]*string) *GetNotifyMeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNotifyMeHeaders) SetAccountContext(v *GetNotifyMeHeadersAccountContext) *GetNotifyMeHeaders {
	s.AccountContext = v
	return s
}

func (s *GetNotifyMeHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNotifyMeHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetNotifyMeHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetNotifyMeHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetNotifyMeHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetNotifyMeHeadersAccountContext) SetAccountId(v string) *GetNotifyMeHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetNotifyMeHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
