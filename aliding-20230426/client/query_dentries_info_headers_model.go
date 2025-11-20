// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDentriesInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryDentriesInfoHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryDentriesInfoHeadersAccountContext) *QueryDentriesInfoHeaders
	GetAccountContext() *QueryDentriesInfoHeadersAccountContext
}

type QueryDentriesInfoHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryDentriesInfoHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryDentriesInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryDentriesInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryDentriesInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryDentriesInfoHeaders) GetAccountContext() *QueryDentriesInfoHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryDentriesInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryDentriesInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDentriesInfoHeaders) SetAccountContext(v *QueryDentriesInfoHeadersAccountContext) *QueryDentriesInfoHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryDentriesInfoHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDentriesInfoHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryDentriesInfoHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryDentriesInfoHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryDentriesInfoHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryDentriesInfoHeadersAccountContext) SetAccountId(v string) *QueryDentriesInfoHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryDentriesInfoHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
