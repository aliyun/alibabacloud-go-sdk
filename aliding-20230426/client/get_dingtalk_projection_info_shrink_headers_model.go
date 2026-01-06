// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkProjectionInfoShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDingtalkProjectionInfoShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetDingtalkProjectionInfoShrinkHeaders
	GetAccountContextShrink() *string
}

type GetDingtalkProjectionInfoShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetDingtalkProjectionInfoShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionInfoShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionInfoShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDingtalkProjectionInfoShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetDingtalkProjectionInfoShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetDingtalkProjectionInfoShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingtalkProjectionInfoShrinkHeaders) SetAccountContextShrink(v string) *GetDingtalkProjectionInfoShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetDingtalkProjectionInfoShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
