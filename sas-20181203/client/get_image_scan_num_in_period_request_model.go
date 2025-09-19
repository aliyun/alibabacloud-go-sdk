// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageScanNumInPeriodRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPastDay(v string) *GetImageScanNumInPeriodRequest
	GetPastDay() *string
}

type GetImageScanNumInPeriodRequest struct {
	// The number of days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	PastDay *string `json:"PastDay,omitempty" xml:"PastDay,omitempty"`
}

func (s GetImageScanNumInPeriodRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageScanNumInPeriodRequest) GoString() string {
	return s.String()
}

func (s *GetImageScanNumInPeriodRequest) GetPastDay() *string {
	return s.PastDay
}

func (s *GetImageScanNumInPeriodRequest) SetPastDay(v string) *GetImageScanNumInPeriodRequest {
	s.PastDay = &v
	return s
}

func (s *GetImageScanNumInPeriodRequest) Validate() error {
	return dara.Validate(s)
}
