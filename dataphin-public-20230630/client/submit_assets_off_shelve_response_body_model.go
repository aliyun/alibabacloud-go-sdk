// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAssetsOffShelveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitAssetsOffShelveResponseBody
	GetCode() *string
	SetData(v *SubmitAssetsOffShelveResponseBodyData) *SubmitAssetsOffShelveResponseBody
	GetData() *SubmitAssetsOffShelveResponseBodyData
	SetHttpStatusCode(v int32) *SubmitAssetsOffShelveResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitAssetsOffShelveResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitAssetsOffShelveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitAssetsOffShelveResponseBody
	GetSuccess() *bool
}

type SubmitAssetsOffShelveResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the asset delisting submission.
	Data *SubmitAssetsOffShelveResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
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
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitAssetsOffShelveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAssetsOffShelveResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAssetsOffShelveResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitAssetsOffShelveResponseBody) GetData() *SubmitAssetsOffShelveResponseBodyData {
	return s.Data
}

func (s *SubmitAssetsOffShelveResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitAssetsOffShelveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitAssetsOffShelveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAssetsOffShelveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitAssetsOffShelveResponseBody) SetCode(v string) *SubmitAssetsOffShelveResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitAssetsOffShelveResponseBody) SetData(v *SubmitAssetsOffShelveResponseBodyData) *SubmitAssetsOffShelveResponseBody {
	s.Data = v
	return s
}

func (s *SubmitAssetsOffShelveResponseBody) SetHttpStatusCode(v int32) *SubmitAssetsOffShelveResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitAssetsOffShelveResponseBody) SetMessage(v string) *SubmitAssetsOffShelveResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitAssetsOffShelveResponseBody) SetRequestId(v string) *SubmitAssetsOffShelveResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAssetsOffShelveResponseBody) SetSuccess(v bool) *SubmitAssetsOffShelveResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitAssetsOffShelveResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitAssetsOffShelveResponseBodyData struct {
	// The number of assets for which the delisting submission failed.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The delisting submission results for each asset.
	ResultList []*SubmitAssetsOffShelveResponseBodyDataResultList `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
	// The number of assets for which the delisting submission succeeded.
	//
	// example:
	//
	// 2
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The total number of assets.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SubmitAssetsOffShelveResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitAssetsOffShelveResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitAssetsOffShelveResponseBodyData) GetFailCount() *int32 {
	return s.FailCount
}

func (s *SubmitAssetsOffShelveResponseBodyData) GetResultList() []*SubmitAssetsOffShelveResponseBodyDataResultList {
	return s.ResultList
}

func (s *SubmitAssetsOffShelveResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *SubmitAssetsOffShelveResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SubmitAssetsOffShelveResponseBodyData) SetFailCount(v int32) *SubmitAssetsOffShelveResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *SubmitAssetsOffShelveResponseBodyData) SetResultList(v []*SubmitAssetsOffShelveResponseBodyDataResultList) *SubmitAssetsOffShelveResponseBodyData {
	s.ResultList = v
	return s
}

func (s *SubmitAssetsOffShelveResponseBodyData) SetSuccessCount(v int32) *SubmitAssetsOffShelveResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *SubmitAssetsOffShelveResponseBodyData) SetTotalCount(v int32) *SubmitAssetsOffShelveResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *SubmitAssetsOffShelveResponseBodyData) Validate() error {
	if s.ResultList != nil {
		for _, item := range s.ResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitAssetsOffShelveResponseBodyDataResultList struct {
	// The error code returned when the submission fails. This value is empty when the submission succeeds.
	//
	// example:
	//
	// OffShelveFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the submission fails. This value is empty when the submission succeeds.
	//
	// example:
	//
	// Asset does not exist: odps.project_a.table_not_exist
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The unique identifier of the asset.
	//
	// example:
	//
	// odps.project_a.table_orders
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// Indicates whether the direct delisting or delisting approval was submitted successfully.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitAssetsOffShelveResponseBodyDataResultList) String() string {
	return dara.Prettify(s)
}

func (s SubmitAssetsOffShelveResponseBodyDataResultList) GoString() string {
	return s.String()
}

func (s *SubmitAssetsOffShelveResponseBodyDataResultList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitAssetsOffShelveResponseBodyDataResultList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SubmitAssetsOffShelveResponseBodyDataResultList) GetGuid() *string {
	return s.Guid
}

func (s *SubmitAssetsOffShelveResponseBodyDataResultList) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitAssetsOffShelveResponseBodyDataResultList) SetErrorCode(v string) *SubmitAssetsOffShelveResponseBodyDataResultList {
	s.ErrorCode = &v
	return s
}

func (s *SubmitAssetsOffShelveResponseBodyDataResultList) SetErrorMessage(v string) *SubmitAssetsOffShelveResponseBodyDataResultList {
	s.ErrorMessage = &v
	return s
}

func (s *SubmitAssetsOffShelveResponseBodyDataResultList) SetGuid(v string) *SubmitAssetsOffShelveResponseBodyDataResultList {
	s.Guid = &v
	return s
}

func (s *SubmitAssetsOffShelveResponseBodyDataResultList) SetSuccess(v bool) *SubmitAssetsOffShelveResponseBodyDataResultList {
	s.Success = &v
	return s
}

func (s *SubmitAssetsOffShelveResponseBodyDataResultList) Validate() error {
	return dara.Validate(s)
}
