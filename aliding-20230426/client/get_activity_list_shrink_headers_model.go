// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetActivityListShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetActivityListShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetActivityListShrinkHeaders
	GetAccountContextShrink() *string
}

type GetActivityListShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetActivityListShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetActivityListShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetActivityListShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetActivityListShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetActivityListShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetActivityListShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetActivityListShrinkHeaders) SetAccountContextShrink(v string) *GetActivityListShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetActivityListShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
