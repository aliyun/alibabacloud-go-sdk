// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterImageSecuritySummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterImageEvent(v *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent) *DescribeClusterImageSecuritySummaryResponseBody
	GetClusterImageEvent() *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent
	SetRequestId(v string) *DescribeClusterImageSecuritySummaryResponseBody
	GetRequestId() *string
}

type DescribeClusterImageSecuritySummaryResponseBody struct {
	// The information about the image-related security events.
	ClusterImageEvent *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent `json:"ClusterImageEvent,omitempty" xml:"ClusterImageEvent,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterImageSecuritySummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterImageSecuritySummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterImageSecuritySummaryResponseBody) GetClusterImageEvent() *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent {
	return s.ClusterImageEvent
}

func (s *DescribeClusterImageSecuritySummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterImageSecuritySummaryResponseBody) SetClusterImageEvent(v *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent) *DescribeClusterImageSecuritySummaryResponseBody {
	s.ClusterImageEvent = v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponseBody) SetRequestId(v string) *DescribeClusterImageSecuritySummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent struct {
	// The information about image baseline risks.
	ImageBaseline []*DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageBaseline `json:"ImageBaseline,omitempty" xml:"ImageBaseline,omitempty" type:"Repeated"`
	// The information about image system vulnerabilities.
	ImageCveVul []*DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageCveVul `json:"ImageCveVul,omitempty" xml:"ImageCveVul,omitempty" type:"Repeated"`
	// The information about malicious image samples.
	ImageMaliciousFile []*DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageMaliciousFile `json:"ImageMaliciousFile,omitempty" xml:"ImageMaliciousFile,omitempty" type:"Repeated"`
	// The information about image application vulnerabilities.
	ImageScaVul []*DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageScaVul `json:"ImageScaVul,omitempty" xml:"ImageScaVul,omitempty" type:"Repeated"`
}

func (s DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent) GoString() string {
	return s.String()
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent) GetImageBaseline() []*DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageBaseline {
	return s.ImageBaseline
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent) GetImageCveVul() []*DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageCveVul {
	return s.ImageCveVul
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent) GetImageMaliciousFile() []*DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageMaliciousFile {
	return s.ImageMaliciousFile
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent) GetImageScaVul() []*DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageScaVul {
	return s.ImageScaVul
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent) SetImageBaseline(v []*DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageBaseline) *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent {
	s.ImageBaseline = v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent) SetImageCveVul(v []*DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageCveVul) *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent {
	s.ImageCveVul = v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent) SetImageMaliciousFile(v []*DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageMaliciousFile) *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent {
	s.ImageMaliciousFile = v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent) SetImageScaVul(v []*DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageScaVul) *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent {
	s.ImageScaVul = v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEvent) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageBaseline struct {
	// The number of baselines.
	//
	// example:
	//
	// 0
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// medium
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageBaseline) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageBaseline) GoString() string {
	return s.String()
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageBaseline) GetCount() *int64 {
	return s.Count
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageBaseline) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageBaseline) SetCount(v int64) *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageBaseline {
	s.Count = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageBaseline) SetRiskLevel(v string) *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageBaseline {
	s.RiskLevel = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageBaseline) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageCveVul struct {
	// The number of vulnerabilities.
	//
	// example:
	//
	// 0
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The alert level. Valid values:
	//
	// 	- **asap**: high. You must fix the vulnerability at the earliest opportunity.
	//
	// 	- **nntf**: medium. You can fix the vulnerability based on your business requirements.
	//
	// 	- **later**: low. You can ignore the vulnerability.
	//
	// example:
	//
	// later
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageCveVul) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageCveVul) GoString() string {
	return s.String()
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageCveVul) GetCount() *int64 {
	return s.Count
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageCveVul) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageCveVul) SetCount(v int64) *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageCveVul {
	s.Count = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageCveVul) SetRiskLevel(v string) *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageCveVul {
	s.RiskLevel = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageCveVul) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageMaliciousFile struct {
	// The number of malicious samples.
	//
	// example:
	//
	// 0
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// medium
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageMaliciousFile) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageMaliciousFile) GoString() string {
	return s.String()
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageMaliciousFile) GetCount() *int64 {
	return s.Count
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageMaliciousFile) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageMaliciousFile) SetCount(v int64) *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageMaliciousFile {
	s.Count = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageMaliciousFile) SetRiskLevel(v string) *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageMaliciousFile {
	s.RiskLevel = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageMaliciousFile) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageScaVul struct {
	// The number of image application vulnerabilities.
	//
	// example:
	//
	// 0
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The alert level. Valid values:
	//
	// 	- **asap**: high. You must fix the vulnerability at the earliest opportunity.
	//
	// 	- **nntf**: medium. You can fix the vulnerability based on your business requirements.
	//
	// 	- **later**: low. You can ignore the vulnerability.
	//
	// example:
	//
	// later
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageScaVul) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageScaVul) GoString() string {
	return s.String()
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageScaVul) GetCount() *int64 {
	return s.Count
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageScaVul) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageScaVul) SetCount(v int64) *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageScaVul {
	s.Count = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageScaVul) SetRiskLevel(v string) *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageScaVul {
	s.RiskLevel = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponseBodyClusterImageEventImageScaVul) Validate() error {
	return dara.Validate(s)
}
