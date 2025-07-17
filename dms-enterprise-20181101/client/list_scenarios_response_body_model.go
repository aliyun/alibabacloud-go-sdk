// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScenariosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListScenariosResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListScenariosResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListScenariosResponseBody
	GetRequestId() *string
	SetScenarioList(v []*ListScenariosResponseBodyScenarioList) *ListScenariosResponseBody
	GetScenarioList() []*ListScenariosResponseBodyScenarioList
	SetSuccess(v bool) *ListScenariosResponseBody
	GetSuccess() *bool
}

type ListScenariosResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7FAD400F-7A5C-4193-8F9A-39D86C4F0231
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the returned business scenarios.
	ScenarioList []*ListScenariosResponseBodyScenarioList `json:"ScenarioList,omitempty" xml:"ScenarioList,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListScenariosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScenariosResponseBody) GoString() string {
	return s.String()
}

func (s *ListScenariosResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListScenariosResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListScenariosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScenariosResponseBody) GetScenarioList() []*ListScenariosResponseBodyScenarioList {
	return s.ScenarioList
}

func (s *ListScenariosResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScenariosResponseBody) SetErrorCode(v string) *ListScenariosResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListScenariosResponseBody) SetErrorMessage(v string) *ListScenariosResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListScenariosResponseBody) SetRequestId(v string) *ListScenariosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScenariosResponseBody) SetScenarioList(v []*ListScenariosResponseBodyScenarioList) *ListScenariosResponseBody {
	s.ScenarioList = v
	return s
}

func (s *ListScenariosResponseBody) SetSuccess(v bool) *ListScenariosResponseBody {
	s.Success = &v
	return s
}

func (s *ListScenariosResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListScenariosResponseBodyScenarioList struct {
	// The ID of the user who created the business scenario.
	//
	// example:
	//
	// 23***
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The description of the business scenario.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the business scenario.
	//
	// example:
	//
	// 41***
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the business scenario.
	//
	// example:
	//
	// test
	ScenarioName *string `json:"ScenarioName,omitempty" xml:"ScenarioName,omitempty"`
}

func (s ListScenariosResponseBodyScenarioList) String() string {
	return dara.Prettify(s)
}

func (s ListScenariosResponseBodyScenarioList) GoString() string {
	return s.String()
}

func (s *ListScenariosResponseBodyScenarioList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListScenariosResponseBodyScenarioList) GetDescription() *string {
	return s.Description
}

func (s *ListScenariosResponseBodyScenarioList) GetId() *int64 {
	return s.Id
}

func (s *ListScenariosResponseBodyScenarioList) GetScenarioName() *string {
	return s.ScenarioName
}

func (s *ListScenariosResponseBodyScenarioList) SetCreatorId(v string) *ListScenariosResponseBodyScenarioList {
	s.CreatorId = &v
	return s
}

func (s *ListScenariosResponseBodyScenarioList) SetDescription(v string) *ListScenariosResponseBodyScenarioList {
	s.Description = &v
	return s
}

func (s *ListScenariosResponseBodyScenarioList) SetId(v int64) *ListScenariosResponseBodyScenarioList {
	s.Id = &v
	return s
}

func (s *ListScenariosResponseBodyScenarioList) SetScenarioName(v string) *ListScenariosResponseBodyScenarioList {
	s.ScenarioName = &v
	return s
}

func (s *ListScenariosResponseBodyScenarioList) Validate() error {
	return dara.Validate(s)
}
