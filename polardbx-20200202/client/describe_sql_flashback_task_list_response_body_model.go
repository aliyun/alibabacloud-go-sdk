// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlFlashbackTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeSqlFlashbackTaskListResponseBodyData) *DescribeSqlFlashbackTaskListResponseBody
	GetData() *DescribeSqlFlashbackTaskListResponseBodyData
	SetMessage(v string) *DescribeSqlFlashbackTaskListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSqlFlashbackTaskListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSqlFlashbackTaskListResponseBody
	GetSuccess() *bool
}

type DescribeSqlFlashbackTaskListResponseBody struct {
	Data *DescribeSqlFlashbackTaskListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSqlFlashbackTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlFlashbackTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbackTaskListResponseBody) GetData() *DescribeSqlFlashbackTaskListResponseBodyData {
	return s.Data
}

func (s *DescribeSqlFlashbackTaskListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSqlFlashbackTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSqlFlashbackTaskListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSqlFlashbackTaskListResponseBody) SetData(v *DescribeSqlFlashbackTaskListResponseBodyData) *DescribeSqlFlashbackTaskListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBody) SetMessage(v string) *DescribeSqlFlashbackTaskListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBody) SetRequestId(v string) *DescribeSqlFlashbackTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBody) SetSuccess(v bool) *DescribeSqlFlashbackTaskListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSqlFlashbackTaskListResponseBodyData struct {
	SqlFlashbackTasks []*DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks `json:"SqlFlashbackTasks,omitempty" xml:"SqlFlashbackTasks,omitempty" type:"Repeated"`
}

func (s DescribeSqlFlashbackTaskListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlFlashbackTaskListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbackTaskListResponseBodyData) GetSqlFlashbackTasks() []*DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	return s.SqlFlashbackTasks
}

func (s *DescribeSqlFlashbackTaskListResponseBodyData) SetSqlFlashbackTasks(v []*DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) *DescribeSqlFlashbackTaskListResponseBodyData {
	s.SqlFlashbackTasks = v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyData) Validate() error {
	if s.SqlFlashbackTasks != nil {
		for _, item := range s.SqlFlashbackTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks struct {
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// http://...
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// example:
	//
	// 1569216270000
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 1568611408000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1568611469000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 111
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// pxc-********
	InstId *string `json:"InstId,omitempty" xml:"InstId,omitempty"`
	// example:
	//
	// 10
	RecallProgress *string `json:"RecallProgress,omitempty" xml:"RecallProgress,omitempty"`
	// example:
	//
	// 0
	RecallRestoreType *string `json:"RecallRestoreType,omitempty" xml:"RecallRestoreType,omitempty"`
	// example:
	//
	// 1
	RecallStatus *string `json:"RecallStatus,omitempty" xml:"RecallStatus,omitempty"`
	// example:
	//
	// 0
	RecallType *string `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// example:
	//
	// 1568609597000
	SearchEndTime *string `json:"SearchEndTime,omitempty" xml:"SearchEndTime,omitempty"`
	// example:
	//
	// 1568609597000
	SearchStartTime *string `json:"SearchStartTime,omitempty" xml:"SearchStartTime,omitempty"`
	// example:
	//
	// 100
	SqlCounter *string `json:"SqlCounter,omitempty" xml:"SqlCounter,omitempty"`
	// example:
	//
	// id
	SqlPk *string `json:"SqlPk,omitempty" xml:"SqlPk,omitempty"`
	// example:
	//
	// INSERT,UPDATE
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 13ed05705f801000
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetDbName() *string {
	return s.DbName
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetId() *string {
	return s.Id
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetInstId() *string {
	return s.InstId
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetRecallProgress() *string {
	return s.RecallProgress
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetRecallRestoreType() *string {
	return s.RecallRestoreType
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetRecallStatus() *string {
	return s.RecallStatus
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetRecallType() *string {
	return s.RecallType
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetSearchEndTime() *string {
	return s.SearchEndTime
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetSearchStartTime() *string {
	return s.SearchStartTime
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetSqlCounter() *string {
	return s.SqlCounter
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetSqlPk() *string {
	return s.SqlPk
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetSqlType() *string {
	return s.SqlType
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetTableName() *string {
	return s.TableName
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetDbName(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.DbName = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetDownloadUrl(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetExpireTime(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.ExpireTime = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetGmtCreate(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetGmtModified(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.GmtModified = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetId(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.Id = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetInstId(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.InstId = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetRecallProgress(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.RecallProgress = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetRecallRestoreType(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.RecallRestoreType = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetRecallStatus(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.RecallStatus = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetRecallType(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.RecallType = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetSearchEndTime(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.SearchEndTime = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetSearchStartTime(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.SearchStartTime = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetSqlCounter(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.SqlCounter = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetSqlPk(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.SqlPk = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetSqlType(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.SqlType = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetTableName(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.TableName = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) SetTraceId(v string) *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks {
	s.TraceId = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponseBodyDataSqlFlashbackTasks) Validate() error {
	return dara.Validate(s)
}
