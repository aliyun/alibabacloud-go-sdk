// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmartSysAvatarModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSmartSysAvatarModelsResponseBody
	GetRequestId() *string
	SetSmartSysAvatarModelList(v []*ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) *ListSmartSysAvatarModelsResponseBody
	GetSmartSysAvatarModelList() []*ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList
	SetTotalCount(v int32) *ListSmartSysAvatarModelsResponseBody
	GetTotalCount() *int32
}

type ListSmartSysAvatarModelsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried digital humans.
	SmartSysAvatarModelList []*ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList `json:"SmartSysAvatarModelList,omitempty" xml:"SmartSysAvatarModelList,omitempty" type:"Repeated"`
	// The total number of system digital human images returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSmartSysAvatarModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSmartSysAvatarModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSmartSysAvatarModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSmartSysAvatarModelsResponseBody) GetSmartSysAvatarModelList() []*ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList {
	return s.SmartSysAvatarModelList
}

func (s *ListSmartSysAvatarModelsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSmartSysAvatarModelsResponseBody) SetRequestId(v string) *ListSmartSysAvatarModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSmartSysAvatarModelsResponseBody) SetSmartSysAvatarModelList(v []*ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) *ListSmartSysAvatarModelsResponseBody {
	s.SmartSysAvatarModelList = v
	return s
}

func (s *ListSmartSysAvatarModelsResponseBody) SetTotalCount(v int32) *ListSmartSysAvatarModelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSmartSysAvatarModelsResponseBody) Validate() error {
	if s.SmartSysAvatarModelList != nil {
		for _, item := range s.SmartSysAvatarModelList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList struct {
	// The ID of the digital human. The ID is required to submit a separate digital human rendering job or use the digital human image in an intelligent timeline.
	//
	// example:
	//
	// yunqiao
	AvatarId *string `json:"AvatarId,omitempty" xml:"AvatarId,omitempty"`
	// The name of the digital human.
	AvatarName *string `json:"AvatarName,omitempty" xml:"AvatarName,omitempty"`
	// The video bitrate.
	//
	// example:
	//
	// 4000
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The sample thumbnail URL of the digital human.
	//
	// example:
	//
	// http://ice-pub-media.myalicdn.com/smart/avatarModel/coverDemo/yunqiao.mp4
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// The video height.
	//
	// example:
	//
	// 1920
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Indicates whether portrait mask rendering is supported.
	//
	// example:
	//
	// false
	OutputMask *bool `json:"OutputMask,omitempty" xml:"OutputMask,omitempty"`
	// The sample video URL of the digital human.
	//
	// example:
	//
	// http://ice-pub-media.myalicdn.com/smart/avatarModel/videoDemo/yunqiao.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	// The video width.
	//
	// example:
	//
	// 1080
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) String() string {
	return dara.Prettify(s)
}

func (s ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) GoString() string {
	return s.String()
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) GetAvatarId() *string {
	return s.AvatarId
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) GetAvatarName() *string {
	return s.AvatarName
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) GetHeight() *int32 {
	return s.Height
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) GetOutputMask() *bool {
	return s.OutputMask
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) GetWidth() *int32 {
	return s.Width
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) SetAvatarId(v string) *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList {
	s.AvatarId = &v
	return s
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) SetAvatarName(v string) *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList {
	s.AvatarName = &v
	return s
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) SetBitrate(v int32) *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList {
	s.Bitrate = &v
	return s
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) SetCoverUrl(v string) *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList {
	s.CoverUrl = &v
	return s
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) SetHeight(v int32) *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList {
	s.Height = &v
	return s
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) SetOutputMask(v bool) *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList {
	s.OutputMask = &v
	return s
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) SetVideoUrl(v string) *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList {
	s.VideoUrl = &v
	return s
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) SetWidth(v int32) *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList {
	s.Width = &v
	return s
}

func (s *ListSmartSysAvatarModelsResponseBodySmartSysAvatarModelList) Validate() error {
	return dara.Validate(s)
}
