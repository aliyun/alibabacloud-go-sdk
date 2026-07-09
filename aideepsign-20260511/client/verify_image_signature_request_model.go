// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyImageSignatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *VerifyImageSignatureRequest
	GetImageUrl() *string
	SetObjectKey(v string) *VerifyImageSignatureRequest
	GetObjectKey() *string
}

type VerifyImageSignatureRequest struct {
	// The URL of the image to verify. Specify either ImageUrl or ObjectKey. At least one of them is required.
	//
	// example:
	//
	// https://example.com/signed-photo.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The ObjectKey of the image in OSS. When you use ObjectKey, ensure that the key belongs to the namespace of the current caller. Specify either ImageUrl or ObjectKey. At least one of them is required.
	//
	// example:
	//
	// deepsign/123456789/image-generation/abc12345-def6-7890-abcd-ef1234567890.png
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
}

func (s VerifyImageSignatureRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyImageSignatureRequest) GoString() string {
	return s.String()
}

func (s *VerifyImageSignatureRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *VerifyImageSignatureRequest) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *VerifyImageSignatureRequest) SetImageUrl(v string) *VerifyImageSignatureRequest {
	s.ImageUrl = &v
	return s
}

func (s *VerifyImageSignatureRequest) SetObjectKey(v string) *VerifyImageSignatureRequest {
	s.ObjectKey = &v
	return s
}

func (s *VerifyImageSignatureRequest) Validate() error {
	return dara.Validate(s)
}
