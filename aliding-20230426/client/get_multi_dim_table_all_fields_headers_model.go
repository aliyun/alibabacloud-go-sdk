// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableAllFieldsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetMultiDimTableAllFieldsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetMultiDimTableAllFieldsHeadersAccountContext) *GetMultiDimTableAllFieldsHeaders
	GetAccountContext() *GetMultiDimTableAllFieldsHeadersAccountContext
}

type GetMultiDimTableAllFieldsHeaders struct {
	CommonHeaders  map[string]*string                              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetMultiDimTableAllFieldsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetMultiDimTableAllFieldsHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllFieldsHeaders) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllFieldsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetMultiDimTableAllFieldsHeaders) GetAccountContext() *GetMultiDimTableAllFieldsHeadersAccountContext {
	return s.AccountContext
}

func (s *GetMultiDimTableAllFieldsHeaders) SetCommonHeaders(v map[string]*string) *GetMultiDimTableAllFieldsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMultiDimTableAllFieldsHeaders) SetAccountContext(v *GetMultiDimTableAllFieldsHeadersAccountContext) *GetMultiDimTableAllFieldsHeaders {
	s.AccountContext = v
	return s
}

func (s *GetMultiDimTableAllFieldsHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMultiDimTableAllFieldsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetMultiDimTableAllFieldsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllFieldsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllFieldsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetMultiDimTableAllFieldsHeadersAccountContext) SetAccountId(v string) *GetMultiDimTableAllFieldsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetMultiDimTableAllFieldsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
