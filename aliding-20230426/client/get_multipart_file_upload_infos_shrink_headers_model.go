// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultipartFileUploadInfosShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetMultipartFileUploadInfosShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetMultipartFileUploadInfosShrinkHeaders
	GetAccountContextShrink() *string
}

type GetMultipartFileUploadInfosShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetMultipartFileUploadInfosShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetMultipartFileUploadInfosShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetMultipartFileUploadInfosShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetMultipartFileUploadInfosShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetMultipartFileUploadInfosShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMultipartFileUploadInfosShrinkHeaders) SetAccountContextShrink(v string) *GetMultipartFileUploadInfosShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetMultipartFileUploadInfosShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
