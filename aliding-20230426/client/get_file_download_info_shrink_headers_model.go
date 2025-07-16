// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileDownloadInfoShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetFileDownloadInfoShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetFileDownloadInfoShrinkHeaders
	GetAccountContextShrink() *string
}

type GetFileDownloadInfoShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetFileDownloadInfoShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetFileDownloadInfoShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetFileDownloadInfoShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetFileDownloadInfoShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetFileDownloadInfoShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFileDownloadInfoShrinkHeaders) SetAccountContextShrink(v string) *GetFileDownloadInfoShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetFileDownloadInfoShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
