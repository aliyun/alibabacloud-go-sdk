// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageScanAuthCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageScan(v *DescribeImageScanAuthCountResponseBodyImageScan) *DescribeImageScanAuthCountResponseBody
	GetImageScan() *DescribeImageScanAuthCountResponseBodyImageScan
	SetRequestId(v string) *DescribeImageScanAuthCountResponseBody
	GetRequestId() *string
}

type DescribeImageScanAuthCountResponseBody struct {
	// The details about the quota for container image scan.
	ImageScan *DescribeImageScanAuthCountResponseBodyImageScan `json:"ImageScan,omitempty" xml:"ImageScan,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 892NYH839-0EDC-4CD0-A2EF-5BD294656C99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageScanAuthCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageScanAuthCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageScanAuthCountResponseBody) GetImageScan() *DescribeImageScanAuthCountResponseBodyImageScan {
	return s.ImageScan
}

func (s *DescribeImageScanAuthCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageScanAuthCountResponseBody) SetImageScan(v *DescribeImageScanAuthCountResponseBodyImageScan) *DescribeImageScanAuthCountResponseBody {
	s.ImageScan = v
	return s
}

func (s *DescribeImageScanAuthCountResponseBody) SetRequestId(v string) *DescribeImageScanAuthCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageScanAuthCountResponseBody) Validate() error {
	if s.ImageScan != nil {
		if err := s.ImageScan.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageScanAuthCountResponseBodyImageScan struct {
	// The quota for container image scan.
	//
	// example:
	//
	// 15340
	ImageScanCapacity *int64 `json:"ImageScanCapacity,omitempty" xml:"ImageScanCapacity,omitempty"`
	// The instance ID of Security Center.
	//
	// example:
	//
	// sas-qdl123412****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The consumed quota for container image scan.
	//
	// example:
	//
	// 5489
	ScanCount *int64 `json:"ScanCount,omitempty" xml:"ScanCount,omitempty"`
}

func (s DescribeImageScanAuthCountResponseBodyImageScan) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageScanAuthCountResponseBodyImageScan) GoString() string {
	return s.String()
}

func (s *DescribeImageScanAuthCountResponseBodyImageScan) GetImageScanCapacity() *int64 {
	return s.ImageScanCapacity
}

func (s *DescribeImageScanAuthCountResponseBodyImageScan) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeImageScanAuthCountResponseBodyImageScan) GetScanCount() *int64 {
	return s.ScanCount
}

func (s *DescribeImageScanAuthCountResponseBodyImageScan) SetImageScanCapacity(v int64) *DescribeImageScanAuthCountResponseBodyImageScan {
	s.ImageScanCapacity = &v
	return s
}

func (s *DescribeImageScanAuthCountResponseBodyImageScan) SetInstanceId(v string) *DescribeImageScanAuthCountResponseBodyImageScan {
	s.InstanceId = &v
	return s
}

func (s *DescribeImageScanAuthCountResponseBodyImageScan) SetScanCount(v int64) *DescribeImageScanAuthCountResponseBodyImageScan {
	s.ScanCount = &v
	return s
}

func (s *DescribeImageScanAuthCountResponseBodyImageScan) Validate() error {
	return dara.Validate(s)
}
