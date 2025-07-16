// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceInfoByRoomCodeShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryConferenceInfoByRoomCodeShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryConferenceInfoByRoomCodeShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryConferenceInfoByRoomCodeShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryConferenceInfoByRoomCodeShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceInfoByRoomCodeShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoByRoomCodeShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryConferenceInfoByRoomCodeShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryConferenceInfoByRoomCodeShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryConferenceInfoByRoomCodeShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryConferenceInfoByRoomCodeShrinkHeaders) SetAccountContextShrink(v string) *QueryConferenceInfoByRoomCodeShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
