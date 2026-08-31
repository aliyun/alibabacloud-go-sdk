// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAssetsOnShelveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitAssetsOnShelveResponseBody
	GetCode() *string
	SetData(v *SubmitAssetsOnShelveResponseBodyData) *SubmitAssetsOnShelveResponseBody
	GetData() *SubmitAssetsOnShelveResponseBodyData
	SetHttpStatusCode(v int32) *SubmitAssetsOnShelveResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitAssetsOnShelveResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitAssetsOnShelveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitAssetsOnShelveResponseBody
	GetSuccess() *bool
}

type SubmitAssetsOnShelveResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the asset listing submission.
	Data *SubmitAssetsOnShelveResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SubmitAssetsOnShelveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAssetsOnShelveResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAssetsOnShelveResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitAssetsOnShelveResponseBody) GetData() *SubmitAssetsOnShelveResponseBodyData {
	return s.Data
}

func (s *SubmitAssetsOnShelveResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitAssetsOnShelveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitAssetsOnShelveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAssetsOnShelveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitAssetsOnShelveResponseBody) SetCode(v string) *SubmitAssetsOnShelveResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitAssetsOnShelveResponseBody) SetData(v *SubmitAssetsOnShelveResponseBodyData) *SubmitAssetsOnShelveResponseBody {
	s.Data = v
	return s
}

func (s *SubmitAssetsOnShelveResponseBody) SetHttpStatusCode(v int32) *SubmitAssetsOnShelveResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitAssetsOnShelveResponseBody) SetMessage(v string) *SubmitAssetsOnShelveResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitAssetsOnShelveResponseBody) SetRequestId(v string) *SubmitAssetsOnShelveResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAssetsOnShelveResponseBody) SetSuccess(v bool) *SubmitAssetsOnShelveResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitAssetsOnShelveResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitAssetsOnShelveResponseBodyData struct {
	// The number of assets that failed to be listed.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The listing results for each asset.
	ResultList []*SubmitAssetsOnShelveResponseBodyDataResultList `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
	// The number of assets that were successfully listed.
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

func (s SubmitAssetsOnShelveResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitAssetsOnShelveResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitAssetsOnShelveResponseBodyData) GetFailCount() *int32 {
	return s.FailCount
}

func (s *SubmitAssetsOnShelveResponseBodyData) GetResultList() []*SubmitAssetsOnShelveResponseBodyDataResultList {
	return s.ResultList
}

func (s *SubmitAssetsOnShelveResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *SubmitAssetsOnShelveResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SubmitAssetsOnShelveResponseBodyData) SetFailCount(v int32) *SubmitAssetsOnShelveResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *SubmitAssetsOnShelveResponseBodyData) SetResultList(v []*SubmitAssetsOnShelveResponseBodyDataResultList) *SubmitAssetsOnShelveResponseBodyData {
	s.ResultList = v
	return s
}

func (s *SubmitAssetsOnShelveResponseBodyData) SetSuccessCount(v int32) *SubmitAssetsOnShelveResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *SubmitAssetsOnShelveResponseBodyData) SetTotalCount(v int32) *SubmitAssetsOnShelveResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *SubmitAssetsOnShelveResponseBodyData) Validate() error {
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

type SubmitAssetsOnShelveResponseBodyDataResultList struct {
	// The error code returned when the listing fails. This parameter is empty if the listing succeeds.
	//
	// example:
	//
	// OnShelveFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the listing fails. This parameter is empty if the listing succeeds.
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
	// Indicates whether the asset was successfully listed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitAssetsOnShelveResponseBodyDataResultList) String() string {
	return dara.Prettify(s)
}

func (s SubmitAssetsOnShelveResponseBodyDataResultList) GoString() string {
	return s.String()
}

func (s *SubmitAssetsOnShelveResponseBodyDataResultList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitAssetsOnShelveResponseBodyDataResultList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SubmitAssetsOnShelveResponseBodyDataResultList) GetGuid() *string {
	return s.Guid
}

func (s *SubmitAssetsOnShelveResponseBodyDataResultList) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitAssetsOnShelveResponseBodyDataResultList) SetErrorCode(v string) *SubmitAssetsOnShelveResponseBodyDataResultList {
	s.ErrorCode = &v
	return s
}

func (s *SubmitAssetsOnShelveResponseBodyDataResultList) SetErrorMessage(v string) *SubmitAssetsOnShelveResponseBodyDataResultList {
	s.ErrorMessage = &v
	return s
}

func (s *SubmitAssetsOnShelveResponseBodyDataResultList) SetGuid(v string) *SubmitAssetsOnShelveResponseBodyDataResultList {
	s.Guid = &v
	return s
}

func (s *SubmitAssetsOnShelveResponseBodyDataResultList) SetSuccess(v bool) *SubmitAssetsOnShelveResponseBodyDataResultList {
	s.Success = &v
	return s
}

func (s *SubmitAssetsOnShelveResponseBodyDataResultList) Validate() error {
	return dara.Validate(s)
}
