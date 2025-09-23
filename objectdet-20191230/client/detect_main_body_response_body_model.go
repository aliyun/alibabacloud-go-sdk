// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectMainBodyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectMainBodyResponseBodyData) *DetectMainBodyResponseBody
	GetData() *DetectMainBodyResponseBodyData
	SetRequestId(v string) *DetectMainBodyResponseBody
	GetRequestId() *string
}

type DetectMainBodyResponseBody struct {
	Data *DetectMainBodyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 2E59C333-5480-4231-A8AB-BEE1001EA7FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectMainBodyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectMainBodyResponseBody) GoString() string {
	return s.String()
}

func (s *DetectMainBodyResponseBody) GetData() *DetectMainBodyResponseBodyData {
	return s.Data
}

func (s *DetectMainBodyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectMainBodyResponseBody) SetData(v *DetectMainBodyResponseBodyData) *DetectMainBodyResponseBody {
	s.Data = v
	return s
}

func (s *DetectMainBodyResponseBody) SetRequestId(v string) *DetectMainBodyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectMainBodyResponseBody) Validate() error {
	return dara.Validate(s)
}

type DetectMainBodyResponseBodyData struct {
	Location *DetectMainBodyResponseBodyDataLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
}

func (s DetectMainBodyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectMainBodyResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectMainBodyResponseBodyData) GetLocation() *DetectMainBodyResponseBodyDataLocation {
	return s.Location
}

func (s *DetectMainBodyResponseBodyData) SetLocation(v *DetectMainBodyResponseBodyDataLocation) *DetectMainBodyResponseBodyData {
	s.Location = v
	return s
}

func (s *DetectMainBodyResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DetectMainBodyResponseBodyDataLocation struct {
	// example:
	//
	// 320
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 583
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 28
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 20
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DetectMainBodyResponseBodyDataLocation) String() string {
	return dara.Prettify(s)
}

func (s DetectMainBodyResponseBodyDataLocation) GoString() string {
	return s.String()
}

func (s *DetectMainBodyResponseBodyDataLocation) GetHeight() *int32 {
	return s.Height
}

func (s *DetectMainBodyResponseBodyDataLocation) GetWidth() *int32 {
	return s.Width
}

func (s *DetectMainBodyResponseBodyDataLocation) GetX() *int32 {
	return s.X
}

func (s *DetectMainBodyResponseBodyDataLocation) GetY() *int32 {
	return s.Y
}

func (s *DetectMainBodyResponseBodyDataLocation) SetHeight(v int32) *DetectMainBodyResponseBodyDataLocation {
	s.Height = &v
	return s
}

func (s *DetectMainBodyResponseBodyDataLocation) SetWidth(v int32) *DetectMainBodyResponseBodyDataLocation {
	s.Width = &v
	return s
}

func (s *DetectMainBodyResponseBodyDataLocation) SetX(v int32) *DetectMainBodyResponseBodyDataLocation {
	s.X = &v
	return s
}

func (s *DetectMainBodyResponseBodyDataLocation) SetY(v int32) *DetectMainBodyResponseBodyDataLocation {
	s.Y = &v
	return s
}

func (s *DetectMainBodyResponseBodyDataLocation) Validate() error {
	return dara.Validate(s)
}
