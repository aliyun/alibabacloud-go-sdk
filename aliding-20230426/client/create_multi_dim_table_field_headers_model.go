// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiDimTableFieldHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateMultiDimTableFieldHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateMultiDimTableFieldHeadersAccountContext) *CreateMultiDimTableFieldHeaders
	GetAccountContext() *CreateMultiDimTableFieldHeadersAccountContext
}

type CreateMultiDimTableFieldHeaders struct {
	CommonHeaders  map[string]*string                             `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateMultiDimTableFieldHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateMultiDimTableFieldHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiDimTableFieldHeaders) GoString() string {
	return s.String()
}

func (s *CreateMultiDimTableFieldHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateMultiDimTableFieldHeaders) GetAccountContext() *CreateMultiDimTableFieldHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateMultiDimTableFieldHeaders) SetCommonHeaders(v map[string]*string) *CreateMultiDimTableFieldHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMultiDimTableFieldHeaders) SetAccountContext(v *CreateMultiDimTableFieldHeadersAccountContext) *CreateMultiDimTableFieldHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateMultiDimTableFieldHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMultiDimTableFieldHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateMultiDimTableFieldHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiDimTableFieldHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateMultiDimTableFieldHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateMultiDimTableFieldHeadersAccountContext) SetAccountId(v string) *CreateMultiDimTableFieldHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateMultiDimTableFieldHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
