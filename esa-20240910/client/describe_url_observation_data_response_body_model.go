// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUrlObservationDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeUrlObservationDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeUrlObservationDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeUrlObservationDataResponseBody
	GetStartTime() *string
	SetUrlDetailData(v []*DescribeUrlObservationDataResponseBodyUrlDetailData) *DescribeUrlObservationDataResponseBody
	GetUrlDetailData() []*DescribeUrlObservationDataResponseBodyUrlDetailData
}

type DescribeUrlObservationDataResponseBody struct {
	// The end of the time range during which data was queried.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The time must be in UTC.
	//
	// example:
	//
	// 2023-04-19T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The create time. The time is in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2022-11-06T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The objects that are returned.
	UrlDetailData []*DescribeUrlObservationDataResponseBodyUrlDetailData `json:"UrlDetailData,omitempty" xml:"UrlDetailData,omitempty" type:"Repeated"`
}

func (s DescribeUrlObservationDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUrlObservationDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUrlObservationDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeUrlObservationDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUrlObservationDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeUrlObservationDataResponseBody) GetUrlDetailData() []*DescribeUrlObservationDataResponseBodyUrlDetailData {
	return s.UrlDetailData
}

func (s *DescribeUrlObservationDataResponseBody) SetEndTime(v string) *DescribeUrlObservationDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeUrlObservationDataResponseBody) SetRequestId(v string) *DescribeUrlObservationDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUrlObservationDataResponseBody) SetStartTime(v string) *DescribeUrlObservationDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeUrlObservationDataResponseBody) SetUrlDetailData(v []*DescribeUrlObservationDataResponseBodyUrlDetailData) *DescribeUrlObservationDataResponseBody {
	s.UrlDetailData = v
	return s
}

func (s *DescribeUrlObservationDataResponseBody) Validate() error {
	if s.UrlDetailData != nil {
		for _, item := range s.UrlDetailData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUrlObservationDataResponseBodyUrlDetailData struct {
	// Measures the maximum layout mutation score for every unexpected layout change that occurs throughout the life of the page.
	//
	// example:
	//
	// 0.5
	CLS *float32 `json:"CLS,omitempty" xml:"CLS,omitempty"`
	// The platform of the device.
	//
	// example:
	//
	// PC
	ClientPlatform *string `json:"ClientPlatform,omitempty" xml:"ClientPlatform,omitempty"`
	// The country or region to which the IP address belongs.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// Measures the time between when the page is loaded and when any part of the page\\"s content is rendered on the screen. Unit: ms.
	//
	// example:
	//
	// 123
	FCP *float32 `json:"FCP,omitempty" xml:"FCP,omitempty"`
	// Measures the time between when the user first interacts with the page and when the browser is actually able to start processing an event handler in response to that interaction. Unit: ms.
	//
	// example:
	//
	// 123
	FID *float32 `json:"FID,omitempty" xml:"FID,omitempty"`
	// Measures the responsiveness of the page, or how long it takes for the page to respond to user input visibly. Unit: ms.
	//
	// example:
	//
	// 123
	INP *float32 `json:"INP,omitempty" xml:"INP,omitempty"`
	// Reports the rendering time of the largest image or text block visible in the viewport. Unit: ms.
	//
	// example:
	//
	// 123
	LCP *float32 `json:"LCP,omitempty" xml:"LCP,omitempty"`
	// This metric measures the time between when a resource initiates a request and when the first byte of the response starts to arrive. Unit: ms.
	//
	// example:
	//
	// 123
	TTFB *float32 `json:"TTFB,omitempty" xml:"TTFB,omitempty"`
	// The URL of the web page to monitor.
	//
	// example:
	//
	// example.com/test
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeUrlObservationDataResponseBodyUrlDetailData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUrlObservationDataResponseBodyUrlDetailData) GoString() string {
	return s.String()
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) GetCLS() *float32 {
	return s.CLS
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) GetClientPlatform() *string {
	return s.ClientPlatform
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) GetCountry() *string {
	return s.Country
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) GetFCP() *float32 {
	return s.FCP
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) GetFID() *float32 {
	return s.FID
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) GetINP() *float32 {
	return s.INP
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) GetLCP() *float32 {
	return s.LCP
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) GetTTFB() *float32 {
	return s.TTFB
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) GetUrl() *string {
	return s.Url
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) SetCLS(v float32) *DescribeUrlObservationDataResponseBodyUrlDetailData {
	s.CLS = &v
	return s
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) SetClientPlatform(v string) *DescribeUrlObservationDataResponseBodyUrlDetailData {
	s.ClientPlatform = &v
	return s
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) SetCountry(v string) *DescribeUrlObservationDataResponseBodyUrlDetailData {
	s.Country = &v
	return s
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) SetFCP(v float32) *DescribeUrlObservationDataResponseBodyUrlDetailData {
	s.FCP = &v
	return s
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) SetFID(v float32) *DescribeUrlObservationDataResponseBodyUrlDetailData {
	s.FID = &v
	return s
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) SetINP(v float32) *DescribeUrlObservationDataResponseBodyUrlDetailData {
	s.INP = &v
	return s
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) SetLCP(v float32) *DescribeUrlObservationDataResponseBodyUrlDetailData {
	s.LCP = &v
	return s
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) SetTTFB(v float32) *DescribeUrlObservationDataResponseBodyUrlDetailData {
	s.TTFB = &v
	return s
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) SetUrl(v string) *DescribeUrlObservationDataResponseBodyUrlDetailData {
	s.Url = &v
	return s
}

func (s *DescribeUrlObservationDataResponseBodyUrlDetailData) Validate() error {
	return dara.Validate(s)
}
