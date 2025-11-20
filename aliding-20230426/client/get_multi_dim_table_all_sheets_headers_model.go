// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableAllSheetsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetMultiDimTableAllSheetsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetMultiDimTableAllSheetsHeadersAccountContext) *GetMultiDimTableAllSheetsHeaders
	GetAccountContext() *GetMultiDimTableAllSheetsHeadersAccountContext
}

type GetMultiDimTableAllSheetsHeaders struct {
	CommonHeaders  map[string]*string                              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetMultiDimTableAllSheetsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetMultiDimTableAllSheetsHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllSheetsHeaders) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllSheetsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetMultiDimTableAllSheetsHeaders) GetAccountContext() *GetMultiDimTableAllSheetsHeadersAccountContext {
	return s.AccountContext
}

func (s *GetMultiDimTableAllSheetsHeaders) SetCommonHeaders(v map[string]*string) *GetMultiDimTableAllSheetsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMultiDimTableAllSheetsHeaders) SetAccountContext(v *GetMultiDimTableAllSheetsHeadersAccountContext) *GetMultiDimTableAllSheetsHeaders {
	s.AccountContext = v
	return s
}

func (s *GetMultiDimTableAllSheetsHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMultiDimTableAllSheetsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetMultiDimTableAllSheetsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllSheetsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllSheetsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetMultiDimTableAllSheetsHeadersAccountContext) SetAccountId(v string) *GetMultiDimTableAllSheetsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetMultiDimTableAllSheetsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
