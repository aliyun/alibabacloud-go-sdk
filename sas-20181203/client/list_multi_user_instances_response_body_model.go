// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiUserInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDaInstance(v *ListMultiUserInstancesResponseBodyDaInstance) *ListMultiUserInstancesResponseBody
	GetDaInstance() *ListMultiUserInstancesResponseBodyDaInstance
	SetPageInfo(v *ListMultiUserInstancesResponseBodyPageInfo) *ListMultiUserInstancesResponseBody
	GetPageInfo() *ListMultiUserInstancesResponseBodyPageInfo
	SetRequestId(v string) *ListMultiUserInstancesResponseBody
	GetRequestId() *string
	SetSaleInstanceList(v []*ListMultiUserInstancesResponseBodySaleInstanceList) *ListMultiUserInstancesResponseBody
	GetSaleInstanceList() []*ListMultiUserInstancesResponseBodySaleInstanceList
}

type ListMultiUserInstancesResponseBody struct {
	DaInstance *ListMultiUserInstancesResponseBodyDaInstance `json:"DaInstance,omitempty" xml:"DaInstance,omitempty" type:"Struct"`
	PageInfo   *ListMultiUserInstancesResponseBodyPageInfo   `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// example:
	//
	// 88F2A6CD-E500-5038-B992-0107B99AA88C
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SaleInstanceList []*ListMultiUserInstancesResponseBodySaleInstanceList `json:"SaleInstanceList,omitempty" xml:"SaleInstanceList,omitempty" type:"Repeated"`
}

func (s ListMultiUserInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBody) GetDaInstance() *ListMultiUserInstancesResponseBodyDaInstance {
	return s.DaInstance
}

func (s *ListMultiUserInstancesResponseBody) GetPageInfo() *ListMultiUserInstancesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListMultiUserInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMultiUserInstancesResponseBody) GetSaleInstanceList() []*ListMultiUserInstancesResponseBodySaleInstanceList {
	return s.SaleInstanceList
}

func (s *ListMultiUserInstancesResponseBody) SetDaInstance(v *ListMultiUserInstancesResponseBodyDaInstance) *ListMultiUserInstancesResponseBody {
	s.DaInstance = v
	return s
}

func (s *ListMultiUserInstancesResponseBody) SetPageInfo(v *ListMultiUserInstancesResponseBodyPageInfo) *ListMultiUserInstancesResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListMultiUserInstancesResponseBody) SetRequestId(v string) *ListMultiUserInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiUserInstancesResponseBody) SetSaleInstanceList(v []*ListMultiUserInstancesResponseBodySaleInstanceList) *ListMultiUserInstancesResponseBody {
	s.SaleInstanceList = v
	return s
}

func (s *ListMultiUserInstancesResponseBody) Validate() error {
	if s.DaInstance != nil {
		if err := s.DaInstance.Validate(); err != nil {
			return err
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SaleInstanceList != nil {
		for _, item := range s.SaleInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMultiUserInstancesResponseBodyDaInstance struct {
	// example:
	//
	// 1766185894104675
	AliUid                 *int64                                                              `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AntiRansomwareCapacity *ListMultiUserInstancesResponseBodyDaInstanceAntiRansomwareCapacity `json:"AntiRansomwareCapacity,omitempty" xml:"AntiRansomwareCapacity,omitempty" type:"Struct"`
	CspmCapacity           *ListMultiUserInstancesResponseBodyDaInstanceCspmCapacity           `json:"CspmCapacity,omitempty" xml:"CspmCapacity,omitempty" type:"Struct"`
	HoneypotCapacity       *ListMultiUserInstancesResponseBodyDaInstanceHoneypotCapacity       `json:"HoneypotCapacity,omitempty" xml:"HoneypotCapacity,omitempty" type:"Struct"`
	ImageScanCapacity      *ListMultiUserInstancesResponseBodyDaInstanceImageScanCapacity      `json:"ImageScanCapacity,omitempty" xml:"ImageScanCapacity,omitempty" type:"Struct"`
	// example:
	//
	// i-bp1gmm4pnacse343nqal
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 0
	InstancePurchaseType *int32                                                    `json:"InstancePurchaseType,omitempty" xml:"InstancePurchaseType,omitempty"`
	RaspCapacity         *ListMultiUserInstancesResponseBodyDaInstanceRaspCapacity `json:"RaspCapacity,omitempty" xml:"RaspCapacity,omitempty" type:"Struct"`
	SdkCapacity          *ListMultiUserInstancesResponseBodyDaInstanceSdkCapacity  `json:"SdkCapacity,omitempty" xml:"SdkCapacity,omitempty" type:"Struct"`
	SlsCapacity          *ListMultiUserInstancesResponseBodyDaInstanceSlsCapacity  `json:"SlsCapacity,omitempty" xml:"SlsCapacity,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Status                 *int32                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	ThreatAnalysisCapacity *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisCapacity `json:"ThreatAnalysisCapacity,omitempty" xml:"ThreatAnalysisCapacity,omitempty" type:"Struct"`
	ThreatAnalysisFlow     *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisFlow     `json:"ThreatAnalysisFlow,omitempty" xml:"ThreatAnalysisFlow,omitempty" type:"Struct"`
	// example:
	//
	// 1
	UserType *int32 `json:"UserType,omitempty" xml:"UserType,omitempty"`
	// example:
	//
	// 3
	Version         *int32                                                        `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionSummary  []*ListMultiUserInstancesResponseBodyDaInstanceVersionSummary `json:"VersionSummary,omitempty" xml:"VersionSummary,omitempty" type:"Repeated"`
	WebLockCapacity *ListMultiUserInstancesResponseBodyDaInstanceWebLockCapacity  `json:"WebLockCapacity,omitempty" xml:"WebLockCapacity,omitempty" type:"Struct"`
}

func (s ListMultiUserInstancesResponseBodyDaInstance) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodyDaInstance) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetAntiRansomwareCapacity() *ListMultiUserInstancesResponseBodyDaInstanceAntiRansomwareCapacity {
	return s.AntiRansomwareCapacity
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetCspmCapacity() *ListMultiUserInstancesResponseBodyDaInstanceCspmCapacity {
	return s.CspmCapacity
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetHoneypotCapacity() *ListMultiUserInstancesResponseBodyDaInstanceHoneypotCapacity {
	return s.HoneypotCapacity
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetImageScanCapacity() *ListMultiUserInstancesResponseBodyDaInstanceImageScanCapacity {
	return s.ImageScanCapacity
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetInstancePurchaseType() *int32 {
	return s.InstancePurchaseType
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetRaspCapacity() *ListMultiUserInstancesResponseBodyDaInstanceRaspCapacity {
	return s.RaspCapacity
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetSdkCapacity() *ListMultiUserInstancesResponseBodyDaInstanceSdkCapacity {
	return s.SdkCapacity
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetSlsCapacity() *ListMultiUserInstancesResponseBodyDaInstanceSlsCapacity {
	return s.SlsCapacity
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetStatus() *int32 {
	return s.Status
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetThreatAnalysisCapacity() *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisCapacity {
	return s.ThreatAnalysisCapacity
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetThreatAnalysisFlow() *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisFlow {
	return s.ThreatAnalysisFlow
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetUserType() *int32 {
	return s.UserType
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetVersion() *int32 {
	return s.Version
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetVersionSummary() []*ListMultiUserInstancesResponseBodyDaInstanceVersionSummary {
	return s.VersionSummary
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) GetWebLockCapacity() *ListMultiUserInstancesResponseBodyDaInstanceWebLockCapacity {
	return s.WebLockCapacity
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetAliUid(v int64) *ListMultiUserInstancesResponseBodyDaInstance {
	s.AliUid = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetAntiRansomwareCapacity(v *ListMultiUserInstancesResponseBodyDaInstanceAntiRansomwareCapacity) *ListMultiUserInstancesResponseBodyDaInstance {
	s.AntiRansomwareCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetCspmCapacity(v *ListMultiUserInstancesResponseBodyDaInstanceCspmCapacity) *ListMultiUserInstancesResponseBodyDaInstance {
	s.CspmCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetHoneypotCapacity(v *ListMultiUserInstancesResponseBodyDaInstanceHoneypotCapacity) *ListMultiUserInstancesResponseBodyDaInstance {
	s.HoneypotCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetImageScanCapacity(v *ListMultiUserInstancesResponseBodyDaInstanceImageScanCapacity) *ListMultiUserInstancesResponseBodyDaInstance {
	s.ImageScanCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetInstanceId(v string) *ListMultiUserInstancesResponseBodyDaInstance {
	s.InstanceId = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetInstancePurchaseType(v int32) *ListMultiUserInstancesResponseBodyDaInstance {
	s.InstancePurchaseType = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetRaspCapacity(v *ListMultiUserInstancesResponseBodyDaInstanceRaspCapacity) *ListMultiUserInstancesResponseBodyDaInstance {
	s.RaspCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetSdkCapacity(v *ListMultiUserInstancesResponseBodyDaInstanceSdkCapacity) *ListMultiUserInstancesResponseBodyDaInstance {
	s.SdkCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetSlsCapacity(v *ListMultiUserInstancesResponseBodyDaInstanceSlsCapacity) *ListMultiUserInstancesResponseBodyDaInstance {
	s.SlsCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetStatus(v int32) *ListMultiUserInstancesResponseBodyDaInstance {
	s.Status = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetThreatAnalysisCapacity(v *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisCapacity) *ListMultiUserInstancesResponseBodyDaInstance {
	s.ThreatAnalysisCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetThreatAnalysisFlow(v *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisFlow) *ListMultiUserInstancesResponseBodyDaInstance {
	s.ThreatAnalysisFlow = v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetUserType(v int32) *ListMultiUserInstancesResponseBodyDaInstance {
	s.UserType = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetVersion(v int32) *ListMultiUserInstancesResponseBodyDaInstance {
	s.Version = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetVersionSummary(v []*ListMultiUserInstancesResponseBodyDaInstanceVersionSummary) *ListMultiUserInstancesResponseBodyDaInstance {
	s.VersionSummary = v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) SetWebLockCapacity(v *ListMultiUserInstancesResponseBodyDaInstanceWebLockCapacity) *ListMultiUserInstancesResponseBodyDaInstance {
	s.WebLockCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstance) Validate() error {
	if s.AntiRansomwareCapacity != nil {
		if err := s.AntiRansomwareCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.CspmCapacity != nil {
		if err := s.CspmCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.HoneypotCapacity != nil {
		if err := s.HoneypotCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.ImageScanCapacity != nil {
		if err := s.ImageScanCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.RaspCapacity != nil {
		if err := s.RaspCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.SdkCapacity != nil {
		if err := s.SdkCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.SlsCapacity != nil {
		if err := s.SlsCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.ThreatAnalysisCapacity != nil {
		if err := s.ThreatAnalysisCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.ThreatAnalysisFlow != nil {
		if err := s.ThreatAnalysisFlow.Validate(); err != nil {
			return err
		}
	}
	if s.VersionSummary != nil {
		for _, item := range s.VersionSummary {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WebLockCapacity != nil {
		if err := s.WebLockCapacity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMultiUserInstancesResponseBodyDaInstanceAntiRansomwareCapacity struct {
	// example:
	//
	// 0
	Assigned *int64 `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	// example:
	//
	// 14
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 2
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodyDaInstanceAntiRansomwareCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodyDaInstanceAntiRansomwareCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceAntiRansomwareCapacity) GetAssigned() *int64 {
	return s.Assigned
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceAntiRansomwareCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceAntiRansomwareCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceAntiRansomwareCapacity) SetAssigned(v int64) *ListMultiUserInstancesResponseBodyDaInstanceAntiRansomwareCapacity {
	s.Assigned = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceAntiRansomwareCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodyDaInstanceAntiRansomwareCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceAntiRansomwareCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodyDaInstanceAntiRansomwareCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceAntiRansomwareCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodyDaInstanceCspmCapacity struct {
	// example:
	//
	// 0
	Assigned *int64 `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	// example:
	//
	// 180000
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 31569
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodyDaInstanceCspmCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodyDaInstanceCspmCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceCspmCapacity) GetAssigned() *int64 {
	return s.Assigned
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceCspmCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceCspmCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceCspmCapacity) SetAssigned(v int64) *ListMultiUserInstancesResponseBodyDaInstanceCspmCapacity {
	s.Assigned = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceCspmCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodyDaInstanceCspmCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceCspmCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodyDaInstanceCspmCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceCspmCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodyDaInstanceHoneypotCapacity struct {
	// example:
	//
	// 0
	Assigned *int64 `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	// example:
	//
	// 45
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 9
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodyDaInstanceHoneypotCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodyDaInstanceHoneypotCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceHoneypotCapacity) GetAssigned() *int64 {
	return s.Assigned
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceHoneypotCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceHoneypotCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceHoneypotCapacity) SetAssigned(v int64) *ListMultiUserInstancesResponseBodyDaInstanceHoneypotCapacity {
	s.Assigned = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceHoneypotCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodyDaInstanceHoneypotCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceHoneypotCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodyDaInstanceHoneypotCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceHoneypotCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodyDaInstanceImageScanCapacity struct {
	// example:
	//
	// 0
	Assigned *int64 `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	// example:
	//
	// 60
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodyDaInstanceImageScanCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodyDaInstanceImageScanCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceImageScanCapacity) GetAssigned() *int64 {
	return s.Assigned
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceImageScanCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceImageScanCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceImageScanCapacity) SetAssigned(v int64) *ListMultiUserInstancesResponseBodyDaInstanceImageScanCapacity {
	s.Assigned = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceImageScanCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodyDaInstanceImageScanCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceImageScanCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodyDaInstanceImageScanCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceImageScanCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodyDaInstanceRaspCapacity struct {
	// example:
	//
	// 0
	Assigned *int64 `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	// example:
	//
	// 7
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 6
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodyDaInstanceRaspCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodyDaInstanceRaspCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceRaspCapacity) GetAssigned() *int64 {
	return s.Assigned
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceRaspCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceRaspCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceRaspCapacity) SetAssigned(v int64) *ListMultiUserInstancesResponseBodyDaInstanceRaspCapacity {
	s.Assigned = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceRaspCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodyDaInstanceRaspCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceRaspCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodyDaInstanceRaspCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceRaspCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodyDaInstanceSdkCapacity struct {
	// example:
	//
	// 0
	Assigned *int64 `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	// example:
	//
	// 50
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodyDaInstanceSdkCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodyDaInstanceSdkCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceSdkCapacity) GetAssigned() *int64 {
	return s.Assigned
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceSdkCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceSdkCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceSdkCapacity) SetAssigned(v int64) *ListMultiUserInstancesResponseBodyDaInstanceSdkCapacity {
	s.Assigned = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceSdkCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodyDaInstanceSdkCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceSdkCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodyDaInstanceSdkCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceSdkCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodyDaInstanceSlsCapacity struct {
	// example:
	//
	// 10
	Assigned *int64 `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	// example:
	//
	// 150
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 5
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodyDaInstanceSlsCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodyDaInstanceSlsCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceSlsCapacity) GetAssigned() *int64 {
	return s.Assigned
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceSlsCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceSlsCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceSlsCapacity) SetAssigned(v int64) *ListMultiUserInstancesResponseBodyDaInstanceSlsCapacity {
	s.Assigned = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceSlsCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodyDaInstanceSlsCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceSlsCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodyDaInstanceSlsCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceSlsCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisCapacity struct {
	// example:
	//
	// 0
	Assigned *int64 `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	// example:
	//
	// 3000
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1548
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisCapacity) GetAssigned() *int64 {
	return s.Assigned
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisCapacity) SetAssigned(v int64) *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisCapacity {
	s.Assigned = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisFlow struct {
	// example:
	//
	// 0
	Assigned *int64 `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	// example:
	//
	// 300
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisFlow) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisFlow) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisFlow) GetAssigned() *int64 {
	return s.Assigned
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisFlow) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisFlow) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisFlow) SetAssigned(v int64) *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisFlow {
	s.Assigned = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisFlow) SetCount(v int64) *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisFlow {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisFlow) SetUsed(v int64) *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisFlow {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceThreatAnalysisFlow) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodyDaInstanceVersionSummary struct {
	// example:
	//
	// ASSET_AND_CORE
	AuthBindType *string                                                              `json:"AuthBindType,omitempty" xml:"AuthBindType,omitempty"`
	CoreCount    *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryCoreCount `json:"CoreCount,omitempty" xml:"CoreCount,omitempty" type:"Struct"`
	EcsCount     *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryEcsCount  `json:"EcsCount,omitempty" xml:"EcsCount,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListMultiUserInstancesResponseBodyDaInstanceVersionSummary) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodyDaInstanceVersionSummary) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummary) GetAuthBindType() *string {
	return s.AuthBindType
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummary) GetCoreCount() *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryCoreCount {
	return s.CoreCount
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummary) GetEcsCount() *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryEcsCount {
	return s.EcsCount
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummary) GetVersion() *int32 {
	return s.Version
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummary) SetAuthBindType(v string) *ListMultiUserInstancesResponseBodyDaInstanceVersionSummary {
	s.AuthBindType = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummary) SetCoreCount(v *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryCoreCount) *ListMultiUserInstancesResponseBodyDaInstanceVersionSummary {
	s.CoreCount = v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummary) SetEcsCount(v *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryEcsCount) *ListMultiUserInstancesResponseBodyDaInstanceVersionSummary {
	s.EcsCount = v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummary) SetVersion(v int32) *ListMultiUserInstancesResponseBodyDaInstanceVersionSummary {
	s.Version = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummary) Validate() error {
	if s.CoreCount != nil {
		if err := s.CoreCount.Validate(); err != nil {
			return err
		}
	}
	if s.EcsCount != nil {
		if err := s.EcsCount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryCoreCount struct {
	// example:
	//
	// 4
	Assigned *int64 `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	// example:
	//
	// 150
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 68
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryCoreCount) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryCoreCount) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryCoreCount) GetAssigned() *int64 {
	return s.Assigned
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryCoreCount) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryCoreCount) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryCoreCount) SetAssigned(v int64) *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryCoreCount {
	s.Assigned = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryCoreCount) SetCount(v int64) *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryCoreCount {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryCoreCount) SetUsed(v int64) *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryCoreCount {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryCoreCount) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryEcsCount struct {
	// example:
	//
	// 3
	Assigned *int64 `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	// example:
	//
	// 20
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 14
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryEcsCount) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryEcsCount) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryEcsCount) GetAssigned() *int64 {
	return s.Assigned
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryEcsCount) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryEcsCount) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryEcsCount) SetAssigned(v int64) *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryEcsCount {
	s.Assigned = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryEcsCount) SetCount(v int64) *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryEcsCount {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryEcsCount) SetUsed(v int64) *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryEcsCount {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceVersionSummaryEcsCount) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodyDaInstanceWebLockCapacity struct {
	// example:
	//
	// 0
	Assigned *int64 `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	// example:
	//
	// 5
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 3
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodyDaInstanceWebLockCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodyDaInstanceWebLockCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceWebLockCapacity) GetAssigned() *int64 {
	return s.Assigned
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceWebLockCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceWebLockCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceWebLockCapacity) SetAssigned(v int64) *ListMultiUserInstancesResponseBodyDaInstanceWebLockCapacity {
	s.Assigned = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceWebLockCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodyDaInstanceWebLockCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceWebLockCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodyDaInstanceWebLockCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyDaInstanceWebLockCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodyPageInfo struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1000
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// B604532DEF982B875E8360A6EFA3B***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 55
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMultiUserInstancesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListMultiUserInstancesResponseBodyPageInfo) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMultiUserInstancesResponseBodyPageInfo) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiUserInstancesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMultiUserInstancesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMultiUserInstancesResponseBodyPageInfo) SetCurrentPage(v int32) *ListMultiUserInstancesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyPageInfo) SetMaxResults(v int32) *ListMultiUserInstancesResponseBodyPageInfo {
	s.MaxResults = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyPageInfo) SetNextToken(v string) *ListMultiUserInstancesResponseBodyPageInfo {
	s.NextToken = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyPageInfo) SetPageSize(v int32) *ListMultiUserInstancesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyPageInfo) SetTotalCount(v int32) *ListMultiUserInstancesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodySaleInstanceList struct {
	// example:
	//
	// 103784262032
	AliUid                 *int64                                                                    `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AntiRansomwareCapacity *ListMultiUserInstancesResponseBodySaleInstanceListAntiRansomwareCapacity `json:"AntiRansomwareCapacity,omitempty" xml:"AntiRansomwareCapacity,omitempty" type:"Struct"`
	CspmCapacity           *ListMultiUserInstancesResponseBodySaleInstanceListCspmCapacity           `json:"CspmCapacity,omitempty" xml:"CspmCapacity,omitempty" type:"Struct"`
	HoneypotCapacity       *ListMultiUserInstancesResponseBodySaleInstanceListHoneypotCapacity       `json:"HoneypotCapacity,omitempty" xml:"HoneypotCapacity,omitempty" type:"Struct"`
	ImageScanCapacity      *ListMultiUserInstancesResponseBodySaleInstanceListImageScanCapacity      `json:"ImageScanCapacity,omitempty" xml:"ImageScanCapacity,omitempty" type:"Struct"`
	// example:
	//
	// api-service-spec
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	InstancePurchaseType *int32                                                          `json:"InstancePurchaseType,omitempty" xml:"InstancePurchaseType,omitempty"`
	RaspCapacity         *ListMultiUserInstancesResponseBodySaleInstanceListRaspCapacity `json:"RaspCapacity,omitempty" xml:"RaspCapacity,omitempty" type:"Struct"`
	SdkCapacity          *ListMultiUserInstancesResponseBodySaleInstanceListSdkCapacity  `json:"SdkCapacity,omitempty" xml:"SdkCapacity,omitempty" type:"Struct"`
	SlsCapacity          *ListMultiUserInstancesResponseBodySaleInstanceListSlsCapacity  `json:"SlsCapacity,omitempty" xml:"SlsCapacity,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Status                 *int32                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	ThreatAnalysisCapacity *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisCapacity `json:"ThreatAnalysisCapacity,omitempty" xml:"ThreatAnalysisCapacity,omitempty" type:"Struct"`
	ThreatAnalysisFlow     *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisFlow     `json:"ThreatAnalysisFlow,omitempty" xml:"ThreatAnalysisFlow,omitempty" type:"Struct"`
	// example:
	//
	// 2
	UserType *int32 `json:"UserType,omitempty" xml:"UserType,omitempty"`
	// example:
	//
	// 3
	Version         *int32                                                              `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionSummary  []*ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary `json:"VersionSummary,omitempty" xml:"VersionSummary,omitempty" type:"Repeated"`
	WebLockCapacity *ListMultiUserInstancesResponseBodySaleInstanceListWebLockCapacity  `json:"WebLockCapacity,omitempty" xml:"WebLockCapacity,omitempty" type:"Struct"`
}

func (s ListMultiUserInstancesResponseBodySaleInstanceList) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodySaleInstanceList) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetAntiRansomwareCapacity() *ListMultiUserInstancesResponseBodySaleInstanceListAntiRansomwareCapacity {
	return s.AntiRansomwareCapacity
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetCspmCapacity() *ListMultiUserInstancesResponseBodySaleInstanceListCspmCapacity {
	return s.CspmCapacity
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetHoneypotCapacity() *ListMultiUserInstancesResponseBodySaleInstanceListHoneypotCapacity {
	return s.HoneypotCapacity
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetImageScanCapacity() *ListMultiUserInstancesResponseBodySaleInstanceListImageScanCapacity {
	return s.ImageScanCapacity
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetInstancePurchaseType() *int32 {
	return s.InstancePurchaseType
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetRaspCapacity() *ListMultiUserInstancesResponseBodySaleInstanceListRaspCapacity {
	return s.RaspCapacity
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetSdkCapacity() *ListMultiUserInstancesResponseBodySaleInstanceListSdkCapacity {
	return s.SdkCapacity
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetSlsCapacity() *ListMultiUserInstancesResponseBodySaleInstanceListSlsCapacity {
	return s.SlsCapacity
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetStatus() *int32 {
	return s.Status
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetThreatAnalysisCapacity() *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisCapacity {
	return s.ThreatAnalysisCapacity
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetThreatAnalysisFlow() *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisFlow {
	return s.ThreatAnalysisFlow
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetUserType() *int32 {
	return s.UserType
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetVersion() *int32 {
	return s.Version
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetVersionSummary() []*ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary {
	return s.VersionSummary
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) GetWebLockCapacity() *ListMultiUserInstancesResponseBodySaleInstanceListWebLockCapacity {
	return s.WebLockCapacity
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetAliUid(v int64) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.AliUid = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetAntiRansomwareCapacity(v *ListMultiUserInstancesResponseBodySaleInstanceListAntiRansomwareCapacity) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.AntiRansomwareCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetCspmCapacity(v *ListMultiUserInstancesResponseBodySaleInstanceListCspmCapacity) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.CspmCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetHoneypotCapacity(v *ListMultiUserInstancesResponseBodySaleInstanceListHoneypotCapacity) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.HoneypotCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetImageScanCapacity(v *ListMultiUserInstancesResponseBodySaleInstanceListImageScanCapacity) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.ImageScanCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetInstanceId(v string) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.InstanceId = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetInstancePurchaseType(v int32) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.InstancePurchaseType = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetRaspCapacity(v *ListMultiUserInstancesResponseBodySaleInstanceListRaspCapacity) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.RaspCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetSdkCapacity(v *ListMultiUserInstancesResponseBodySaleInstanceListSdkCapacity) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.SdkCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetSlsCapacity(v *ListMultiUserInstancesResponseBodySaleInstanceListSlsCapacity) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.SlsCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetStatus(v int32) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.Status = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetThreatAnalysisCapacity(v *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisCapacity) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.ThreatAnalysisCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetThreatAnalysisFlow(v *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisFlow) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.ThreatAnalysisFlow = v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetUserType(v int32) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.UserType = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetVersion(v int32) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.Version = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetVersionSummary(v []*ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.VersionSummary = v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) SetWebLockCapacity(v *ListMultiUserInstancesResponseBodySaleInstanceListWebLockCapacity) *ListMultiUserInstancesResponseBodySaleInstanceList {
	s.WebLockCapacity = v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceList) Validate() error {
	if s.AntiRansomwareCapacity != nil {
		if err := s.AntiRansomwareCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.CspmCapacity != nil {
		if err := s.CspmCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.HoneypotCapacity != nil {
		if err := s.HoneypotCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.ImageScanCapacity != nil {
		if err := s.ImageScanCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.RaspCapacity != nil {
		if err := s.RaspCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.SdkCapacity != nil {
		if err := s.SdkCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.SlsCapacity != nil {
		if err := s.SlsCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.ThreatAnalysisCapacity != nil {
		if err := s.ThreatAnalysisCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.ThreatAnalysisFlow != nil {
		if err := s.ThreatAnalysisFlow.Validate(); err != nil {
			return err
		}
	}
	if s.VersionSummary != nil {
		for _, item := range s.VersionSummary {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WebLockCapacity != nil {
		if err := s.WebLockCapacity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMultiUserInstancesResponseBodySaleInstanceListAntiRansomwareCapacity struct {
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListAntiRansomwareCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListAntiRansomwareCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListAntiRansomwareCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListAntiRansomwareCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListAntiRansomwareCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListAntiRansomwareCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListAntiRansomwareCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListAntiRansomwareCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListAntiRansomwareCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodySaleInstanceListCspmCapacity struct {
	// example:
	//
	// 0
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListCspmCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListCspmCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListCspmCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListCspmCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListCspmCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListCspmCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListCspmCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListCspmCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListCspmCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodySaleInstanceListHoneypotCapacity struct {
	// example:
	//
	// 0
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListHoneypotCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListHoneypotCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListHoneypotCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListHoneypotCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListHoneypotCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListHoneypotCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListHoneypotCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListHoneypotCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListHoneypotCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodySaleInstanceListImageScanCapacity struct {
	// example:
	//
	// 0
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListImageScanCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListImageScanCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListImageScanCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListImageScanCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListImageScanCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListImageScanCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListImageScanCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListImageScanCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListImageScanCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodySaleInstanceListRaspCapacity struct {
	// example:
	//
	// 0
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListRaspCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListRaspCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListRaspCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListRaspCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListRaspCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListRaspCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListRaspCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListRaspCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListRaspCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodySaleInstanceListSdkCapacity struct {
	// example:
	//
	// 0
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListSdkCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListSdkCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListSdkCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListSdkCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListSdkCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListSdkCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListSdkCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListSdkCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListSdkCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodySaleInstanceListSlsCapacity struct {
	// example:
	//
	// 0
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListSlsCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListSlsCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListSlsCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListSlsCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListSlsCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListSlsCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListSlsCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListSlsCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListSlsCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisCapacity struct {
	// example:
	//
	// 0
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisCapacity) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisFlow struct {
	// example:
	//
	// 0
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisFlow) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisFlow) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisFlow) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisFlow) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisFlow) SetCount(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisFlow {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisFlow) SetUsed(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisFlow {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListThreatAnalysisFlow) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary struct {
	// example:
	//
	// ASSET_AND_CORE
	AuthBindType *string                                                                    `json:"AuthBindType,omitempty" xml:"AuthBindType,omitempty"`
	CoreCount    *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryCoreCount `json:"CoreCount,omitempty" xml:"CoreCount,omitempty" type:"Struct"`
	EcsCount     *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryEcsCount  `json:"EcsCount,omitempty" xml:"EcsCount,omitempty" type:"Struct"`
	// example:
	//
	// 5
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary) GetAuthBindType() *string {
	return s.AuthBindType
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary) GetCoreCount() *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryCoreCount {
	return s.CoreCount
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary) GetEcsCount() *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryEcsCount {
	return s.EcsCount
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary) GetVersion() *int32 {
	return s.Version
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary) SetAuthBindType(v string) *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary {
	s.AuthBindType = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary) SetCoreCount(v *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryCoreCount) *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary {
	s.CoreCount = v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary) SetEcsCount(v *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryEcsCount) *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary {
	s.EcsCount = v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary) SetVersion(v int32) *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary {
	s.Version = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummary) Validate() error {
	if s.CoreCount != nil {
		if err := s.CoreCount.Validate(); err != nil {
			return err
		}
	}
	if s.EcsCount != nil {
		if err := s.EcsCount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryCoreCount struct {
	// example:
	//
	// 0
	Assigned *int64 `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	// example:
	//
	// 4
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryCoreCount) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryCoreCount) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryCoreCount) GetAssigned() *int64 {
	return s.Assigned
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryCoreCount) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryCoreCount) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryCoreCount) SetAssigned(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryCoreCount {
	s.Assigned = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryCoreCount) SetCount(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryCoreCount {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryCoreCount) SetUsed(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryCoreCount {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryCoreCount) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryEcsCount struct {
	// example:
	//
	// 0
	Assigned *int64 `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	// example:
	//
	// 0
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryEcsCount) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryEcsCount) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryEcsCount) GetAssigned() *int64 {
	return s.Assigned
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryEcsCount) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryEcsCount) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryEcsCount) SetAssigned(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryEcsCount {
	s.Assigned = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryEcsCount) SetCount(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryEcsCount {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryEcsCount) SetUsed(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryEcsCount {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListVersionSummaryEcsCount) Validate() error {
	return dara.Validate(s)
}

type ListMultiUserInstancesResponseBodySaleInstanceListWebLockCapacity struct {
	// example:
	//
	// 0
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListWebLockCapacity) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponseBodySaleInstanceListWebLockCapacity) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListWebLockCapacity) GetCount() *int64 {
	return s.Count
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListWebLockCapacity) GetUsed() *int64 {
	return s.Used
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListWebLockCapacity) SetCount(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListWebLockCapacity {
	s.Count = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListWebLockCapacity) SetUsed(v int64) *ListMultiUserInstancesResponseBodySaleInstanceListWebLockCapacity {
	s.Used = &v
	return s
}

func (s *ListMultiUserInstancesResponseBodySaleInstanceListWebLockCapacity) Validate() error {
	return dara.Validate(s)
}
