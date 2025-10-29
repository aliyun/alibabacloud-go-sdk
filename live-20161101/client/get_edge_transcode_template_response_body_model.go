// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeTranscodeTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetEdgeTranscodeTemplateResponseBody
	GetRequestId() *string
	SetTemplate(v *GetEdgeTranscodeTemplateResponseBodyTemplate) *GetEdgeTranscodeTemplateResponseBody
	GetTemplate() *GetEdgeTranscodeTemplateResponseBodyTemplate
}

type GetEdgeTranscodeTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the edge transcoding template.
	Template *GetEdgeTranscodeTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s GetEdgeTranscodeTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeTranscodeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetEdgeTranscodeTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEdgeTranscodeTemplateResponseBody) GetTemplate() *GetEdgeTranscodeTemplateResponseBodyTemplate {
	return s.Template
}

func (s *GetEdgeTranscodeTemplateResponseBody) SetRequestId(v string) *GetEdgeTranscodeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEdgeTranscodeTemplateResponseBody) SetTemplate(v *GetEdgeTranscodeTemplateResponseBodyTemplate) *GetEdgeTranscodeTemplateResponseBody {
	s.Template = v
	return s
}

func (s *GetEdgeTranscodeTemplateResponseBody) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEdgeTranscodeTemplateResponseBodyTemplate struct {
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
	// The time when the template was created.
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
	// >  If the width value is -1, the width of the output stream is adapted to the height. If the height value is -2, the height of the output stream is adapted to the width.
	//
	// example:
	//
	// 1920*1080
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The template ID.
	//
	// example:
	//
	// 9b1571b513cb44f7a1ba6ae561ff****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The type of edge transcoding.
	//
	// example:
	//
	// common
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetEdgeTranscodeTemplateResponseBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeTranscodeTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) GetCodec() *string {
	return s.Codec
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) GetFps() *string {
	return s.Fps
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) GetGop() *string {
	return s.Gop
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) GetName() *string {
	return s.Name
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) GetResolution() *string {
	return s.Resolution
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) GetType() *string {
	return s.Type
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) SetBitrate(v string) *GetEdgeTranscodeTemplateResponseBodyTemplate {
	s.Bitrate = &v
	return s
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) SetCodec(v string) *GetEdgeTranscodeTemplateResponseBodyTemplate {
	s.Codec = &v
	return s
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) SetCreateTime(v string) *GetEdgeTranscodeTemplateResponseBodyTemplate {
	s.CreateTime = &v
	return s
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) SetFps(v string) *GetEdgeTranscodeTemplateResponseBodyTemplate {
	s.Fps = &v
	return s
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) SetGop(v string) *GetEdgeTranscodeTemplateResponseBodyTemplate {
	s.Gop = &v
	return s
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) SetName(v string) *GetEdgeTranscodeTemplateResponseBodyTemplate {
	s.Name = &v
	return s
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) SetResolution(v string) *GetEdgeTranscodeTemplateResponseBodyTemplate {
	s.Resolution = &v
	return s
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) SetTemplateId(v string) *GetEdgeTranscodeTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) SetType(v string) *GetEdgeTranscodeTemplateResponseBodyTemplate {
	s.Type = &v
	return s
}

func (s *GetEdgeTranscodeTemplateResponseBodyTemplate) Validate() error {
	return dara.Validate(s)
}
