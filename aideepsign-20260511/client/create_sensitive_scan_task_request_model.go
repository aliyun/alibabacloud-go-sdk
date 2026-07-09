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
	SetObjectKey(v string) *CreateSensitiveScanTaskRequest
	GetObjectKey() *string
}

type CreateSensitiveScanTaskRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The URL of the image to scan. Only HTTP and HTTPS protocols are supported. The image size cannot exceed 10 MB. You must specify at least one of ImageUrl and ObjectKey.
	//
	// example:
	//
	// https://example.com/document.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The ObjectKey of the image to scan in OSS. When you use ObjectKey, make sure that the key belongs to the namespace of the current caller. You must specify at least one of ImageUrl and ObjectKey.
	//
	// example:
	//
	// deepsign/123456789/scan/abc12345.jpg
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
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

func (s *CreateSensitiveScanTaskRequest) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *CreateSensitiveScanTaskRequest) SetClientToken(v string) *CreateSensitiveScanTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSensitiveScanTaskRequest) SetImageUrl(v string) *CreateSensitiveScanTaskRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateSensitiveScanTaskRequest) SetObjectKey(v string) *CreateSensitiveScanTaskRequest {
	s.ObjectKey = &v
	return s
}

func (s *CreateSensitiveScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
