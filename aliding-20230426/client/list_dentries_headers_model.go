// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDentriesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListDentriesHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListDentriesHeadersAccountContext) *ListDentriesHeaders
	GetAccountContext() *ListDentriesHeadersAccountContext
}

type ListDentriesHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListDentriesHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListDentriesHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDentriesHeaders) GoString() string {
	return s.String()
}

func (s *ListDentriesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListDentriesHeaders) GetAccountContext() *ListDentriesHeadersAccountContext {
	return s.AccountContext
}

func (s *ListDentriesHeaders) SetCommonHeaders(v map[string]*string) *ListDentriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDentriesHeaders) SetAccountContext(v *ListDentriesHeadersAccountContext) *ListDentriesHeaders {
	s.AccountContext = v
	return s
}

func (s *ListDentriesHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDentriesHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListDentriesHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListDentriesHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListDentriesHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListDentriesHeadersAccountContext) SetAccountId(v string) *ListDentriesHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListDentriesHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
