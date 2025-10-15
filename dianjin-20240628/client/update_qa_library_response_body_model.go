// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQaLibraryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *UpdateQaLibraryResponseBody
	GetCost() *int64
	SetData(v *UpdateQaLibraryResponseBodyData) *UpdateQaLibraryResponseBody
	GetData() *UpdateQaLibraryResponseBodyData
	SetDataType(v string) *UpdateQaLibraryResponseBody
	GetDataType() *string
	SetErrCode(v string) *UpdateQaLibraryResponseBody
	GetErrCode() *string
	SetMessage(v string) *UpdateQaLibraryResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateQaLibraryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateQaLibraryResponseBody
	GetSuccess() *bool
	SetTime(v string) *UpdateQaLibraryResponseBody
	GetTime() *string
}

type UpdateQaLibraryResponseBody struct {
	// example:
	//
	// null
	Cost *int64                           `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *UpdateQaLibraryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s UpdateQaLibraryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateQaLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQaLibraryResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *UpdateQaLibraryResponseBody) GetData() *UpdateQaLibraryResponseBodyData {
	return s.Data
}

func (s *UpdateQaLibraryResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *UpdateQaLibraryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateQaLibraryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateQaLibraryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateQaLibraryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateQaLibraryResponseBody) GetTime() *string {
	return s.Time
}

func (s *UpdateQaLibraryResponseBody) SetCost(v int64) *UpdateQaLibraryResponseBody {
	s.Cost = &v
	return s
}

func (s *UpdateQaLibraryResponseBody) SetData(v *UpdateQaLibraryResponseBodyData) *UpdateQaLibraryResponseBody {
	s.Data = v
	return s
}

func (s *UpdateQaLibraryResponseBody) SetDataType(v string) *UpdateQaLibraryResponseBody {
	s.DataType = &v
	return s
}

func (s *UpdateQaLibraryResponseBody) SetErrCode(v string) *UpdateQaLibraryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateQaLibraryResponseBody) SetMessage(v string) *UpdateQaLibraryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateQaLibraryResponseBody) SetRequestId(v string) *UpdateQaLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQaLibraryResponseBody) SetSuccess(v bool) *UpdateQaLibraryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateQaLibraryResponseBody) SetTime(v string) *UpdateQaLibraryResponseBody {
	s.Time = &v
	return s
}

func (s *UpdateQaLibraryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateQaLibraryResponseBodyData struct {
	// example:
	//
	// 6jh378d
	QaLibraryId *string `json:"qaLibraryId,omitempty" xml:"qaLibraryId,omitempty"`
}

func (s UpdateQaLibraryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateQaLibraryResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateQaLibraryResponseBodyData) GetQaLibraryId() *string {
	return s.QaLibraryId
}

func (s *UpdateQaLibraryResponseBodyData) SetQaLibraryId(v string) *UpdateQaLibraryResponseBodyData {
	s.QaLibraryId = &v
	return s
}

func (s *UpdateQaLibraryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
