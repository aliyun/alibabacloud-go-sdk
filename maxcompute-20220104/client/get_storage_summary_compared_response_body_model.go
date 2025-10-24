// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageSummaryComparedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetStorageSummaryComparedResponseBodyData) *GetStorageSummaryComparedResponseBody
	GetData() *GetStorageSummaryComparedResponseBodyData
	SetErrorCode(v string) *GetStorageSummaryComparedResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetStorageSummaryComparedResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *GetStorageSummaryComparedResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetStorageSummaryComparedResponseBody
	GetRequestId() *string
}

type GetStorageSummaryComparedResponseBody struct {
	Data *GetStorageSummaryComparedResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// plan \\"***\\" does not exist
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 0bc3b4b016674434996033675e71ee
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetStorageSummaryComparedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStorageSummaryComparedResponseBody) GoString() string {
	return s.String()
}

func (s *GetStorageSummaryComparedResponseBody) GetData() *GetStorageSummaryComparedResponseBodyData {
	return s.Data
}

func (s *GetStorageSummaryComparedResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetStorageSummaryComparedResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetStorageSummaryComparedResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetStorageSummaryComparedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStorageSummaryComparedResponseBody) SetData(v *GetStorageSummaryComparedResponseBodyData) *GetStorageSummaryComparedResponseBody {
	s.Data = v
	return s
}

func (s *GetStorageSummaryComparedResponseBody) SetErrorCode(v string) *GetStorageSummaryComparedResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetStorageSummaryComparedResponseBody) SetErrorMsg(v string) *GetStorageSummaryComparedResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetStorageSummaryComparedResponseBody) SetHttpCode(v int32) *GetStorageSummaryComparedResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetStorageSummaryComparedResponseBody) SetRequestId(v string) *GetStorageSummaryComparedResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStorageSummaryComparedResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStorageSummaryComparedResponseBodyData struct {
	// example:
	//
	// 20250601
	BeginDate *string `json:"beginDate,omitempty" xml:"beginDate,omitempty"`
	// example:
	//
	// 20250604
	EndDate *string             `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Rate    map[string]*float64 `json:"rate,omitempty" xml:"rate,omitempty"`
	Unit    map[string]*string  `json:"unit,omitempty" xml:"unit,omitempty"`
	Value   map[string]*float64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetStorageSummaryComparedResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetStorageSummaryComparedResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStorageSummaryComparedResponseBodyData) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetStorageSummaryComparedResponseBodyData) GetEndDate() *string {
	return s.EndDate
}

func (s *GetStorageSummaryComparedResponseBodyData) GetRate() map[string]*float64 {
	return s.Rate
}

func (s *GetStorageSummaryComparedResponseBodyData) GetUnit() map[string]*string {
	return s.Unit
}

func (s *GetStorageSummaryComparedResponseBodyData) GetValue() map[string]*float64 {
	return s.Value
}

func (s *GetStorageSummaryComparedResponseBodyData) SetBeginDate(v string) *GetStorageSummaryComparedResponseBodyData {
	s.BeginDate = &v
	return s
}

func (s *GetStorageSummaryComparedResponseBodyData) SetEndDate(v string) *GetStorageSummaryComparedResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *GetStorageSummaryComparedResponseBodyData) SetRate(v map[string]*float64) *GetStorageSummaryComparedResponseBodyData {
	s.Rate = v
	return s
}

func (s *GetStorageSummaryComparedResponseBodyData) SetUnit(v map[string]*string) *GetStorageSummaryComparedResponseBodyData {
	s.Unit = v
	return s
}

func (s *GetStorageSummaryComparedResponseBodyData) SetValue(v map[string]*float64) *GetStorageSummaryComparedResponseBodyData {
	s.Value = v
	return s
}

func (s *GetStorageSummaryComparedResponseBodyData) Validate() error {
	return dara.Validate(s)
}
