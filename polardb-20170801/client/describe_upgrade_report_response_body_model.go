// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpgradeReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetails(v string) *DescribeUpgradeReportResponseBody
	GetDetails() *string
	SetDstDBType(v string) *DescribeUpgradeReportResponseBody
	GetDstDBType() *string
	SetItems(v []*DescribeUpgradeReportResponseBodyItems) *DescribeUpgradeReportResponseBody
	GetItems() []*DescribeUpgradeReportResponseBodyItems
	SetItemsSize(v int64) *DescribeUpgradeReportResponseBody
	GetItemsSize() *int64
	SetRequestId(v string) *DescribeUpgradeReportResponseBody
	GetRequestId() *string
	SetSourceDBClusterId(v string) *DescribeUpgradeReportResponseBody
	GetSourceDBClusterId() *string
	SetSrcDBType(v string) *DescribeUpgradeReportResponseBody
	GetSrcDBType() *string
	SetSrcDeleted(v string) *DescribeUpgradeReportResponseBody
	GetSrcDeleted() *string
	SetTotalSize(v int64) *DescribeUpgradeReportResponseBody
	GetTotalSize() *int64
	SetType(v string) *DescribeUpgradeReportResponseBody
	GetType() *string
	SetUpgradeReportList(v []*DescribeUpgradeReportResponseBodyUpgradeReportList) *DescribeUpgradeReportResponseBody
	GetUpgradeReportList() []*DescribeUpgradeReportResponseBodyUpgradeReportList
}

type DescribeUpgradeReportResponseBody struct {
	// example:
	//
	// []
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// MySQL
	DstDBType *string                                   `json:"DstDBType,omitempty" xml:"DstDBType,omitempty"`
	Items     []*DescribeUpgradeReportResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	ItemsSize *int64 `json:"ItemsSize,omitempty" xml:"ItemsSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rm-2zen5pe5vi56447d0
	SourceDBClusterId *string `json:"SourceDBClusterId,omitempty" xml:"SourceDBClusterId,omitempty"`
	// example:
	//
	// MySQL
	SrcDBType *string `json:"SrcDBType,omitempty" xml:"SrcDBType,omitempty"`
	// example:
	//
	// 1
	SrcDeleted *string `json:"SrcDeleted,omitempty" xml:"SrcDeleted,omitempty"`
	// example:
	//
	// 137
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
	// example:
	//
	// tair
	Type              *string                                               `json:"Type,omitempty" xml:"Type,omitempty"`
	UpgradeReportList []*DescribeUpgradeReportResponseBodyUpgradeReportList `json:"UpgradeReportList,omitempty" xml:"UpgradeReportList,omitempty" type:"Repeated"`
}

func (s DescribeUpgradeReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeReportResponseBody) GetDetails() *string {
	return s.Details
}

func (s *DescribeUpgradeReportResponseBody) GetDstDBType() *string {
	return s.DstDBType
}

func (s *DescribeUpgradeReportResponseBody) GetItems() []*DescribeUpgradeReportResponseBodyItems {
	return s.Items
}

func (s *DescribeUpgradeReportResponseBody) GetItemsSize() *int64 {
	return s.ItemsSize
}

func (s *DescribeUpgradeReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUpgradeReportResponseBody) GetSourceDBClusterId() *string {
	return s.SourceDBClusterId
}

func (s *DescribeUpgradeReportResponseBody) GetSrcDBType() *string {
	return s.SrcDBType
}

func (s *DescribeUpgradeReportResponseBody) GetSrcDeleted() *string {
	return s.SrcDeleted
}

func (s *DescribeUpgradeReportResponseBody) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DescribeUpgradeReportResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeUpgradeReportResponseBody) GetUpgradeReportList() []*DescribeUpgradeReportResponseBodyUpgradeReportList {
	return s.UpgradeReportList
}

func (s *DescribeUpgradeReportResponseBody) SetDetails(v string) *DescribeUpgradeReportResponseBody {
	s.Details = &v
	return s
}

func (s *DescribeUpgradeReportResponseBody) SetDstDBType(v string) *DescribeUpgradeReportResponseBody {
	s.DstDBType = &v
	return s
}

func (s *DescribeUpgradeReportResponseBody) SetItems(v []*DescribeUpgradeReportResponseBodyItems) *DescribeUpgradeReportResponseBody {
	s.Items = v
	return s
}

func (s *DescribeUpgradeReportResponseBody) SetItemsSize(v int64) *DescribeUpgradeReportResponseBody {
	s.ItemsSize = &v
	return s
}

func (s *DescribeUpgradeReportResponseBody) SetRequestId(v string) *DescribeUpgradeReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUpgradeReportResponseBody) SetSourceDBClusterId(v string) *DescribeUpgradeReportResponseBody {
	s.SourceDBClusterId = &v
	return s
}

func (s *DescribeUpgradeReportResponseBody) SetSrcDBType(v string) *DescribeUpgradeReportResponseBody {
	s.SrcDBType = &v
	return s
}

func (s *DescribeUpgradeReportResponseBody) SetSrcDeleted(v string) *DescribeUpgradeReportResponseBody {
	s.SrcDeleted = &v
	return s
}

func (s *DescribeUpgradeReportResponseBody) SetTotalSize(v int64) *DescribeUpgradeReportResponseBody {
	s.TotalSize = &v
	return s
}

func (s *DescribeUpgradeReportResponseBody) SetType(v string) *DescribeUpgradeReportResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeUpgradeReportResponseBody) SetUpgradeReportList(v []*DescribeUpgradeReportResponseBodyUpgradeReportList) *DescribeUpgradeReportResponseBody {
	s.UpgradeReportList = v
	return s
}

func (s *DescribeUpgradeReportResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UpgradeReportList != nil {
		for _, item := range s.UpgradeReportList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUpgradeReportResponseBodyItems struct {
	// example:
	//
	// CREATE XXXX
	DDL *string `json:"DDL,omitempty" xml:"DDL,omitempty"`
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ny_openapi
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// orca
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeUpgradeReportResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeReportResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeReportResponseBodyItems) GetDDL() *string {
	return s.DDL
}

func (s *DescribeUpgradeReportResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeUpgradeReportResponseBodyItems) GetSchema() *string {
	return s.Schema
}

func (s *DescribeUpgradeReportResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeUpgradeReportResponseBodyItems) GetType() *string {
	return s.Type
}

func (s *DescribeUpgradeReportResponseBodyItems) SetDDL(v string) *DescribeUpgradeReportResponseBodyItems {
	s.DDL = &v
	return s
}

func (s *DescribeUpgradeReportResponseBodyItems) SetName(v string) *DescribeUpgradeReportResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeUpgradeReportResponseBodyItems) SetSchema(v string) *DescribeUpgradeReportResponseBodyItems {
	s.Schema = &v
	return s
}

func (s *DescribeUpgradeReportResponseBodyItems) SetStatus(v string) *DescribeUpgradeReportResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeUpgradeReportResponseBodyItems) SetType(v string) *DescribeUpgradeReportResponseBodyItems {
	s.Type = &v
	return s
}

func (s *DescribeUpgradeReportResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type DescribeUpgradeReportResponseBodyUpgradeReportList struct {
	// example:
	//
	// 2024-03-15T06:48:44Z
	CheckTime *string `json:"CheckTime,omitempty" xml:"CheckTime,omitempty"`
	// example:
	//
	// 8.0
	DstVersion *string `json:"DstVersion,omitempty" xml:"DstVersion,omitempty"`
	// example:
	//
	// 2024-03-08T06:48:44Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// example:
	//
	// 2024-03-08T06:48:44Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// running
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// pc-2ze54671qoz830za9
	SrcInsName *string `json:"SrcInsName,omitempty" xml:"SrcInsName,omitempty"`
	// example:
	//
	// 5.7
	SrcVersion *string `json:"SrcVersion,omitempty" xml:"SrcVersion,omitempty"`
	// example:
	//
	// 2025-07-05T01:56:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 275948
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// rds2polar_pengine_with_dts
	UpgradeMode *string `json:"UpgradeMode,omitempty" xml:"UpgradeMode,omitempty"`
}

func (s DescribeUpgradeReportResponseBodyUpgradeReportList) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeReportResponseBodyUpgradeReportList) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) GetCheckTime() *string {
	return s.CheckTime
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) GetDstVersion() *string {
	return s.DstVersion
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) GetResult() *string {
	return s.Result
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) GetSrcInsName() *string {
	return s.SrcInsName
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) GetSrcVersion() *string {
	return s.SrcVersion
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) GetUpgradeMode() *string {
	return s.UpgradeMode
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) SetCheckTime(v string) *DescribeUpgradeReportResponseBodyUpgradeReportList {
	s.CheckTime = &v
	return s
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) SetDstVersion(v string) *DescribeUpgradeReportResponseBodyUpgradeReportList {
	s.DstVersion = &v
	return s
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) SetEffectiveTime(v string) *DescribeUpgradeReportResponseBodyUpgradeReportList {
	s.EffectiveTime = &v
	return s
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) SetEndTime(v string) *DescribeUpgradeReportResponseBodyUpgradeReportList {
	s.EndTime = &v
	return s
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) SetResult(v string) *DescribeUpgradeReportResponseBodyUpgradeReportList {
	s.Result = &v
	return s
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) SetSrcInsName(v string) *DescribeUpgradeReportResponseBodyUpgradeReportList {
	s.SrcInsName = &v
	return s
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) SetSrcVersion(v string) *DescribeUpgradeReportResponseBodyUpgradeReportList {
	s.SrcVersion = &v
	return s
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) SetStartTime(v string) *DescribeUpgradeReportResponseBodyUpgradeReportList {
	s.StartTime = &v
	return s
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) SetTaskId(v string) *DescribeUpgradeReportResponseBodyUpgradeReportList {
	s.TaskId = &v
	return s
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) SetUpgradeMode(v string) *DescribeUpgradeReportResponseBodyUpgradeReportList {
	s.UpgradeMode = &v
	return s
}

func (s *DescribeUpgradeReportResponseBodyUpgradeReportList) Validate() error {
	return dara.Validate(s)
}
