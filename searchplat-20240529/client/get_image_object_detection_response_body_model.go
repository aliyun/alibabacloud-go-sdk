// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageObjectDetectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetImageObjectDetectionResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetImageObjectDetectionResponseBody
	GetRequestId() *string
	SetResult(v *GetImageObjectDetectionResponseBodyResult) *GetImageObjectDetectionResponseBody
	GetResult() *GetImageObjectDetectionResponseBodyResult
	SetUsage(v *GetImageObjectDetectionResponseBodyUsage) *GetImageObjectDetectionResponseBody
	GetUsage() *GetImageObjectDetectionResponseBodyUsage
}

type GetImageObjectDetectionResponseBody struct {
	Latency   *int32                                     `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                    `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetImageObjectDetectionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetImageObjectDetectionResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetImageObjectDetectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageObjectDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageObjectDetectionResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetImageObjectDetectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageObjectDetectionResponseBody) GetResult() *GetImageObjectDetectionResponseBodyResult {
	return s.Result
}

func (s *GetImageObjectDetectionResponseBody) GetUsage() *GetImageObjectDetectionResponseBodyUsage {
	return s.Usage
}

func (s *GetImageObjectDetectionResponseBody) SetLatency(v int32) *GetImageObjectDetectionResponseBody {
	s.Latency = &v
	return s
}

func (s *GetImageObjectDetectionResponseBody) SetRequestId(v string) *GetImageObjectDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageObjectDetectionResponseBody) SetResult(v *GetImageObjectDetectionResponseBodyResult) *GetImageObjectDetectionResponseBody {
	s.Result = v
	return s
}

func (s *GetImageObjectDetectionResponseBody) SetUsage(v *GetImageObjectDetectionResponseBodyUsage) *GetImageObjectDetectionResponseBody {
	s.Usage = v
	return s
}

func (s *GetImageObjectDetectionResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageObjectDetectionResponseBodyResult struct {
	Objects []*GetImageObjectDetectionResponseBodyResultObjects `json:"objects,omitempty" xml:"objects,omitempty" type:"Repeated"`
}

func (s GetImageObjectDetectionResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetImageObjectDetectionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetImageObjectDetectionResponseBodyResult) GetObjects() []*GetImageObjectDetectionResponseBodyResultObjects {
	return s.Objects
}

func (s *GetImageObjectDetectionResponseBodyResult) SetObjects(v []*GetImageObjectDetectionResponseBodyResultObjects) *GetImageObjectDetectionResponseBodyResult {
	s.Objects = v
	return s
}

func (s *GetImageObjectDetectionResponseBodyResult) Validate() error {
	if s.Objects != nil {
		for _, item := range s.Objects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetImageObjectDetectionResponseBodyResultObjects struct {
	Confidence *string                                                   `json:"confidence,omitempty" xml:"confidence,omitempty"`
	Location   *GetImageObjectDetectionResponseBodyResultObjectsLocation `json:"location,omitempty" xml:"location,omitempty" type:"Struct"`
}

func (s GetImageObjectDetectionResponseBodyResultObjects) String() string {
	return dara.Prettify(s)
}

func (s GetImageObjectDetectionResponseBodyResultObjects) GoString() string {
	return s.String()
}

func (s *GetImageObjectDetectionResponseBodyResultObjects) GetConfidence() *string {
	return s.Confidence
}

func (s *GetImageObjectDetectionResponseBodyResultObjects) GetLocation() *GetImageObjectDetectionResponseBodyResultObjectsLocation {
	return s.Location
}

func (s *GetImageObjectDetectionResponseBodyResultObjects) SetConfidence(v string) *GetImageObjectDetectionResponseBodyResultObjects {
	s.Confidence = &v
	return s
}

func (s *GetImageObjectDetectionResponseBodyResultObjects) SetLocation(v *GetImageObjectDetectionResponseBodyResultObjectsLocation) *GetImageObjectDetectionResponseBodyResultObjects {
	s.Location = v
	return s
}

func (s *GetImageObjectDetectionResponseBodyResultObjects) Validate() error {
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageObjectDetectionResponseBodyResultObjectsLocation struct {
	X      *int32 `json:"X,omitempty" xml:"X,omitempty"`
	Y      *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	Width  *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s GetImageObjectDetectionResponseBodyResultObjectsLocation) String() string {
	return dara.Prettify(s)
}

func (s GetImageObjectDetectionResponseBodyResultObjectsLocation) GoString() string {
	return s.String()
}

func (s *GetImageObjectDetectionResponseBodyResultObjectsLocation) GetX() *int32 {
	return s.X
}

func (s *GetImageObjectDetectionResponseBodyResultObjectsLocation) GetY() *int32 {
	return s.Y
}

func (s *GetImageObjectDetectionResponseBodyResultObjectsLocation) GetHeight() *int32 {
	return s.Height
}

func (s *GetImageObjectDetectionResponseBodyResultObjectsLocation) GetWidth() *int32 {
	return s.Width
}

func (s *GetImageObjectDetectionResponseBodyResultObjectsLocation) SetX(v int32) *GetImageObjectDetectionResponseBodyResultObjectsLocation {
	s.X = &v
	return s
}

func (s *GetImageObjectDetectionResponseBodyResultObjectsLocation) SetY(v int32) *GetImageObjectDetectionResponseBodyResultObjectsLocation {
	s.Y = &v
	return s
}

func (s *GetImageObjectDetectionResponseBodyResultObjectsLocation) SetHeight(v int32) *GetImageObjectDetectionResponseBodyResultObjectsLocation {
	s.Height = &v
	return s
}

func (s *GetImageObjectDetectionResponseBodyResultObjectsLocation) SetWidth(v int32) *GetImageObjectDetectionResponseBodyResultObjectsLocation {
	s.Width = &v
	return s
}

func (s *GetImageObjectDetectionResponseBodyResultObjectsLocation) Validate() error {
	return dara.Validate(s)
}

type GetImageObjectDetectionResponseBodyUsage struct {
	Image *int64 `json:"image,omitempty" xml:"image,omitempty"`
}

func (s GetImageObjectDetectionResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetImageObjectDetectionResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetImageObjectDetectionResponseBodyUsage) GetImage() *int64 {
	return s.Image
}

func (s *GetImageObjectDetectionResponseBodyUsage) SetImage(v int64) *GetImageObjectDetectionResponseBodyUsage {
	s.Image = &v
	return s
}

func (s *GetImageObjectDetectionResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
