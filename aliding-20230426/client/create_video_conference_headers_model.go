// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoConferenceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateVideoConferenceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateVideoConferenceHeadersAccountContext) *CreateVideoConferenceHeaders
	GetAccountContext() *CreateVideoConferenceHeadersAccountContext
}

type CreateVideoConferenceHeaders struct {
	CommonHeaders  map[string]*string                          `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateVideoConferenceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateVideoConferenceHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoConferenceHeaders) GoString() string {
	return s.String()
}

func (s *CreateVideoConferenceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateVideoConferenceHeaders) GetAccountContext() *CreateVideoConferenceHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateVideoConferenceHeaders) SetCommonHeaders(v map[string]*string) *CreateVideoConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateVideoConferenceHeaders) SetAccountContext(v *CreateVideoConferenceHeadersAccountContext) *CreateVideoConferenceHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateVideoConferenceHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateVideoConferenceHeadersAccountContext struct {
	// This parameter is required.
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateVideoConferenceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoConferenceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateVideoConferenceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateVideoConferenceHeadersAccountContext) SetAccountId(v string) *CreateVideoConferenceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateVideoConferenceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
