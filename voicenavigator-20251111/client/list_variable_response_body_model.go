// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVariableResponseBody
	GetCode() *string
	SetData(v *ListVariableResponseBodyData) *ListVariableResponseBody
	GetData() *ListVariableResponseBodyData
	SetHttpStatusCode(v int32) *ListVariableResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListVariableResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListVariableResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListVariableResponseBody
	GetRequestId() *string
}

type ListVariableResponseBody struct {
	// example:
	//
	// OK
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListVariableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// d74d6290-7cbe-4436-b5d7-014ebb0f4060
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVariableResponseBody) GoString() string {
	return s.String()
}

func (s *ListVariableResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVariableResponseBody) GetData() *ListVariableResponseBodyData {
	return s.Data
}

func (s *ListVariableResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListVariableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVariableResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVariableResponseBody) SetCode(v string) *ListVariableResponseBody {
	s.Code = &v
	return s
}

func (s *ListVariableResponseBody) SetData(v *ListVariableResponseBodyData) *ListVariableResponseBody {
	s.Data = v
	return s
}

func (s *ListVariableResponseBody) SetHttpStatusCode(v int32) *ListVariableResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListVariableResponseBody) SetMessage(v string) *ListVariableResponseBody {
	s.Message = &v
	return s
}

func (s *ListVariableResponseBody) SetParams(v []*string) *ListVariableResponseBody {
	s.Params = v
	return s
}

func (s *ListVariableResponseBody) SetRequestId(v string) *ListVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVariableResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVariableResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Variables  []*ListVariableResponseBodyDataVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListVariableResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVariableResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVariableResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVariableResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVariableResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVariableResponseBodyData) GetVariables() []*ListVariableResponseBodyDataVariables {
	return s.Variables
}

func (s *ListVariableResponseBodyData) SetPageNumber(v int32) *ListVariableResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListVariableResponseBodyData) SetPageSize(v int32) *ListVariableResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListVariableResponseBodyData) SetTotalCount(v int32) *ListVariableResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListVariableResponseBodyData) SetVariables(v []*ListVariableResponseBodyDataVariables) *ListVariableResponseBodyData {
	s.Variables = v
	return s
}

func (s *ListVariableResponseBodyData) Validate() error {
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVariableResponseBodyDataVariables struct {
	// example:
	//
	// 1754013825102
	CreatedTime *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66c
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// age
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1308144684576655
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 1754013825102
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	VariableId *string `json:"VariableId,omitempty" xml:"VariableId,omitempty"`
}

func (s ListVariableResponseBodyDataVariables) String() string {
	return dara.Prettify(s)
}

func (s ListVariableResponseBodyDataVariables) GoString() string {
	return s.String()
}

func (s *ListVariableResponseBodyDataVariables) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListVariableResponseBodyDataVariables) GetDescription() *string {
	return s.Description
}

func (s *ListVariableResponseBodyDataVariables) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListVariableResponseBodyDataVariables) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVariableResponseBodyDataVariables) GetName() *string {
	return s.Name
}

func (s *ListVariableResponseBodyDataVariables) GetTenantId() *string {
	return s.TenantId
}

func (s *ListVariableResponseBodyDataVariables) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListVariableResponseBodyDataVariables) GetVariableId() *string {
	return s.VariableId
}

func (s *ListVariableResponseBodyDataVariables) SetCreatedTime(v int64) *ListVariableResponseBodyDataVariables {
	s.CreatedTime = &v
	return s
}

func (s *ListVariableResponseBodyDataVariables) SetDescription(v string) *ListVariableResponseBodyDataVariables {
	s.Description = &v
	return s
}

func (s *ListVariableResponseBodyDataVariables) SetDisplayName(v string) *ListVariableResponseBodyDataVariables {
	s.DisplayName = &v
	return s
}

func (s *ListVariableResponseBodyDataVariables) SetInstanceId(v string) *ListVariableResponseBodyDataVariables {
	s.InstanceId = &v
	return s
}

func (s *ListVariableResponseBodyDataVariables) SetName(v string) *ListVariableResponseBodyDataVariables {
	s.Name = &v
	return s
}

func (s *ListVariableResponseBodyDataVariables) SetTenantId(v string) *ListVariableResponseBodyDataVariables {
	s.TenantId = &v
	return s
}

func (s *ListVariableResponseBodyDataVariables) SetUpdatedTime(v int64) *ListVariableResponseBodyDataVariables {
	s.UpdatedTime = &v
	return s
}

func (s *ListVariableResponseBodyDataVariables) SetVariableId(v string) *ListVariableResponseBodyDataVariables {
	s.VariableId = &v
	return s
}

func (s *ListVariableResponseBodyDataVariables) Validate() error {
	return dara.Validate(s)
}
