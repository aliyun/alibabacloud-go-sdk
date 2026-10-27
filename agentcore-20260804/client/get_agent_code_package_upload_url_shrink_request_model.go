// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentCodePackageUploadUrlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *GetAgentCodePackageUploadUrlShrinkRequest
	GetBodyShrink() *string
}

type GetAgentCodePackageUploadUrlShrinkRequest struct {
	// The request body.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentCodePackageUploadUrlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentCodePackageUploadUrlShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAgentCodePackageUploadUrlShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *GetAgentCodePackageUploadUrlShrinkRequest) SetBodyShrink(v string) *GetAgentCodePackageUploadUrlShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *GetAgentCodePackageUploadUrlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
