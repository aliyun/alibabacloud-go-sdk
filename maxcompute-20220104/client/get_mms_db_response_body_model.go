// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsDbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMmsDbResponseBodyData) *GetMmsDbResponseBody
	GetData() *GetMmsDbResponseBodyData
	SetRequestId(v string) *GetMmsDbResponseBody
	GetRequestId() *string
}

type GetMmsDbResponseBody struct {
	Data *GetMmsDbResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 90D64EB6-2962-5B1C-A039-BC41C8176C7F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsDbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMmsDbResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsDbResponseBody) GetData() *GetMmsDbResponseBodyData {
	return s.Data
}

func (s *GetMmsDbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMmsDbResponseBody) SetData(v *GetMmsDbResponseBodyData) *GetMmsDbResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsDbResponseBody) SetRequestId(v string) *GetMmsDbResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMmsDbResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMmsDbResponseBodyData struct {
	// example:
	//
	// for mms_test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// default
	DstName *string `json:"dstName,omitempty" xml:"dstName,omitempty"`
	// example:
	//
	// mma_test
	DstProjectName *string `json:"dstProjectName,omitempty" xml:"dstProjectName,omitempty"`
	// example:
	//
	// {}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// example:
	//
	// 63
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// last ddl time
	//
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTime *string `json:"lastDdlTime,omitempty" xml:"lastDdlTime,omitempty"`
	// example:
	//
	// hdfs://master-1-1.c-6fc187819ed6bae0.cn-shanghai.emr.aliyuncs.com:9000/user/hive/warehouse
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// mms_test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2323
	NumRows *int64 `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// example:
	//
	// System user
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// 2000
	Partitions *int32 `json:"partitions,omitempty" xml:"partitions,omitempty"`
	// example:
	//
	// 200
	PartitionsDoing *int32 `json:"partitionsDoing,omitempty" xml:"partitionsDoing,omitempty"`
	// example:
	//
	// 1400
	PartitionsDone *int32 `json:"partitionsDone,omitempty" xml:"partitionsDone,omitempty"`
	// example:
	//
	// 400
	PartitionsFailed *int32 `json:"partitionsFailed,omitempty" xml:"partitionsFailed,omitempty"`
	// example:
	//
	// 323232332
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 2000017
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// example:
	//
	// DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 200
	Tables *int32 `json:"tables,omitempty" xml:"tables,omitempty"`
	// example:
	//
	// 20
	TablesDoing *int32 `json:"tablesDoing,omitempty" xml:"tablesDoing,omitempty"`
	// example:
	//
	// 120
	TablesDone *int32 `json:"tablesDone,omitempty" xml:"tablesDone,omitempty"`
	// example:
	//
	// 20
	TablesFailed *int32 `json:"tablesFailed,omitempty" xml:"tablesFailed,omitempty"`
	// example:
	//
	// 20
	TablesPartDone *int32 `json:"tablesPartDone,omitempty" xml:"tablesPartDone,omitempty"`
	// example:
	//
	// true
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s GetMmsDbResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMmsDbResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsDbResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetMmsDbResponseBodyData) GetDstName() *string {
	return s.DstName
}

func (s *GetMmsDbResponseBodyData) GetDstProjectName() *string {
	return s.DstProjectName
}

func (s *GetMmsDbResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *GetMmsDbResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetMmsDbResponseBodyData) GetLastDdlTime() *string {
	return s.LastDdlTime
}

func (s *GetMmsDbResponseBodyData) GetLocation() *string {
	return s.Location
}

func (s *GetMmsDbResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetMmsDbResponseBodyData) GetNumRows() *int64 {
	return s.NumRows
}

func (s *GetMmsDbResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetMmsDbResponseBodyData) GetPartitions() *int32 {
	return s.Partitions
}

func (s *GetMmsDbResponseBodyData) GetPartitionsDoing() *int32 {
	return s.PartitionsDoing
}

func (s *GetMmsDbResponseBodyData) GetPartitionsDone() *int32 {
	return s.PartitionsDone
}

func (s *GetMmsDbResponseBodyData) GetPartitionsFailed() *int32 {
	return s.PartitionsFailed
}

func (s *GetMmsDbResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *GetMmsDbResponseBodyData) GetSourceId() *int64 {
	return s.SourceId
}

func (s *GetMmsDbResponseBodyData) GetSourceName() *string {
	return s.SourceName
}

func (s *GetMmsDbResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetMmsDbResponseBodyData) GetTables() *int32 {
	return s.Tables
}

func (s *GetMmsDbResponseBodyData) GetTablesDoing() *int32 {
	return s.TablesDoing
}

func (s *GetMmsDbResponseBodyData) GetTablesDone() *int32 {
	return s.TablesDone
}

func (s *GetMmsDbResponseBodyData) GetTablesFailed() *int32 {
	return s.TablesFailed
}

func (s *GetMmsDbResponseBodyData) GetTablesPartDone() *int32 {
	return s.TablesPartDone
}

func (s *GetMmsDbResponseBodyData) GetUpdated() *bool {
	return s.Updated
}

func (s *GetMmsDbResponseBodyData) SetDescription(v string) *GetMmsDbResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetDstName(v string) *GetMmsDbResponseBodyData {
	s.DstName = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetDstProjectName(v string) *GetMmsDbResponseBodyData {
	s.DstProjectName = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetExtra(v string) *GetMmsDbResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetId(v int64) *GetMmsDbResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetLastDdlTime(v string) *GetMmsDbResponseBodyData {
	s.LastDdlTime = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetLocation(v string) *GetMmsDbResponseBodyData {
	s.Location = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetName(v string) *GetMmsDbResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetNumRows(v int64) *GetMmsDbResponseBodyData {
	s.NumRows = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetOwner(v string) *GetMmsDbResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetPartitions(v int32) *GetMmsDbResponseBodyData {
	s.Partitions = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetPartitionsDoing(v int32) *GetMmsDbResponseBodyData {
	s.PartitionsDoing = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetPartitionsDone(v int32) *GetMmsDbResponseBodyData {
	s.PartitionsDone = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetPartitionsFailed(v int32) *GetMmsDbResponseBodyData {
	s.PartitionsFailed = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetSize(v int64) *GetMmsDbResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetSourceId(v int64) *GetMmsDbResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetSourceName(v string) *GetMmsDbResponseBodyData {
	s.SourceName = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetStatus(v string) *GetMmsDbResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetTables(v int32) *GetMmsDbResponseBodyData {
	s.Tables = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetTablesDoing(v int32) *GetMmsDbResponseBodyData {
	s.TablesDoing = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetTablesDone(v int32) *GetMmsDbResponseBodyData {
	s.TablesDone = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetTablesFailed(v int32) *GetMmsDbResponseBodyData {
	s.TablesFailed = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetTablesPartDone(v int32) *GetMmsDbResponseBodyData {
	s.TablesPartDone = &v
	return s
}

func (s *GetMmsDbResponseBodyData) SetUpdated(v bool) *GetMmsDbResponseBodyData {
	s.Updated = &v
	return s
}

func (s *GetMmsDbResponseBodyData) Validate() error {
	return dara.Validate(s)
}
