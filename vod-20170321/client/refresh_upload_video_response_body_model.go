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
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A43-7DF6-D7393642****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The upload URL.
	//
	// >  The returned upload URL is a Base64-encoded URL. You must decode the Base64-encoded upload URL before you use an SDK or call an API operation to upload media files. You need to parse UploadAddress only if you use the OSS SDK or call an OSS API operation to upload media files.
	//
	// example:
	//
	// eyJTZWN1cml0eVRiQ0FJU3p3TjFxNkZ0NUIyeW****
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	// The upload credential.
	//
	// >  The returned upload credential is a Base64-encoded value. You must decode the Base64-encoded upload URL before you use an SDK or call an API operation to upload media files. You need to parse UploadAuth only if you use the OSS SDK or call an OSS API operation to upload media files.
	//
	// example:
	//
	// FJU3p3TZ0NUIyeW****
	UploadAuth *string `json:"UploadAuth,omitempty" xml:"UploadAuth,omitempty"`
	// The ID of the audio or video file.
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
