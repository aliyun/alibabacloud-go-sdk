// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEvaluateAndImportTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeEvaluateAndImportTasksResponseBodyData) *DescribeEvaluateAndImportTasksResponseBody
	GetData() []*DescribeEvaluateAndImportTasksResponseBodyData
	SetMessage(v string) *DescribeEvaluateAndImportTasksResponseBody
	GetMessage() *string
	SetPageNumber(v int64) *DescribeEvaluateAndImportTasksResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeEvaluateAndImportTasksResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeEvaluateAndImportTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeEvaluateAndImportTasksResponseBody
	GetSuccess() *bool
	SetTotalNumber(v int64) *DescribeEvaluateAndImportTasksResponseBody
	GetTotalNumber() *int64
}

type DescribeEvaluateAndImportTasksResponseBody struct {
	// The task details.
	Data []*DescribeEvaluateAndImportTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message. > This parameter is empty when the request succeeds. When the request fails, exception information (such as error codes) is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 0
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 81519FDE-713B-****-B6F1-231479A4C9DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 0
	TotalNumber *int64 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s DescribeEvaluateAndImportTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEvaluateAndImportTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEvaluateAndImportTasksResponseBody) GetData() []*DescribeEvaluateAndImportTasksResponseBodyData {
	return s.Data
}

func (s *DescribeEvaluateAndImportTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEvaluateAndImportTasksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeEvaluateAndImportTasksResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeEvaluateAndImportTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEvaluateAndImportTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEvaluateAndImportTasksResponseBody) GetTotalNumber() *int64 {
	return s.TotalNumber
}

func (s *DescribeEvaluateAndImportTasksResponseBody) SetData(v []*DescribeEvaluateAndImportTasksResponseBodyData) *DescribeEvaluateAndImportTasksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBody) SetMessage(v string) *DescribeEvaluateAndImportTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBody) SetPageNumber(v int64) *DescribeEvaluateAndImportTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBody) SetPageSize(v int64) *DescribeEvaluateAndImportTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBody) SetRequestId(v string) *DescribeEvaluateAndImportTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBody) SetSuccess(v bool) *DescribeEvaluateAndImportTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBody) SetTotalNumber(v int64) *DescribeEvaluateAndImportTasksResponseBody {
	s.TotalNumber = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBody) Validate() error {
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

type DescribeEvaluateAndImportTasksResponseBodyData struct {
	// The business ID of the import task.
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The creator of the import task.
	//
	// example:
	//
	// 346650
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Indicates whether the task is deleted.
	//
	// example:
	//
	// fase
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// The task creation timestamp, in milliseconds.
	//
	// example:
	//
	// 2025-01-08T15:00Z
	GmtCreated *int64 `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The modification time of the metadata information.
	//
	// example:
	//
	// 2025-04-02T02:10:59Z
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ap-northeast-1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the slink destination database.
	//
	// example:
	//
	// e3plus_*****_prod
	SlinkDstDb *string `json:"SlinkDstDb,omitempty" xml:"SlinkDstDb,omitempty"`
	// The instance ID of the destination.
	//
	// example:
	//
	// pxc-hzrehc****5zye
	SlinkDstResId *string `json:"SlinkDstResId,omitempty" xml:"SlinkDstResId,omitempty"`
	// The username for the data connection of the destination.
	//
	// example:
	//
	// dts_temp_2_23
	SlinkDstUserName *string `json:"SlinkDstUserName,omitempty" xml:"SlinkDstUserName,omitempty"`
	// The name of the slink source database.
	//
	// example:
	//
	// stock
	SlinkSrcDb *string `json:"SlinkSrcDb,omitempty" xml:"SlinkSrcDb,omitempty"`
	// The instance ID of the source.
	//
	// example:
	//
	// drdshbga*****w4p
	SlinkSrcResId *string `json:"SlinkSrcResId,omitempty" xml:"SlinkSrcResId,omitempty"`
	// The resource type of the source.
	//
	// example:
	//
	// RDS_MYSQL
	SlinkSrcResType *string `json:"SlinkSrcResType,omitempty" xml:"SlinkSrcResType,omitempty"`
	// The username for the data connection of the source.
	//
	// example:
	//
	// dtdba
	SlinkSrcUserName *string `json:"SlinkSrcUserName,omitempty" xml:"SlinkSrcUserName,omitempty"`
	// The stage of the synchronization task.
	//
	// example:
	//
	// DATA_IMPORT_RUNNING
	SlinkStage *string `json:"SlinkStage,omitempty" xml:"SlinkStage,omitempty"`
	// The running status of the synchronization task.
	//
	// example:
	//
	// IMPORT_NOT_BEGIN
	SlinkStatus *string `json:"SlinkStatus,omitempty" xml:"SlinkStatus,omitempty"`
	// The description of the slink task.
	//
	// example:
	//
	// test
	SlinkTaskDesc *string `json:"SlinkTaskDesc,omitempty" xml:"SlinkTaskDesc,omitempty"`
	// The slink task ID.
	//
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
	// The synchronization type.
	//
	// example:
	//
	// evaluate_import
	SlinkType *string `json:"SlinkType,omitempty" xml:"SlinkType,omitempty"`
}

func (s DescribeEvaluateAndImportTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEvaluateAndImportTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetBid() *string {
	return s.Bid
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetDeleted() *bool {
	return s.Deleted
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetGmtCreated() *int64 {
	return s.GmtCreated
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetSlinkDstDb() *string {
	return s.SlinkDstDb
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetSlinkDstResId() *string {
	return s.SlinkDstResId
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetSlinkDstUserName() *string {
	return s.SlinkDstUserName
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetSlinkSrcDb() *string {
	return s.SlinkSrcDb
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetSlinkSrcResId() *string {
	return s.SlinkSrcResId
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetSlinkSrcResType() *string {
	return s.SlinkSrcResType
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetSlinkSrcUserName() *string {
	return s.SlinkSrcUserName
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetSlinkStage() *string {
	return s.SlinkStage
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetSlinkStatus() *string {
	return s.SlinkStatus
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetSlinkTaskDesc() *string {
	return s.SlinkTaskDesc
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) GetSlinkType() *string {
	return s.SlinkType
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetBid(v string) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.Bid = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetCreator(v string) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.Creator = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetDeleted(v bool) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetGmtCreated(v int64) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetGmtModified(v int64) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetId(v int64) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.Id = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetRegionId(v string) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetSlinkDstDb(v string) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.SlinkDstDb = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetSlinkDstResId(v string) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.SlinkDstResId = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetSlinkDstUserName(v string) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.SlinkDstUserName = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetSlinkSrcDb(v string) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.SlinkSrcDb = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetSlinkSrcResId(v string) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.SlinkSrcResId = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetSlinkSrcResType(v string) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.SlinkSrcResType = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetSlinkSrcUserName(v string) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.SlinkSrcUserName = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetSlinkStage(v string) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.SlinkStage = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetSlinkStatus(v string) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.SlinkStatus = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetSlinkTaskDesc(v string) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.SlinkTaskDesc = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetSlinkTaskId(v string) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.SlinkTaskId = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) SetSlinkType(v string) *DescribeEvaluateAndImportTasksResponseBodyData {
	s.SlinkType = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
