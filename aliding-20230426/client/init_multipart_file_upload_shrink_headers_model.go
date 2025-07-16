// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitMultipartFileUploadShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InitMultipartFileUploadShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *InitMultipartFileUploadShrinkHeaders
	GetAccountContextShrink() *string
}

type InitMultipartFileUploadShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s InitMultipartFileUploadShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s InitMultipartFileUploadShrinkHeaders) GoString() string {
	return s.String()
}

func (s *InitMultipartFileUploadShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InitMultipartFileUploadShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *InitMultipartFileUploadShrinkHeaders) SetCommonHeaders(v map[string]*string) *InitMultipartFileUploadShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InitMultipartFileUploadShrinkHeaders) SetAccountContextShrink(v string) *InitMultipartFileUploadShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *InitMultipartFileUploadShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
