// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignUserImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SignUserImageRequest
	GetClientToken() *string
	SetImageUrl(v string) *SignUserImageRequest
	GetImageUrl() *string
	SetObjectKey(v string) *SignUserImageRequest
	GetObjectKey() *string
}

type SignUserImageRequest struct {
	// The client token that is used to ensure the idempotence of the request. The client generates this value. Make sure the value is unique across different requests. ClientToken supports only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The URL of the image to be signed. HTTP and HTTPS URLs are supported. Specify at least one of `ImageUrl` and `ObjectKey`.
	//
	// example:
	//
	// https://example.com/photo.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The ObjectKey of the image to be signed in OSS. When you use `ObjectKey`, make sure the key belongs to the namespace of the current caller. Specify at least one of `ImageUrl` and `ObjectKey`.
	//
	// example:
	//
	// deepsign/123456789/image-generation/abc12345-def6-7890-abcd-ef1234567890.png
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
}

func (s SignUserImageRequest) String() string {
	return dara.Prettify(s)
}

func (s SignUserImageRequest) GoString() string {
	return s.String()
}

func (s *SignUserImageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SignUserImageRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *SignUserImageRequest) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *SignUserImageRequest) SetClientToken(v string) *SignUserImageRequest {
	s.ClientToken = &v
	return s
}

func (s *SignUserImageRequest) SetImageUrl(v string) *SignUserImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *SignUserImageRequest) SetObjectKey(v string) *SignUserImageRequest {
	s.ObjectKey = &v
	return s
}

func (s *SignUserImageRequest) Validate() error {
	return dara.Validate(s)
}
