// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLibraryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *CreateLibraryResponseBody
	GetCost() *int64
	SetData(v string) *CreateLibraryResponseBody
	GetData() *string
	SetDataType(v string) *CreateLibraryResponseBody
	GetDataType() *string
	SetErrCode(v string) *CreateLibraryResponseBody
	GetErrCode() *string
	SetMessage(v string) *CreateLibraryResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateLibraryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateLibraryResponseBody
	GetSuccess() *bool
	SetTime(v string) *CreateLibraryResponseBody
	GetTime() *string
}

type CreateLibraryResponseBody struct {
	// example:
	//
	// 300
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// a1b2c3
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// null
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s CreateLibraryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLibraryResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *CreateLibraryResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateLibraryResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *CreateLibraryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateLibraryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateLibraryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLibraryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateLibraryResponseBody) GetTime() *string {
	return s.Time
}

func (s *CreateLibraryResponseBody) SetCost(v int64) *CreateLibraryResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateLibraryResponseBody) SetData(v string) *CreateLibraryResponseBody {
	s.Data = &v
	return s
}

func (s *CreateLibraryResponseBody) SetDataType(v string) *CreateLibraryResponseBody {
	s.DataType = &v
	return s
}

func (s *CreateLibraryResponseBody) SetErrCode(v string) *CreateLibraryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateLibraryResponseBody) SetMessage(v string) *CreateLibraryResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLibraryResponseBody) SetRequestId(v string) *CreateLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLibraryResponseBody) SetSuccess(v bool) *CreateLibraryResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLibraryResponseBody) SetTime(v string) *CreateLibraryResponseBody {
	s.Time = &v
	return s
}

func (s *CreateLibraryResponseBody) Validate() error {
	return dara.Validate(s)
}
