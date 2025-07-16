// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScheduleConferenceShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryScheduleConferenceShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryScheduleConferenceShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryScheduleConferenceShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryScheduleConferenceShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryScheduleConferenceShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryScheduleConferenceShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryScheduleConferenceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryScheduleConferenceShrinkHeaders) SetAccountContextShrink(v string) *QueryScheduleConferenceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryScheduleConferenceShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
