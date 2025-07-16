// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveMeetingRoomsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RemoveMeetingRoomsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *RemoveMeetingRoomsShrinkHeaders
	GetAccountContextShrink() *string
}

type RemoveMeetingRoomsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s RemoveMeetingRoomsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s RemoveMeetingRoomsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *RemoveMeetingRoomsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RemoveMeetingRoomsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *RemoveMeetingRoomsShrinkHeaders) SetCommonHeaders(v map[string]*string) *RemoveMeetingRoomsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveMeetingRoomsShrinkHeaders) SetAccountContextShrink(v string) *RemoveMeetingRoomsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *RemoveMeetingRoomsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
