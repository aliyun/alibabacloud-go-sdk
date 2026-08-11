// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloneVoiceModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCloneVoiceModelsResponseBody
	GetCode() *string
	SetData(v *ListCloneVoiceModelsResponseBodyData) *ListCloneVoiceModelsResponseBody
	GetData() *ListCloneVoiceModelsResponseBodyData
	SetHttpStatusCode(v int32) *ListCloneVoiceModelsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCloneVoiceModelsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListCloneVoiceModelsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListCloneVoiceModelsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCloneVoiceModelsResponseBody
	GetSuccess() *bool
}

type ListCloneVoiceModelsResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *ListCloneVoiceModelsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance does not exist. Instance=392db13c-8901-4a25-b566-91d0d8114cec
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 019FDAC7-13C5-1B64-A853-999DF105B9EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCloneVoiceModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloneVoiceModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloneVoiceModelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCloneVoiceModelsResponseBody) GetData() *ListCloneVoiceModelsResponseBodyData {
	return s.Data
}

func (s *ListCloneVoiceModelsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCloneVoiceModelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCloneVoiceModelsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListCloneVoiceModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloneVoiceModelsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCloneVoiceModelsResponseBody) SetCode(v string) *ListCloneVoiceModelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCloneVoiceModelsResponseBody) SetData(v *ListCloneVoiceModelsResponseBodyData) *ListCloneVoiceModelsResponseBody {
	s.Data = v
	return s
}

func (s *ListCloneVoiceModelsResponseBody) SetHttpStatusCode(v int32) *ListCloneVoiceModelsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCloneVoiceModelsResponseBody) SetMessage(v string) *ListCloneVoiceModelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCloneVoiceModelsResponseBody) SetParams(v []*string) *ListCloneVoiceModelsResponseBody {
	s.Params = v
	return s
}

func (s *ListCloneVoiceModelsResponseBody) SetRequestId(v string) *ListCloneVoiceModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloneVoiceModelsResponseBody) SetSuccess(v bool) *ListCloneVoiceModelsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCloneVoiceModelsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloneVoiceModelsResponseBodyData struct {
	// The list of voice cloning models.
	CloneVoiceModels []*ListCloneVoiceModelsResponseBodyDataCloneVoiceModels `json:"CloneVoiceModels,omitempty" xml:"CloneVoiceModels,omitempty" type:"Repeated"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records that match the conditions.
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloneVoiceModelsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCloneVoiceModelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCloneVoiceModelsResponseBodyData) GetCloneVoiceModels() []*ListCloneVoiceModelsResponseBodyDataCloneVoiceModels {
	return s.CloneVoiceModels
}

func (s *ListCloneVoiceModelsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloneVoiceModelsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloneVoiceModelsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCloneVoiceModelsResponseBodyData) SetCloneVoiceModels(v []*ListCloneVoiceModelsResponseBodyDataCloneVoiceModels) *ListCloneVoiceModelsResponseBodyData {
	s.CloneVoiceModels = v
	return s
}

func (s *ListCloneVoiceModelsResponseBodyData) SetPageNumber(v int32) *ListCloneVoiceModelsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCloneVoiceModelsResponseBodyData) SetPageSize(v int32) *ListCloneVoiceModelsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCloneVoiceModelsResponseBodyData) SetTotalCount(v int32) *ListCloneVoiceModelsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCloneVoiceModelsResponseBodyData) Validate() error {
	if s.CloneVoiceModels != nil {
		for _, item := range s.CloneVoiceModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloneVoiceModelsResponseBodyDataCloneVoiceModels struct {
	// The description.
	//
	// example:
	//
	// Responsive, adaptive tone adjustment, effective complex text processing
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name.
	//
	// example:
	//
	// Qwen
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The model name.
	//
	// example:
	//
	// QwenVc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListCloneVoiceModelsResponseBodyDataCloneVoiceModels) String() string {
	return dara.Prettify(s)
}

func (s ListCloneVoiceModelsResponseBodyDataCloneVoiceModels) GoString() string {
	return s.String()
}

func (s *ListCloneVoiceModelsResponseBodyDataCloneVoiceModels) GetDescription() *string {
	return s.Description
}

func (s *ListCloneVoiceModelsResponseBodyDataCloneVoiceModels) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListCloneVoiceModelsResponseBodyDataCloneVoiceModels) GetName() *string {
	return s.Name
}

func (s *ListCloneVoiceModelsResponseBodyDataCloneVoiceModels) SetDescription(v string) *ListCloneVoiceModelsResponseBodyDataCloneVoiceModels {
	s.Description = &v
	return s
}

func (s *ListCloneVoiceModelsResponseBodyDataCloneVoiceModels) SetDisplayName(v string) *ListCloneVoiceModelsResponseBodyDataCloneVoiceModels {
	s.DisplayName = &v
	return s
}

func (s *ListCloneVoiceModelsResponseBodyDataCloneVoiceModels) SetName(v string) *ListCloneVoiceModelsResponseBodyDataCloneVoiceModels {
	s.Name = &v
	return s
}

func (s *ListCloneVoiceModelsResponseBodyDataCloneVoiceModels) Validate() error {
	return dara.Validate(s)
}
