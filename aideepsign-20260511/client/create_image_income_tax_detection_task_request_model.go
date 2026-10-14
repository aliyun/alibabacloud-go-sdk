// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageIncomeTaxDetectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateImageIncomeTaxDetectionTaskRequest
	GetClientToken() *string
	SetImageUrl(v string) *CreateImageIncomeTaxDetectionTaskRequest
	GetImageUrl() *string
}

type CreateImageIncomeTaxDetectionTaskRequest struct {
	// The idempotency parameter.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The URL of the income tax screenshot to detect. Only HTTP/HTTPS URLs are supported. The image size cannot exceed 10 MB. External images are not supported. Only URLs accessible from the Chinese mainland public network are supported.
	//
	// example:
	//
	// https://example.com/incometax.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s CreateImageIncomeTaxDetectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageIncomeTaxDetectionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageIncomeTaxDetectionTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateImageIncomeTaxDetectionTaskRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateImageIncomeTaxDetectionTaskRequest) SetClientToken(v string) *CreateImageIncomeTaxDetectionTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageIncomeTaxDetectionTaskRequest) SetImageUrl(v string) *CreateImageIncomeTaxDetectionTaskRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateImageIncomeTaxDetectionTaskRequest) Validate() error {
	return dara.Validate(s)
}
