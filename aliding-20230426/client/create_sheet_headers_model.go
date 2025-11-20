// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSheetHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateSheetHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateSheetHeadersAccountContext) *CreateSheetHeaders
	GetAccountContext() *CreateSheetHeadersAccountContext
}

type CreateSheetHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateSheetHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateSheetHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateSheetHeaders) GoString() string {
	return s.String()
}

func (s *CreateSheetHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateSheetHeaders) GetAccountContext() *CreateSheetHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateSheetHeaders) SetCommonHeaders(v map[string]*string) *CreateSheetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSheetHeaders) SetAccountContext(v *CreateSheetHeadersAccountContext) *CreateSheetHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateSheetHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSheetHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateSheetHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateSheetHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateSheetHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateSheetHeadersAccountContext) SetAccountId(v string) *CreateSheetHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateSheetHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
