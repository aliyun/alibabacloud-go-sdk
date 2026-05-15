// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteAllShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MuteAllShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *MuteAllShrinkHeaders
	GetAccountContextShrink() *string
}

type MuteAllShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s MuteAllShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s MuteAllShrinkHeaders) GoString() string {
	return s.String()
}

func (s *MuteAllShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MuteAllShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *MuteAllShrinkHeaders) SetCommonHeaders(v map[string]*string) *MuteAllShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MuteAllShrinkHeaders) SetAccountContextShrink(v string) *MuteAllShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *MuteAllShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
