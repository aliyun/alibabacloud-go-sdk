// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStructSyncJobAnalyzeResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetStructSyncJobAnalyzeResultResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetStructSyncJobAnalyzeResultResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetStructSyncJobAnalyzeResultResponseBody
	GetRequestId() *string
	SetStructSyncJobAnalyzeResult(v *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult) *GetStructSyncJobAnalyzeResultResponseBody
	GetStructSyncJobAnalyzeResult() *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult
	SetSuccess(v bool) *GetStructSyncJobAnalyzeResultResponseBody
	GetSuccess() *bool
}

type GetStructSyncJobAnalyzeResultResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1810E635-A2D7-428B-BAA9-85DAEB9B1A77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The analysis result of the schema synchronization task.
	StructSyncJobAnalyzeResult *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult `json:"StructSyncJobAnalyzeResult,omitempty" xml:"StructSyncJobAnalyzeResult,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStructSyncJobAnalyzeResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncJobAnalyzeResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobAnalyzeResultResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetStructSyncJobAnalyzeResultResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetStructSyncJobAnalyzeResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStructSyncJobAnalyzeResultResponseBody) GetStructSyncJobAnalyzeResult() *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult {
	return s.StructSyncJobAnalyzeResult
}

func (s *GetStructSyncJobAnalyzeResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStructSyncJobAnalyzeResultResponseBody) SetErrorCode(v string) *GetStructSyncJobAnalyzeResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBody) SetErrorMessage(v string) *GetStructSyncJobAnalyzeResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBody) SetRequestId(v string) *GetStructSyncJobAnalyzeResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBody) SetStructSyncJobAnalyzeResult(v *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult) *GetStructSyncJobAnalyzeResultResponseBody {
	s.StructSyncJobAnalyzeResult = v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBody) SetSuccess(v bool) *GetStructSyncJobAnalyzeResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBody) Validate() error {
	if s.StructSyncJobAnalyzeResult != nil {
		if err := s.StructSyncJobAnalyzeResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult struct {
	// The details of the analysis results.
	ResultList []*GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
	// The statistics on the analysis results.
	SummaryList []*GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList `json:"SummaryList,omitempty" xml:"SummaryList,omitempty" type:"Repeated"`
}

func (s GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult) GetResultList() []*GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList {
	return s.ResultList
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult) GetSummaryList() []*GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList {
	return s.SummaryList
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult) SetResultList(v []*GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList) *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult {
	s.ResultList = v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult) SetSummaryList(v []*GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList) *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult {
	s.SummaryList = v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult) Validate() error {
	if s.ResultList != nil {
		for _, item := range s.ResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SummaryList != nil {
		for _, item := range s.SummaryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList struct {
	// The SQL script.
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// The name of the source table.
	//
	// example:
	//
	// helloz_bak
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	// The name of the destination table.
	//
	// example:
	//
	// helloz_bak
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
}

func (s GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList) GetScript() *string {
	return s.Script
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList) GetSourceTableName() *string {
	return s.SourceTableName
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList) GetTargetTableName() *string {
	return s.TargetTableName
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList) SetScript(v string) *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList {
	s.Script = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList) SetSourceTableName(v string) *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList {
	s.SourceTableName = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList) SetTargetTableName(v string) *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList {
	s.TargetTableName = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList) Validate() error {
	return dara.Validate(s)
}

type GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList struct {
	// The type of the comparison. Valid values:
	//
	// 	- **CREATE_TABLE**: compares the created tables.
	//
	// 	- **ALTER_TABLE**: compares the modified tables.
	//
	// 	- **EQUAL_TABLE**: compares the identical tables.
	//
	// 	- **PASS_TABLE**: compares the tables that are skipped during schema synchronization.
	//
	// 	- **NOT_COMPARE**: does not compare tables.
	//
	// example:
	//
	// CREATE_TABLE
	CompareType *string `json:"CompareType,omitempty" xml:"CompareType,omitempty"`
	// The number of tables.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList) GetCompareType() *string {
	return s.CompareType
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList) GetCount() *int64 {
	return s.Count
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList) SetCompareType(v string) *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList {
	s.CompareType = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList) SetCount(v int64) *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList {
	s.Count = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList) Validate() error {
	return dara.Validate(s)
}
