// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptPublishHistoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListScriptPublishHistoriesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListScriptPublishHistoriesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListScriptPublishHistoriesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListScriptPublishHistoriesResponseBody
	GetRequestId() *string
	SetScriptPublishHistories(v *ListScriptPublishHistoriesResponseBodyScriptPublishHistories) *ListScriptPublishHistoriesResponseBody
	GetScriptPublishHistories() *ListScriptPublishHistoriesResponseBodyScriptPublishHistories
	SetSuccess(v bool) *ListScriptPublishHistoriesResponseBody
	GetSuccess() *bool
}

type ListScriptPublishHistoriesResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// History list
	ScriptPublishHistories *ListScriptPublishHistoriesResponseBodyScriptPublishHistories `json:"ScriptPublishHistories,omitempty" xml:"ScriptPublishHistories,omitempty" type:"Struct"`
	// is Succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListScriptPublishHistoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScriptPublishHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListScriptPublishHistoriesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListScriptPublishHistoriesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListScriptPublishHistoriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScriptPublishHistoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScriptPublishHistoriesResponseBody) GetScriptPublishHistories() *ListScriptPublishHistoriesResponseBodyScriptPublishHistories {
	return s.ScriptPublishHistories
}

func (s *ListScriptPublishHistoriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScriptPublishHistoriesResponseBody) SetCode(v string) *ListScriptPublishHistoriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListScriptPublishHistoriesResponseBody) SetHttpStatusCode(v int32) *ListScriptPublishHistoriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListScriptPublishHistoriesResponseBody) SetMessage(v string) *ListScriptPublishHistoriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListScriptPublishHistoriesResponseBody) SetRequestId(v string) *ListScriptPublishHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScriptPublishHistoriesResponseBody) SetScriptPublishHistories(v *ListScriptPublishHistoriesResponseBodyScriptPublishHistories) *ListScriptPublishHistoriesResponseBody {
	s.ScriptPublishHistories = v
	return s
}

func (s *ListScriptPublishHistoriesResponseBody) SetSuccess(v bool) *ListScriptPublishHistoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListScriptPublishHistoriesResponseBody) Validate() error {
	if s.ScriptPublishHistories != nil {
		if err := s.ScriptPublishHistories.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListScriptPublishHistoriesResponseBodyScriptPublishHistories struct {
	// Data array
	List []*ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Count per page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total count
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScriptPublishHistoriesResponseBodyScriptPublishHistories) String() string {
	return dara.Prettify(s)
}

func (s ListScriptPublishHistoriesResponseBodyScriptPublishHistories) GoString() string {
	return s.String()
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistories) GetList() []*ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList {
	return s.List
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistories) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistories) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistories) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistories) SetList(v []*ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList) *ListScriptPublishHistoriesResponseBodyScriptPublishHistories {
	s.List = v
	return s
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistories) SetPageNumber(v int32) *ListScriptPublishHistoriesResponseBodyScriptPublishHistories {
	s.PageNumber = &v
	return s
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistories) SetPageSize(v int32) *ListScriptPublishHistoriesResponseBodyScriptPublishHistories {
	s.PageSize = &v
	return s
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistories) SetTotalCount(v int32) *ListScriptPublishHistoriesResponseBodyScriptPublishHistories {
	s.TotalCount = &v
	return s
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistories) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList struct {
	// Publish description
	//
	// example:
	//
	// 第一次发布
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Instance ID
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Published At
	//
	// example:
	//
	// 1578965079000
	PublishTime *int64 `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// Scenario ID
	//
	// example:
	//
	// 29420f65-8f1f-4009-b2f8-f4f7b5d59090
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Script version
	//
	// example:
	//
	// 1578965079000
	ScriptVersion *string `json:"ScriptVersion,omitempty" xml:"ScriptVersion,omitempty"`
}

func (s ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList) String() string {
	return dara.Prettify(s)
}

func (s ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList) GoString() string {
	return s.String()
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList) GetDescription() *string {
	return s.Description
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList) GetPublishTime() *int64 {
	return s.PublishTime
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList) GetScriptVersion() *string {
	return s.ScriptVersion
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList) SetDescription(v string) *ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList {
	s.Description = &v
	return s
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList) SetInstanceId(v string) *ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList {
	s.InstanceId = &v
	return s
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList) SetPublishTime(v int64) *ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList {
	s.PublishTime = &v
	return s
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList) SetScriptId(v string) *ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList {
	s.ScriptId = &v
	return s
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList) SetScriptVersion(v string) *ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList {
	s.ScriptVersion = &v
	return s
}

func (s *ListScriptPublishHistoriesResponseBodyScriptPublishHistoriesList) Validate() error {
	return dara.Validate(s)
}
