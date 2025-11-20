// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserIdHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetUserIdHeadersAccountContext) *GetUserIdHeaders
	GetAccountContext() *GetUserIdHeadersAccountContext
}

type GetUserIdHeaders struct {
	CommonHeaders  map[string]*string              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetUserIdHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetUserIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdHeaders) GoString() string {
	return s.String()
}

func (s *GetUserIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserIdHeaders) GetAccountContext() *GetUserIdHeadersAccountContext {
	return s.AccountContext
}

func (s *GetUserIdHeaders) SetCommonHeaders(v map[string]*string) *GetUserIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserIdHeaders) SetAccountContext(v *GetUserIdHeadersAccountContext) *GetUserIdHeaders {
	s.AccountContext = v
	return s
}

func (s *GetUserIdHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserIdHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetUserIdHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetUserIdHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetUserIdHeadersAccountContext) SetAccountId(v string) *GetUserIdHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetUserIdHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
