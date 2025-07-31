// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUploadUrl(v string) *GetUploadLinkRequest
	GetUploadUrl() *string
}

type GetUploadLinkRequest struct {
	// example:
	//
	// http://www.aliyuncs.com/test.mp3
	UploadUrl *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
}

func (s GetUploadLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUploadLinkRequest) GoString() string {
	return s.String()
}

func (s *GetUploadLinkRequest) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *GetUploadLinkRequest) SetUploadUrl(v string) *GetUploadLinkRequest {
	s.UploadUrl = &v
	return s
}

func (s *GetUploadLinkRequest) Validate() error {
	return dara.Validate(s)
}
