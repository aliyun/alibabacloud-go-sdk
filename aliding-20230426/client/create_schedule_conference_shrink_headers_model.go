// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleConferenceShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateScheduleConferenceShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateScheduleConferenceShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateScheduleConferenceShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateScheduleConferenceShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleConferenceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateScheduleConferenceShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateScheduleConferenceShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateScheduleConferenceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateScheduleConferenceShrinkHeaders) SetAccountContextShrink(v string) *CreateScheduleConferenceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateScheduleConferenceShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
