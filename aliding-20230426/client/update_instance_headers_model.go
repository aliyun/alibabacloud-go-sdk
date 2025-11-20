// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateInstanceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateInstanceHeadersAccountContext) *UpdateInstanceHeaders
	GetAccountContext() *UpdateInstanceHeadersAccountContext
}

type UpdateInstanceHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateInstanceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateInstanceHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInstanceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateInstanceHeaders) GetAccountContext() *UpdateInstanceHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateInstanceHeaders) SetCommonHeaders(v map[string]*string) *UpdateInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInstanceHeaders) SetAccountContext(v *UpdateInstanceHeadersAccountContext) *UpdateInstanceHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateInstanceHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateInstanceHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateInstanceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateInstanceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateInstanceHeadersAccountContext) SetAccountId(v string) *UpdateInstanceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateInstanceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
