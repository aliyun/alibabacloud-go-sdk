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
	// The OSS URL of the image file. The URL does not contain authentication information.
	//
	// When you add an image watermark template, you can use this URL as the value of the `FileUrl` request parameter for the [AddWatermark](~~AddWatermark~~) operation.
	//
	// example:
	//
	// http://example.aliyundoc.com/cover/2017-34DB-4F4C-9373-003AA060****.png
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The image ID. You can use it as a request parameter for operations such as [GetImageInfo](~~GetImageInfo~~), [GetImageInfos](~~GetImageInfos~~), [UpdateImageInfos](~~UpdateImageInfos~~), and [DeleteImage](~~DeleteImage~~).
	//
	// example:
	//
	// 93ab850b4f6f46e91d24d81d4****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The access URL of the image.
	//
	// > If the returned ImageURL cannot be accessed in a browser and a 403 error is returned, it is because you have enabled URL authentication for the ApsaraVideo VOD domain name. You can disable [URL authentication](https://help.aliyun.com/document_detail/86090.html) or [generate an authentication signature](https://help.aliyun.com/document_detail/57007.html) by yourself.
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
	// > The upload URL returned by the operation is Base64-encoded. When you call an SDK or API operation to upload media assets, you must Base64-decode the URL before use. You need to parse UploadAddress only when you use the native OSS SDK or OSS API to upload media assets.
	//
	// example:
	//
	// eyJTZWN1cmuIjoiQ0FJU3p3TjF****
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	// The upload credential.
	//
	// > The upload credential returned by the operation is Base64-encoded. When you call an SDK or API operation to upload media assets, you must Base64-decode the credential before use. You need to parse UploadAuth only when you use the native OSS SDK or OSS API to upload media assets.
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
