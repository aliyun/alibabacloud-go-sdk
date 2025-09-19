// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageScanNumInPeriodResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageScanData(v *GetImageScanNumInPeriodResponseBodyImageScanData) *GetImageScanNumInPeriodResponseBody
	GetImageScanData() *GetImageScanNumInPeriodResponseBodyImageScanData
	SetRequestId(v string) *GetImageScanNumInPeriodResponseBody
	GetRequestId() *string
}

type GetImageScanNumInPeriodResponseBody struct {
	// The data returned.
	ImageScanData *GetImageScanNumInPeriodResponseBodyImageScanData `json:"ImageScanData,omitempty" xml:"ImageScanData,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageScanNumInPeriodResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageScanNumInPeriodResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageScanNumInPeriodResponseBody) GetImageScanData() *GetImageScanNumInPeriodResponseBodyImageScanData {
	return s.ImageScanData
}

func (s *GetImageScanNumInPeriodResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageScanNumInPeriodResponseBody) SetImageScanData(v *GetImageScanNumInPeriodResponseBodyImageScanData) *GetImageScanNumInPeriodResponseBody {
	s.ImageScanData = v
	return s
}

func (s *GetImageScanNumInPeriodResponseBody) SetRequestId(v string) *GetImageScanNumInPeriodResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageScanNumInPeriodResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetImageScanNumInPeriodResponseBodyImageScanData struct {
	// The number of image scans.
	//
	// example:
	//
	// 150
	ImageScanCount *int32 `json:"ImageScanCount,omitempty" xml:"ImageScanCount,omitempty"`
}

func (s GetImageScanNumInPeriodResponseBodyImageScanData) String() string {
	return dara.Prettify(s)
}

func (s GetImageScanNumInPeriodResponseBodyImageScanData) GoString() string {
	return s.String()
}

func (s *GetImageScanNumInPeriodResponseBodyImageScanData) GetImageScanCount() *int32 {
	return s.ImageScanCount
}

func (s *GetImageScanNumInPeriodResponseBodyImageScanData) SetImageScanCount(v int32) *GetImageScanNumInPeriodResponseBodyImageScanData {
	s.ImageScanCount = &v
	return s
}

func (s *GetImageScanNumInPeriodResponseBodyImageScanData) Validate() error {
	return dara.Validate(s)
}
