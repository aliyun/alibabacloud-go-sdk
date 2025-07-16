// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMinutesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StartMinutesHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *StartMinutesHeadersAccountContext) *StartMinutesHeaders
	GetAccountContext() *StartMinutesHeadersAccountContext
}

type StartMinutesHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *StartMinutesHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s StartMinutesHeaders) String() string {
	return dara.Prettify(s)
}

func (s StartMinutesHeaders) GoString() string {
	return s.String()
}

func (s *StartMinutesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StartMinutesHeaders) GetAccountContext() *StartMinutesHeadersAccountContext {
	return s.AccountContext
}

func (s *StartMinutesHeaders) SetCommonHeaders(v map[string]*string) *StartMinutesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartMinutesHeaders) SetAccountContext(v *StartMinutesHeadersAccountContext) *StartMinutesHeaders {
	s.AccountContext = v
	return s
}

func (s *StartMinutesHeaders) Validate() error {
	return dara.Validate(s)
}

type StartMinutesHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s StartMinutesHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s StartMinutesHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *StartMinutesHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *StartMinutesHeadersAccountContext) SetAccountId(v string) *StartMinutesHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *StartMinutesHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
