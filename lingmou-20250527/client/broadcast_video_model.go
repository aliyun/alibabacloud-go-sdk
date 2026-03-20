// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBroadcastVideo interface {
	dara.Model
	String() string
	GoString() string
	SetAlignmentFileURL(v string) *BroadcastVideo
	GetAlignmentFileURL() *string
	SetCaptionURL(v string) *BroadcastVideo
	GetCaptionURL() *string
	SetCoverURL(v string) *BroadcastVideo
	GetCoverURL() *string
	SetCreateTime(v string) *BroadcastVideo
	GetCreateTime() *string
	SetId(v string) *BroadcastVideo
	GetId() *string
	SetModifiedTime(v string) *BroadcastVideo
	GetModifiedTime() *string
	SetName(v string) *BroadcastVideo
	GetName() *string
	SetStatus(v string) *BroadcastVideo
	GetStatus() *string
	SetVideoURL(v string) *BroadcastVideo
	GetVideoURL() *string
}

type BroadcastVideo struct {
	// example:
	//
	// https://online-avatar-property.oss-cn-beijing.aliyuncs.com/aigc_material/OUTPUT_BROADCAST_SHORT_VIDEO/Mt.CQEYXYQW4MQU2/alignment.json
	AlignmentFileURL *string `json:"alignmentFileURL,omitempty" xml:"alignmentFileURL,omitempty"`
	// example:
	//
	// https://online-avatar-property.oss-cn-beijing.aliyuncs.com/aigc_material/OUTPUT_BROADCAST_SHORT_VIDEO/Mt.CQEYXYQW4MQU2/result.srt
	CaptionURL *string `json:"captionURL,omitempty" xml:"captionURL,omitempty"`
	// example:
	//
	// https://online-avatar-property.oss-cn-beijing.aliyuncs.com/aigc_material/OUTPUT_BROADCAST_SHORT_VIDEO/Mt.CQEYXYQW4MQU2/cover.jpg
	CoverURL *string `json:"coverURL,omitempty" xml:"coverURL,omitempty"`
	// example:
	//
	// 2025-11-28T13:40:33
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// M1k3So6n9IlrDV69sr3jDa3g
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2025-11-28T13:41:31
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// https://online-avatar-property.oss-cn-beijing.aliyuncs.com/aigc_material/OUTPUT_BROADCAST_SHORT_VIDEO/Mt.CQEYXYQW4MQU2/result.mp4
	VideoURL *string `json:"videoURL,omitempty" xml:"videoURL,omitempty"`
}

func (s BroadcastVideo) String() string {
	return dara.Prettify(s)
}

func (s BroadcastVideo) GoString() string {
	return s.String()
}

func (s *BroadcastVideo) GetAlignmentFileURL() *string {
	return s.AlignmentFileURL
}

func (s *BroadcastVideo) GetCaptionURL() *string {
	return s.CaptionURL
}

func (s *BroadcastVideo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *BroadcastVideo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BroadcastVideo) GetId() *string {
	return s.Id
}

func (s *BroadcastVideo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *BroadcastVideo) GetName() *string {
	return s.Name
}

func (s *BroadcastVideo) GetStatus() *string {
	return s.Status
}

func (s *BroadcastVideo) GetVideoURL() *string {
	return s.VideoURL
}

func (s *BroadcastVideo) SetAlignmentFileURL(v string) *BroadcastVideo {
	s.AlignmentFileURL = &v
	return s
}

func (s *BroadcastVideo) SetCaptionURL(v string) *BroadcastVideo {
	s.CaptionURL = &v
	return s
}

func (s *BroadcastVideo) SetCoverURL(v string) *BroadcastVideo {
	s.CoverURL = &v
	return s
}

func (s *BroadcastVideo) SetCreateTime(v string) *BroadcastVideo {
	s.CreateTime = &v
	return s
}

func (s *BroadcastVideo) SetId(v string) *BroadcastVideo {
	s.Id = &v
	return s
}

func (s *BroadcastVideo) SetModifiedTime(v string) *BroadcastVideo {
	s.ModifiedTime = &v
	return s
}

func (s *BroadcastVideo) SetName(v string) *BroadcastVideo {
	s.Name = &v
	return s
}

func (s *BroadcastVideo) SetStatus(v string) *BroadcastVideo {
	s.Status = &v
	return s
}

func (s *BroadcastVideo) SetVideoURL(v string) *BroadcastVideo {
	s.VideoURL = &v
	return s
}

func (s *BroadcastVideo) Validate() error {
	return dara.Validate(s)
}
