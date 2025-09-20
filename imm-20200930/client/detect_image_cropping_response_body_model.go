// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageCroppingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCroppings(v []*CroppingSuggestion) *DetectImageCroppingResponseBody
	GetCroppings() []*CroppingSuggestion
	SetRequestId(v string) *DetectImageCroppingResponseBody
	GetRequestId() *string
}

type DetectImageCroppingResponseBody struct {
	// The image cropping suggestions.
	Croppings []*CroppingSuggestion `json:"Croppings,omitempty" xml:"Croppings,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 91AC8C98-0F36-49D2-8290-742E24D*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageCroppingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectImageCroppingResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageCroppingResponseBody) GetCroppings() []*CroppingSuggestion {
	return s.Croppings
}

func (s *DetectImageCroppingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectImageCroppingResponseBody) SetCroppings(v []*CroppingSuggestion) *DetectImageCroppingResponseBody {
	s.Croppings = v
	return s
}

func (s *DetectImageCroppingResponseBody) SetRequestId(v string) *DetectImageCroppingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectImageCroppingResponseBody) Validate() error {
	return dara.Validate(s)
}
