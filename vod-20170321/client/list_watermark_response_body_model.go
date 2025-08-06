// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWatermarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWatermarkResponseBody
	GetRequestId() *string
	SetWatermarkInfos(v []*ListWatermarkResponseBodyWatermarkInfos) *ListWatermarkResponseBody
	GetWatermarkInfos() []*ListWatermarkResponseBodyWatermarkInfos
}

type ListWatermarkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the watermark template.
	WatermarkInfos []*ListWatermarkResponseBodyWatermarkInfos `json:"WatermarkInfos,omitempty" xml:"WatermarkInfos,omitempty" type:"Repeated"`
}

func (s ListWatermarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *ListWatermarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWatermarkResponseBody) GetWatermarkInfos() []*ListWatermarkResponseBodyWatermarkInfos {
	return s.WatermarkInfos
}

func (s *ListWatermarkResponseBody) SetRequestId(v string) *ListWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWatermarkResponseBody) SetWatermarkInfos(v []*ListWatermarkResponseBodyWatermarkInfos) *ListWatermarkResponseBody {
	s.WatermarkInfos = v
	return s
}

func (s *ListWatermarkResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListWatermarkResponseBodyWatermarkInfos struct {
	// The ID of the application.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the watermark template was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-11-07T09:05:52Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The URL of the watermark file. The URL is an Object Storage Service (OSS) URL or an Alibaba Cloud CDN URL.
	//
	// >  This parameter is returned only for image watermark templates.
	//
	// example:
	//
	// https://outin-3262681cd*****89f4b3e7.oss-cn-shanghai.aliyuncs.com/image/cover/8CC8B715E6F8A72EC6B-6-2.png?Expires=1541600583&OSSAccessKeyId=****&Signature=gmf1eYMoDVg%2BHQCb4UGozB****
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// Indicates whether the watermark template is the default one. Valid values:
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
	// testName
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
	// {"FontColor": "Blue","FontSize": 80,"Content": "test watermark"}
	WatermarkConfig *string `json:"WatermarkConfig,omitempty" xml:"WatermarkConfig,omitempty"`
	// The ID of the watermark template.
	//
	// example:
	//
	// 9bcc8bfadb843*****109a2671d0df97
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s ListWatermarkResponseBodyWatermarkInfos) String() string {
	return dara.Prettify(s)
}

func (s ListWatermarkResponseBodyWatermarkInfos) GoString() string {
	return s.String()
}

func (s *ListWatermarkResponseBodyWatermarkInfos) GetAppId() *string {
	return s.AppId
}

func (s *ListWatermarkResponseBodyWatermarkInfos) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListWatermarkResponseBodyWatermarkInfos) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ListWatermarkResponseBodyWatermarkInfos) GetIsDefault() *string {
	return s.IsDefault
}

func (s *ListWatermarkResponseBodyWatermarkInfos) GetName() *string {
	return s.Name
}

func (s *ListWatermarkResponseBodyWatermarkInfos) GetType() *string {
	return s.Type
}

func (s *ListWatermarkResponseBodyWatermarkInfos) GetWatermarkConfig() *string {
	return s.WatermarkConfig
}

func (s *ListWatermarkResponseBodyWatermarkInfos) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *ListWatermarkResponseBodyWatermarkInfos) SetAppId(v string) *ListWatermarkResponseBodyWatermarkInfos {
	s.AppId = &v
	return s
}

func (s *ListWatermarkResponseBodyWatermarkInfos) SetCreationTime(v string) *ListWatermarkResponseBodyWatermarkInfos {
	s.CreationTime = &v
	return s
}

func (s *ListWatermarkResponseBodyWatermarkInfos) SetFileUrl(v string) *ListWatermarkResponseBodyWatermarkInfos {
	s.FileUrl = &v
	return s
}

func (s *ListWatermarkResponseBodyWatermarkInfos) SetIsDefault(v string) *ListWatermarkResponseBodyWatermarkInfos {
	s.IsDefault = &v
	return s
}

func (s *ListWatermarkResponseBodyWatermarkInfos) SetName(v string) *ListWatermarkResponseBodyWatermarkInfos {
	s.Name = &v
	return s
}

func (s *ListWatermarkResponseBodyWatermarkInfos) SetType(v string) *ListWatermarkResponseBodyWatermarkInfos {
	s.Type = &v
	return s
}

func (s *ListWatermarkResponseBodyWatermarkInfos) SetWatermarkConfig(v string) *ListWatermarkResponseBodyWatermarkInfos {
	s.WatermarkConfig = &v
	return s
}

func (s *ListWatermarkResponseBodyWatermarkInfos) SetWatermarkId(v string) *ListWatermarkResponseBodyWatermarkInfos {
	s.WatermarkId = &v
	return s
}

func (s *ListWatermarkResponseBodyWatermarkInfos) Validate() error {
	return dara.Validate(s)
}
