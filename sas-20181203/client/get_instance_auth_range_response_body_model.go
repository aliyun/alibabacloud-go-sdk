// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceAuthRangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceAuthRange(v *GetInstanceAuthRangeResponseBodyInstanceAuthRange) *GetInstanceAuthRangeResponseBody
	GetInstanceAuthRange() *GetInstanceAuthRangeResponseBodyInstanceAuthRange
	SetRequestId(v string) *GetInstanceAuthRangeResponseBody
	GetRequestId() *string
}

type GetInstanceAuthRangeResponseBody struct {
	InstanceAuthRange *GetInstanceAuthRangeResponseBodyInstanceAuthRange `json:"InstanceAuthRange,omitempty" xml:"InstanceAuthRange,omitempty" type:"Struct"`
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceAuthRangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAuthRangeResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceAuthRangeResponseBody) GetInstanceAuthRange() *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	return s.InstanceAuthRange
}

func (s *GetInstanceAuthRangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceAuthRangeResponseBody) SetInstanceAuthRange(v *GetInstanceAuthRangeResponseBodyInstanceAuthRange) *GetInstanceAuthRangeResponseBody {
	s.InstanceAuthRange = v
	return s
}

func (s *GetInstanceAuthRangeResponseBody) SetRequestId(v string) *GetInstanceAuthRangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBody) Validate() error {
	if s.InstanceAuthRange != nil {
		if err := s.InstanceAuthRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceAuthRangeResponseBodyInstanceAuthRange struct {
	// example:
	//
	// 1-2000000000:1
	AdvancedCount *string `json:"AdvancedCount,omitempty" xml:"AdvancedCount,omitempty"`
	// example:
	//
	// 0-9000000000:10
	AntiRansomwareCapacity *string `json:"AntiRansomwareCapacity,omitempty" xml:"AntiRansomwareCapacity,omitempty"`
	// example:
	//
	// 1
	AntiRansomwareService *int32 `json:"AntiRansomwareService,omitempty" xml:"AntiRansomwareService,omitempty"`
	// example:
	//
	// 1-2000000000:1
	AntiVirusCore *string `json:"AntiVirusCore,omitempty" xml:"AntiVirusCore,omitempty"`
	// example:
	//
	// 1-2000000000:1
	ContainerCore *string `json:"ContainerCore,omitempty" xml:"ContainerCore,omitempty"`
	// example:
	//
	// 1-2000000000:1
	ContainerCount *string `json:"ContainerCount,omitempty" xml:"ContainerCount,omitempty"`
	// example:
	//
	// 15000-9999999999:55000
	CspmCapacity *string `json:"CspmCapacity,omitempty" xml:"CspmCapacity,omitempty"`
	// example:
	//
	// 1-2000000000:1
	EnterpriseCount *string `json:"EnterpriseCount,omitempty" xml:"EnterpriseCount,omitempty"`
	// example:
	//
	// 20-500:1
	HoneypotCapacity *string `json:"HoneypotCapacity,omitempty" xml:"HoneypotCapacity,omitempty"`
	// example:
	//
	// 0-200000:20
	ImageScanCapacity *string `json:"ImageScanCapacity,omitempty" xml:"ImageScanCapacity,omitempty"`
	// example:
	//
	// 0-100000000:1
	RaspCapacity *string `json:"RaspCapacity,omitempty" xml:"RaspCapacity,omitempty"`
	// example:
	//
	// 10-9999999999:10
	SdkCapacity *string `json:"SdkCapacity,omitempty" xml:"SdkCapacity,omitempty"`
	// example:
	//
	// 0-600000000:10
	SlsCapacity *string `json:"SlsCapacity,omitempty" xml:"SlsCapacity,omitempty"`
	// example:
	//
	// 0-9999999999:1000
	ThreatAnalysisCapacity *string `json:"ThreatAnalysisCapacity,omitempty" xml:"ThreatAnalysisCapacity,omitempty"`
	// example:
	//
	// 0-9999999999:100
	ThreatAnalysisFlow *string `json:"ThreatAnalysisFlow,omitempty" xml:"ThreatAnalysisFlow,omitempty"`
	// example:
	//
	// 0-9999:1
	WebLockCapacity *string `json:"WebLockCapacity,omitempty" xml:"WebLockCapacity,omitempty"`
}

func (s GetInstanceAuthRangeResponseBodyInstanceAuthRange) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAuthRangeResponseBodyInstanceAuthRange) GoString() string {
	return s.String()
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) GetAdvancedCount() *string {
	return s.AdvancedCount
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) GetAntiRansomwareCapacity() *string {
	return s.AntiRansomwareCapacity
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) GetAntiRansomwareService() *int32 {
	return s.AntiRansomwareService
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) GetAntiVirusCore() *string {
	return s.AntiVirusCore
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) GetContainerCore() *string {
	return s.ContainerCore
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) GetContainerCount() *string {
	return s.ContainerCount
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) GetCspmCapacity() *string {
	return s.CspmCapacity
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) GetEnterpriseCount() *string {
	return s.EnterpriseCount
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) GetHoneypotCapacity() *string {
	return s.HoneypotCapacity
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) GetImageScanCapacity() *string {
	return s.ImageScanCapacity
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) GetRaspCapacity() *string {
	return s.RaspCapacity
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) GetSdkCapacity() *string {
	return s.SdkCapacity
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) GetSlsCapacity() *string {
	return s.SlsCapacity
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) GetThreatAnalysisCapacity() *string {
	return s.ThreatAnalysisCapacity
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) GetThreatAnalysisFlow() *string {
	return s.ThreatAnalysisFlow
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) GetWebLockCapacity() *string {
	return s.WebLockCapacity
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) SetAdvancedCount(v string) *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	s.AdvancedCount = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) SetAntiRansomwareCapacity(v string) *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	s.AntiRansomwareCapacity = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) SetAntiRansomwareService(v int32) *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	s.AntiRansomwareService = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) SetAntiVirusCore(v string) *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	s.AntiVirusCore = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) SetContainerCore(v string) *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	s.ContainerCore = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) SetContainerCount(v string) *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	s.ContainerCount = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) SetCspmCapacity(v string) *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	s.CspmCapacity = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) SetEnterpriseCount(v string) *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	s.EnterpriseCount = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) SetHoneypotCapacity(v string) *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	s.HoneypotCapacity = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) SetImageScanCapacity(v string) *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	s.ImageScanCapacity = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) SetRaspCapacity(v string) *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	s.RaspCapacity = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) SetSdkCapacity(v string) *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	s.SdkCapacity = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) SetSlsCapacity(v string) *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	s.SlsCapacity = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) SetThreatAnalysisCapacity(v string) *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	s.ThreatAnalysisCapacity = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) SetThreatAnalysisFlow(v string) *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	s.ThreatAnalysisFlow = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) SetWebLockCapacity(v string) *GetInstanceAuthRangeResponseBodyInstanceAuthRange {
	s.WebLockCapacity = &v
	return s
}

func (s *GetInstanceAuthRangeResponseBodyInstanceAuthRange) Validate() error {
	return dara.Validate(s)
}
