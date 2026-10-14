// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSensitiveScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSensitiveScanTaskRequest
	GetClientToken() *string
	SetImageUrl(v string) *CreateSensitiveScanTaskRequest
	GetImageUrl() *string
}

type CreateSensitiveScanTaskRequest struct {
	// Ensures the idempotence of the request. The client generates this parameter value, which must be unique across different requests. `ClientToken` supports only ASCII characters and cannot exceed 64 characters.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The URL of the image to be scanned. Only HTTP/HTTPS protocols are supported. The image size cannot exceed 10 MB. You must specify either `ImageUrl` or `ObjectKey`.
	//
	// example:
	//
	// https://example.com/document.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s CreateSensitiveScanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSensitiveScanTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSensitiveScanTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSensitiveScanTaskRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateSensitiveScanTaskRequest) SetClientToken(v string) *CreateSensitiveScanTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSensitiveScanTaskRequest) SetImageUrl(v string) *CreateSensitiveScanTaskRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateSensitiveScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
