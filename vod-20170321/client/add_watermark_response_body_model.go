// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWatermarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddWatermarkResponseBody
	GetRequestId() *string
	SetWatermarkInfo(v *AddWatermarkResponseBodyWatermarkInfo) *AddWatermarkResponseBody
	GetWatermarkInfo() *AddWatermarkResponseBodyWatermarkInfo
}

type AddWatermarkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The watermark template information.
	WatermarkInfo *AddWatermarkResponseBodyWatermarkInfo `json:"WatermarkInfo,omitempty" xml:"WatermarkInfo,omitempty" type:"Struct"`
}

func (s AddWatermarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *AddWatermarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddWatermarkResponseBody) GetWatermarkInfo() *AddWatermarkResponseBodyWatermarkInfo {
	return s.WatermarkInfo
}

func (s *AddWatermarkResponseBody) SetRequestId(v string) *AddWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWatermarkResponseBody) SetWatermarkInfo(v *AddWatermarkResponseBodyWatermarkInfo) *AddWatermarkResponseBody {
	s.WatermarkInfo = v
	return s
}

func (s *AddWatermarkResponseBody) Validate() error {
	if s.WatermarkInfo != nil {
		if err := s.WatermarkInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddWatermarkResponseBodyWatermarkInfo struct {
	// The time when the watermark template was created. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2018-11-07T09:05:52Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The URL of the watermark file (OSS URL or CDN URL).
	//
	// > This parameter is returned only for image watermark templates.
	//
	// example:
	//
	// https://outin-3262*****9f4b3e7.oss-cn-shanghai.aliyuncs.com/image/cover/E6C3448CC8B715E6F8A72EC6B-6-2.png?Expires=1541600583&OSSAccessKeyId=****&Signature=gmf1eYMoDVg%2BHQCb4UGozBW****
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// Indicates whether the watermark template is the default template. Valid values:
	//
	// - **Default**: the default watermark template.
	//
	// - **NotDefault**: not the default watermark template.
	//
	// example:
	//
	// NotDefault
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The name of the watermark template.
	//
	// example:
	//
	// Image watermark template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the watermark template. Valid values:
	//
	// - **Image**: image watermark template.
	//
	// - **Text**: text watermark template.
	//
	// example:
	//
	// Image
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The configuration information of the watermark (JSON string), including the display position and effect of the watermark. The configuration parameters differ between image watermarks and text watermarks. For more information about the parameter structure, see [WatermarkConfig](~~98618#section-h01-44s-2lr~~).
	//
	// example:
	//
	// {
	//
	//       "Width": "55",
	//
	//       "Height": "55",
	//
	//       "Dx": "9",
	//
	//       "Dy": "9",
	//
	//       "ReferPos": "BottomLeft"
	//
	// }
	WatermarkConfig *string `json:"WatermarkConfig,omitempty" xml:"WatermarkConfig,omitempty"`
	// The ID of the watermark template. You can use this watermark template ID to associate the template with a transcoding template group, or to query, modify, delete, or set the template as the default watermark template.
	//
	// example:
	//
	// 9bcc8bfadb84*****109a2671d0df97
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s AddWatermarkResponseBodyWatermarkInfo) String() string {
	return dara.Prettify(s)
}

func (s AddWatermarkResponseBodyWatermarkInfo) GoString() string {
	return s.String()
}

func (s *AddWatermarkResponseBodyWatermarkInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *AddWatermarkResponseBodyWatermarkInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *AddWatermarkResponseBodyWatermarkInfo) GetIsDefault() *string {
	return s.IsDefault
}

func (s *AddWatermarkResponseBodyWatermarkInfo) GetName() *string {
	return s.Name
}

func (s *AddWatermarkResponseBodyWatermarkInfo) GetType() *string {
	return s.Type
}

func (s *AddWatermarkResponseBodyWatermarkInfo) GetWatermarkConfig() *string {
	return s.WatermarkConfig
}

func (s *AddWatermarkResponseBodyWatermarkInfo) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *AddWatermarkResponseBodyWatermarkInfo) SetCreationTime(v string) *AddWatermarkResponseBodyWatermarkInfo {
	s.CreationTime = &v
	return s
}

func (s *AddWatermarkResponseBodyWatermarkInfo) SetFileUrl(v string) *AddWatermarkResponseBodyWatermarkInfo {
	s.FileUrl = &v
	return s
}

func (s *AddWatermarkResponseBodyWatermarkInfo) SetIsDefault(v string) *AddWatermarkResponseBodyWatermarkInfo {
	s.IsDefault = &v
	return s
}

func (s *AddWatermarkResponseBodyWatermarkInfo) SetName(v string) *AddWatermarkResponseBodyWatermarkInfo {
	s.Name = &v
	return s
}

func (s *AddWatermarkResponseBodyWatermarkInfo) SetType(v string) *AddWatermarkResponseBodyWatermarkInfo {
	s.Type = &v
	return s
}

func (s *AddWatermarkResponseBodyWatermarkInfo) SetWatermarkConfig(v string) *AddWatermarkResponseBodyWatermarkInfo {
	s.WatermarkConfig = &v
	return s
}

func (s *AddWatermarkResponseBodyWatermarkInfo) SetWatermarkId(v string) *AddWatermarkResponseBodyWatermarkInfo {
	s.WatermarkId = &v
	return s
}

func (s *AddWatermarkResponseBodyWatermarkInfo) Validate() error {
	return dara.Validate(s)
}
