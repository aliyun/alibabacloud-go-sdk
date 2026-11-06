// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpCodePackageUploadUrlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *GetMcpCodePackageUploadUrlShrinkRequest
	GetBodyShrink() *string
}

type GetMcpCodePackageUploadUrlShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcpCodePackageUploadUrlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcpCodePackageUploadUrlShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMcpCodePackageUploadUrlShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *GetMcpCodePackageUploadUrlShrinkRequest) SetBodyShrink(v string) *GetMcpCodePackageUploadUrlShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *GetMcpCodePackageUploadUrlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
