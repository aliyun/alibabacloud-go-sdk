// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterStorageLimitationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeDBClusterStorageLimitationResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *DescribeDBClusterStorageLimitationResponseBodyData) *DescribeDBClusterStorageLimitationResponseBody
	GetData() *DescribeDBClusterStorageLimitationResponseBodyData
	SetDynamicCode(v string) *DescribeDBClusterStorageLimitationResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeDBClusterStorageLimitationResponseBody
	GetDynamicMessage() *string
	SetRequestId(v string) *DescribeDBClusterStorageLimitationResponseBody
	GetRequestId() *string
}

type DescribeDBClusterStorageLimitationResponseBody struct {
	// example:
	//
	// failed
	AccessDeniedDetail *string                                             `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *DescribeDBClusterStorageLimitationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// An error occurred while processing your request.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// 4773E4EC-025D-509F-AEA9-D53123FDFB0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterStorageLimitationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterStorageLimitationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStorageLimitationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeDBClusterStorageLimitationResponseBody) GetData() *DescribeDBClusterStorageLimitationResponseBodyData {
	return s.Data
}

func (s *DescribeDBClusterStorageLimitationResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeDBClusterStorageLimitationResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeDBClusterStorageLimitationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterStorageLimitationResponseBody) SetAccessDeniedDetail(v string) *DescribeDBClusterStorageLimitationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeDBClusterStorageLimitationResponseBody) SetData(v *DescribeDBClusterStorageLimitationResponseBodyData) *DescribeDBClusterStorageLimitationResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBClusterStorageLimitationResponseBody) SetDynamicCode(v string) *DescribeDBClusterStorageLimitationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDBClusterStorageLimitationResponseBody) SetDynamicMessage(v string) *DescribeDBClusterStorageLimitationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDBClusterStorageLimitationResponseBody) SetRequestId(v string) *DescribeDBClusterStorageLimitationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterStorageLimitationResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterStorageLimitationResponseBodyData struct {
	ClassCodeList []*DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList `json:"ClassCodeList,omitempty" xml:"ClassCodeList,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterStorageLimitationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterStorageLimitationResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStorageLimitationResponseBodyData) GetClassCodeList() []*DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList {
	return s.ClassCodeList
}

func (s *DescribeDBClusterStorageLimitationResponseBodyData) SetClassCodeList(v []*DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) *DescribeDBClusterStorageLimitationResponseBodyData {
	s.ClassCodeList = v
	return s
}

func (s *DescribeDBClusterStorageLimitationResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList struct {
	// example:
	//
	// selectdb.xlarge
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// example:
	//
	// 4
	CpuCores *int32 `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	// example:
	//
	// 8
	DefaultStorageInGB *int32 `json:"DefaultStorageInGB,omitempty" xml:"DefaultStorageInGB,omitempty"`
	// example:
	//
	// 16
	MaxStorageInGB *int32 `json:"MaxStorageInGB,omitempty" xml:"MaxStorageInGB,omitempty"`
	// example:
	//
	// 16
	MemoryInGB *int32 `json:"MemoryInGB,omitempty" xml:"MemoryInGB,omitempty"`
	// example:
	//
	// 2
	MinStorageInGB *int32 `json:"MinStorageInGB,omitempty" xml:"MinStorageInGB,omitempty"`
	// example:
	//
	// 1
	StepStorageInGB *int32 `json:"StepStorageInGB,omitempty" xml:"StepStorageInGB,omitempty"`
}

func (s DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) GetCpuCores() *int32 {
	return s.CpuCores
}

func (s *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) GetDefaultStorageInGB() *int32 {
	return s.DefaultStorageInGB
}

func (s *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) GetMaxStorageInGB() *int32 {
	return s.MaxStorageInGB
}

func (s *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) GetMemoryInGB() *int32 {
	return s.MemoryInGB
}

func (s *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) GetMinStorageInGB() *int32 {
	return s.MinStorageInGB
}

func (s *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) GetStepStorageInGB() *int32 {
	return s.StepStorageInGB
}

func (s *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) SetClassCode(v string) *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList {
	s.ClassCode = &v
	return s
}

func (s *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) SetCpuCores(v int32) *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList {
	s.CpuCores = &v
	return s
}

func (s *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) SetDefaultStorageInGB(v int32) *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList {
	s.DefaultStorageInGB = &v
	return s
}

func (s *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) SetMaxStorageInGB(v int32) *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList {
	s.MaxStorageInGB = &v
	return s
}

func (s *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) SetMemoryInGB(v int32) *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList {
	s.MemoryInGB = &v
	return s
}

func (s *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) SetMinStorageInGB(v int32) *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList {
	s.MinStorageInGB = &v
	return s
}

func (s *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) SetStepStorageInGB(v int32) *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList {
	s.StepStorageInGB = &v
	return s
}

func (s *DescribeDBClusterStorageLimitationResponseBodyDataClassCodeList) Validate() error {
	return dara.Validate(s)
}
