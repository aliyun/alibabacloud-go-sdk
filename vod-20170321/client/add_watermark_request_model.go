// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWatermarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *AddWatermarkRequest
	GetAppId() *string
	SetFileUrl(v string) *AddWatermarkRequest
	GetFileUrl() *string
	SetName(v string) *AddWatermarkRequest
	GetName() *string
	SetType(v string) *AddWatermarkRequest
	GetType() *string
	SetWatermarkConfig(v string) *AddWatermarkRequest
	GetWatermarkConfig() *string
}

type AddWatermarkRequest struct {
	// The ID of the application. Default value: **app-1000000**. If you have activated the multi-application service, specify the ID of the application to add the watermark template in the specified application. For more information, see [Overview](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The URL of the watermark file. The URL must be an Object Storage Service (OSS) URL and cannot contain the information used for URL signing.
	//
	// > 	- This parameter is required if you set `Type` to `Image`.
	//
	// > 	- You can obtain the URL from the `FileURL` parameter in the response to the [CreateUploadAttachedMedia](~~CreateUploadAttachedMedia~~) operation that you call to upload the watermark image to ApsaraVideo VOD.
	//
	// example:
	//
	// http://outin-326268*****63e1403e7.oss-cn-shanghai.aliyuncs.com/image/cover/C99345*****E7FDEC-6-2.png
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The name of the watermark template.
	//
	// 	- Only letters and digits are supported.
	//
	// 	- The name cannot exceed 128 bytes.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// This parameter is required.
	//
	// example:
	//
	// watermark
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the watermark template. Valid values:
	//
	// 	- **Image*	- (default): image watermark template
	//
	// 	- **Text**: text watermark template
	//
	// This parameter is required.
	//
	// example:
	//
	// Text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The configuration information of the watermark such as the display position and special effects. The value must be a JSON string. The configuration parameters for image and text watermarks are different. For more information about the parameter structure, see [WatermarkConfig](~~98618#section-h01-44s-2lr~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Width":"55","Height":"55","Dx":"9","Dy":"9","ReferPos":"BottonLeft"}
	WatermarkConfig *string `json:"WatermarkConfig,omitempty" xml:"WatermarkConfig,omitempty"`
}

func (s AddWatermarkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWatermarkRequest) GoString() string {
	return s.String()
}

func (s *AddWatermarkRequest) GetAppId() *string {
	return s.AppId
}

func (s *AddWatermarkRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *AddWatermarkRequest) GetName() *string {
	return s.Name
}

func (s *AddWatermarkRequest) GetType() *string {
	return s.Type
}

func (s *AddWatermarkRequest) GetWatermarkConfig() *string {
	return s.WatermarkConfig
}

func (s *AddWatermarkRequest) SetAppId(v string) *AddWatermarkRequest {
	s.AppId = &v
	return s
}

func (s *AddWatermarkRequest) SetFileUrl(v string) *AddWatermarkRequest {
	s.FileUrl = &v
	return s
}

func (s *AddWatermarkRequest) SetName(v string) *AddWatermarkRequest {
	s.Name = &v
	return s
}

func (s *AddWatermarkRequest) SetType(v string) *AddWatermarkRequest {
	s.Type = &v
	return s
}

func (s *AddWatermarkRequest) SetWatermarkConfig(v string) *AddWatermarkRequest {
	s.WatermarkConfig = &v
	return s
}

func (s *AddWatermarkRequest) Validate() error {
	return dara.Validate(s)
}
