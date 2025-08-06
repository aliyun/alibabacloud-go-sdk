// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadVideoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateUploadVideoResponseBody
	GetRequestId() *string
	SetUploadAddress(v string) *CreateUploadVideoResponseBody
	GetUploadAddress() *string
	SetUploadAuth(v string) *CreateUploadVideoResponseBody
	GetUploadAuth() *string
	SetVideoId(v string) *CreateUploadVideoResponseBody
	GetVideoId() *string
}

type CreateUploadVideoResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-04D5-D7393642****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The upload URL.
	//
	// > The returned upload URL is a Base64-encoded URL. You must decode the Base64-encoded URL before you use an SDK or call an API operation to upload media files. You need to parse UploadAddress only if you use the Object Storage Service (OSS) SDK or call an OSS API operation to upload media files.
	//
	// example:
	//
	// eyJTZWN1cml0a2VuIjoiQ0FJU3p3TjF****
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	// The upload credential.
	//
	// > The returned upload credential is a Base64-encoded value. You must decode the Base64-encoded credential before you use an SDK or call an API operation to upload media files. You need to parse UploadAuth only if you use the OSS SDK or call an OSS API operation to upload media files.
	//
	// example:
	//
	// eyJFbmRwb2ludCI6Imm****
	UploadAuth *string `json:"UploadAuth,omitempty" xml:"UploadAuth,omitempty"`
	// The ID of the audio or video file. VideoId can be used as a request parameter when you call an operation for media asset management, media processing, or media review.
	//
	// example:
	//
	// 93ab850b4f6f54b6e91d24d81d44****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s CreateUploadVideoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadVideoResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadVideoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUploadVideoResponseBody) GetUploadAddress() *string {
	return s.UploadAddress
}

func (s *CreateUploadVideoResponseBody) GetUploadAuth() *string {
	return s.UploadAuth
}

func (s *CreateUploadVideoResponseBody) GetVideoId() *string {
	return s.VideoId
}

func (s *CreateUploadVideoResponseBody) SetRequestId(v string) *CreateUploadVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUploadVideoResponseBody) SetUploadAddress(v string) *CreateUploadVideoResponseBody {
	s.UploadAddress = &v
	return s
}

func (s *CreateUploadVideoResponseBody) SetUploadAuth(v string) *CreateUploadVideoResponseBody {
	s.UploadAuth = &v
	return s
}

func (s *CreateUploadVideoResponseBody) SetVideoId(v string) *CreateUploadVideoResponseBody {
	s.VideoId = &v
	return s
}

func (s *CreateUploadVideoResponseBody) Validate() error {
	return dara.Validate(s)
}
