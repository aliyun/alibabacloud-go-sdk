// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVideoShotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectVideoShotResponseBodyData) *DetectVideoShotResponseBody
	GetData() *DetectVideoShotResponseBodyData
	SetMessage(v string) *DetectVideoShotResponseBody
	GetMessage() *string
	SetRequestId(v string) *DetectVideoShotResponseBody
	GetRequestId() *string
}

type DetectVideoShotResponseBody struct {
	Data    *DetectVideoShotResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0033B795-09C7-4EB9-A33C-EBA325192B0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectVideoShotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoShotResponseBody) GoString() string {
	return s.String()
}

func (s *DetectVideoShotResponseBody) GetData() *DetectVideoShotResponseBodyData {
	return s.Data
}

func (s *DetectVideoShotResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetectVideoShotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectVideoShotResponseBody) SetData(v *DetectVideoShotResponseBodyData) *DetectVideoShotResponseBody {
	s.Data = v
	return s
}

func (s *DetectVideoShotResponseBody) SetMessage(v string) *DetectVideoShotResponseBody {
	s.Message = &v
	return s
}

func (s *DetectVideoShotResponseBody) SetRequestId(v string) *DetectVideoShotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectVideoShotResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectVideoShotResponseBodyData struct {
	// 1
	ShotFrameIds []*int32 `json:"ShotFrameIds,omitempty" xml:"ShotFrameIds,omitempty" type:"Repeated"`
}

func (s DetectVideoShotResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoShotResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectVideoShotResponseBodyData) GetShotFrameIds() []*int32 {
	return s.ShotFrameIds
}

func (s *DetectVideoShotResponseBodyData) SetShotFrameIds(v []*int32) *DetectVideoShotResponseBodyData {
	s.ShotFrameIds = v
	return s
}

func (s *DetectVideoShotResponseBodyData) Validate() error {
	return dara.Validate(s)
}
