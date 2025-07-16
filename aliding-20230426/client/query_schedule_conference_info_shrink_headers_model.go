// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScheduleConferenceInfoShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryScheduleConferenceInfoShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryScheduleConferenceInfoShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryScheduleConferenceInfoShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryScheduleConferenceInfoShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceInfoShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceInfoShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryScheduleConferenceInfoShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryScheduleConferenceInfoShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryScheduleConferenceInfoShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryScheduleConferenceInfoShrinkHeaders) SetAccountContextShrink(v string) *QueryScheduleConferenceInfoShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryScheduleConferenceInfoShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
