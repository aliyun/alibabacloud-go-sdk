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
}

type SignUserImageRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The URL of the image to be signed. HTTP/HTTPS URLs are supported. Specify either `ImageUrl` or `ObjectKey`. At least one of them is required.
	//
	// example:
	//
	// https://example.com/photo.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
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

func (s *SignUserImageRequest) SetClientToken(v string) *SignUserImageRequest {
	s.ClientToken = &v
	return s
}

func (s *SignUserImageRequest) SetImageUrl(v string) *SignUserImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *SignUserImageRequest) Validate() error {
	return dara.Validate(s)
}
