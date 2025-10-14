// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransitVisaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TransitVisaResponseBody
	GetRequestId() *string
	SetData(v []*TransitVisaResponseBodyData) *TransitVisaResponseBody
	GetData() []*TransitVisaResponseBodyData
	SetErrorCode(v string) *TransitVisaResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *TransitVisaResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *TransitVisaResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *TransitVisaResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *TransitVisaResponseBody
	GetSuccess() *bool
}

type TransitVisaResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*TransitVisaResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s TransitVisaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransitVisaResponseBody) GoString() string {
	return s.String()
}

func (s *TransitVisaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransitVisaResponseBody) GetData() []*TransitVisaResponseBodyData {
	return s.Data
}

func (s *TransitVisaResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TransitVisaResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *TransitVisaResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TransitVisaResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *TransitVisaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TransitVisaResponseBody) SetRequestId(v string) *TransitVisaResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransitVisaResponseBody) SetData(v []*TransitVisaResponseBodyData) *TransitVisaResponseBody {
	s.Data = v
	return s
}

func (s *TransitVisaResponseBody) SetErrorCode(v string) *TransitVisaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TransitVisaResponseBody) SetErrorData(v interface{}) *TransitVisaResponseBody {
	s.ErrorData = v
	return s
}

func (s *TransitVisaResponseBody) SetErrorMsg(v string) *TransitVisaResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TransitVisaResponseBody) SetStatus(v int32) *TransitVisaResponseBody {
	s.Status = &v
	return s
}

func (s *TransitVisaResponseBody) SetSuccess(v bool) *TransitVisaResponseBody {
	s.Success = &v
	return s
}

func (s *TransitVisaResponseBody) Validate() error {
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

type TransitVisaResponseBodyData struct {
	// example:
	//
	// HGH
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// example:
	//
	// 1
	VisaType *int32 `json:"visa_type,omitempty" xml:"visa_type,omitempty"`
}

func (s TransitVisaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TransitVisaResponseBodyData) GoString() string {
	return s.String()
}

func (s *TransitVisaResponseBodyData) GetCityCode() *string {
	return s.CityCode
}

func (s *TransitVisaResponseBodyData) GetVisaType() *int32 {
	return s.VisaType
}

func (s *TransitVisaResponseBodyData) SetCityCode(v string) *TransitVisaResponseBodyData {
	s.CityCode = &v
	return s
}

func (s *TransitVisaResponseBodyData) SetVisaType(v int32) *TransitVisaResponseBodyData {
	s.VisaType = &v
	return s
}

func (s *TransitVisaResponseBodyData) Validate() error {
	return dara.Validate(s)
}
