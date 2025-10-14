// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStructureImportTaskInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeStructureImportTaskInfoResponseBodyData) *DescribeStructureImportTaskInfoResponseBody
	GetData() *DescribeStructureImportTaskInfoResponseBodyData
	SetMessage(v string) *DescribeStructureImportTaskInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeStructureImportTaskInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeStructureImportTaskInfoResponseBody
	GetSuccess() *bool
}

type DescribeStructureImportTaskInfoResponseBody struct {
	Data *DescribeStructureImportTaskInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeStructureImportTaskInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStructureImportTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStructureImportTaskInfoResponseBody) GetData() *DescribeStructureImportTaskInfoResponseBodyData {
	return s.Data
}

func (s *DescribeStructureImportTaskInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeStructureImportTaskInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStructureImportTaskInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeStructureImportTaskInfoResponseBody) SetData(v *DescribeStructureImportTaskInfoResponseBodyData) *DescribeStructureImportTaskInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeStructureImportTaskInfoResponseBody) SetMessage(v string) *DescribeStructureImportTaskInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeStructureImportTaskInfoResponseBody) SetRequestId(v string) *DescribeStructureImportTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStructureImportTaskInfoResponseBody) SetSuccess(v bool) *DescribeStructureImportTaskInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeStructureImportTaskInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeStructureImportTaskInfoResponseBodyData struct {
	// example:
	//
	// STRUCTURE_IMPORT
	SlinkStage            *string                                                               `json:"SlinkStage,omitempty" xml:"SlinkStage,omitempty"`
	StructureImportResult *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult `json:"StructureImportResult,omitempty" xml:"StructureImportResult,omitempty" type:"Struct"`
}

func (s DescribeStructureImportTaskInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeStructureImportTaskInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeStructureImportTaskInfoResponseBodyData) GetSlinkStage() *string {
	return s.SlinkStage
}

func (s *DescribeStructureImportTaskInfoResponseBodyData) GetStructureImportResult() *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult {
	return s.StructureImportResult
}

func (s *DescribeStructureImportTaskInfoResponseBodyData) SetSlinkStage(v string) *DescribeStructureImportTaskInfoResponseBodyData {
	s.SlinkStage = &v
	return s
}

func (s *DescribeStructureImportTaskInfoResponseBodyData) SetStructureImportResult(v *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult) *DescribeStructureImportTaskInfoResponseBodyData {
	s.StructureImportResult = v
	return s
}

func (s *DescribeStructureImportTaskInfoResponseBodyData) Validate() error {
	if s.StructureImportResult != nil {
		if err := s.StructureImportResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult struct {
	// example:
	//
	// java.sql.SQLException: [1a7a5f22aa403000][10.239.190.4:3058][hash_realtime_new]ERR-CODE: [TDDL-5123][ERR_INSTANCE_READ_ONLY_OPTION_NOT_SUPPORT] server is running with the instance-read-only option so it cannot execute this statement
	ExceptionDetail *string `json:"ExceptionDetail,omitempty" xml:"ExceptionDetail,omitempty"`
	// example:
	//
	// hash_realtime_new.wm_in_job_et
	ExceptionFullTableName *string `json:"ExceptionFullTableName,omitempty" xml:"ExceptionFullTableName,omitempty"`
	// example:
	//
	// 118
	FinishedNum *int32 `json:"FinishedNum,omitempty" xml:"FinishedNum,omitempty"`
	// example:
	//
	// 100
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// example:
	//
	// finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult) GoString() string {
	return s.String()
}

func (s *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult) GetExceptionDetail() *string {
	return s.ExceptionDetail
}

func (s *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult) GetExceptionFullTableName() *string {
	return s.ExceptionFullTableName
}

func (s *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult) GetFinishedNum() *int32 {
	return s.FinishedNum
}

func (s *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult) GetPercentage() *int32 {
	return s.Percentage
}

func (s *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult) GetStatus() *string {
	return s.Status
}

func (s *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult) SetExceptionDetail(v string) *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult {
	s.ExceptionDetail = &v
	return s
}

func (s *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult) SetExceptionFullTableName(v string) *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult {
	s.ExceptionFullTableName = &v
	return s
}

func (s *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult) SetFinishedNum(v int32) *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult {
	s.FinishedNum = &v
	return s
}

func (s *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult) SetPercentage(v int32) *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult {
	s.Percentage = &v
	return s
}

func (s *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult) SetStatus(v string) *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult {
	s.Status = &v
	return s
}

func (s *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult) SetTotalNum(v int32) *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult {
	s.TotalNum = &v
	return s
}

func (s *DescribeStructureImportTaskInfoResponseBodyDataStructureImportResult) Validate() error {
	return dara.Validate(s)
}
