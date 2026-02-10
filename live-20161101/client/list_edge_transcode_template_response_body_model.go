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
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Bitrate    *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Codec      *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Fps        *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Gop        *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
