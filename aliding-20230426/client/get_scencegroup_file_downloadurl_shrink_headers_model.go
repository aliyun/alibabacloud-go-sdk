// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScencegroupFileDownloadurlShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetScencegroupFileDownloadurlShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetScencegroupFileDownloadurlShrinkHeaders
	GetAccountContextShrink() *string
}

type GetScencegroupFileDownloadurlShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetScencegroupFileDownloadurlShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetScencegroupFileDownloadurlShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetScencegroupFileDownloadurlShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetScencegroupFileDownloadurlShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetScencegroupFileDownloadurlShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetScencegroupFileDownloadurlShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetScencegroupFileDownloadurlShrinkHeaders) SetAccountContextShrink(v string) *GetScencegroupFileDownloadurlShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetScencegroupFileDownloadurlShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
