// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageAmountSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetStorageAmountSummaryResponseBodyData) *GetStorageAmountSummaryResponseBody
	GetData() *GetStorageAmountSummaryResponseBodyData
	SetErrorCode(v string) *GetStorageAmountSummaryResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetStorageAmountSummaryResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *GetStorageAmountSummaryResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetStorageAmountSummaryResponseBody
	GetRequestId() *string
}

type GetStorageAmountSummaryResponseBody struct {
	Data *GetStorageAmountSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// this quota is not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 688003E1-D1B4-5468-957E-2FFB3AC8D79B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetStorageAmountSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStorageAmountSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetStorageAmountSummaryResponseBody) GetData() *GetStorageAmountSummaryResponseBodyData {
	return s.Data
}

func (s *GetStorageAmountSummaryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetStorageAmountSummaryResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetStorageAmountSummaryResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetStorageAmountSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStorageAmountSummaryResponseBody) SetData(v *GetStorageAmountSummaryResponseBodyData) *GetStorageAmountSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetStorageAmountSummaryResponseBody) SetErrorCode(v string) *GetStorageAmountSummaryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetStorageAmountSummaryResponseBody) SetErrorMsg(v string) *GetStorageAmountSummaryResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetStorageAmountSummaryResponseBody) SetHttpCode(v int32) *GetStorageAmountSummaryResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetStorageAmountSummaryResponseBody) SetRequestId(v string) *GetStorageAmountSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStorageAmountSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetStorageAmountSummaryResponseBodyData struct {
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// -
	Timestamp *int64             `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Unit      map[string]*string `json:"unit,omitempty" xml:"unit,omitempty"`
	Value     map[string]*int64  `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetStorageAmountSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetStorageAmountSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStorageAmountSummaryResponseBodyData) GetDate() *string {
	return s.Date
}

func (s *GetStorageAmountSummaryResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetStorageAmountSummaryResponseBodyData) GetUnit() map[string]*string {
	return s.Unit
}

func (s *GetStorageAmountSummaryResponseBodyData) GetValue() map[string]*int64 {
	return s.Value
}

func (s *GetStorageAmountSummaryResponseBodyData) SetDate(v string) *GetStorageAmountSummaryResponseBodyData {
	s.Date = &v
	return s
}

func (s *GetStorageAmountSummaryResponseBodyData) SetTimestamp(v int64) *GetStorageAmountSummaryResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *GetStorageAmountSummaryResponseBodyData) SetUnit(v map[string]*string) *GetStorageAmountSummaryResponseBodyData {
	s.Unit = v
	return s
}

func (s *GetStorageAmountSummaryResponseBodyData) SetValue(v map[string]*int64) *GetStorageAmountSummaryResponseBodyData {
	s.Value = v
	return s
}

func (s *GetStorageAmountSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
