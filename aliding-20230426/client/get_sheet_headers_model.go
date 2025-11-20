// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSheetHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetSheetHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetSheetHeadersAccountContext) *GetSheetHeaders
	GetAccountContext() *GetSheetHeadersAccountContext
}

type GetSheetHeaders struct {
	CommonHeaders  map[string]*string             `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetSheetHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetSheetHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetSheetHeaders) GoString() string {
	return s.String()
}

func (s *GetSheetHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetSheetHeaders) GetAccountContext() *GetSheetHeadersAccountContext {
	return s.AccountContext
}

func (s *GetSheetHeaders) SetCommonHeaders(v map[string]*string) *GetSheetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSheetHeaders) SetAccountContext(v *GetSheetHeadersAccountContext) *GetSheetHeaders {
	s.AccountContext = v
	return s
}

func (s *GetSheetHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSheetHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetSheetHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetSheetHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetSheetHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetSheetHeadersAccountContext) SetAccountId(v string) *GetSheetHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetSheetHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
