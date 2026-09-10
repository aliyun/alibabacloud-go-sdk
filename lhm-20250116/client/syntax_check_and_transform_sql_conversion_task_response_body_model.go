// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyntaxCheckAndTransformSqlConversionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SyntaxCheckAndTransformSqlConversionTaskResponseBodyData) *SyntaxCheckAndTransformSqlConversionTaskResponseBody
	GetData() *SyntaxCheckAndTransformSqlConversionTaskResponseBodyData
	SetErrCode(v string) *SyntaxCheckAndTransformSqlConversionTaskResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *SyntaxCheckAndTransformSqlConversionTaskResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *SyntaxCheckAndTransformSqlConversionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyntaxCheckAndTransformSqlConversionTaskResponseBody
	GetSuccess() *bool
}

type SyntaxCheckAndTransformSqlConversionTaskResponseBody struct {
	Data *SyntaxCheckAndTransformSqlConversionTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyntaxCheckAndTransformSqlConversionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyntaxCheckAndTransformSqlConversionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponseBody) GetData() *SyntaxCheckAndTransformSqlConversionTaskResponseBodyData {
	return s.Data
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponseBody) SetData(v *SyntaxCheckAndTransformSqlConversionTaskResponseBodyData) *SyntaxCheckAndTransformSqlConversionTaskResponseBody {
	s.Data = v
	return s
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponseBody) SetErrCode(v string) *SyntaxCheckAndTransformSqlConversionTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponseBody) SetErrMessage(v string) *SyntaxCheckAndTransformSqlConversionTaskResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponseBody) SetRequestId(v string) *SyntaxCheckAndTransformSqlConversionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponseBody) SetSuccess(v bool) *SyntaxCheckAndTransformSqlConversionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SyntaxCheckAndTransformSqlConversionTaskResponseBodyData struct {
	// example:
	//
	// 10001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SyntaxCheckAndTransformSqlConversionTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SyntaxCheckAndTransformSqlConversionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponseBodyData) SetTaskId(v int64) *SyntaxCheckAndTransformSqlConversionTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
