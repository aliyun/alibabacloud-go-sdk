// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeSpecResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeSpecResponseBody
	GetRequestId() *string
	SetSpecInfoModel(v []*DescribeSpecResponseBodySpecInfoModel) *DescribeSpecResponseBody
	GetSpecInfoModel() []*DescribeSpecResponseBodySpecInfoModel
	SetTotalCount(v int32) *DescribeSpecResponseBody
	GetTotalCount() *int32
}

type DescribeSpecResponseBody struct {
	// Indicates the current read position returned by this call. An empty value means that all data has been read.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kw9dGL5jves2FS9RLq****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// D9888DAD-331E-5FBC-B5A0-F2445115****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The specifications.
	SpecInfoModel []*DescribeSpecResponseBodySpecInfoModel `json:"SpecInfoModel,omitempty" xml:"SpecInfoModel,omitempty" type:"Repeated"`
	// Total number of items.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpecResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSpecResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSpecResponseBody) GetSpecInfoModel() []*DescribeSpecResponseBodySpecInfoModel {
	return s.SpecInfoModel
}

func (s *DescribeSpecResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSpecResponseBody) SetNextToken(v string) *DescribeSpecResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSpecResponseBody) SetRequestId(v string) *DescribeSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSpecResponseBody) SetSpecInfoModel(v []*DescribeSpecResponseBodySpecInfoModel) *DescribeSpecResponseBody {
	s.SpecInfoModel = v
	return s
}

func (s *DescribeSpecResponseBody) SetTotalCount(v int32) *DescribeSpecResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSpecResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSpecResponseBodySpecInfoModel struct {
	// Number of CPU cores.
	//
	// example:
	//
	// 8
	Core *int32 `json:"Core,omitempty" xml:"Core,omitempty"`
	// The maximum number of cloud phone instances.
	//
	// example:
	//
	// 40
	MaxPhoneCount *string `json:"MaxPhoneCount,omitempty" xml:"MaxPhoneCount,omitempty"`
	// Memory size.
	//
	// example:
	//
	// 16
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The minimum number of cloud phone instances.
	//
	// example:
	//
	// 4
	MinPhoneCount *string `json:"MinPhoneCount,omitempty" xml:"MinPhoneCount,omitempty"`
	// example:
	//
	// 2
	PhoneCount *string `json:"PhoneCount,omitempty" xml:"PhoneCount,omitempty"`
	// example:
	//
	// 1920*1080
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// Specification ID.
	//
	// example:
	//
	// acp.basic.small
	SpecId *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	// Specification status.
	//
	// example:
	//
	// Available
	SpecStatus *string `json:"SpecStatus,omitempty" xml:"SpecStatus,omitempty"`
	// Specification type.
	//
	// example:
	//
	// ARM
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// System disk size, in GB.
	//
	// example:
	//
	// 32
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeSpecResponseBodySpecInfoModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpecResponseBodySpecInfoModel) GoString() string {
	return s.String()
}

func (s *DescribeSpecResponseBodySpecInfoModel) GetCore() *int32 {
	return s.Core
}

func (s *DescribeSpecResponseBodySpecInfoModel) GetMaxPhoneCount() *string {
	return s.MaxPhoneCount
}

func (s *DescribeSpecResponseBodySpecInfoModel) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeSpecResponseBodySpecInfoModel) GetMinPhoneCount() *string {
	return s.MinPhoneCount
}

func (s *DescribeSpecResponseBodySpecInfoModel) GetPhoneCount() *string {
	return s.PhoneCount
}

func (s *DescribeSpecResponseBodySpecInfoModel) GetResolution() *string {
	return s.Resolution
}

func (s *DescribeSpecResponseBodySpecInfoModel) GetSpecId() *string {
	return s.SpecId
}

func (s *DescribeSpecResponseBodySpecInfoModel) GetSpecStatus() *string {
	return s.SpecStatus
}

func (s *DescribeSpecResponseBodySpecInfoModel) GetSpecType() *string {
	return s.SpecType
}

func (s *DescribeSpecResponseBodySpecInfoModel) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetCore(v int32) *DescribeSpecResponseBodySpecInfoModel {
	s.Core = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetMaxPhoneCount(v string) *DescribeSpecResponseBodySpecInfoModel {
	s.MaxPhoneCount = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetMemory(v int32) *DescribeSpecResponseBodySpecInfoModel {
	s.Memory = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetMinPhoneCount(v string) *DescribeSpecResponseBodySpecInfoModel {
	s.MinPhoneCount = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetPhoneCount(v string) *DescribeSpecResponseBodySpecInfoModel {
	s.PhoneCount = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetResolution(v string) *DescribeSpecResponseBodySpecInfoModel {
	s.Resolution = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetSpecId(v string) *DescribeSpecResponseBodySpecInfoModel {
	s.SpecId = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetSpecStatus(v string) *DescribeSpecResponseBodySpecInfoModel {
	s.SpecStatus = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetSpecType(v string) *DescribeSpecResponseBodySpecInfoModel {
	s.SpecType = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) SetSystemDiskSize(v int32) *DescribeSpecResponseBodySpecInfoModel {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeSpecResponseBodySpecInfoModel) Validate() error {
	return dara.Validate(s)
}
