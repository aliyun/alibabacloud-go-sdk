// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiCallTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataServiceApiCallTrendResponseBody
	GetCode() *string
	SetData(v *GetDataServiceApiCallTrendResponseBodyData) *GetDataServiceApiCallTrendResponseBody
	GetData() *GetDataServiceApiCallTrendResponseBodyData
	SetHttpStatusCode(v int32) *GetDataServiceApiCallTrendResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataServiceApiCallTrendResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDataServiceApiCallTrendResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataServiceApiCallTrendResponseBody
	GetSuccess() *bool
}

type GetDataServiceApiCallTrendResponseBody struct {
	// example:
	//
	// OK
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDataServiceApiCallTrendResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetDataServiceApiCallTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiCallTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiCallTrendResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataServiceApiCallTrendResponseBody) GetData() *GetDataServiceApiCallTrendResponseBodyData {
	return s.Data
}

func (s *GetDataServiceApiCallTrendResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataServiceApiCallTrendResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataServiceApiCallTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceApiCallTrendResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataServiceApiCallTrendResponseBody) SetCode(v string) *GetDataServiceApiCallTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataServiceApiCallTrendResponseBody) SetData(v *GetDataServiceApiCallTrendResponseBodyData) *GetDataServiceApiCallTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetDataServiceApiCallTrendResponseBody) SetHttpStatusCode(v int32) *GetDataServiceApiCallTrendResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataServiceApiCallTrendResponseBody) SetMessage(v string) *GetDataServiceApiCallTrendResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataServiceApiCallTrendResponseBody) SetRequestId(v string) *GetDataServiceApiCallTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceApiCallTrendResponseBody) SetSuccess(v bool) *GetDataServiceApiCallTrendResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataServiceApiCallTrendResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataServiceApiCallTrendResponseBodyData struct {
	CallErrorImpactTrendList []*GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList `json:"CallErrorImpactTrendList,omitempty" xml:"CallErrorImpactTrendList,omitempty" type:"Repeated"`
	CallErrorTrendList       []*GetDataServiceApiCallTrendResponseBodyDataCallErrorTrendList       `json:"CallErrorTrendList,omitempty" xml:"CallErrorTrendList,omitempty" type:"Repeated"`
}

func (s GetDataServiceApiCallTrendResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiCallTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiCallTrendResponseBodyData) GetCallErrorImpactTrendList() []*GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList {
	return s.CallErrorImpactTrendList
}

func (s *GetDataServiceApiCallTrendResponseBodyData) GetCallErrorTrendList() []*GetDataServiceApiCallTrendResponseBodyDataCallErrorTrendList {
	return s.CallErrorTrendList
}

func (s *GetDataServiceApiCallTrendResponseBodyData) SetCallErrorImpactTrendList(v []*GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList) *GetDataServiceApiCallTrendResponseBodyData {
	s.CallErrorImpactTrendList = v
	return s
}

func (s *GetDataServiceApiCallTrendResponseBodyData) SetCallErrorTrendList(v []*GetDataServiceApiCallTrendResponseBodyDataCallErrorTrendList) *GetDataServiceApiCallTrendResponseBodyData {
	s.CallErrorTrendList = v
	return s
}

func (s *GetDataServiceApiCallTrendResponseBodyData) Validate() error {
	if s.CallErrorImpactTrendList != nil {
		for _, item := range s.CallErrorImpactTrendList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CallErrorTrendList != nil {
		for _, item := range s.CallErrorTrendList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList struct {
	ApiIdList []*int64 `json:"ApiIdList,omitempty" xml:"ApiIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	ErrorApiCount *int32 `json:"ErrorApiCount,omitempty" xml:"ErrorApiCount,omitempty"`
	// example:
	//
	// 1
	ErrorAppCount *int32 `json:"ErrorAppCount,omitempty" xml:"ErrorAppCount,omitempty"`
	// example:
	//
	// 2025-06-30 08:00
	Minute *string `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList) GetApiIdList() []*int64 {
	return s.ApiIdList
}

func (s *GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList) GetErrorApiCount() *int32 {
	return s.ErrorApiCount
}

func (s *GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList) GetErrorAppCount() *int32 {
	return s.ErrorAppCount
}

func (s *GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList) GetMinute() *string {
	return s.Minute
}

func (s *GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList) SetApiIdList(v []*int64) *GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList {
	s.ApiIdList = v
	return s
}

func (s *GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList) SetErrorApiCount(v int32) *GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList {
	s.ErrorApiCount = &v
	return s
}

func (s *GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList) SetErrorAppCount(v int32) *GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList {
	s.ErrorAppCount = &v
	return s
}

func (s *GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList) SetMinute(v string) *GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList {
	s.Minute = &v
	return s
}

func (s *GetDataServiceApiCallTrendResponseBodyDataCallErrorImpactTrendList) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceApiCallTrendResponseBodyDataCallErrorTrendList struct {
	// example:
	//
	// 1021
	CallCount *int64 `json:"CallCount,omitempty" xml:"CallCount,omitempty"`
	// example:
	//
	// 102
	ErrorCount *int64 `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	// example:
	//
	// 3
	Minute *string `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s GetDataServiceApiCallTrendResponseBodyDataCallErrorTrendList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiCallTrendResponseBodyDataCallErrorTrendList) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiCallTrendResponseBodyDataCallErrorTrendList) GetCallCount() *int64 {
	return s.CallCount
}

func (s *GetDataServiceApiCallTrendResponseBodyDataCallErrorTrendList) GetErrorCount() *int64 {
	return s.ErrorCount
}

func (s *GetDataServiceApiCallTrendResponseBodyDataCallErrorTrendList) GetMinute() *string {
	return s.Minute
}

func (s *GetDataServiceApiCallTrendResponseBodyDataCallErrorTrendList) SetCallCount(v int64) *GetDataServiceApiCallTrendResponseBodyDataCallErrorTrendList {
	s.CallCount = &v
	return s
}

func (s *GetDataServiceApiCallTrendResponseBodyDataCallErrorTrendList) SetErrorCount(v int64) *GetDataServiceApiCallTrendResponseBodyDataCallErrorTrendList {
	s.ErrorCount = &v
	return s
}

func (s *GetDataServiceApiCallTrendResponseBodyDataCallErrorTrendList) SetMinute(v string) *GetDataServiceApiCallTrendResponseBodyDataCallErrorTrendList {
	s.Minute = &v
	return s
}

func (s *GetDataServiceApiCallTrendResponseBodyDataCallErrorTrendList) Validate() error {
	return dara.Validate(s)
}
