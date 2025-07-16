// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoConferenceSettingShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateVideoConferenceSettingShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateVideoConferenceSettingShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateVideoConferenceSettingShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateVideoConferenceSettingShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoConferenceSettingShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceSettingShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateVideoConferenceSettingShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateVideoConferenceSettingShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateVideoConferenceSettingShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateVideoConferenceSettingShrinkHeaders) SetAccountContextShrink(v string) *UpdateVideoConferenceSettingShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateVideoConferenceSettingShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
