// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshUploadVideoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RefreshUploadVideoResponseBody
	GetRequestId() *string
	SetUploadAddress(v string) *RefreshUploadVideoResponseBody
	GetUploadAddress() *string
	SetUploadAuth(v string) *RefreshUploadVideoResponseBody
	GetUploadAuth() *string
	SetVideoId(v string) *RefreshUploadVideoResponseBody
	GetVideoId() *string
}

type RefreshUploadVideoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A43-7DF6-D7393642****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The upload URL.
	//
	// >The upload URL returned by this operation is a Base64-encoded value. When you use an SDK or API to upload media assets, you must decode the value in Base64 before use. You need to parse UploadAddress only if you use the China (China) native OSS SDK or OSS API for upload.
	//
	// example:
	//
	// eyJTZWN1cml0eVRiQ0FJU3p3TjFxNkZ0NUIyeW****
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	// The upload credential.
	//
	// >The upload credential returned by this operation is a Base64-encoded value. When you use an SDK or API to upload media assets, you must decode the value in Base64 before use. You need to parse UploadAuth only if you use the native OSS SDK or OSS API for upload.
	//
	// example:
	//
	// FJU3p3TZ0NUIyeW****
	UploadAuth *string `json:"UploadAuth,omitempty" xml:"UploadAuth,omitempty"`
	// The audio or video ID.
	//
	// example:
	//
	// c6a23a870c8c4ffcd40cbd381333****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s RefreshUploadVideoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshUploadVideoResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshUploadVideoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshUploadVideoResponseBody) GetUploadAddress() *string {
	return s.UploadAddress
}

func (s *RefreshUploadVideoResponseBody) GetUploadAuth() *string {
	return s.UploadAuth
}

func (s *RefreshUploadVideoResponseBody) GetVideoId() *string {
	return s.VideoId
}

func (s *RefreshUploadVideoResponseBody) SetRequestId(v string) *RefreshUploadVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshUploadVideoResponseBody) SetUploadAddress(v string) *RefreshUploadVideoResponseBody {
	s.UploadAddress = &v
	return s
}

func (s *RefreshUploadVideoResponseBody) SetUploadAuth(v string) *RefreshUploadVideoResponseBody {
	s.UploadAuth = &v
	return s
}

func (s *RefreshUploadVideoResponseBody) SetVideoId(v string) *RefreshUploadVideoResponseBody {
	s.VideoId = &v
	return s
}

func (s *RefreshUploadVideoResponseBody) Validate() error {
	return dara.Validate(s)
}
