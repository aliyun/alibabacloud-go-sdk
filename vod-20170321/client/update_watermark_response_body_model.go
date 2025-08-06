// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWatermarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWatermarkResponseBody
	GetRequestId() *string
	SetWatermarkInfo(v *UpdateWatermarkResponseBodyWatermarkInfo) *UpdateWatermarkResponseBody
	GetWatermarkInfo() *UpdateWatermarkResponseBodyWatermarkInfo
}

type UpdateWatermarkResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the watermark template.
	WatermarkInfo *UpdateWatermarkResponseBodyWatermarkInfo `json:"WatermarkInfo,omitempty" xml:"WatermarkInfo,omitempty" type:"Struct"`
}

func (s UpdateWatermarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWatermarkResponseBody) GetWatermarkInfo() *UpdateWatermarkResponseBodyWatermarkInfo {
	return s.WatermarkInfo
}

func (s *UpdateWatermarkResponseBody) SetRequestId(v string) *UpdateWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWatermarkResponseBody) SetWatermarkInfo(v *UpdateWatermarkResponseBodyWatermarkInfo) *UpdateWatermarkResponseBody {
	s.WatermarkInfo = v
	return s
}

func (s *UpdateWatermarkResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateWatermarkResponseBodyWatermarkInfo struct {
	// The time when the watermark template was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-11-06T08:03:17Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The URL of the watermark file. The URL is an Object Storage Service (OSS) URL or an Alibaba Cloud CDN URL.
	//
	// >  This parameter is returned only for image watermark templates.
	//
	// example:
	//
	// https://outin-32****9f4b3e7.oss-cn-shanghai.aliyuncs.com/image/cover/E6C3448CC8B715E6F8A72EC6B-6-2.png?Expires=1541600583&OSSAccessKeyId=****&Signature=gmf1eYMoDVg%2BHQCb4UGozBW****
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// Indicates whether the watermark template is the default one.
	//
	// 	- **Default**
	//
	// 	- **NotDefault**
	//
	// example:
	//
	// NotDefault
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The name of the watermark template.
	//
	// example:
	//
	// image watermark test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the watermark template.
	//
	// 	- **Image**: image watermark template
	//
	// 	- **Text**: text watermark template
	//
	// example:
	//
	// Text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The configuration information of the watermark such as the display position and special effects. The value is a JSON string. The configuration parameters for image and text watermarks are different. For more information about the parameter structure, see [WatermarkConfig](~~98618#section-h01-44s-2lr~~).
	//
	// example:
	//
	// {"Width":"55","Height":"55","Dx":"9","Dy":"9","ReferPos":"BottonLeft","Type":"Image"}
	WatermarkConfig *string `json:"WatermarkConfig,omitempty" xml:"WatermarkConfig,omitempty"`
	// The ID of the watermark template.
	//
	// example:
	//
	// 505e2e287ea*****ecfddd386d384
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s UpdateWatermarkResponseBodyWatermarkInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateWatermarkResponseBodyWatermarkInfo) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) GetIsDefault() *string {
	return s.IsDefault
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) GetName() *string {
	return s.Name
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) GetType() *string {
	return s.Type
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) GetWatermarkConfig() *string {
	return s.WatermarkConfig
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) SetCreationTime(v string) *UpdateWatermarkResponseBodyWatermarkInfo {
	s.CreationTime = &v
	return s
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) SetFileUrl(v string) *UpdateWatermarkResponseBodyWatermarkInfo {
	s.FileUrl = &v
	return s
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) SetIsDefault(v string) *UpdateWatermarkResponseBodyWatermarkInfo {
	s.IsDefault = &v
	return s
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) SetName(v string) *UpdateWatermarkResponseBodyWatermarkInfo {
	s.Name = &v
	return s
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) SetType(v string) *UpdateWatermarkResponseBodyWatermarkInfo {
	s.Type = &v
	return s
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) SetWatermarkConfig(v string) *UpdateWatermarkResponseBodyWatermarkInfo {
	s.WatermarkConfig = &v
	return s
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) SetWatermarkId(v string) *UpdateWatermarkResponseBodyWatermarkInfo {
	s.WatermarkId = &v
	return s
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) Validate() error {
	return dara.Validate(s)
}
