// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScencegroupFileDownloadurlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadCode(v string) *GetScencegroupFileDownloadurlShrinkRequest
	GetDownloadCode() *string
	SetTenantContextShrink(v string) *GetScencegroupFileDownloadurlShrinkRequest
	GetTenantContextShrink() *string
}

type GetScencegroupFileDownloadurlShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// abc*****xyz
	DownloadCode        *string `json:"DownloadCode,omitempty" xml:"DownloadCode,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetScencegroupFileDownloadurlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScencegroupFileDownloadurlShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetScencegroupFileDownloadurlShrinkRequest) GetDownloadCode() *string {
	return s.DownloadCode
}

func (s *GetScencegroupFileDownloadurlShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetScencegroupFileDownloadurlShrinkRequest) SetDownloadCode(v string) *GetScencegroupFileDownloadurlShrinkRequest {
	s.DownloadCode = &v
	return s
}

func (s *GetScencegroupFileDownloadurlShrinkRequest) SetTenantContextShrink(v string) *GetScencegroupFileDownloadurlShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetScencegroupFileDownloadurlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
