// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageBasicInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *DetectImageBasicInfoRequest
	GetImageUrl() *string
	SetObjectKey(v string) *DetectImageBasicInfoRequest
	GetObjectKey() *string
}

type DetectImageBasicInfoRequest struct {
	// The URL of the image. Only HTTP and HTTPS protocols are supported. Specify either ImageUrl or ObjectKey. At least one of them is required.
	//
	// example:
	//
	// https://example.com/photo.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The ObjectKey of the image in OSS. When using ObjectKey, ensure that the key belongs to the namespace of the current caller. Specify either ImageUrl or ObjectKey. At least one of them is required.
	//
	// example:
	//
	// deepsign/123456789/image-generation/abc12345-def6-7890-abcd-ef1234567890.png
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
}

func (s DetectImageBasicInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectImageBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *DetectImageBasicInfoRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DetectImageBasicInfoRequest) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *DetectImageBasicInfoRequest) SetImageUrl(v string) *DetectImageBasicInfoRequest {
	s.ImageUrl = &v
	return s
}

func (s *DetectImageBasicInfoRequest) SetObjectKey(v string) *DetectImageBasicInfoRequest {
	s.ObjectKey = &v
	return s
}

func (s *DetectImageBasicInfoRequest) Validate() error {
	return dara.Validate(s)
}
