// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordFiles(v []*DescribeRecordFileResponseBodyRecordFiles) *DescribeRecordFileResponseBody
	GetRecordFiles() []*DescribeRecordFileResponseBodyRecordFiles
	SetRequestId(v string) *DescribeRecordFileResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeRecordFileResponseBody
	GetTotalCount() *int32
}

type DescribeRecordFileResponseBody struct {
	RecordFiles []*DescribeRecordFileResponseBodyRecordFiles `json:"RecordFiles,omitempty" xml:"RecordFiles,omitempty" type:"Repeated"`
	// RequestId
	//
	// example:
	//
	// 05C2791F-41A7-5E7C-B5E4-1401FD0E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRecordFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordFileResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordFileResponseBody) GetRecordFiles() []*DescribeRecordFileResponseBodyRecordFiles {
	return s.RecordFiles
}

func (s *DescribeRecordFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecordFileResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRecordFileResponseBody) SetRecordFiles(v []*DescribeRecordFileResponseBodyRecordFiles) *DescribeRecordFileResponseBody {
	s.RecordFiles = v
	return s
}

func (s *DescribeRecordFileResponseBody) SetRequestId(v string) *DescribeRecordFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordFileResponseBody) SetTotalCount(v int32) *DescribeRecordFileResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRecordFileResponseBody) Validate() error {
	if s.RecordFiles != nil {
		for _, item := range s.RecordFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRecordFileResponseBodyRecordFiles struct {
	// example:
	//
	// ecd-7yulhw1g1attet7d2
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// example:
	//
	// fvt-ecd
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// example:
	//
	// Alice
	EndUserId    *string                                                  `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EventDetails []*DescribeRecordFileResponseBodyRecordFilesEventDetails `json:"EventDetails,omitempty" xml:"EventDetails,omitempty" type:"Repeated"`
	// example:
	//
	// Task3
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 139
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// pg-**
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 2025-12-18T06:02:25Z
	RecordEndTime *string `json:"RecordEndTime,omitempty" xml:"RecordEndTime,omitempty"`
	// example:
	//
	// 1774656000000
	RecordExpire *int64 `json:"RecordExpire,omitempty" xml:"RecordExpire,omitempty"`
	// example:
	//
	// 2025-12-18T06:02:25Z
	RecordStartTime *string `json:"RecordStartTime,omitempty" xml:"RecordStartTime,omitempty"`
	// example:
	//
	// period
	RecordType *int32 `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-f3s3dgt8dtb0vlqc8
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// resource
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRecordFileResponseBodyRecordFiles) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordFileResponseBodyRecordFiles) GoString() string {
	return s.String()
}

func (s *DescribeRecordFileResponseBodyRecordFiles) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeRecordFileResponseBodyRecordFiles) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeRecordFileResponseBodyRecordFiles) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeRecordFileResponseBodyRecordFiles) GetEventDetails() []*DescribeRecordFileResponseBodyRecordFilesEventDetails {
	return s.EventDetails
}

func (s *DescribeRecordFileResponseBodyRecordFiles) GetFileName() *string {
	return s.FileName
}

func (s *DescribeRecordFileResponseBodyRecordFiles) GetFileSize() *int64 {
	return s.FileSize
}

func (s *DescribeRecordFileResponseBodyRecordFiles) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeRecordFileResponseBodyRecordFiles) GetRecordEndTime() *string {
	return s.RecordEndTime
}

func (s *DescribeRecordFileResponseBodyRecordFiles) GetRecordExpire() *int64 {
	return s.RecordExpire
}

func (s *DescribeRecordFileResponseBodyRecordFiles) GetRecordStartTime() *string {
	return s.RecordStartTime
}

func (s *DescribeRecordFileResponseBodyRecordFiles) GetRecordType() *int32 {
	return s.RecordType
}

func (s *DescribeRecordFileResponseBodyRecordFiles) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRecordFileResponseBodyRecordFiles) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRecordFileResponseBodyRecordFiles) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeRecordFileResponseBodyRecordFiles) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeRecordFileResponseBodyRecordFiles) SetDesktopId(v string) *DescribeRecordFileResponseBodyRecordFiles {
	s.DesktopId = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFiles) SetDesktopName(v string) *DescribeRecordFileResponseBodyRecordFiles {
	s.DesktopName = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFiles) SetEndUserId(v string) *DescribeRecordFileResponseBodyRecordFiles {
	s.EndUserId = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFiles) SetEventDetails(v []*DescribeRecordFileResponseBodyRecordFilesEventDetails) *DescribeRecordFileResponseBodyRecordFiles {
	s.EventDetails = v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFiles) SetFileName(v string) *DescribeRecordFileResponseBodyRecordFiles {
	s.FileName = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFiles) SetFileSize(v int64) *DescribeRecordFileResponseBodyRecordFiles {
	s.FileSize = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFiles) SetPolicyId(v string) *DescribeRecordFileResponseBodyRecordFiles {
	s.PolicyId = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFiles) SetRecordEndTime(v string) *DescribeRecordFileResponseBodyRecordFiles {
	s.RecordEndTime = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFiles) SetRecordExpire(v int64) *DescribeRecordFileResponseBodyRecordFiles {
	s.RecordExpire = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFiles) SetRecordStartTime(v string) *DescribeRecordFileResponseBodyRecordFiles {
	s.RecordStartTime = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFiles) SetRecordType(v int32) *DescribeRecordFileResponseBodyRecordFiles {
	s.RecordType = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFiles) SetRegionId(v string) *DescribeRecordFileResponseBodyRecordFiles {
	s.RegionId = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFiles) SetResourceGroupId(v string) *DescribeRecordFileResponseBodyRecordFiles {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFiles) SetResourceGroupName(v string) *DescribeRecordFileResponseBodyRecordFiles {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFiles) SetStatus(v int32) *DescribeRecordFileResponseBodyRecordFiles {
	s.Status = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFiles) Validate() error {
	if s.EventDetails != nil {
		for _, item := range s.EventDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRecordFileResponseBodyRecordFilesEventDetails struct {
	// example:
	//
	// {\\n\\t\\"messageInfo\\" : \\"mouse_move\\",\\n\\t\\"messageType\\" : \\"UserInput\\"\\n}\\n"
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// 60
	EventTime *int32 `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// example:
	//
	// UserInput
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
}

func (s DescribeRecordFileResponseBodyRecordFilesEventDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordFileResponseBodyRecordFilesEventDetails) GoString() string {
	return s.String()
}

func (s *DescribeRecordFileResponseBodyRecordFilesEventDetails) GetEventName() *string {
	return s.EventName
}

func (s *DescribeRecordFileResponseBodyRecordFilesEventDetails) GetEventTime() *int32 {
	return s.EventTime
}

func (s *DescribeRecordFileResponseBodyRecordFilesEventDetails) GetEventType() *string {
	return s.EventType
}

func (s *DescribeRecordFileResponseBodyRecordFilesEventDetails) SetEventName(v string) *DescribeRecordFileResponseBodyRecordFilesEventDetails {
	s.EventName = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFilesEventDetails) SetEventTime(v int32) *DescribeRecordFileResponseBodyRecordFilesEventDetails {
	s.EventTime = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFilesEventDetails) SetEventType(v string) *DescribeRecordFileResponseBodyRecordFilesEventDetails {
	s.EventType = &v
	return s
}

func (s *DescribeRecordFileResponseBodyRecordFilesEventDetails) Validate() error {
	return dara.Validate(s)
}
