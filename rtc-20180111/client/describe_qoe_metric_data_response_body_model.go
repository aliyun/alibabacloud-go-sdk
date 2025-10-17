// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQoeMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudioData(v []*DescribeQoeMetricDataResponseBodyAudioData) *DescribeQoeMetricDataResponseBody
	GetAudioData() []*DescribeQoeMetricDataResponseBodyAudioData
	SetRequestId(v string) *DescribeQoeMetricDataResponseBody
	GetRequestId() *string
	SetVideoData(v []*DescribeQoeMetricDataResponseBodyVideoData) *DescribeQoeMetricDataResponseBody
	GetVideoData() []*DescribeQoeMetricDataResponseBodyVideoData
}

type DescribeQoeMetricDataResponseBody struct {
	AudioData []*DescribeQoeMetricDataResponseBodyAudioData `json:"AudioData,omitempty" xml:"AudioData,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VideoData []*DescribeQoeMetricDataResponseBodyVideoData `json:"VideoData,omitempty" xml:"VideoData,omitempty" type:"Repeated"`
}

func (s DescribeQoeMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQoeMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponseBody) GetAudioData() []*DescribeQoeMetricDataResponseBodyAudioData {
	return s.AudioData
}

func (s *DescribeQoeMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQoeMetricDataResponseBody) GetVideoData() []*DescribeQoeMetricDataResponseBodyVideoData {
	return s.VideoData
}

func (s *DescribeQoeMetricDataResponseBody) SetAudioData(v []*DescribeQoeMetricDataResponseBodyAudioData) *DescribeQoeMetricDataResponseBody {
	s.AudioData = v
	return s
}

func (s *DescribeQoeMetricDataResponseBody) SetRequestId(v string) *DescribeQoeMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBody) SetVideoData(v []*DescribeQoeMetricDataResponseBodyVideoData) *DescribeQoeMetricDataResponseBody {
	s.VideoData = v
	return s
}

func (s *DescribeQoeMetricDataResponseBody) Validate() error {
	if s.AudioData != nil {
		for _, item := range s.AudioData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoData != nil {
		for _, item := range s.VideoData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeQoeMetricDataResponseBodyAudioData struct {
	Nodes []*DescribeQoeMetricDataResponseBodyAudioDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// AUDIO
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeQoeMetricDataResponseBodyAudioData) String() string {
	return dara.Prettify(s)
}

func (s DescribeQoeMetricDataResponseBodyAudioData) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponseBodyAudioData) GetNodes() []*DescribeQoeMetricDataResponseBodyAudioDataNodes {
	return s.Nodes
}

func (s *DescribeQoeMetricDataResponseBodyAudioData) GetType() *string {
	return s.Type
}

func (s *DescribeQoeMetricDataResponseBodyAudioData) GetUserId() *string {
	return s.UserId
}

func (s *DescribeQoeMetricDataResponseBodyAudioData) SetNodes(v []*DescribeQoeMetricDataResponseBodyAudioDataNodes) *DescribeQoeMetricDataResponseBodyAudioData {
	s.Nodes = v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyAudioData) SetType(v string) *DescribeQoeMetricDataResponseBodyAudioData {
	s.Type = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyAudioData) SetUserId(v string) *DescribeQoeMetricDataResponseBodyAudioData {
	s.UserId = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyAudioData) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeQoeMetricDataResponseBodyAudioDataNodes struct {
	// example:
	//
	// 1548670256
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 123
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeQoeMetricDataResponseBodyAudioDataNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeQoeMetricDataResponseBodyAudioDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponseBodyAudioDataNodes) GetX() *string {
	return s.X
}

func (s *DescribeQoeMetricDataResponseBodyAudioDataNodes) GetY() *string {
	return s.Y
}

func (s *DescribeQoeMetricDataResponseBodyAudioDataNodes) SetX(v string) *DescribeQoeMetricDataResponseBodyAudioDataNodes {
	s.X = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyAudioDataNodes) SetY(v string) *DescribeQoeMetricDataResponseBodyAudioDataNodes {
	s.Y = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyAudioDataNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeQoeMetricDataResponseBodyVideoData struct {
	Nodes []*DescribeQoeMetricDataResponseBodyVideoDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// VIDEO_CAMERA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeQoeMetricDataResponseBodyVideoData) String() string {
	return dara.Prettify(s)
}

func (s DescribeQoeMetricDataResponseBodyVideoData) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponseBodyVideoData) GetNodes() []*DescribeQoeMetricDataResponseBodyVideoDataNodes {
	return s.Nodes
}

func (s *DescribeQoeMetricDataResponseBodyVideoData) GetType() *string {
	return s.Type
}

func (s *DescribeQoeMetricDataResponseBodyVideoData) GetUserId() *string {
	return s.UserId
}

func (s *DescribeQoeMetricDataResponseBodyVideoData) SetNodes(v []*DescribeQoeMetricDataResponseBodyVideoDataNodes) *DescribeQoeMetricDataResponseBodyVideoData {
	s.Nodes = v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyVideoData) SetType(v string) *DescribeQoeMetricDataResponseBodyVideoData {
	s.Type = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyVideoData) SetUserId(v string) *DescribeQoeMetricDataResponseBodyVideoData {
	s.UserId = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyVideoData) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeQoeMetricDataResponseBodyVideoDataNodes struct {
	// example:
	//
	// 1548670256
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 123
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeQoeMetricDataResponseBodyVideoDataNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeQoeMetricDataResponseBodyVideoDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponseBodyVideoDataNodes) GetX() *string {
	return s.X
}

func (s *DescribeQoeMetricDataResponseBodyVideoDataNodes) GetY() *string {
	return s.Y
}

func (s *DescribeQoeMetricDataResponseBodyVideoDataNodes) SetX(v string) *DescribeQoeMetricDataResponseBodyVideoDataNodes {
	s.X = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyVideoDataNodes) SetY(v string) *DescribeQoeMetricDataResponseBodyVideoDataNodes {
	s.Y = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyVideoDataNodes) Validate() error {
	return dara.Validate(s)
}
