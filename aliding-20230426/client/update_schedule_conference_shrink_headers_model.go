// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleConferenceShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateScheduleConferenceShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateScheduleConferenceShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateScheduleConferenceShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateScheduleConferenceShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConferenceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConferenceShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateScheduleConferenceShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateScheduleConferenceShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateScheduleConferenceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateScheduleConferenceShrinkHeaders) SetAccountContextShrink(v string) *UpdateScheduleConferenceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateScheduleConferenceShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
