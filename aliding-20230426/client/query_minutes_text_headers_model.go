// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesTextHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryMinutesTextHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryMinutesTextHeadersAccountContext) *QueryMinutesTextHeaders
	GetAccountContext() *QueryMinutesTextHeadersAccountContext
}

type QueryMinutesTextHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryMinutesTextHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryMinutesTextHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesTextHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryMinutesTextHeaders) GetAccountContext() *QueryMinutesTextHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryMinutesTextHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesTextHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesTextHeaders) SetAccountContext(v *QueryMinutesTextHeadersAccountContext) *QueryMinutesTextHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryMinutesTextHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryMinutesTextHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryMinutesTextHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesTextHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryMinutesTextHeadersAccountContext) SetAccountId(v string) *QueryMinutesTextHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryMinutesTextHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
