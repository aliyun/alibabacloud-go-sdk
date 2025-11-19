// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWatermarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWatermarkResponseBody
	GetRequestId() *string
	SetWatermarkInfo(v *GetWatermarkResponseBodyWatermarkInfo) *GetWatermarkResponseBody
	GetWatermarkInfo() *GetWatermarkResponseBodyWatermarkInfo
}

type GetWatermarkResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the watermark template.
	WatermarkInfo *GetWatermarkResponseBodyWatermarkInfo `json:"WatermarkInfo,omitempty" xml:"WatermarkInfo,omitempty" type:"Struct"`
}

func (s GetWatermarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *GetWatermarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWatermarkResponseBody) GetWatermarkInfo() *GetWatermarkResponseBodyWatermarkInfo {
	return s.WatermarkInfo
}

func (s *GetWatermarkResponseBody) SetRequestId(v string) *GetWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWatermarkResponseBody) SetWatermarkInfo(v *GetWatermarkResponseBodyWatermarkInfo) *GetWatermarkResponseBody {
	s.WatermarkInfo = v
	return s
}

func (s *GetWatermarkResponseBody) Validate() error {
	if s.WatermarkInfo != nil {
		if err := s.WatermarkInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWatermarkResponseBodyWatermarkInfo struct {
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
	// 2018-11-06T08:03:17Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The URL of the watermark file. The URL is an Object Storage Service (OSS) URL or an Alibaba Cloud CDN URL.
	//
	// >  This parameter is returned only for image watermark templates.
	//
	// example:
	//
	// https://outin-32*****f4b3e7.oss-cn-shanghai.aliyuncs.com/image/cover/F85529C8B715E6F8A72EC6B-6-2.png?Expires=1541600583&OSSAccessKeyId=****&Signature=gmf1eYMoDVg%2BHQCb4UGozBW****
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
	// image watermark test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the watermark template. Valid values:
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
	// {"ReferPos": "BottomRight","Height": "55","Width": "55","Dx": "8","Dy": "8" }
	WatermarkConfig *string `json:"WatermarkConfig,omitempty" xml:"WatermarkConfig,omitempty"`
	// The ID of the watermark template.
	//
	// example:
	//
	// 505e2e287ea*****ecfddd386d384
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s GetWatermarkResponseBodyWatermarkInfo) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarkResponseBodyWatermarkInfo) GoString() string {
	return s.String()
}

func (s *GetWatermarkResponseBodyWatermarkInfo) GetAppId() *string {
	return s.AppId
}

func (s *GetWatermarkResponseBodyWatermarkInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetWatermarkResponseBodyWatermarkInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetWatermarkResponseBodyWatermarkInfo) GetIsDefault() *string {
	return s.IsDefault
}

func (s *GetWatermarkResponseBodyWatermarkInfo) GetName() *string {
	return s.Name
}

func (s *GetWatermarkResponseBodyWatermarkInfo) GetType() *string {
	return s.Type
}

func (s *GetWatermarkResponseBodyWatermarkInfo) GetWatermarkConfig() *string {
	return s.WatermarkConfig
}

func (s *GetWatermarkResponseBodyWatermarkInfo) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *GetWatermarkResponseBodyWatermarkInfo) SetAppId(v string) *GetWatermarkResponseBodyWatermarkInfo {
	s.AppId = &v
	return s
}

func (s *GetWatermarkResponseBodyWatermarkInfo) SetCreationTime(v string) *GetWatermarkResponseBodyWatermarkInfo {
	s.CreationTime = &v
	return s
}

func (s *GetWatermarkResponseBodyWatermarkInfo) SetFileUrl(v string) *GetWatermarkResponseBodyWatermarkInfo {
	s.FileUrl = &v
	return s
}

func (s *GetWatermarkResponseBodyWatermarkInfo) SetIsDefault(v string) *GetWatermarkResponseBodyWatermarkInfo {
	s.IsDefault = &v
	return s
}

func (s *GetWatermarkResponseBodyWatermarkInfo) SetName(v string) *GetWatermarkResponseBodyWatermarkInfo {
	s.Name = &v
	return s
}

func (s *GetWatermarkResponseBodyWatermarkInfo) SetType(v string) *GetWatermarkResponseBodyWatermarkInfo {
	s.Type = &v
	return s
}

func (s *GetWatermarkResponseBodyWatermarkInfo) SetWatermarkConfig(v string) *GetWatermarkResponseBodyWatermarkInfo {
	s.WatermarkConfig = &v
	return s
}

func (s *GetWatermarkResponseBodyWatermarkInfo) SetWatermarkId(v string) *GetWatermarkResponseBodyWatermarkInfo {
	s.WatermarkId = &v
	return s
}

func (s *GetWatermarkResponseBodyWatermarkInfo) Validate() error {
	return dara.Validate(s)
}
