// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeFaceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAge(v bool) *RecognizeFaceAdvanceRequest
	GetAge() *bool
	SetBeauty(v bool) *RecognizeFaceAdvanceRequest
	GetBeauty() *bool
	SetExpression(v bool) *RecognizeFaceAdvanceRequest
	GetExpression() *bool
	SetGender(v bool) *RecognizeFaceAdvanceRequest
	GetGender() *bool
	SetGlass(v bool) *RecognizeFaceAdvanceRequest
	GetGlass() *bool
	SetHat(v bool) *RecognizeFaceAdvanceRequest
	GetHat() *bool
	SetImageURLObject(v io.Reader) *RecognizeFaceAdvanceRequest
	GetImageURLObject() io.Reader
	SetMask(v bool) *RecognizeFaceAdvanceRequest
	GetMask() *bool
	SetMaxFaceNumber(v int64) *RecognizeFaceAdvanceRequest
	GetMaxFaceNumber() *int64
	SetQuality(v bool) *RecognizeFaceAdvanceRequest
	GetQuality() *bool
}

type RecognizeFaceAdvanceRequest struct {
	// example:
	//
	// false
	Age *bool `json:"Age,omitempty" xml:"Age,omitempty"`
	// example:
	//
	// false
	Beauty *bool `json:"Beauty,omitempty" xml:"Beauty,omitempty"`
	// example:
	//
	// false
	Expression *bool `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// false
	Gender *bool `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// example:
	//
	// false
	Glass *bool `json:"Glass,omitempty" xml:"Glass,omitempty"`
	// example:
	//
	// false
	Hat *bool `json:"Hat,omitempty" xml:"Hat,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/RecognizeFace/RecognizeFace1.png
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// false
	Mask *bool `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// example:
	//
	// 1
	MaxFaceNumber *int64 `json:"MaxFaceNumber,omitempty" xml:"MaxFaceNumber,omitempty"`
	// example:
	//
	// false
	Quality *bool `json:"Quality,omitempty" xml:"Quality,omitempty"`
}

func (s RecognizeFaceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFaceAdvanceRequest) GetAge() *bool {
	return s.Age
}

func (s *RecognizeFaceAdvanceRequest) GetBeauty() *bool {
	return s.Beauty
}

func (s *RecognizeFaceAdvanceRequest) GetExpression() *bool {
	return s.Expression
}

func (s *RecognizeFaceAdvanceRequest) GetGender() *bool {
	return s.Gender
}

func (s *RecognizeFaceAdvanceRequest) GetGlass() *bool {
	return s.Glass
}

func (s *RecognizeFaceAdvanceRequest) GetHat() *bool {
	return s.Hat
}

func (s *RecognizeFaceAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RecognizeFaceAdvanceRequest) GetMask() *bool {
	return s.Mask
}

func (s *RecognizeFaceAdvanceRequest) GetMaxFaceNumber() *int64 {
	return s.MaxFaceNumber
}

func (s *RecognizeFaceAdvanceRequest) GetQuality() *bool {
	return s.Quality
}

func (s *RecognizeFaceAdvanceRequest) SetAge(v bool) *RecognizeFaceAdvanceRequest {
	s.Age = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetBeauty(v bool) *RecognizeFaceAdvanceRequest {
	s.Beauty = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetExpression(v bool) *RecognizeFaceAdvanceRequest {
	s.Expression = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetGender(v bool) *RecognizeFaceAdvanceRequest {
	s.Gender = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetGlass(v bool) *RecognizeFaceAdvanceRequest {
	s.Glass = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetHat(v bool) *RecognizeFaceAdvanceRequest {
	s.Hat = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetMask(v bool) *RecognizeFaceAdvanceRequest {
	s.Mask = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetMaxFaceNumber(v int64) *RecognizeFaceAdvanceRequest {
	s.MaxFaceNumber = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetQuality(v bool) *RecognizeFaceAdvanceRequest {
	s.Quality = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
