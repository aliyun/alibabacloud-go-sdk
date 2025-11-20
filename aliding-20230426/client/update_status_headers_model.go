// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStatusHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateStatusHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateStatusHeadersAccountContext) *UpdateStatusHeaders
	GetAccountContext() *UpdateStatusHeadersAccountContext
}

type UpdateStatusHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateStatusHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateStatusHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateStatusHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateStatusHeaders) GetAccountContext() *UpdateStatusHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateStatusHeaders) SetAccountContext(v *UpdateStatusHeadersAccountContext) *UpdateStatusHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateStatusHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStatusHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateStatusHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateStatusHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateStatusHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateStatusHeadersAccountContext) SetAccountId(v string) *UpdateStatusHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateStatusHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
