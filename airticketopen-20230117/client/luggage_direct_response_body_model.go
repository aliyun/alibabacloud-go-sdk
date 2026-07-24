// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLuggageDirectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *LuggageDirectResponseBody
	GetRequestId() *string
	SetData(v []*LuggageDirectResponseBodyData) *LuggageDirectResponseBody
	GetData() []*LuggageDirectResponseBodyData
	SetErrorCode(v string) *LuggageDirectResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *LuggageDirectResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *LuggageDirectResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *LuggageDirectResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *LuggageDirectResponseBody
	GetSuccess() *bool
}

type LuggageDirectResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data returned for a successful request.
	Data []*LuggageDirectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The business error code.
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// The data returned with the error response.
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// The error message.
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// The HTTP status code. The value is always 200 for successful requests.
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s LuggageDirectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LuggageDirectResponseBody) GoString() string {
	return s.String()
}

func (s *LuggageDirectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LuggageDirectResponseBody) GetData() []*LuggageDirectResponseBodyData {
	return s.Data
}

func (s *LuggageDirectResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *LuggageDirectResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *LuggageDirectResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *LuggageDirectResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *LuggageDirectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *LuggageDirectResponseBody) SetRequestId(v string) *LuggageDirectResponseBody {
	s.RequestId = &v
	return s
}

func (s *LuggageDirectResponseBody) SetData(v []*LuggageDirectResponseBodyData) *LuggageDirectResponseBody {
	s.Data = v
	return s
}

func (s *LuggageDirectResponseBody) SetErrorCode(v string) *LuggageDirectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *LuggageDirectResponseBody) SetErrorData(v interface{}) *LuggageDirectResponseBody {
	s.ErrorData = v
	return s
}

func (s *LuggageDirectResponseBody) SetErrorMsg(v string) *LuggageDirectResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *LuggageDirectResponseBody) SetStatus(v int32) *LuggageDirectResponseBody {
	s.Status = &v
	return s
}

func (s *LuggageDirectResponseBody) SetSuccess(v bool) *LuggageDirectResponseBody {
	s.Success = &v
	return s
}

func (s *LuggageDirectResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type LuggageDirectResponseBodyData struct {
	// The three-letter IATA code of the city.
	//
	// example:
	//
	// BJS
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// The luggage through-check rule type. Valid values:
	//
	// - 0: luggage through-check is not supported.
	//
	// - 1: luggage through-check is supported.
	//
	// example:
	//
	// 1
	DirectType *int32 `json:"direct_type,omitempty" xml:"direct_type,omitempty"`
}

func (s LuggageDirectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s LuggageDirectResponseBodyData) GoString() string {
	return s.String()
}

func (s *LuggageDirectResponseBodyData) GetCityCode() *string {
	return s.CityCode
}

func (s *LuggageDirectResponseBodyData) GetDirectType() *int32 {
	return s.DirectType
}

func (s *LuggageDirectResponseBodyData) SetCityCode(v string) *LuggageDirectResponseBodyData {
	s.CityCode = &v
	return s
}

func (s *LuggageDirectResponseBodyData) SetDirectType(v int32) *LuggageDirectResponseBodyData {
	s.DirectType = &v
	return s
}

func (s *LuggageDirectResponseBodyData) Validate() error {
	return dara.Validate(s)
}
