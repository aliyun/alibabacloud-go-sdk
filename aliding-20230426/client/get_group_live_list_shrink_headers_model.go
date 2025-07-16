// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupLiveListShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetGroupLiveListShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetGroupLiveListShrinkHeaders
	GetAccountContextShrink() *string
}

type GetGroupLiveListShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetGroupLiveListShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetGroupLiveListShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetGroupLiveListShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetGroupLiveListShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetGroupLiveListShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetGroupLiveListShrinkHeaders) SetAccountContextShrink(v string) *GetGroupLiveListShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetGroupLiveListShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
