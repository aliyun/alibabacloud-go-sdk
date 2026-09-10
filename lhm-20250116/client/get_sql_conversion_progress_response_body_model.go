// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlConversionProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSqlConversionProgressResponseBodyData) *GetSqlConversionProgressResponseBody
	GetData() *GetSqlConversionProgressResponseBodyData
	SetErrCode(v string) *GetSqlConversionProgressResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetSqlConversionProgressResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetSqlConversionProgressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSqlConversionProgressResponseBody
	GetSuccess() *bool
}

type GetSqlConversionProgressResponseBody struct {
	// The data body returned by the operation. For the field structure, see the child field descriptions.
	Data *GetSqlConversionProgressResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues with this call.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed. Check errCode and errMessage for troubleshooting.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSqlConversionProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConversionProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlConversionProgressResponseBody) GetData() *GetSqlConversionProgressResponseBodyData {
	return s.Data
}

func (s *GetSqlConversionProgressResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetSqlConversionProgressResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetSqlConversionProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSqlConversionProgressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSqlConversionProgressResponseBody) SetData(v *GetSqlConversionProgressResponseBodyData) *GetSqlConversionProgressResponseBody {
	s.Data = v
	return s
}

func (s *GetSqlConversionProgressResponseBody) SetErrCode(v string) *GetSqlConversionProgressResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetSqlConversionProgressResponseBody) SetErrMessage(v string) *GetSqlConversionProgressResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetSqlConversionProgressResponseBody) SetRequestId(v string) *GetSqlConversionProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlConversionProgressResponseBody) SetSuccess(v bool) *GetSqlConversionProgressResponseBody {
	s.Success = &v
	return s
}

func (s *GetSqlConversionProgressResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSqlConversionProgressResponseBodyData struct {
	// The number of failed scripts.
	//
	// example:
	//
	// 0
	Fail *int64 `json:"fail,omitempty" xml:"fail,omitempty"`
	// The number of completed scripts.
	//
	// example:
	//
	// 8
	Finish *int64 `json:"finish,omitempty" xml:"finish,omitempty"`
	// The completion percentage.
	//
	// example:
	//
	// 66.67
	Percent *float64 `json:"percent,omitempty" xml:"percent,omitempty"`
	// The number of scripts being converted.
	//
	// example:
	//
	// 4
	Running *int64 `json:"running,omitempty" xml:"running,omitempty"`
	// The total number of scripts.
	//
	// example:
	//
	// 12
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetSqlConversionProgressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConversionProgressResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSqlConversionProgressResponseBodyData) GetFail() *int64 {
	return s.Fail
}

func (s *GetSqlConversionProgressResponseBodyData) GetFinish() *int64 {
	return s.Finish
}

func (s *GetSqlConversionProgressResponseBodyData) GetPercent() *float64 {
	return s.Percent
}

func (s *GetSqlConversionProgressResponseBodyData) GetRunning() *int64 {
	return s.Running
}

func (s *GetSqlConversionProgressResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetSqlConversionProgressResponseBodyData) SetFail(v int64) *GetSqlConversionProgressResponseBodyData {
	s.Fail = &v
	return s
}

func (s *GetSqlConversionProgressResponseBodyData) SetFinish(v int64) *GetSqlConversionProgressResponseBodyData {
	s.Finish = &v
	return s
}

func (s *GetSqlConversionProgressResponseBodyData) SetPercent(v float64) *GetSqlConversionProgressResponseBodyData {
	s.Percent = &v
	return s
}

func (s *GetSqlConversionProgressResponseBodyData) SetRunning(v int64) *GetSqlConversionProgressResponseBodyData {
	s.Running = &v
	return s
}

func (s *GetSqlConversionProgressResponseBodyData) SetTotal(v int64) *GetSqlConversionProgressResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetSqlConversionProgressResponseBodyData) Validate() error {
	return dara.Validate(s)
}
