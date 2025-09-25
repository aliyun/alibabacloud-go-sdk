// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDialogAnalysisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *CreateDialogAnalysisTaskResponseBody
	GetCost() *int64
	SetData(v []*string) *CreateDialogAnalysisTaskResponseBody
	GetData() []*string
	SetDataType(v string) *CreateDialogAnalysisTaskResponseBody
	GetDataType() *string
	SetErrCode(v string) *CreateDialogAnalysisTaskResponseBody
	GetErrCode() *string
	SetMessage(v string) *CreateDialogAnalysisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDialogAnalysisTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDialogAnalysisTaskResponseBody
	GetSuccess() *bool
	SetTime(v string) *CreateDialogAnalysisTaskResponseBody
	GetTime() *string
}

type CreateDialogAnalysisTaskResponseBody struct {
	// example:
	//
	// null
	Cost *int64    `json:"cost,omitempty" xml:"cost,omitempty"`
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// EF4B5C9B-3BC8-5171-A47B-4C5CF3DC3258
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s CreateDialogAnalysisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDialogAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDialogAnalysisTaskResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *CreateDialogAnalysisTaskResponseBody) GetData() []*string {
	return s.Data
}

func (s *CreateDialogAnalysisTaskResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *CreateDialogAnalysisTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateDialogAnalysisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDialogAnalysisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDialogAnalysisTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDialogAnalysisTaskResponseBody) GetTime() *string {
	return s.Time
}

func (s *CreateDialogAnalysisTaskResponseBody) SetCost(v int64) *CreateDialogAnalysisTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateDialogAnalysisTaskResponseBody) SetData(v []*string) *CreateDialogAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateDialogAnalysisTaskResponseBody) SetDataType(v string) *CreateDialogAnalysisTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *CreateDialogAnalysisTaskResponseBody) SetErrCode(v string) *CreateDialogAnalysisTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateDialogAnalysisTaskResponseBody) SetMessage(v string) *CreateDialogAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDialogAnalysisTaskResponseBody) SetRequestId(v string) *CreateDialogAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDialogAnalysisTaskResponseBody) SetSuccess(v bool) *CreateDialogAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDialogAnalysisTaskResponseBody) SetTime(v string) *CreateDialogAnalysisTaskResponseBody {
	s.Time = &v
	return s
}

func (s *CreateDialogAnalysisTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
