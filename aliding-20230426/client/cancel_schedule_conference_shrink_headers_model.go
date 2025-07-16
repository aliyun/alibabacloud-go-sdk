// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelScheduleConferenceShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CancelScheduleConferenceShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CancelScheduleConferenceShrinkHeaders
	GetAccountContextShrink() *string
}

type CancelScheduleConferenceShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CancelScheduleConferenceShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CancelScheduleConferenceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CancelScheduleConferenceShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CancelScheduleConferenceShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CancelScheduleConferenceShrinkHeaders) SetCommonHeaders(v map[string]*string) *CancelScheduleConferenceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CancelScheduleConferenceShrinkHeaders) SetAccountContextShrink(v string) *CancelScheduleConferenceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CancelScheduleConferenceShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
