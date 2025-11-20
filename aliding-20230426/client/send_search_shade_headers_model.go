// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSearchShadeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SendSearchShadeHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SendSearchShadeHeadersAccountContext) *SendSearchShadeHeaders
	GetAccountContext() *SendSearchShadeHeadersAccountContext
}

type SendSearchShadeHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SendSearchShadeHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SendSearchShadeHeaders) String() string {
	return dara.Prettify(s)
}

func (s SendSearchShadeHeaders) GoString() string {
	return s.String()
}

func (s *SendSearchShadeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SendSearchShadeHeaders) GetAccountContext() *SendSearchShadeHeadersAccountContext {
	return s.AccountContext
}

func (s *SendSearchShadeHeaders) SetCommonHeaders(v map[string]*string) *SendSearchShadeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendSearchShadeHeaders) SetAccountContext(v *SendSearchShadeHeadersAccountContext) *SendSearchShadeHeaders {
	s.AccountContext = v
	return s
}

func (s *SendSearchShadeHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendSearchShadeHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SendSearchShadeHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SendSearchShadeHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SendSearchShadeHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SendSearchShadeHeadersAccountContext) SetAccountId(v string) *SendSearchShadeHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SendSearchShadeHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
