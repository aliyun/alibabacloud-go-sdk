// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeTranscodeTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListEdgeTranscodeTemplateResponseBody
	GetRequestId() *string
	SetTemplateList(v *ListEdgeTranscodeTemplateResponseBodyTemplateList) *ListEdgeTranscodeTemplateResponseBody
	GetTemplateList() *ListEdgeTranscodeTemplateResponseBodyTemplateList
	SetTotalCount(v int32) *ListEdgeTranscodeTemplateResponseBody
	GetTotalCount() *int32
}

type ListEdgeTranscodeTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the edge transcoding templates.
	TemplateList *ListEdgeTranscodeTemplateResponseBodyTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Struct"`
	// The total number of templates returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEdgeTranscodeTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeTranscodeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListEdgeTranscodeTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEdgeTranscodeTemplateResponseBody) GetTemplateList() *ListEdgeTranscodeTemplateResponseBodyTemplateList {
	return s.TemplateList
}

func (s *ListEdgeTranscodeTemplateResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEdgeTranscodeTemplateResponseBody) SetRequestId(v string) *ListEdgeTranscodeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEdgeTranscodeTemplateResponseBody) SetTemplateList(v *ListEdgeTranscodeTemplateResponseBodyTemplateList) *ListEdgeTranscodeTemplateResponseBody {
	s.TemplateList = v
	return s
}

func (s *ListEdgeTranscodeTemplateResponseBody) SetTotalCount(v int32) *ListEdgeTranscodeTemplateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEdgeTranscodeTemplateResponseBody) Validate() error {
	if s.TemplateList != nil {
		if err := s.TemplateList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEdgeTranscodeTemplateResponseBodyTemplateList struct {
	Template []*ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s ListEdgeTranscodeTemplateResponseBodyTemplateList) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeTranscodeTemplateResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateList) GetTemplate() []*ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate {
	return s.Template
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateList) SetTemplate(v []*ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) *ListEdgeTranscodeTemplateResponseBodyTemplateList {
	s.Template = v
	return s
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateList) Validate() error {
	if s.Template != nil {
		for _, item := range s.Template {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate struct {
	// The bitrate. If a numeric value is returned, a fixed bitrate is configured for the output stream. If ws is returned, the output stream maintains the same bitrate as the input stream.
	//
	// example:
	//
	// 3000
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The video encoding format. Valid values:
	//
	// 	- H.264
	//
	// 	- H.265
	//
	// example:
	//
	// H.264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The time when the image template was created.
	//
	// example:
	//
	// 2023-07-25T02:48:58Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The frame rate. If a numeric value is returned, a fixed frame rate is configured for the output stream. If ws is returned, the output stream maintains the same frame rate as the input stream.
	//
	// example:
	//
	// 30
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The group of pictures (GOP) size. The GOP size can be defined by the number of frames or the time interval between I-frames. If ws is returned, the output stream maintains the same GOP size as the input stream.
	//
	// example:
	//
	// 2s
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The template name.
	//
	// example:
	//
	// my_template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resolution. If width and height values are returned, a fixed resolution is configured for the output stream. If ws is returned, the output stream maintains the same resolution as the input stream.
	//
	// >  If the width value is -1, the width of the output video is adapted to a fixed height. If the height value is -2, the height of the output video is adapted to a fixed width.
	//
	// example:
	//
	// 1920*1080
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The ID of the edge transcoding template.
	//
	// example:
	//
	// 9b1571b513cb44f7a1ba6ae561ff46f7
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The type of edge transcoding.
	//
	// example:
	//
	// common
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) GoString() string {
	return s.String()
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) GetBitrate() *string {
	return s.Bitrate
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) GetCodec() *string {
	return s.Codec
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) GetFps() *string {
	return s.Fps
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) GetGop() *string {
	return s.Gop
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) GetName() *string {
	return s.Name
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) GetResolution() *string {
	return s.Resolution
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) GetType() *string {
	return s.Type
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) SetBitrate(v string) *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate {
	s.Bitrate = &v
	return s
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) SetCodec(v string) *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate {
	s.Codec = &v
	return s
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) SetCreateTime(v string) *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate {
	s.CreateTime = &v
	return s
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) SetFps(v string) *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate {
	s.Fps = &v
	return s
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) SetGop(v string) *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate {
	s.Gop = &v
	return s
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) SetName(v string) *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate {
	s.Name = &v
	return s
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) SetResolution(v string) *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate {
	s.Resolution = &v
	return s
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) SetTemplateId(v string) *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate {
	s.TemplateId = &v
	return s
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) SetType(v string) *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate {
	s.Type = &v
	return s
}

func (s *ListEdgeTranscodeTemplateResponseBodyTemplateListTemplate) Validate() error {
	return dara.Validate(s)
}
