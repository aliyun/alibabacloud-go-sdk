// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageScoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageScore(v *DetectImageScoreResponseBodyImageScore) *DetectImageScoreResponseBody
	GetImageScore() *DetectImageScoreResponseBodyImageScore
	SetRequestId(v string) *DetectImageScoreResponseBody
	GetRequestId() *string
}

type DetectImageScoreResponseBody struct {
	// The quality score of the image.
	ImageScore *DetectImageScoreResponseBodyImageScore `json:"ImageScore,omitempty" xml:"ImageScore,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6E93D6C9-5AC0-49F9-914D-E02678D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageScoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectImageScoreResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageScoreResponseBody) GetImageScore() *DetectImageScoreResponseBodyImageScore {
	return s.ImageScore
}

func (s *DetectImageScoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectImageScoreResponseBody) SetImageScore(v *DetectImageScoreResponseBodyImageScore) *DetectImageScoreResponseBody {
	s.ImageScore = v
	return s
}

func (s *DetectImageScoreResponseBody) SetRequestId(v string) *DetectImageScoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectImageScoreResponseBody) Validate() error {
	return dara.Validate(s)
}

type DetectImageScoreResponseBodyImageScore struct {
	// The overall quality score.
	//
	// example:
	//
	// 0.6
	OverallQualityScore *float32 `json:"OverallQualityScore,omitempty" xml:"OverallQualityScore,omitempty"`
}

func (s DetectImageScoreResponseBodyImageScore) String() string {
	return dara.Prettify(s)
}

func (s DetectImageScoreResponseBodyImageScore) GoString() string {
	return s.String()
}

func (s *DetectImageScoreResponseBodyImageScore) GetOverallQualityScore() *float32 {
	return s.OverallQualityScore
}

func (s *DetectImageScoreResponseBodyImageScore) SetOverallQualityScore(v float32) *DetectImageScoreResponseBodyImageScore {
	s.OverallQualityScore = &v
	return s
}

func (s *DetectImageScoreResponseBodyImageScore) Validate() error {
	return dara.Validate(s)
}
