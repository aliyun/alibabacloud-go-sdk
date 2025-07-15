// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopOversoldGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeDesktopOversoldGroupResponseBody
	GetCount() *int32
	SetData(v []*DescribeDesktopOversoldGroupResponseBodyData) *DescribeDesktopOversoldGroupResponseBody
	GetData() []*DescribeDesktopOversoldGroupResponseBodyData
	SetNextToken(v string) *DescribeDesktopOversoldGroupResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDesktopOversoldGroupResponseBody
	GetRequestId() *string
}

type DescribeDesktopOversoldGroupResponseBody struct {
	Count     *int32                                          `json:"Count,omitempty" xml:"Count,omitempty"`
	Data      []*DescribeDesktopOversoldGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	NextToken *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDesktopOversoldGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopOversoldGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopOversoldGroupResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeDesktopOversoldGroupResponseBody) GetData() []*DescribeDesktopOversoldGroupResponseBodyData {
	return s.Data
}

func (s *DescribeDesktopOversoldGroupResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopOversoldGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDesktopOversoldGroupResponseBody) SetCount(v int32) *DescribeDesktopOversoldGroupResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBody) SetData(v []*DescribeDesktopOversoldGroupResponseBodyData) *DescribeDesktopOversoldGroupResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBody) SetNextToken(v string) *DescribeDesktopOversoldGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBody) SetRequestId(v string) *DescribeDesktopOversoldGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDesktopOversoldGroupResponseBodyData struct {
	ConcurrenceCount       *int32  `json:"ConcurrenceCount,omitempty" xml:"ConcurrenceCount,omitempty"`
	CurConcurrenceCount    *int32  `json:"CurConcurrenceCount,omitempty" xml:"CurConcurrenceCount,omitempty"`
	DataDiskSize           *int32  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopType            *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	DirectoryId            *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	ExpireTime             *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	IdleDisconnectDuration *string `json:"IdleDisconnectDuration,omitempty" xml:"IdleDisconnectDuration,omitempty"`
	ImageId                *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	KeepDuration           *string `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OversoldGroupId        *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
	OversoldUserCount      *int32  `json:"OversoldUserCount,omitempty" xml:"OversoldUserCount,omitempty"`
	OversoldWarn           *int32  `json:"OversoldWarn,omitempty" xml:"OversoldWarn,omitempty"`
	PolicyGroupId          *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	SaleStatus             *string `json:"SaleStatus,omitempty" xml:"SaleStatus,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StopDuration           *int64  `json:"StopDuration,omitempty" xml:"StopDuration,omitempty"`
	SystemDiskSize         *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeDesktopOversoldGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopOversoldGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetConcurrenceCount() *int32 {
	return s.ConcurrenceCount
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetCurConcurrenceCount() *int32 {
	return s.CurConcurrenceCount
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetDesktopType() *string {
	return s.DesktopType
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetIdleDisconnectDuration() *string {
	return s.IdleDisconnectDuration
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetKeepDuration() *string {
	return s.KeepDuration
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetOversoldUserCount() *int32 {
	return s.OversoldUserCount
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetOversoldWarn() *int32 {
	return s.OversoldWarn
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetSaleStatus() *string {
	return s.SaleStatus
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetStopDuration() *int64 {
	return s.StopDuration
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetConcurrenceCount(v int32) *DescribeDesktopOversoldGroupResponseBodyData {
	s.ConcurrenceCount = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetCurConcurrenceCount(v int32) *DescribeDesktopOversoldGroupResponseBodyData {
	s.CurConcurrenceCount = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetDataDiskSize(v int32) *DescribeDesktopOversoldGroupResponseBodyData {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetDescription(v string) *DescribeDesktopOversoldGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetDesktopType(v string) *DescribeDesktopOversoldGroupResponseBodyData {
	s.DesktopType = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetDirectoryId(v string) *DescribeDesktopOversoldGroupResponseBodyData {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetExpireTime(v string) *DescribeDesktopOversoldGroupResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetIdleDisconnectDuration(v string) *DescribeDesktopOversoldGroupResponseBodyData {
	s.IdleDisconnectDuration = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetImageId(v string) *DescribeDesktopOversoldGroupResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetKeepDuration(v string) *DescribeDesktopOversoldGroupResponseBodyData {
	s.KeepDuration = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetName(v string) *DescribeDesktopOversoldGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetOversoldGroupId(v string) *DescribeDesktopOversoldGroupResponseBodyData {
	s.OversoldGroupId = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetOversoldUserCount(v int32) *DescribeDesktopOversoldGroupResponseBodyData {
	s.OversoldUserCount = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetOversoldWarn(v int32) *DescribeDesktopOversoldGroupResponseBodyData {
	s.OversoldWarn = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetPolicyGroupId(v string) *DescribeDesktopOversoldGroupResponseBodyData {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetSaleStatus(v string) *DescribeDesktopOversoldGroupResponseBodyData {
	s.SaleStatus = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetStatus(v string) *DescribeDesktopOversoldGroupResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetStopDuration(v int64) *DescribeDesktopOversoldGroupResponseBodyData {
	s.StopDuration = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) SetSystemDiskSize(v int32) *DescribeDesktopOversoldGroupResponseBodyData {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopOversoldGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
