// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiCallSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataServiceApiCallSummaryResponseBody
	GetCode() *string
	SetData(v *GetDataServiceApiCallSummaryResponseBodyData) *GetDataServiceApiCallSummaryResponseBody
	GetData() *GetDataServiceApiCallSummaryResponseBodyData
	SetHttpStatusCode(v int32) *GetDataServiceApiCallSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataServiceApiCallSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDataServiceApiCallSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataServiceApiCallSummaryResponseBody
	GetSuccess() *bool
}

type GetDataServiceApiCallSummaryResponseBody struct {
	// example:
	//
	// OK
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDataServiceApiCallSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataServiceApiCallSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiCallSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiCallSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataServiceApiCallSummaryResponseBody) GetData() *GetDataServiceApiCallSummaryResponseBodyData {
	return s.Data
}

func (s *GetDataServiceApiCallSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataServiceApiCallSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataServiceApiCallSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceApiCallSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataServiceApiCallSummaryResponseBody) SetCode(v string) *GetDataServiceApiCallSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataServiceApiCallSummaryResponseBody) SetData(v *GetDataServiceApiCallSummaryResponseBodyData) *GetDataServiceApiCallSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetDataServiceApiCallSummaryResponseBody) SetHttpStatusCode(v int32) *GetDataServiceApiCallSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataServiceApiCallSummaryResponseBody) SetMessage(v string) *GetDataServiceApiCallSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataServiceApiCallSummaryResponseBody) SetRequestId(v string) *GetDataServiceApiCallSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceApiCallSummaryResponseBody) SetSuccess(v bool) *GetDataServiceApiCallSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataServiceApiCallSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataServiceApiCallSummaryResponseBodyData struct {
	// example:
	//
	// 1021
	CallCount *int64 `json:"CallCount,omitempty" xml:"CallCount,omitempty"`
	// example:
	//
	// 8
	ErrorApiCount *int64 `json:"ErrorApiCount,omitempty" xml:"ErrorApiCount,omitempty"`
	// example:
	//
	// 2
	ErrorAppCount *int64 `json:"ErrorAppCount,omitempty" xml:"ErrorAppCount,omitempty"`
	// example:
	//
	// 102
	ErrorCount *int64 `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	// example:
	//
	// 10.01
	ErrorRate *float64 `json:"ErrorRate,omitempty" xml:"ErrorRate,omitempty"`
	// example:
	//
	// 2.03
	OfflineRate *float64 `json:"OfflineRate,omitempty" xml:"OfflineRate,omitempty"`
}

func (s GetDataServiceApiCallSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiCallSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiCallSummaryResponseBodyData) GetCallCount() *int64 {
	return s.CallCount
}

func (s *GetDataServiceApiCallSummaryResponseBodyData) GetErrorApiCount() *int64 {
	return s.ErrorApiCount
}

func (s *GetDataServiceApiCallSummaryResponseBodyData) GetErrorAppCount() *int64 {
	return s.ErrorAppCount
}

func (s *GetDataServiceApiCallSummaryResponseBodyData) GetErrorCount() *int64 {
	return s.ErrorCount
}

func (s *GetDataServiceApiCallSummaryResponseBodyData) GetErrorRate() *float64 {
	return s.ErrorRate
}

func (s *GetDataServiceApiCallSummaryResponseBodyData) GetOfflineRate() *float64 {
	return s.OfflineRate
}

func (s *GetDataServiceApiCallSummaryResponseBodyData) SetCallCount(v int64) *GetDataServiceApiCallSummaryResponseBodyData {
	s.CallCount = &v
	return s
}

func (s *GetDataServiceApiCallSummaryResponseBodyData) SetErrorApiCount(v int64) *GetDataServiceApiCallSummaryResponseBodyData {
	s.ErrorApiCount = &v
	return s
}

func (s *GetDataServiceApiCallSummaryResponseBodyData) SetErrorAppCount(v int64) *GetDataServiceApiCallSummaryResponseBodyData {
	s.ErrorAppCount = &v
	return s
}

func (s *GetDataServiceApiCallSummaryResponseBodyData) SetErrorCount(v int64) *GetDataServiceApiCallSummaryResponseBodyData {
	s.ErrorCount = &v
	return s
}

func (s *GetDataServiceApiCallSummaryResponseBodyData) SetErrorRate(v float64) *GetDataServiceApiCallSummaryResponseBodyData {
	s.ErrorRate = &v
	return s
}

func (s *GetDataServiceApiCallSummaryResponseBodyData) SetOfflineRate(v float64) *GetDataServiceApiCallSummaryResponseBodyData {
	s.OfflineRate = &v
	return s
}

func (s *GetDataServiceApiCallSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
