// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDentryByNodeIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CopyDentryByNodeIdHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CopyDentryByNodeIdHeadersAccountContext) *CopyDentryByNodeIdHeaders
	GetAccountContext() *CopyDentryByNodeIdHeadersAccountContext
}

type CopyDentryByNodeIdHeaders struct {
	CommonHeaders  map[string]*string                       `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CopyDentryByNodeIdHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CopyDentryByNodeIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryByNodeIdHeaders) GoString() string {
	return s.String()
}

func (s *CopyDentryByNodeIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CopyDentryByNodeIdHeaders) GetAccountContext() *CopyDentryByNodeIdHeadersAccountContext {
	return s.AccountContext
}

func (s *CopyDentryByNodeIdHeaders) SetCommonHeaders(v map[string]*string) *CopyDentryByNodeIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CopyDentryByNodeIdHeaders) SetAccountContext(v *CopyDentryByNodeIdHeadersAccountContext) *CopyDentryByNodeIdHeaders {
	s.AccountContext = v
	return s
}

func (s *CopyDentryByNodeIdHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CopyDentryByNodeIdHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CopyDentryByNodeIdHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryByNodeIdHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CopyDentryByNodeIdHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CopyDentryByNodeIdHeadersAccountContext) SetAccountId(v string) *CopyDentryByNodeIdHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CopyDentryByNodeIdHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
