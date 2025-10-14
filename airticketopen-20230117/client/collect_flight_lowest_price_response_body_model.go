// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCollectFlightLowestPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CollectFlightLowestPriceResponseBody
	GetRequestId() *string
	SetData(v map[string]interface{}) *CollectFlightLowestPriceResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v string) *CollectFlightLowestPriceResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *CollectFlightLowestPriceResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *CollectFlightLowestPriceResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *CollectFlightLowestPriceResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *CollectFlightLowestPriceResponseBody
	GetSuccess() *bool
}

type CollectFlightLowestPriceResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// null
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CollectFlightLowestPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CollectFlightLowestPriceResponseBody) GoString() string {
	return s.String()
}

func (s *CollectFlightLowestPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CollectFlightLowestPriceResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *CollectFlightLowestPriceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CollectFlightLowestPriceResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *CollectFlightLowestPriceResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CollectFlightLowestPriceResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *CollectFlightLowestPriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CollectFlightLowestPriceResponseBody) SetRequestId(v string) *CollectFlightLowestPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CollectFlightLowestPriceResponseBody) SetData(v map[string]interface{}) *CollectFlightLowestPriceResponseBody {
	s.Data = v
	return s
}

func (s *CollectFlightLowestPriceResponseBody) SetErrorCode(v string) *CollectFlightLowestPriceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CollectFlightLowestPriceResponseBody) SetErrorData(v interface{}) *CollectFlightLowestPriceResponseBody {
	s.ErrorData = v
	return s
}

func (s *CollectFlightLowestPriceResponseBody) SetErrorMsg(v string) *CollectFlightLowestPriceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CollectFlightLowestPriceResponseBody) SetStatus(v int32) *CollectFlightLowestPriceResponseBody {
	s.Status = &v
	return s
}

func (s *CollectFlightLowestPriceResponseBody) SetSuccess(v bool) *CollectFlightLowestPriceResponseBody {
	s.Success = &v
	return s
}

func (s *CollectFlightLowestPriceResponseBody) Validate() error {
	return dara.Validate(s)
}
