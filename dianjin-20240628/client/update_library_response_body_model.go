// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLibraryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *UpdateLibraryResponseBody
	GetCost() *int64
	SetData(v string) *UpdateLibraryResponseBody
	GetData() *string
	SetDataType(v string) *UpdateLibraryResponseBody
	GetDataType() *string
	SetErrCode(v string) *UpdateLibraryResponseBody
	GetErrCode() *string
	SetMessage(v string) *UpdateLibraryResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateLibraryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateLibraryResponseBody
	GetSuccess() *bool
	SetTime(v string) *UpdateLibraryResponseBody
	GetTime() *string
}

type UpdateLibraryResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// null
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
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
	// eb2b6139-ddf1-91a0-a47f-df7617ae9032
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

func (s UpdateLibraryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLibraryResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *UpdateLibraryResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateLibraryResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *UpdateLibraryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateLibraryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateLibraryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLibraryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateLibraryResponseBody) GetTime() *string {
	return s.Time
}

func (s *UpdateLibraryResponseBody) SetCost(v int64) *UpdateLibraryResponseBody {
	s.Cost = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetData(v string) *UpdateLibraryResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetDataType(v string) *UpdateLibraryResponseBody {
	s.DataType = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetErrCode(v string) *UpdateLibraryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetMessage(v string) *UpdateLibraryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetRequestId(v string) *UpdateLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetSuccess(v bool) *UpdateLibraryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetTime(v string) *UpdateLibraryResponseBody {
	s.Time = &v
	return s
}

func (s *UpdateLibraryResponseBody) Validate() error {
	return dara.Validate(s)
}
