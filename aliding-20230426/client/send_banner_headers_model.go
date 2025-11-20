// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendBannerHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SendBannerHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SendBannerHeadersAccountContext) *SendBannerHeaders
	GetAccountContext() *SendBannerHeadersAccountContext
}

type SendBannerHeaders struct {
	CommonHeaders  map[string]*string               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SendBannerHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SendBannerHeaders) String() string {
	return dara.Prettify(s)
}

func (s SendBannerHeaders) GoString() string {
	return s.String()
}

func (s *SendBannerHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SendBannerHeaders) GetAccountContext() *SendBannerHeadersAccountContext {
	return s.AccountContext
}

func (s *SendBannerHeaders) SetCommonHeaders(v map[string]*string) *SendBannerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendBannerHeaders) SetAccountContext(v *SendBannerHeadersAccountContext) *SendBannerHeaders {
	s.AccountContext = v
	return s
}

func (s *SendBannerHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendBannerHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SendBannerHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SendBannerHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SendBannerHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SendBannerHeadersAccountContext) SetAccountId(v string) *SendBannerHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SendBannerHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
