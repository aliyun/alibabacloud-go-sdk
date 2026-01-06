// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkProjectionListShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDingtalkProjectionListShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetDingtalkProjectionListShrinkHeaders
	GetAccountContextShrink() *string
}

type GetDingtalkProjectionListShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetDingtalkProjectionListShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionListShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionListShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDingtalkProjectionListShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetDingtalkProjectionListShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetDingtalkProjectionListShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingtalkProjectionListShrinkHeaders) SetAccountContextShrink(v string) *GetDingtalkProjectionListShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetDingtalkProjectionListShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
