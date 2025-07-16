// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleConfSettingsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateScheduleConfSettingsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateScheduleConfSettingsShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateScheduleConfSettingsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateScheduleConfSettingsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConfSettingsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateScheduleConfSettingsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateScheduleConfSettingsShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateScheduleConfSettingsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateScheduleConfSettingsShrinkHeaders) SetAccountContextShrink(v string) *UpdateScheduleConfSettingsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateScheduleConfSettingsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
