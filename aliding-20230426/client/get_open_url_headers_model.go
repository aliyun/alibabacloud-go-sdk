// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpenUrlHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetOpenUrlHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetOpenUrlHeadersAccountContext) *GetOpenUrlHeaders
	GetAccountContext() *GetOpenUrlHeadersAccountContext
}

type GetOpenUrlHeaders struct {
	CommonHeaders  map[string]*string               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetOpenUrlHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetOpenUrlHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetOpenUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetOpenUrlHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetOpenUrlHeaders) GetAccountContext() *GetOpenUrlHeadersAccountContext {
	return s.AccountContext
}

func (s *GetOpenUrlHeaders) SetCommonHeaders(v map[string]*string) *GetOpenUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOpenUrlHeaders) SetAccountContext(v *GetOpenUrlHeadersAccountContext) *GetOpenUrlHeaders {
	s.AccountContext = v
	return s
}

func (s *GetOpenUrlHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOpenUrlHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetOpenUrlHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetOpenUrlHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetOpenUrlHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetOpenUrlHeadersAccountContext) SetAccountId(v string) *GetOpenUrlHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetOpenUrlHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
