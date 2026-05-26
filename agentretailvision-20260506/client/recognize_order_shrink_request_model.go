// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackUrl(v string) *RecognizeOrderShrinkRequest
	GetCallbackUrl() *string
	SetCandidateItemsShrink(v string) *RecognizeOrderShrinkRequest
	GetCandidateItemsShrink() *string
	SetDeviceId(v string) *RecognizeOrderShrinkRequest
	GetDeviceId() *string
	SetOrderUniqueId(v string) *RecognizeOrderShrinkRequest
	GetOrderUniqueId() *string
	SetVideoUrlsShrink(v string) *RecognizeOrderShrinkRequest
	GetVideoUrlsShrink() *string
}

type RecognizeOrderShrinkRequest struct {
	// example:
	//
	// https://example.com/callback
	CallbackUrl          *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	CandidateItemsShrink *string `json:"CandidateItems,omitempty" xml:"CandidateItems,omitempty"`
	// example:
	//
	// DEVICE_001
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// ORDER_001
	OrderUniqueId *string `json:"OrderUniqueId,omitempty" xml:"OrderUniqueId,omitempty"`
	// example:
	//
	// ["https://oss.example.com/video1.mp4"]
	VideoUrlsShrink *string `json:"VideoUrls,omitempty" xml:"VideoUrls,omitempty"`
}

func (s RecognizeOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *RecognizeOrderShrinkRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *RecognizeOrderShrinkRequest) GetCandidateItemsShrink() *string {
	return s.CandidateItemsShrink
}

func (s *RecognizeOrderShrinkRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *RecognizeOrderShrinkRequest) GetOrderUniqueId() *string {
	return s.OrderUniqueId
}

func (s *RecognizeOrderShrinkRequest) GetVideoUrlsShrink() *string {
	return s.VideoUrlsShrink
}

func (s *RecognizeOrderShrinkRequest) SetCallbackUrl(v string) *RecognizeOrderShrinkRequest {
	s.CallbackUrl = &v
	return s
}

func (s *RecognizeOrderShrinkRequest) SetCandidateItemsShrink(v string) *RecognizeOrderShrinkRequest {
	s.CandidateItemsShrink = &v
	return s
}

func (s *RecognizeOrderShrinkRequest) SetDeviceId(v string) *RecognizeOrderShrinkRequest {
	s.DeviceId = &v
	return s
}

func (s *RecognizeOrderShrinkRequest) SetOrderUniqueId(v string) *RecognizeOrderShrinkRequest {
	s.OrderUniqueId = &v
	return s
}

func (s *RecognizeOrderShrinkRequest) SetVideoUrlsShrink(v string) *RecognizeOrderShrinkRequest {
	s.VideoUrlsShrink = &v
	return s
}

func (s *RecognizeOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
