// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnderstandVideoContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UnderstandVideoContentResponseBodyData) *UnderstandVideoContentResponseBody
	GetData() *UnderstandVideoContentResponseBodyData
	SetMessage(v string) *UnderstandVideoContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnderstandVideoContentResponseBody
	GetRequestId() *string
}

type UnderstandVideoContentResponseBody struct {
	Data    *UnderstandVideoContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 71EC3F13-F0CA-4558-AC7F-A351106F59F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnderstandVideoContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnderstandVideoContentResponseBody) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentResponseBody) GetData() *UnderstandVideoContentResponseBodyData {
	return s.Data
}

func (s *UnderstandVideoContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnderstandVideoContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnderstandVideoContentResponseBody) SetData(v *UnderstandVideoContentResponseBodyData) *UnderstandVideoContentResponseBody {
	s.Data = v
	return s
}

func (s *UnderstandVideoContentResponseBody) SetMessage(v string) *UnderstandVideoContentResponseBody {
	s.Message = &v
	return s
}

func (s *UnderstandVideoContentResponseBody) SetRequestId(v string) *UnderstandVideoContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnderstandVideoContentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UnderstandVideoContentResponseBodyData struct {
	TagInfo   map[string]interface{}                           `json:"TagInfo,omitempty" xml:"TagInfo,omitempty"`
	VideoInfo *UnderstandVideoContentResponseBodyDataVideoInfo `json:"VideoInfo,omitempty" xml:"VideoInfo,omitempty" type:"Struct"`
}

func (s UnderstandVideoContentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UnderstandVideoContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentResponseBodyData) GetTagInfo() map[string]interface{} {
	return s.TagInfo
}

func (s *UnderstandVideoContentResponseBodyData) GetVideoInfo() *UnderstandVideoContentResponseBodyDataVideoInfo {
	return s.VideoInfo
}

func (s *UnderstandVideoContentResponseBodyData) SetTagInfo(v map[string]interface{}) *UnderstandVideoContentResponseBodyData {
	s.TagInfo = v
	return s
}

func (s *UnderstandVideoContentResponseBodyData) SetVideoInfo(v *UnderstandVideoContentResponseBodyDataVideoInfo) *UnderstandVideoContentResponseBodyData {
	s.VideoInfo = v
	return s
}

func (s *UnderstandVideoContentResponseBodyData) Validate() error {
	if s.VideoInfo != nil {
		if err := s.VideoInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UnderstandVideoContentResponseBodyDataVideoInfo struct {
	// example:
	//
	// 43380
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 25.0
	Fps *float32 `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// example:
	//
	// 1280
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 720
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s UnderstandVideoContentResponseBodyDataVideoInfo) String() string {
	return dara.Prettify(s)
}

func (s UnderstandVideoContentResponseBodyDataVideoInfo) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) GetDuration() *int64 {
	return s.Duration
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) GetFps() *float32 {
	return s.Fps
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) GetHeight() *int64 {
	return s.Height
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) GetWidth() *int64 {
	return s.Width
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) SetDuration(v int64) *UnderstandVideoContentResponseBodyDataVideoInfo {
	s.Duration = &v
	return s
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) SetFps(v float32) *UnderstandVideoContentResponseBodyDataVideoInfo {
	s.Fps = &v
	return s
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) SetHeight(v int64) *UnderstandVideoContentResponseBodyDataVideoInfo {
	s.Height = &v
	return s
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) SetWidth(v int64) *UnderstandVideoContentResponseBodyDataVideoInfo {
	s.Width = &v
	return s
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) Validate() error {
	return dara.Validate(s)
}
