// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataExportPreCheckDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetDataExportPreCheckDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataExportPreCheckDetailResponseBody
	GetErrorMessage() *string
	SetPreCheckResult(v *GetDataExportPreCheckDetailResponseBodyPreCheckResult) *GetDataExportPreCheckDetailResponseBody
	GetPreCheckResult() *GetDataExportPreCheckDetailResponseBodyPreCheckResult
	SetRequestId(v string) *GetDataExportPreCheckDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataExportPreCheckDetailResponseBody
	GetSuccess() *bool
}

type GetDataExportPreCheckDetailResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates the result of the precheck task.
	PreCheckResult *GetDataExportPreCheckDetailResponseBodyPreCheckResult `json:"PreCheckResult,omitempty" xml:"PreCheckResult,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C1D39814-9808-47F8-AFE0-AF167239AC9B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataExportPreCheckDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataExportPreCheckDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataExportPreCheckDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataExportPreCheckDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataExportPreCheckDetailResponseBody) GetPreCheckResult() *GetDataExportPreCheckDetailResponseBodyPreCheckResult {
	return s.PreCheckResult
}

func (s *GetDataExportPreCheckDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataExportPreCheckDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataExportPreCheckDetailResponseBody) SetErrorCode(v string) *GetDataExportPreCheckDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataExportPreCheckDetailResponseBody) SetErrorMessage(v string) *GetDataExportPreCheckDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataExportPreCheckDetailResponseBody) SetPreCheckResult(v *GetDataExportPreCheckDetailResponseBodyPreCheckResult) *GetDataExportPreCheckDetailResponseBody {
	s.PreCheckResult = v
	return s
}

func (s *GetDataExportPreCheckDetailResponseBody) SetRequestId(v string) *GetDataExportPreCheckDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataExportPreCheckDetailResponseBody) SetSuccess(v bool) *GetDataExportPreCheckDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataExportPreCheckDetailResponseBody) Validate() error {
	if s.PreCheckResult != nil {
		if err := s.PreCheckResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataExportPreCheckDetailResponseBodyPreCheckResult struct {
	// Specifies whether to skip verification. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	IgnoreAffectRows *bool `json:"IgnoreAffectRows,omitempty" xml:"IgnoreAffectRows,omitempty"`
	// The list of pre-check details.
	PreCheckDetailList *GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailList `json:"PreCheckDetailList,omitempty" xml:"PreCheckDetailList,omitempty" type:"Struct"`
}

func (s GetDataExportPreCheckDetailResponseBodyPreCheckResult) String() string {
	return dara.Prettify(s)
}

func (s GetDataExportPreCheckDetailResponseBodyPreCheckResult) GoString() string {
	return s.String()
}

func (s *GetDataExportPreCheckDetailResponseBodyPreCheckResult) GetIgnoreAffectRows() *bool {
	return s.IgnoreAffectRows
}

func (s *GetDataExportPreCheckDetailResponseBodyPreCheckResult) GetPreCheckDetailList() *GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailList {
	return s.PreCheckDetailList
}

func (s *GetDataExportPreCheckDetailResponseBodyPreCheckResult) SetIgnoreAffectRows(v bool) *GetDataExportPreCheckDetailResponseBodyPreCheckResult {
	s.IgnoreAffectRows = &v
	return s
}

func (s *GetDataExportPreCheckDetailResponseBodyPreCheckResult) SetPreCheckDetailList(v *GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailList) *GetDataExportPreCheckDetailResponseBodyPreCheckResult {
	s.PreCheckDetailList = v
	return s
}

func (s *GetDataExportPreCheckDetailResponseBodyPreCheckResult) Validate() error {
	if s.PreCheckDetailList != nil {
		if err := s.PreCheckDetailList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailList struct {
	PreCheckDetailList []*GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailListPreCheckDetailList `json:"PreCheckDetailList,omitempty" xml:"PreCheckDetailList,omitempty" type:"Repeated"`
}

func (s GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailList) String() string {
	return dara.Prettify(s)
}

func (s GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailList) GoString() string {
	return s.String()
}

func (s *GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailList) GetPreCheckDetailList() []*GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailListPreCheckDetailList {
	return s.PreCheckDetailList
}

func (s *GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailList) SetPreCheckDetailList(v []*GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailListPreCheckDetailList) *GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailList {
	s.PreCheckDetailList = v
	return s
}

func (s *GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailList) Validate() error {
	if s.PreCheckDetailList != nil {
		for _, item := range s.PreCheckDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailListPreCheckDetailList struct {
	// The estimated number of data rows to be affected.
	//
	// example:
	//
	// 1
	AffectRows *int64 `json:"AffectRows,omitempty" xml:"AffectRows,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// SELECT 	- FROM tmp_table LIMIT 1
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
}

func (s GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailListPreCheckDetailList) String() string {
	return dara.Prettify(s)
}

func (s GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailListPreCheckDetailList) GoString() string {
	return s.String()
}

func (s *GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailListPreCheckDetailList) GetAffectRows() *int64 {
	return s.AffectRows
}

func (s *GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailListPreCheckDetailList) GetSQL() *string {
	return s.SQL
}

func (s *GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailListPreCheckDetailList) SetAffectRows(v int64) *GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailListPreCheckDetailList {
	s.AffectRows = &v
	return s
}

func (s *GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailListPreCheckDetailList) SetSQL(v string) *GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailListPreCheckDetailList {
	s.SQL = &v
	return s
}

func (s *GetDataExportPreCheckDetailResponseBodyPreCheckResultPreCheckDetailListPreCheckDetailList) Validate() error {
	return dara.Validate(s)
}
