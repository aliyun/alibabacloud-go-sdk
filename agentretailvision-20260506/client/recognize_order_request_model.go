// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackUrl(v string) *RecognizeOrderRequest
	GetCallbackUrl() *string
	SetCandidateItems(v []*string) *RecognizeOrderRequest
	GetCandidateItems() []*string
	SetDeviceId(v string) *RecognizeOrderRequest
	GetDeviceId() *string
	SetOrderUniqueId(v string) *RecognizeOrderRequest
	GetOrderUniqueId() *string
	SetVideoUrls(v []*string) *RecognizeOrderRequest
	GetVideoUrls() []*string
}

type RecognizeOrderRequest struct {
	// example:
	//
	// https://example.com/callback
	CallbackUrl    *string   `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	CandidateItems []*string `json:"CandidateItems,omitempty" xml:"CandidateItems,omitempty" type:"Repeated"`
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
	VideoUrls []*string `json:"VideoUrls,omitempty" xml:"VideoUrls,omitempty" type:"Repeated"`
}

func (s RecognizeOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeOrderRequest) GoString() string {
	return s.String()
}

func (s *RecognizeOrderRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *RecognizeOrderRequest) GetCandidateItems() []*string {
	return s.CandidateItems
}

func (s *RecognizeOrderRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *RecognizeOrderRequest) GetOrderUniqueId() *string {
	return s.OrderUniqueId
}

func (s *RecognizeOrderRequest) GetVideoUrls() []*string {
	return s.VideoUrls
}

func (s *RecognizeOrderRequest) SetCallbackUrl(v string) *RecognizeOrderRequest {
	s.CallbackUrl = &v
	return s
}

func (s *RecognizeOrderRequest) SetCandidateItems(v []*string) *RecognizeOrderRequest {
	s.CandidateItems = v
	return s
}

func (s *RecognizeOrderRequest) SetDeviceId(v string) *RecognizeOrderRequest {
	s.DeviceId = &v
	return s
}

func (s *RecognizeOrderRequest) SetOrderUniqueId(v string) *RecognizeOrderRequest {
	s.OrderUniqueId = &v
	return s
}

func (s *RecognizeOrderRequest) SetVideoUrls(v []*string) *RecognizeOrderRequest {
	s.VideoUrls = v
	return s
}

func (s *RecognizeOrderRequest) Validate() error {
	return dara.Validate(s)
}
