// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertDropDownListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsertDropDownListHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *InsertDropDownListHeadersAccountContext) *InsertDropDownListHeaders
	GetAccountContext() *InsertDropDownListHeadersAccountContext
}

type InsertDropDownListHeaders struct {
	CommonHeaders  map[string]*string                       `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *InsertDropDownListHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s InsertDropDownListHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsertDropDownListHeaders) GoString() string {
	return s.String()
}

func (s *InsertDropDownListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsertDropDownListHeaders) GetAccountContext() *InsertDropDownListHeadersAccountContext {
	return s.AccountContext
}

func (s *InsertDropDownListHeaders) SetCommonHeaders(v map[string]*string) *InsertDropDownListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertDropDownListHeaders) SetAccountContext(v *InsertDropDownListHeadersAccountContext) *InsertDropDownListHeaders {
	s.AccountContext = v
	return s
}

func (s *InsertDropDownListHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertDropDownListHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s InsertDropDownListHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s InsertDropDownListHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *InsertDropDownListHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *InsertDropDownListHeadersAccountContext) SetAccountId(v string) *InsertDropDownListHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *InsertDropDownListHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
