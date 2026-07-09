// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageDetectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateImageDetectionTaskRequest
	GetClientToken() *string
	SetCredType(v string) *CreateImageDetectionTaskRequest
	GetCredType() *string
	SetDetectType(v string) *CreateImageDetectionTaskRequest
	GetDetectType() *string
	SetImageUrl(v string) *CreateImageDetectionTaskRequest
	GetImageUrl() *string
	SetObjectKey(v string) *CreateImageDetectionTaskRequest
	GetObjectKey() *string
}

type CreateImageDetectionTaskRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The credential type code. This parameter is required when `DetectType` is set to `credential`. Valid values: `0101` (ID card), `0102` (bank card), `0104` (teacher qualification certificate), `0107` (student ID), `0108` (driver license), `0201` (storefront photo), `0202` (counter photo), `0203` (scene photo), `0301` (business license).
	//
	// example:
	//
	// 0101
	CredType *string `json:"CredType,omitempty" xml:"CredType,omitempty"`
	// The detection type. Valid values: `auto` (automatic, default), `aigc` (AIGC detection only), `credential` (credential detection only).
	//
	// example:
	//
	// auto
	DetectType *string `json:"DetectType,omitempty" xml:"DetectType,omitempty"`
	// The URL of the image to be detected. Only HTTP and HTTPS protocols are supported. You must specify at least one of `ImageUrl` and `ObjectKey`.
	//
	// example:
	//
	// https://example.com/id-card.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The `ObjectKey` of the image to be detected in OSS. When you use `ObjectKey`, make sure that the key belongs to the namespace of the current caller. You must specify at least one of `ImageUrl` and `ObjectKey`.
	//
	// example:
	//
	// deepsign/123456789/scan/abc12345.jpg
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
}

func (s CreateImageDetectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageDetectionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageDetectionTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateImageDetectionTaskRequest) GetCredType() *string {
	return s.CredType
}

func (s *CreateImageDetectionTaskRequest) GetDetectType() *string {
	return s.DetectType
}

func (s *CreateImageDetectionTaskRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateImageDetectionTaskRequest) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *CreateImageDetectionTaskRequest) SetClientToken(v string) *CreateImageDetectionTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageDetectionTaskRequest) SetCredType(v string) *CreateImageDetectionTaskRequest {
	s.CredType = &v
	return s
}

func (s *CreateImageDetectionTaskRequest) SetDetectType(v string) *CreateImageDetectionTaskRequest {
	s.DetectType = &v
	return s
}

func (s *CreateImageDetectionTaskRequest) SetImageUrl(v string) *CreateImageDetectionTaskRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateImageDetectionTaskRequest) SetObjectKey(v string) *CreateImageDetectionTaskRequest {
	s.ObjectKey = &v
	return s
}

func (s *CreateImageDetectionTaskRequest) Validate() error {
	return dara.Validate(s)
}
