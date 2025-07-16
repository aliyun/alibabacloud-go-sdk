// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateLiveHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateLiveHeadersAccountContext) *CreateLiveHeaders
	GetAccountContext() *CreateLiveHeadersAccountContext
}

type CreateLiveHeaders struct {
	CommonHeaders  map[string]*string               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateLiveHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateLiveHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveHeaders) GoString() string {
	return s.String()
}

func (s *CreateLiveHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateLiveHeaders) GetAccountContext() *CreateLiveHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateLiveHeaders) SetCommonHeaders(v map[string]*string) *CreateLiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateLiveHeaders) SetAccountContext(v *CreateLiveHeadersAccountContext) *CreateLiveHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateLiveHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateLiveHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateLiveHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateLiveHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateLiveHeadersAccountContext) SetAccountId(v string) *CreateLiveHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateLiveHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
