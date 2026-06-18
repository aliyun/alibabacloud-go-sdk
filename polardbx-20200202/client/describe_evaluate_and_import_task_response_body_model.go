// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEvaluateAndImportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeEvaluateAndImportTaskResponseBodyData) *DescribeEvaluateAndImportTaskResponseBody
	GetData() *DescribeEvaluateAndImportTaskResponseBodyData
	SetMessage(v string) *DescribeEvaluateAndImportTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEvaluateAndImportTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeEvaluateAndImportTaskResponseBody
	GetSuccess() *bool
}

type DescribeEvaluateAndImportTaskResponseBody struct {
	// The returned result.
	Data *DescribeEvaluateAndImportTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message. This parameter is empty when the request succeeds. If the request fails, an exception message is returned, such as an error code.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 73559800-****-11ec-bd40-99cfcff3fe1e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEvaluateAndImportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEvaluateAndImportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEvaluateAndImportTaskResponseBody) GetData() *DescribeEvaluateAndImportTaskResponseBodyData {
	return s.Data
}

func (s *DescribeEvaluateAndImportTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEvaluateAndImportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEvaluateAndImportTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEvaluateAndImportTaskResponseBody) SetData(v *DescribeEvaluateAndImportTaskResponseBodyData) *DescribeEvaluateAndImportTaskResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBody) SetMessage(v string) *DescribeEvaluateAndImportTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBody) SetRequestId(v string) *DescribeEvaluateAndImportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBody) SetSuccess(v bool) *DescribeEvaluateAndImportTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEvaluateAndImportTaskResponseBodyData struct {
	// The business ID of the import task.
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The summary information.
	//
	// example:
	//
	// "{\\"structureImportResult\\":{\\"status\\":\\"finish\\",\\"finishedNum\\":382,\\"totalNum\\":382,\\"percentage\\":100},\\"switchTips\\":\\"正向同步延迟较小时才允许开始正向切换\\",\\"autoTableSuggestionSwitch\\":false,\\"structureAlterationMap\\":{},\\"acceptFilterMap\\":{\\"fr_import\\":[\\"dy_dr_new\\"],\\"e3plus_drp_prod_new\\":[],\\"qbtong_prod\\":[],\\"e3plus_account_prod\\":[],\\"e3plus_drp_prod\\":[]},\\"structChunkMap\\":{},\\"count\\":1,\\"requestId\\":\\"9380D025-C9B1-51D6-A4B9-733C4879CC4A\\"}",
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// The creator of the import task.
	//
	// example:
	//
	// 346650
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Indicates whether the import task is successful.
	//
	// example:
	//
	// false
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// The timestamp when the task was created.
	//
	// example:
	//
	// 2025-06-08T15:00Z
	GmtCreated *int64 `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The modification time in timestamp format.
	//
	// example:
	//
	// 1605080537000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 76251
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the slink destination database.
	//
	// example:
	//
	// fr_import
	SlinkDstDb *string `json:"SlinkDstDb,omitempty" xml:"SlinkDstDb,omitempty"`
	// The instance ID of the destination.
	//
	// example:
	//
	// pxc-hzrehc60jz5zye
	SlinkDstResId *string `json:"SlinkDstResId,omitempty" xml:"SlinkDstResId,omitempty"`
	// The username for the data connection of the destination.
	//
	// example:
	//
	// dts_dtdba
	SlinkDstUserName *string `json:"SlinkDstUserName,omitempty" xml:"SlinkDstUserName,omitempty"`
	// The name of the slink source database.
	//
	// example:
	//
	// fr_import
	SlinkSrcDb *string `json:"SlinkSrcDb,omitempty" xml:"SlinkSrcDb,omitempty"`
	// The instance ID of the source.
	//
	// example:
	//
	// drdshbga69378w4p
	SlinkSrcResId *string `json:"SlinkSrcResId,omitempty" xml:"SlinkSrcResId,omitempty"`
	// The resource type of the source.
	//
	// example:
	//
	// POLARX1
	SlinkSrcResType *string `json:"SlinkSrcResType,omitempty" xml:"SlinkSrcResType,omitempty"`
	// The username for the data connection of the source.
	//
	// example:
	//
	// dtdba
	SlinkSrcUserName *string `json:"SlinkSrcUserName,omitempty" xml:"SlinkSrcUserName,omitempty"`
	// The stage of the synchronization.
	//
	// example:
	//
	// DATA_IMPORT_RUNNING
	SlinkStage *string `json:"SlinkStage,omitempty" xml:"SlinkStage,omitempty"`
	// The running status of the synchronization.
	//
	// example:
	//
	// IMPORT_RUNNING
	SlinkStatus *string `json:"SlinkStatus,omitempty" xml:"SlinkStatus,omitempty"`
	// The description of the slink task.
	//
	// example:
	//
	// drdsprod_to_2.0
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

func (s DescribeEvaluateAndImportTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEvaluateAndImportTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetBid() *string {
	return s.Bid
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetContext() *string {
	return s.Context
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetDeleted() *bool {
	return s.Deleted
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetGmtCreated() *int64 {
	return s.GmtCreated
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetSlinkDstDb() *string {
	return s.SlinkDstDb
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetSlinkDstResId() *string {
	return s.SlinkDstResId
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetSlinkDstUserName() *string {
	return s.SlinkDstUserName
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetSlinkSrcDb() *string {
	return s.SlinkSrcDb
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetSlinkSrcResId() *string {
	return s.SlinkSrcResId
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetSlinkSrcResType() *string {
	return s.SlinkSrcResType
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetSlinkSrcUserName() *string {
	return s.SlinkSrcUserName
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetSlinkStage() *string {
	return s.SlinkStage
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetSlinkStatus() *string {
	return s.SlinkStatus
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetSlinkTaskDesc() *string {
	return s.SlinkTaskDesc
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) GetSlinkType() *string {
	return s.SlinkType
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetBid(v string) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.Bid = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetContext(v string) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.Context = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetCreator(v string) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.Creator = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetDeleted(v bool) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetGmtCreated(v int64) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetGmtModified(v int64) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetId(v int64) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetRegionId(v string) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetSlinkDstDb(v string) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.SlinkDstDb = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetSlinkDstResId(v string) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.SlinkDstResId = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetSlinkDstUserName(v string) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.SlinkDstUserName = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetSlinkSrcDb(v string) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.SlinkSrcDb = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetSlinkSrcResId(v string) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.SlinkSrcResId = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetSlinkSrcResType(v string) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.SlinkSrcResType = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetSlinkSrcUserName(v string) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.SlinkSrcUserName = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetSlinkStage(v string) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.SlinkStage = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetSlinkStatus(v string) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.SlinkStatus = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetSlinkTaskDesc(v string) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.SlinkTaskDesc = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetSlinkTaskId(v string) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.SlinkTaskId = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) SetSlinkType(v string) *DescribeEvaluateAndImportTaskResponseBodyData {
	s.SlinkType = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
