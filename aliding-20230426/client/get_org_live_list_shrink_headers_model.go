// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgLiveListShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetOrgLiveListShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetOrgLiveListShrinkHeaders
	GetAccountContextShrink() *string
}

type GetOrgLiveListShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetOrgLiveListShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetOrgLiveListShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetOrgLiveListShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetOrgLiveListShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetOrgLiveListShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrgLiveListShrinkHeaders) SetAccountContextShrink(v string) *GetOrgLiveListShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetOrgLiveListShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
