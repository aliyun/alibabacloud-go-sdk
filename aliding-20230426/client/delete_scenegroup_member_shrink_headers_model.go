// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScenegroupMemberShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteScenegroupMemberShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteScenegroupMemberShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteScenegroupMemberShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteScenegroupMemberShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteScenegroupMemberShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteScenegroupMemberShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteScenegroupMemberShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteScenegroupMemberShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteScenegroupMemberShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteScenegroupMemberShrinkHeaders) SetAccountContextShrink(v string) *DeleteScenegroupMemberShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteScenegroupMemberShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
