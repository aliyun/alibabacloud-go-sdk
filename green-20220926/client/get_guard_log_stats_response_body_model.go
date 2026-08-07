// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGuardLogStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetGuardLogStatsResponseBodyData) *GetGuardLogStatsResponseBody
	GetData() []*GetGuardLogStatsResponseBodyData
	SetRequestId(v string) *GetGuardLogStatsResponseBody
	GetRequestId() *string
}

type GetGuardLogStatsResponseBody struct {
	// The data.
	Data []*GetGuardLogStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGuardLogStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGuardLogStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetGuardLogStatsResponseBody) GetData() []*GetGuardLogStatsResponseBodyData {
	return s.Data
}

func (s *GetGuardLogStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGuardLogStatsResponseBody) SetData(v []*GetGuardLogStatsResponseBodyData) *GetGuardLogStatsResponseBody {
	s.Data = v
	return s
}

func (s *GetGuardLogStatsResponseBody) SetRequestId(v string) *GetGuardLogStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGuardLogStatsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGuardLogStatsResponseBodyData struct {
	// The delivery region.
	//
	// example:
	//
	// cn-beijing
	DeliveryRegion *string `json:"DeliveryRegion,omitempty" xml:"DeliveryRegion,omitempty"`
	// Indicates whether the feature is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The log analysis configuration.
	LogAnalysisConfig map[string]interface{} `json:"LogAnalysisConfig,omitempty" xml:"LogAnalysisConfig,omitempty"`
	// The name of the Simple Log Service Logstore.
	//
	// example:
	//
	// test003x
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The remaining storage space. Unit: TB.
	//
	// example:
	//
	// 1
	PendingStorage *int64 `json:"PendingStorage,omitempty" xml:"PendingStorage,omitempty"`
	// The reserved storage. Unit: bytes.
	//
	// example:
	//
	// 1
	PreserveStorage *int64 `json:"PreserveStorage,omitempty" xml:"PreserveStorage,omitempty"`
	// The project space.
	//
	// example:
	//
	// xxx_log
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The total storage space. Unit: TB.
	//
	// example:
	//
	// 3
	TotalStorage *int64 `json:"TotalStorage,omitempty" xml:"TotalStorage,omitempty"`
	// The number of days for which data is retained.
	//
	// example:
	//
	// 30
	Ttl *int64 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type.
	//
	// example:
	//
	// guard_meta_log
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// UID。
	//
	// example:
	//
	// 1643953****74290
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The used storage. Unit: bytes.
	//
	// example:
	//
	// 1
	UsedStorage *int64 `json:"UsedStorage,omitempty" xml:"UsedStorage,omitempty"`
}

func (s GetGuardLogStatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGuardLogStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGuardLogStatsResponseBodyData) GetDeliveryRegion() *string {
	return s.DeliveryRegion
}

func (s *GetGuardLogStatsResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *GetGuardLogStatsResponseBodyData) GetLogAnalysisConfig() map[string]interface{} {
	return s.LogAnalysisConfig
}

func (s *GetGuardLogStatsResponseBodyData) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *GetGuardLogStatsResponseBodyData) GetPendingStorage() *int64 {
	return s.PendingStorage
}

func (s *GetGuardLogStatsResponseBodyData) GetPreserveStorage() *int64 {
	return s.PreserveStorage
}

func (s *GetGuardLogStatsResponseBodyData) GetProject() *string {
	return s.Project
}

func (s *GetGuardLogStatsResponseBodyData) GetTotalStorage() *int64 {
	return s.TotalStorage
}

func (s *GetGuardLogStatsResponseBodyData) GetTtl() *int64 {
	return s.Ttl
}

func (s *GetGuardLogStatsResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetGuardLogStatsResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *GetGuardLogStatsResponseBodyData) GetUsedStorage() *int64 {
	return s.UsedStorage
}

func (s *GetGuardLogStatsResponseBodyData) SetDeliveryRegion(v string) *GetGuardLogStatsResponseBodyData {
	s.DeliveryRegion = &v
	return s
}

func (s *GetGuardLogStatsResponseBodyData) SetEnable(v bool) *GetGuardLogStatsResponseBodyData {
	s.Enable = &v
	return s
}

func (s *GetGuardLogStatsResponseBodyData) SetLogAnalysisConfig(v map[string]interface{}) *GetGuardLogStatsResponseBodyData {
	s.LogAnalysisConfig = v
	return s
}

func (s *GetGuardLogStatsResponseBodyData) SetLogStoreName(v string) *GetGuardLogStatsResponseBodyData {
	s.LogStoreName = &v
	return s
}

func (s *GetGuardLogStatsResponseBodyData) SetPendingStorage(v int64) *GetGuardLogStatsResponseBodyData {
	s.PendingStorage = &v
	return s
}

func (s *GetGuardLogStatsResponseBodyData) SetPreserveStorage(v int64) *GetGuardLogStatsResponseBodyData {
	s.PreserveStorage = &v
	return s
}

func (s *GetGuardLogStatsResponseBodyData) SetProject(v string) *GetGuardLogStatsResponseBodyData {
	s.Project = &v
	return s
}

func (s *GetGuardLogStatsResponseBodyData) SetTotalStorage(v int64) *GetGuardLogStatsResponseBodyData {
	s.TotalStorage = &v
	return s
}

func (s *GetGuardLogStatsResponseBodyData) SetTtl(v int64) *GetGuardLogStatsResponseBodyData {
	s.Ttl = &v
	return s
}

func (s *GetGuardLogStatsResponseBodyData) SetType(v string) *GetGuardLogStatsResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetGuardLogStatsResponseBodyData) SetUid(v string) *GetGuardLogStatsResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetGuardLogStatsResponseBodyData) SetUsedStorage(v int64) *GetGuardLogStatsResponseBodyData {
	s.UsedStorage = &v
	return s
}

func (s *GetGuardLogStatsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
