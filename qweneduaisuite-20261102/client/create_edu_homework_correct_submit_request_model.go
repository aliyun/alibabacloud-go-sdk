// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEduHomeworkCorrectSubmitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDewarp(v bool) *CreateEduHomeworkCorrectSubmitRequest
	GetDewarp() *bool
	SetImageBase64(v string) *CreateEduHomeworkCorrectSubmitRequest
	GetImageBase64() *string
	SetImageUrl(v string) *CreateEduHomeworkCorrectSubmitRequest
	GetImageUrl() *string
	SetRequestId(v string) *CreateEduHomeworkCorrectSubmitRequest
	GetRequestId() *string
}

type CreateEduHomeworkCorrectSubmitRequest struct {
	// Specifies whether to correct the image distortion. Default value: `true`. If you set this parameter to `false`, the image is not corrected, `dewarp_url` returns the original image, and the coordinates are based on `dewarp_url`.
	Dewarp *bool `json:"dewarp,omitempty" xml:"dewarp,omitempty"`
	// The pure Base64-encoded string of the image, excluding the `data:` prefix and MIME declaration. The encoded size cannot exceed 15 MB. If both parameters are specified, `image_base64` takes precedence.
	//
	// example:
	//
	// The pure Base64-encoded string of the image, excluding the data: prefix and MIME declaration. The encoded size cannot exceed 15 MB. If both parameters are specified, image_base64 takes precedence
	ImageBase64 *string `json:"imageBase64,omitempty" xml:"imageBase64,omitempty"`
	// The publicly accessible HTTP or HTTPS URL of the image. The server downloads the image from this URL.
	//
	// example:
	//
	// https://example.com/homework.jpg
	ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	// The request ID on the caller side. Use a UUID for reconciliation and troubleshooting. If this parameter is not specified, the server generates a request ID and returns it in the response.
	//
	// example:
	//
	// req-20261116-0001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateEduHomeworkCorrectSubmitRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEduHomeworkCorrectSubmitRequest) GoString() string {
	return s.String()
}

func (s *CreateEduHomeworkCorrectSubmitRequest) GetDewarp() *bool {
	return s.Dewarp
}

func (s *CreateEduHomeworkCorrectSubmitRequest) GetImageBase64() *string {
	return s.ImageBase64
}

func (s *CreateEduHomeworkCorrectSubmitRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateEduHomeworkCorrectSubmitRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEduHomeworkCorrectSubmitRequest) SetDewarp(v bool) *CreateEduHomeworkCorrectSubmitRequest {
	s.Dewarp = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitRequest) SetImageBase64(v string) *CreateEduHomeworkCorrectSubmitRequest {
	s.ImageBase64 = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitRequest) SetImageUrl(v string) *CreateEduHomeworkCorrectSubmitRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitRequest) SetRequestId(v string) *CreateEduHomeworkCorrectSubmitRequest {
	s.RequestId = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitRequest) Validate() error {
	return dara.Validate(s)
}
