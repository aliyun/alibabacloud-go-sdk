// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSecurityScanCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeImageSecurityScanCountResponseBodyData) *DescribeImageSecurityScanCountResponseBody
	GetData() *DescribeImageSecurityScanCountResponseBodyData
	SetRequestId(v string) *DescribeImageSecurityScanCountResponseBody
	GetRequestId() *string
}

type DescribeImageSecurityScanCountResponseBody struct {
	// The data returned.
	Data *DescribeImageSecurityScanCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C699E4E4-F2F4-58FC-A949-457FFE59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageSecurityScanCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSecurityScanCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageSecurityScanCountResponseBody) GetData() *DescribeImageSecurityScanCountResponseBodyData {
	return s.Data
}

func (s *DescribeImageSecurityScanCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageSecurityScanCountResponseBody) SetData(v *DescribeImageSecurityScanCountResponseBodyData) *DescribeImageSecurityScanCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImageSecurityScanCountResponseBody) SetRequestId(v string) *DescribeImageSecurityScanCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageSecurityScanCountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageSecurityScanCountResponseBodyData struct {
	// The number of image baseline risks detected on the current asset.
	//
	// example:
	//
	// 0
	ImageBaselineCount *int32 `json:"ImageBaselineCount,omitempty" xml:"ImageBaselineCount,omitempty"`
	// The number of image system vulnerabilities returned on the current page.
	//
	// example:
	//
	// 0
	ImageCveVulCount *int32 `json:"ImageCveVulCount,omitempty" xml:"ImageCveVulCount,omitempty"`
	// The number of malicious image samples returned on the current page.
	//
	// example:
	//
	// 0
	ImageMaliciousFileCount *int32 `json:"ImageMaliciousFileCount,omitempty" xml:"ImageMaliciousFileCount,omitempty"`
	// The number of image application vulnerabilities returned on the current page.
	//
	// example:
	//
	// 0
	ImageScaVulCount *int32 `json:"ImageScaVulCount,omitempty" xml:"ImageScaVulCount,omitempty"`
}

func (s DescribeImageSecurityScanCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSecurityScanCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImageSecurityScanCountResponseBodyData) GetImageBaselineCount() *int32 {
	return s.ImageBaselineCount
}

func (s *DescribeImageSecurityScanCountResponseBodyData) GetImageCveVulCount() *int32 {
	return s.ImageCveVulCount
}

func (s *DescribeImageSecurityScanCountResponseBodyData) GetImageMaliciousFileCount() *int32 {
	return s.ImageMaliciousFileCount
}

func (s *DescribeImageSecurityScanCountResponseBodyData) GetImageScaVulCount() *int32 {
	return s.ImageScaVulCount
}

func (s *DescribeImageSecurityScanCountResponseBodyData) SetImageBaselineCount(v int32) *DescribeImageSecurityScanCountResponseBodyData {
	s.ImageBaselineCount = &v
	return s
}

func (s *DescribeImageSecurityScanCountResponseBodyData) SetImageCveVulCount(v int32) *DescribeImageSecurityScanCountResponseBodyData {
	s.ImageCveVulCount = &v
	return s
}

func (s *DescribeImageSecurityScanCountResponseBodyData) SetImageMaliciousFileCount(v int32) *DescribeImageSecurityScanCountResponseBodyData {
	s.ImageMaliciousFileCount = &v
	return s
}

func (s *DescribeImageSecurityScanCountResponseBodyData) SetImageScaVulCount(v int32) *DescribeImageSecurityScanCountResponseBodyData {
	s.ImageScaVulCount = &v
	return s
}

func (s *DescribeImageSecurityScanCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
