// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StartInstanceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *StartInstanceHeadersAccountContext) *StartInstanceHeaders
	GetAccountContext() *StartInstanceHeadersAccountContext
}

type StartInstanceHeaders struct {
	CommonHeaders  map[string]*string                  `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *StartInstanceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s StartInstanceHeaders) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceHeaders) GoString() string {
	return s.String()
}

func (s *StartInstanceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StartInstanceHeaders) GetAccountContext() *StartInstanceHeadersAccountContext {
	return s.AccountContext
}

func (s *StartInstanceHeaders) SetCommonHeaders(v map[string]*string) *StartInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartInstanceHeaders) SetAccountContext(v *StartInstanceHeadersAccountContext) *StartInstanceHeaders {
	s.AccountContext = v
	return s
}

func (s *StartInstanceHeaders) Validate() error {
	return dara.Validate(s)
}

type StartInstanceHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s StartInstanceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *StartInstanceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *StartInstanceHeadersAccountContext) SetAccountId(v string) *StartInstanceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *StartInstanceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
