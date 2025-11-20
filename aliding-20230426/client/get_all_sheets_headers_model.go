// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllSheetsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAllSheetsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetAllSheetsHeadersAccountContext) *GetAllSheetsHeaders
	GetAccountContext() *GetAllSheetsHeadersAccountContext
}

type GetAllSheetsHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetAllSheetsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetAllSheetsHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAllSheetsHeaders) GoString() string {
	return s.String()
}

func (s *GetAllSheetsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAllSheetsHeaders) GetAccountContext() *GetAllSheetsHeadersAccountContext {
	return s.AccountContext
}

func (s *GetAllSheetsHeaders) SetCommonHeaders(v map[string]*string) *GetAllSheetsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAllSheetsHeaders) SetAccountContext(v *GetAllSheetsHeadersAccountContext) *GetAllSheetsHeaders {
	s.AccountContext = v
	return s
}

func (s *GetAllSheetsHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAllSheetsHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetAllSheetsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetAllSheetsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetAllSheetsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetAllSheetsHeadersAccountContext) SetAccountId(v string) *GetAllSheetsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetAllSheetsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
