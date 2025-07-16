// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupSetListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeBackupSetListResponseBodyData) *DescribeBackupSetListResponseBody
	GetData() []*DescribeBackupSetListResponseBodyData
	SetMessage(v string) *DescribeBackupSetListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeBackupSetListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackupSetListResponseBody
	GetSuccess() *bool
}

type DescribeBackupSetListResponseBody struct {
	Data []*DescribeBackupSetListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1A6D328C-84B8-40DC-BF49-6C73984D7494
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupSetListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetListResponseBody) GetData() []*DescribeBackupSetListResponseBodyData {
	return s.Data
}

func (s *DescribeBackupSetListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeBackupSetListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupSetListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupSetListResponseBody) SetData(v []*DescribeBackupSetListResponseBodyData) *DescribeBackupSetListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBackupSetListResponseBody) SetMessage(v string) *DescribeBackupSetListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBackupSetListResponseBody) SetRequestId(v string) *DescribeBackupSetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupSetListResponseBody) SetSuccess(v bool) *DescribeBackupSetListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupSetListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupSetListResponseBodyData struct {
	// example:
	//
	// 0
	BackupModel *int32 `json:"BackupModel,omitempty" xml:"BackupModel,omitempty"`
	// example:
	//
	// 111
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// example:
	//
	// 88803195
	BackupSetSize *int64 `json:"BackupSetSize,omitempty" xml:"BackupSetSize,omitempty"`
	// example:
	//
	// 1
	BackupType *int32 `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// example:
	//
	// 1635706960956
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 1635706960956
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBackupSetListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetListResponseBodyData) GetBackupModel() *int32 {
	return s.BackupModel
}

func (s *DescribeBackupSetListResponseBodyData) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *DescribeBackupSetListResponseBodyData) GetBackupSetSize() *int64 {
	return s.BackupSetSize
}

func (s *DescribeBackupSetListResponseBodyData) GetBackupType() *int32 {
	return s.BackupType
}

func (s *DescribeBackupSetListResponseBodyData) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeBackupSetListResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeBackupSetListResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeBackupSetListResponseBodyData) SetBackupModel(v int32) *DescribeBackupSetListResponseBodyData {
	s.BackupModel = &v
	return s
}

func (s *DescribeBackupSetListResponseBodyData) SetBackupSetId(v string) *DescribeBackupSetListResponseBodyData {
	s.BackupSetId = &v
	return s
}

func (s *DescribeBackupSetListResponseBodyData) SetBackupSetSize(v int64) *DescribeBackupSetListResponseBodyData {
	s.BackupSetSize = &v
	return s
}

func (s *DescribeBackupSetListResponseBodyData) SetBackupType(v int32) *DescribeBackupSetListResponseBodyData {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupSetListResponseBodyData) SetBeginTime(v int64) *DescribeBackupSetListResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *DescribeBackupSetListResponseBodyData) SetEndTime(v int64) *DescribeBackupSetListResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupSetListResponseBodyData) SetStatus(v int32) *DescribeBackupSetListResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeBackupSetListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
