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
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4AF6-04D5-D7393642****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The upload URL.
	//
	// > The upload URL returned by the operation is Base64-encoded. When you call an SDK or API operation to upload the media asset, you must decode the URL before use. You only need to parse UploadAddress when you use the native OSS SDK or OSS API for upload.
	//
	// example:
	//
	// eyJTZWN1cml0a2VuIjoiQ0FJU3p3TjF****
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	// The upload credential.
	//
	// > The upload credential returned by the operation is Base64-encoded. When you call an SDK or API operation to upload the media asset, you must decode the credential before use. You only need to parse UploadAuth when you use the native OSS SDK or OSS API for upload.
	//
	// example:
	//
	// eyJFbmRwb2ludCI6Imm****
	UploadAuth *string `json:"UploadAuth,omitempty" xml:"UploadAuth,omitempty"`
	// The audio or video ID. You can use this ID as a request parameter for media asset management, media processing, and media review operations.
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
