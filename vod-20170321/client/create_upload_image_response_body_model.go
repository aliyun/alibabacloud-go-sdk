// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileURL(v string) *CreateUploadImageResponseBody
	GetFileURL() *string
	SetImageId(v string) *CreateUploadImageResponseBody
	GetImageId() *string
	SetImageURL(v string) *CreateUploadImageResponseBody
	GetImageURL() *string
	SetRequestId(v string) *CreateUploadImageResponseBody
	GetRequestId() *string
	SetUploadAddress(v string) *CreateUploadImageResponseBody
	GetUploadAddress() *string
	SetUploadAuth(v string) *CreateUploadImageResponseBody
	GetUploadAuth() *string
}

type CreateUploadImageResponseBody struct {
	// The OSS URL of the image file (without authentication).
	//
	// When you add an image watermark template, this URL can be used as the `FileUrl` request parameter of the [AddWatermark](~~AddWatermark~~) operation.
	//
	// example:
	//
	// http://example.aliyundoc.com/cover/2017-34DB-4F4C-9373-003AA060****.png
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The image ID. This ID can be used as a request parameter for operations such as [GetImageInfo](~~GetImageInfo~~), [GetImageInfos](~~GetImageInfos~~), [UpdateImageInfos](~~UpdateImageInfos~~), and [DeleteImage](~~DeleteImage~~).
	//
	// example:
	//
	// 93ab850b4f6f46e91d24d81d4****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The access URL of the image.
	//
	// > If the returned ImageURL is inaccessible in a browser (403 error), URL authentication is enabled for your VOD domain name. Disable [URL authentication](https://help.aliyun.com/document_detail/86090.html) or [generate a signed URL](https://help.aliyun.com/document_detail/57007.html).
	//
	// example:
	//
	// http://example.aliyundoc.com/cover/2017-34DB-4F4C-9373-003AA060****.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-AEF6-D7393642****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The upload URL.
	//
	// > The upload URL returned by this operation is a Base64-encoded value. When you use an SDK or API to upload media assets, decode the value in Base64 before use. Only uploads by using the OSS native SDK or OSS API require you to parse UploadAddress.
	//
	// example:
	//
	// eyJTZWN1cmuIjoiQ0FJU3p3TjF****
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	// The upload credential.
	//
	// > The upload credential returned by this operation is a Base64-encoded value. When you use an SDK or API to upload media assets, decode the value in Base64 before use. Only uploads by using the OSS native SDK or OSS API require you to parse UploadAuth.
	//
	// example:
	//
	// eyJFbmmRCI6Im****
	UploadAuth *string `json:"UploadAuth,omitempty" xml:"UploadAuth,omitempty"`
}

func (s CreateUploadImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadImageResponseBody) GetFileURL() *string {
	return s.FileURL
}

func (s *CreateUploadImageResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *CreateUploadImageResponseBody) GetImageURL() *string {
	return s.ImageURL
}

func (s *CreateUploadImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUploadImageResponseBody) GetUploadAddress() *string {
	return s.UploadAddress
}

func (s *CreateUploadImageResponseBody) GetUploadAuth() *string {
	return s.UploadAuth
}

func (s *CreateUploadImageResponseBody) SetFileURL(v string) *CreateUploadImageResponseBody {
	s.FileURL = &v
	return s
}

func (s *CreateUploadImageResponseBody) SetImageId(v string) *CreateUploadImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateUploadImageResponseBody) SetImageURL(v string) *CreateUploadImageResponseBody {
	s.ImageURL = &v
	return s
}

func (s *CreateUploadImageResponseBody) SetRequestId(v string) *CreateUploadImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUploadImageResponseBody) SetUploadAddress(v string) *CreateUploadImageResponseBody {
	s.UploadAddress = &v
	return s
}

func (s *CreateUploadImageResponseBody) SetUploadAuth(v string) *CreateUploadImageResponseBody {
	s.UploadAuth = &v
	return s
}

func (s *CreateUploadImageResponseBody) Validate() error {
	return dara.Validate(s)
}
