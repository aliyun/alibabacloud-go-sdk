// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateLiveHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateLiveHeadersAccountContext) *UpdateLiveHeaders
	GetAccountContext() *UpdateLiveHeadersAccountContext
}

type UpdateLiveHeaders struct {
	CommonHeaders  map[string]*string               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateLiveHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateLiveHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveHeaders) GoString() string {
	return s.String()
}

func (s *UpdateLiveHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateLiveHeaders) GetAccountContext() *UpdateLiveHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateLiveHeaders) SetCommonHeaders(v map[string]*string) *UpdateLiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateLiveHeaders) SetAccountContext(v *UpdateLiveHeadersAccountContext) *UpdateLiveHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateLiveHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateLiveHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateLiveHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateLiveHeadersAccountContext) SetAccountId(v string) *UpdateLiveHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateLiveHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
