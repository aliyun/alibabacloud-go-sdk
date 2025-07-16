// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoConferenceSettingHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateVideoConferenceSettingHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateVideoConferenceSettingHeadersAccountContext) *UpdateVideoConferenceSettingHeaders
	GetAccountContext() *UpdateVideoConferenceSettingHeadersAccountContext
}

type UpdateVideoConferenceSettingHeaders struct {
	CommonHeaders  map[string]*string                                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateVideoConferenceSettingHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateVideoConferenceSettingHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoConferenceSettingHeaders) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceSettingHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateVideoConferenceSettingHeaders) GetAccountContext() *UpdateVideoConferenceSettingHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateVideoConferenceSettingHeaders) SetCommonHeaders(v map[string]*string) *UpdateVideoConferenceSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateVideoConferenceSettingHeaders) SetAccountContext(v *UpdateVideoConferenceSettingHeadersAccountContext) *UpdateVideoConferenceSettingHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateVideoConferenceSettingHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateVideoConferenceSettingHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateVideoConferenceSettingHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoConferenceSettingHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceSettingHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateVideoConferenceSettingHeadersAccountContext) SetAccountId(v string) *UpdateVideoConferenceSettingHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateVideoConferenceSettingHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
