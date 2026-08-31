// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAssetAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAssetAttributesResponseBody
	GetCode() *string
	SetData(v *UpdateAssetAttributesResponseBodyData) *UpdateAssetAttributesResponseBody
	GetData() *UpdateAssetAttributesResponseBodyData
	SetHttpStatusCode(v int32) *UpdateAssetAttributesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateAssetAttributesResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAssetAttributesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAssetAttributesResponseBody
	GetSuccess() *bool
}

type UpdateAssetAttributesResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of updating asset property values.
	Data *UpdateAssetAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s UpdateAssetAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssetAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAssetAttributesResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAssetAttributesResponseBody) GetData() *UpdateAssetAttributesResponseBodyData {
	return s.Data
}

func (s *UpdateAssetAttributesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateAssetAttributesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAssetAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAssetAttributesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAssetAttributesResponseBody) SetCode(v string) *UpdateAssetAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAssetAttributesResponseBody) SetData(v *UpdateAssetAttributesResponseBodyData) *UpdateAssetAttributesResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAssetAttributesResponseBody) SetHttpStatusCode(v int32) *UpdateAssetAttributesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAssetAttributesResponseBody) SetMessage(v string) *UpdateAssetAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAssetAttributesResponseBody) SetRequestId(v string) *UpdateAssetAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAssetAttributesResponseBody) SetSuccess(v bool) *UpdateAssetAttributesResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAssetAttributesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAssetAttributesResponseBodyData struct {
	// The number of assets that failed to be updated.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The update results for each asset.
	ResultList []*UpdateAssetAttributesResponseBodyDataResultList `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
	// The number of assets that were updated successfully.
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

func (s UpdateAssetAttributesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssetAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAssetAttributesResponseBodyData) GetFailCount() *int32 {
	return s.FailCount
}

func (s *UpdateAssetAttributesResponseBodyData) GetResultList() []*UpdateAssetAttributesResponseBodyDataResultList {
	return s.ResultList
}

func (s *UpdateAssetAttributesResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *UpdateAssetAttributesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *UpdateAssetAttributesResponseBodyData) SetFailCount(v int32) *UpdateAssetAttributesResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *UpdateAssetAttributesResponseBodyData) SetResultList(v []*UpdateAssetAttributesResponseBodyDataResultList) *UpdateAssetAttributesResponseBodyData {
	s.ResultList = v
	return s
}

func (s *UpdateAssetAttributesResponseBodyData) SetSuccessCount(v int32) *UpdateAssetAttributesResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *UpdateAssetAttributesResponseBodyData) SetTotalCount(v int32) *UpdateAssetAttributesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *UpdateAssetAttributesResponseBodyData) Validate() error {
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

type UpdateAssetAttributesResponseBodyDataResultList struct {
	// The error code returned when the update fails. This value is empty when the update succeeds.
	//
	// example:
	//
	// AssetNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the update fails. This value is empty when the update succeeds.
	//
	// example:
	//
	// Asset does not exist: odps.project_a.table_not_exist
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The globally unique identifier (GUID) of the asset.
	//
	// example:
	//
	// odps.project_a.table_orders
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// Indicates whether the asset was updated successfully.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAssetAttributesResponseBodyDataResultList) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssetAttributesResponseBodyDataResultList) GoString() string {
	return s.String()
}

func (s *UpdateAssetAttributesResponseBodyDataResultList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateAssetAttributesResponseBodyDataResultList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateAssetAttributesResponseBodyDataResultList) GetGuid() *string {
	return s.Guid
}

func (s *UpdateAssetAttributesResponseBodyDataResultList) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAssetAttributesResponseBodyDataResultList) SetErrorCode(v string) *UpdateAssetAttributesResponseBodyDataResultList {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAssetAttributesResponseBodyDataResultList) SetErrorMessage(v string) *UpdateAssetAttributesResponseBodyDataResultList {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateAssetAttributesResponseBodyDataResultList) SetGuid(v string) *UpdateAssetAttributesResponseBodyDataResultList {
	s.Guid = &v
	return s
}

func (s *UpdateAssetAttributesResponseBodyDataResultList) SetSuccess(v bool) *UpdateAssetAttributesResponseBodyDataResultList {
	s.Success = &v
	return s
}

func (s *UpdateAssetAttributesResponseBodyDataResultList) Validate() error {
	return dara.Validate(s)
}
