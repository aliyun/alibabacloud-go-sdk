// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAge(v bool) *RecognizeFaceRequest
	GetAge() *bool
	SetBeauty(v bool) *RecognizeFaceRequest
	GetBeauty() *bool
	SetExpression(v bool) *RecognizeFaceRequest
	GetExpression() *bool
	SetGender(v bool) *RecognizeFaceRequest
	GetGender() *bool
	SetGlass(v bool) *RecognizeFaceRequest
	GetGlass() *bool
	SetHat(v bool) *RecognizeFaceRequest
	GetHat() *bool
	SetImageURL(v string) *RecognizeFaceRequest
	GetImageURL() *string
	SetMask(v bool) *RecognizeFaceRequest
	GetMask() *bool
	SetMaxFaceNumber(v int64) *RecognizeFaceRequest
	GetMaxFaceNumber() *int64
	SetQuality(v bool) *RecognizeFaceRequest
	GetQuality() *bool
}

type RecognizeFaceRequest struct {
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
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

func (s RecognizeFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeFaceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFaceRequest) GetAge() *bool {
	return s.Age
}

func (s *RecognizeFaceRequest) GetBeauty() *bool {
	return s.Beauty
}

func (s *RecognizeFaceRequest) GetExpression() *bool {
	return s.Expression
}

func (s *RecognizeFaceRequest) GetGender() *bool {
	return s.Gender
}

func (s *RecognizeFaceRequest) GetGlass() *bool {
	return s.Glass
}

func (s *RecognizeFaceRequest) GetHat() *bool {
	return s.Hat
}

func (s *RecognizeFaceRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeFaceRequest) GetMask() *bool {
	return s.Mask
}

func (s *RecognizeFaceRequest) GetMaxFaceNumber() *int64 {
	return s.MaxFaceNumber
}

func (s *RecognizeFaceRequest) GetQuality() *bool {
	return s.Quality
}

func (s *RecognizeFaceRequest) SetAge(v bool) *RecognizeFaceRequest {
	s.Age = &v
	return s
}

func (s *RecognizeFaceRequest) SetBeauty(v bool) *RecognizeFaceRequest {
	s.Beauty = &v
	return s
}

func (s *RecognizeFaceRequest) SetExpression(v bool) *RecognizeFaceRequest {
	s.Expression = &v
	return s
}

func (s *RecognizeFaceRequest) SetGender(v bool) *RecognizeFaceRequest {
	s.Gender = &v
	return s
}

func (s *RecognizeFaceRequest) SetGlass(v bool) *RecognizeFaceRequest {
	s.Glass = &v
	return s
}

func (s *RecognizeFaceRequest) SetHat(v bool) *RecognizeFaceRequest {
	s.Hat = &v
	return s
}

func (s *RecognizeFaceRequest) SetImageURL(v string) *RecognizeFaceRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeFaceRequest) SetMask(v bool) *RecognizeFaceRequest {
	s.Mask = &v
	return s
}

func (s *RecognizeFaceRequest) SetMaxFaceNumber(v int64) *RecognizeFaceRequest {
	s.MaxFaceNumber = &v
	return s
}

func (s *RecognizeFaceRequest) SetQuality(v bool) *RecognizeFaceRequest {
	s.Quality = &v
	return s
}

func (s *RecognizeFaceRequest) Validate() error {
	return dara.Validate(s)
}
