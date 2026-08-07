// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUploadUrl(v string) *GetUploadContentRequest
	GetUploadUrl() *string
}

type GetUploadContentRequest struct {
	// The OSS URL of the uploaded file.
	//
	// example:
	//
	// https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/image/upload/test_text.txt
	UploadUrl *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
}

func (s GetUploadContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUploadContentRequest) GoString() string {
	return s.String()
}

func (s *GetUploadContentRequest) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *GetUploadContentRequest) SetUploadUrl(v string) *GetUploadContentRequest {
	s.UploadUrl = &v
	return s
}

func (s *GetUploadContentRequest) Validate() error {
	return dara.Validate(s)
}
