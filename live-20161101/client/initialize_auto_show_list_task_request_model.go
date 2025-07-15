// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeAutoShowListTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallBackUrl(v string) *InitializeAutoShowListTaskRequest
	GetCallBackUrl() *string
	SetCasterConfig(v string) *InitializeAutoShowListTaskRequest
	GetCasterConfig() *string
	SetDomainName(v string) *InitializeAutoShowListTaskRequest
	GetDomainName() *string
	SetEndTime(v int64) *InitializeAutoShowListTaskRequest
	GetEndTime() *int64
	SetOwnerId(v int64) *InitializeAutoShowListTaskRequest
	GetOwnerId() *int64
	SetRegionId(v string) *InitializeAutoShowListTaskRequest
	GetRegionId() *string
	SetResourceIds(v string) *InitializeAutoShowListTaskRequest
	GetResourceIds() *string
	SetStartTime(v int64) *InitializeAutoShowListTaskRequest
	GetStartTime() *int64
}

type InitializeAutoShowListTaskRequest struct {
	// The callback URL.
	//
	// example:
	//
	// http://***.com/callback
	CallBackUrl *string `json:"CallBackUrl,omitempty" xml:"CallBackUrl,omitempty"`
	// The configurations of the production studio. The following configurations are involved:
	//
	// 	- CasterTemplate: required. The output resolution.
	//
	// 	- LiveTemplate: optional. The templates to be used for transcoding.
	//
	// >  Set the value to a JSON string. Use upper camel case for fields of the string.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"CasterTemplate": "lp_ld","LiveTemplates":["lhd", "lsd","lud"]}
	CasterConfig *string `json:"CasterConfig,omitempty" xml:"CasterConfig,omitempty"`
	// The output streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1645688994000
	EndTime  *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of on-demand media asset files in the playlist. Only on-demand video files are supported. You can specify up to three video files in the playlist. The video files in the playlist are automatically played in sequence. The playback stops at the point in time specified by the EndTime parameter.
	//
	// >
	//
	// 	- You can obtain the ID of a video file in the ApsaraVideo Live console or by calling an API operation. For more information, see [Media asset management](https://help.aliyun.com/document_detail/86057.html) or [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html). - If the video files are all played before the time specified by EndTime, the final frame of the final video file is played until the time specified by EndTime arrives.
	//
	// example:
	//
	// ["89e02xxxxfb349axxxxa0c350d***	- ","6ae0xxxxxb349axxxxa0c350a****"]
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The start timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1645688994000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s InitializeAutoShowListTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s InitializeAutoShowListTaskRequest) GoString() string {
	return s.String()
}

func (s *InitializeAutoShowListTaskRequest) GetCallBackUrl() *string {
	return s.CallBackUrl
}

func (s *InitializeAutoShowListTaskRequest) GetCasterConfig() *string {
	return s.CasterConfig
}

func (s *InitializeAutoShowListTaskRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *InitializeAutoShowListTaskRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *InitializeAutoShowListTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *InitializeAutoShowListTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InitializeAutoShowListTaskRequest) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *InitializeAutoShowListTaskRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *InitializeAutoShowListTaskRequest) SetCallBackUrl(v string) *InitializeAutoShowListTaskRequest {
	s.CallBackUrl = &v
	return s
}

func (s *InitializeAutoShowListTaskRequest) SetCasterConfig(v string) *InitializeAutoShowListTaskRequest {
	s.CasterConfig = &v
	return s
}

func (s *InitializeAutoShowListTaskRequest) SetDomainName(v string) *InitializeAutoShowListTaskRequest {
	s.DomainName = &v
	return s
}

func (s *InitializeAutoShowListTaskRequest) SetEndTime(v int64) *InitializeAutoShowListTaskRequest {
	s.EndTime = &v
	return s
}

func (s *InitializeAutoShowListTaskRequest) SetOwnerId(v int64) *InitializeAutoShowListTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *InitializeAutoShowListTaskRequest) SetRegionId(v string) *InitializeAutoShowListTaskRequest {
	s.RegionId = &v
	return s
}

func (s *InitializeAutoShowListTaskRequest) SetResourceIds(v string) *InitializeAutoShowListTaskRequest {
	s.ResourceIds = &v
	return s
}

func (s *InitializeAutoShowListTaskRequest) SetStartTime(v int64) *InitializeAutoShowListTaskRequest {
	s.StartTime = &v
	return s
}

func (s *InitializeAutoShowListTaskRequest) Validate() error {
	return dara.Validate(s)
}
