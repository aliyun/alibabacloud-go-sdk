// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableDataByFormInstanceIdTableIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListTableDataByFormInstanceIdTableIdHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListTableDataByFormInstanceIdTableIdHeadersAccountContext) *ListTableDataByFormInstanceIdTableIdHeaders
	GetAccountContext() *ListTableDataByFormInstanceIdTableIdHeadersAccountContext
}

type ListTableDataByFormInstanceIdTableIdHeaders struct {
	CommonHeaders  map[string]*string                                         `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListTableDataByFormInstanceIdTableIdHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListTableDataByFormInstanceIdTableIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListTableDataByFormInstanceIdTableIdHeaders) GoString() string {
	return s.String()
}

func (s *ListTableDataByFormInstanceIdTableIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListTableDataByFormInstanceIdTableIdHeaders) GetAccountContext() *ListTableDataByFormInstanceIdTableIdHeadersAccountContext {
	return s.AccountContext
}

func (s *ListTableDataByFormInstanceIdTableIdHeaders) SetCommonHeaders(v map[string]*string) *ListTableDataByFormInstanceIdTableIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdHeaders) SetAccountContext(v *ListTableDataByFormInstanceIdTableIdHeadersAccountContext) *ListTableDataByFormInstanceIdTableIdHeaders {
	s.AccountContext = v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTableDataByFormInstanceIdTableIdHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListTableDataByFormInstanceIdTableIdHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListTableDataByFormInstanceIdTableIdHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListTableDataByFormInstanceIdTableIdHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListTableDataByFormInstanceIdTableIdHeadersAccountContext) SetAccountId(v string) *ListTableDataByFormInstanceIdTableIdHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
