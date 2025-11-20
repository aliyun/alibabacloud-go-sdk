// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConvExtensionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateConvExtensionHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateConvExtensionHeadersAccountContext) *UpdateConvExtensionHeaders
	GetAccountContext() *UpdateConvExtensionHeadersAccountContext
}

type UpdateConvExtensionHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateConvExtensionHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateConvExtensionHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateConvExtensionHeaders) GoString() string {
	return s.String()
}

func (s *UpdateConvExtensionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateConvExtensionHeaders) GetAccountContext() *UpdateConvExtensionHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateConvExtensionHeaders) SetCommonHeaders(v map[string]*string) *UpdateConvExtensionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateConvExtensionHeaders) SetAccountContext(v *UpdateConvExtensionHeadersAccountContext) *UpdateConvExtensionHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateConvExtensionHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateConvExtensionHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateConvExtensionHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateConvExtensionHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateConvExtensionHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateConvExtensionHeadersAccountContext) SetAccountId(v string) *UpdateConvExtensionHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateConvExtensionHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
