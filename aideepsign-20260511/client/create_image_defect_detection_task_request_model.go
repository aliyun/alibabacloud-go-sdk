// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageDefectDetectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateImageDefectDetectionTaskRequest
	GetClientToken() *string
	SetImageUrl(v string) *CreateImageDefectDetectionTaskRequest
	GetImageUrl() *string
}

type CreateImageDefectDetectionTaskRequest struct {
	// The idempotency token.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The URL of the image to be detected. Only HTTP/HTTPS protocols are supported.
	//
	// example:
	//
	// https://example.com/product.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s CreateImageDefectDetectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageDefectDetectionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageDefectDetectionTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateImageDefectDetectionTaskRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateImageDefectDetectionTaskRequest) SetClientToken(v string) *CreateImageDefectDetectionTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageDefectDetectionTaskRequest) SetImageUrl(v string) *CreateImageDefectDetectionTaskRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateImageDefectDetectionTaskRequest) Validate() error {
	return dara.Validate(s)
}
