// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFormDataHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteFormDataHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteFormDataHeadersAccountContext) *DeleteFormDataHeaders
	GetAccountContext() *DeleteFormDataHeadersAccountContext
}

type DeleteFormDataHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteFormDataHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteFormDataHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteFormDataHeaders) GoString() string {
	return s.String()
}

func (s *DeleteFormDataHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteFormDataHeaders) GetAccountContext() *DeleteFormDataHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteFormDataHeaders) SetCommonHeaders(v map[string]*string) *DeleteFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteFormDataHeaders) SetAccountContext(v *DeleteFormDataHeadersAccountContext) *DeleteFormDataHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteFormDataHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteFormDataHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteFormDataHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteFormDataHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteFormDataHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteFormDataHeadersAccountContext) SetAccountId(v string) *DeleteFormDataHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteFormDataHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
