// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileUploadInfoShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetFileUploadInfoShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetFileUploadInfoShrinkHeaders
	GetAccountContextShrink() *string
}

type GetFileUploadInfoShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetFileUploadInfoShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadInfoShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetFileUploadInfoShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetFileUploadInfoShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetFileUploadInfoShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFileUploadInfoShrinkHeaders) SetAccountContextShrink(v string) *GetFileUploadInfoShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetFileUploadInfoShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
