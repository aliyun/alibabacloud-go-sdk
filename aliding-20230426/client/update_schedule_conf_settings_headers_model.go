// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleConfSettingsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateScheduleConfSettingsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateScheduleConfSettingsHeadersAccountContext) *UpdateScheduleConfSettingsHeaders
	GetAccountContext() *UpdateScheduleConfSettingsHeadersAccountContext
}

type UpdateScheduleConfSettingsHeaders struct {
	CommonHeaders  map[string]*string                               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateScheduleConfSettingsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateScheduleConfSettingsHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConfSettingsHeaders) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateScheduleConfSettingsHeaders) GetAccountContext() *UpdateScheduleConfSettingsHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateScheduleConfSettingsHeaders) SetCommonHeaders(v map[string]*string) *UpdateScheduleConfSettingsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateScheduleConfSettingsHeaders) SetAccountContext(v *UpdateScheduleConfSettingsHeadersAccountContext) *UpdateScheduleConfSettingsHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateScheduleConfSettingsHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateScheduleConfSettingsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateScheduleConfSettingsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConfSettingsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateScheduleConfSettingsHeadersAccountContext) SetAccountId(v string) *UpdateScheduleConfSettingsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateScheduleConfSettingsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
