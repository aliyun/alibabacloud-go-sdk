// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMediaShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UploadMediaShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UploadMediaShrinkHeaders
	GetAccountContextShrink() *string
}

type UploadMediaShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UploadMediaShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UploadMediaShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UploadMediaShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UploadMediaShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UploadMediaShrinkHeaders) SetCommonHeaders(v map[string]*string) *UploadMediaShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UploadMediaShrinkHeaders) SetAccountContextShrink(v string) *UploadMediaShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UploadMediaShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
