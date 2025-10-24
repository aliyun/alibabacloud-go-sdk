// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageSizeSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetStorageSizeSummaryResponseBodyData) *GetStorageSizeSummaryResponseBody
	GetData() *GetStorageSizeSummaryResponseBodyData
	SetErrorCode(v string) *GetStorageSizeSummaryResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetStorageSizeSummaryResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *GetStorageSizeSummaryResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetStorageSizeSummaryResponseBody
	GetRequestId() *string
}

type GetStorageSizeSummaryResponseBody struct {
	Data *GetStorageSizeSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 0be3e0aa16667684362147582e038f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetStorageSizeSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStorageSizeSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetStorageSizeSummaryResponseBody) GetData() *GetStorageSizeSummaryResponseBodyData {
	return s.Data
}

func (s *GetStorageSizeSummaryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetStorageSizeSummaryResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetStorageSizeSummaryResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetStorageSizeSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStorageSizeSummaryResponseBody) SetData(v *GetStorageSizeSummaryResponseBodyData) *GetStorageSizeSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetStorageSizeSummaryResponseBody) SetErrorCode(v string) *GetStorageSizeSummaryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetStorageSizeSummaryResponseBody) SetErrorMsg(v string) *GetStorageSizeSummaryResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetStorageSizeSummaryResponseBody) SetHttpCode(v int32) *GetStorageSizeSummaryResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetStorageSizeSummaryResponseBody) SetRequestId(v string) *GetStorageSizeSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStorageSizeSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStorageSizeSummaryResponseBodyData struct {
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// 1749090705919
	Timestamp *int64              `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Unit      map[string]*string  `json:"unit,omitempty" xml:"unit,omitempty"`
	Value     map[string]*float64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetStorageSizeSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetStorageSizeSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStorageSizeSummaryResponseBodyData) GetDate() *string {
	return s.Date
}

func (s *GetStorageSizeSummaryResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetStorageSizeSummaryResponseBodyData) GetUnit() map[string]*string {
	return s.Unit
}

func (s *GetStorageSizeSummaryResponseBodyData) GetValue() map[string]*float64 {
	return s.Value
}

func (s *GetStorageSizeSummaryResponseBodyData) SetDate(v string) *GetStorageSizeSummaryResponseBodyData {
	s.Date = &v
	return s
}

func (s *GetStorageSizeSummaryResponseBodyData) SetTimestamp(v int64) *GetStorageSizeSummaryResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *GetStorageSizeSummaryResponseBodyData) SetUnit(v map[string]*string) *GetStorageSizeSummaryResponseBodyData {
	s.Unit = v
	return s
}

func (s *GetStorageSizeSummaryResponseBodyData) SetValue(v map[string]*float64) *GetStorageSizeSummaryResponseBodyData {
	s.Value = v
	return s
}

func (s *GetStorageSizeSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
