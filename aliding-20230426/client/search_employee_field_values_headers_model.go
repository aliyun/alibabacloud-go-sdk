// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchEmployeeFieldValuesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SearchEmployeeFieldValuesHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SearchEmployeeFieldValuesHeadersAccountContext) *SearchEmployeeFieldValuesHeaders
	GetAccountContext() *SearchEmployeeFieldValuesHeadersAccountContext
}

type SearchEmployeeFieldValuesHeaders struct {
	CommonHeaders  map[string]*string                              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SearchEmployeeFieldValuesHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SearchEmployeeFieldValuesHeaders) String() string {
	return dara.Prettify(s)
}

func (s SearchEmployeeFieldValuesHeaders) GoString() string {
	return s.String()
}

func (s *SearchEmployeeFieldValuesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SearchEmployeeFieldValuesHeaders) GetAccountContext() *SearchEmployeeFieldValuesHeadersAccountContext {
	return s.AccountContext
}

func (s *SearchEmployeeFieldValuesHeaders) SetCommonHeaders(v map[string]*string) *SearchEmployeeFieldValuesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchEmployeeFieldValuesHeaders) SetAccountContext(v *SearchEmployeeFieldValuesHeadersAccountContext) *SearchEmployeeFieldValuesHeaders {
	s.AccountContext = v
	return s
}

func (s *SearchEmployeeFieldValuesHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchEmployeeFieldValuesHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SearchEmployeeFieldValuesHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SearchEmployeeFieldValuesHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SearchEmployeeFieldValuesHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SearchEmployeeFieldValuesHeadersAccountContext) SetAccountId(v string) *SearchEmployeeFieldValuesHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SearchEmployeeFieldValuesHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
