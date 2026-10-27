// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderSyncScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetIdentityProviderSyncScopeResponseBody
	GetCode() *string
	SetData(v *GetIdentityProviderSyncScopeResponseBodyData) *GetIdentityProviderSyncScopeResponseBody
	GetData() *GetIdentityProviderSyncScopeResponseBodyData
	SetHttpStatusCode(v int32) *GetIdentityProviderSyncScopeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetIdentityProviderSyncScopeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIdentityProviderSyncScopeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetIdentityProviderSyncScopeResponseBody
	GetSuccess() *bool
}

type GetIdentityProviderSyncScopeResponseBody struct {
	// The business status code of the API request. SUCCESS indicates that the request is successful, but does not indicate that the asynchronous synchronization task is completed.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The business data returned by the request.
	//
	// example:
	//
	// {"syncScope":{"mode":"SelectedDepartments","departmentIds":["10001"]},"selectedDepartments":[{"id":"10001","name":"R&D Department","exists":true}],"estimateStatus":"Ready","expandedDepartmentCount":2,"estimatedUserCount":12}
	Data *GetIdentityProviderSyncScopeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code of the API request. It does not indicate the execution result of the asynchronous task.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message of the API request.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique identifier of the request. Provide this value when troubleshooting issues.
	//
	// example:
	//
	// 01234567-89AB-CDEF-0123-456789ABCDEF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API request is successful. To check whether the asynchronous task is completed, query the task status.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetIdentityProviderSyncScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSyncScopeResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSyncScopeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIdentityProviderSyncScopeResponseBody) GetData() *GetIdentityProviderSyncScopeResponseBodyData {
	return s.Data
}

func (s *GetIdentityProviderSyncScopeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetIdentityProviderSyncScopeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIdentityProviderSyncScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentityProviderSyncScopeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIdentityProviderSyncScopeResponseBody) SetCode(v string) *GetIdentityProviderSyncScopeResponseBody {
	s.Code = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBody) SetData(v *GetIdentityProviderSyncScopeResponseBodyData) *GetIdentityProviderSyncScopeResponseBody {
	s.Data = v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBody) SetHttpStatusCode(v int32) *GetIdentityProviderSyncScopeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBody) SetMessage(v string) *GetIdentityProviderSyncScopeResponseBody {
	s.Message = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBody) SetRequestId(v string) *GetIdentityProviderSyncScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBody) SetSuccess(v bool) *GetIdentityProviderSyncScopeResponseBody {
	s.Success = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderSyncScopeResponseBodyData struct {
	// The estimation status of the user count for the synchronization scope. The count is available only when the status is Ready. A missing count in other states cannot be considered as 0.
	//
	// example:
	//
	// Ready
	EstimateStatus *string `json:"estimateStatus,omitempty" xml:"estimateStatus,omitempty"`
	// The estimated number of users in the saved synchronization scope. This parameter is returned only when estimateStatus is Ready and the downstream system provides the value. It does not represent the number of users that have been synchronized.
	//
	// example:
	//
	// 12
	EstimatedUserCount *int64 `json:"estimatedUserCount,omitempty" xml:"estimatedUserCount,omitempty"`
	// The estimated number of departments including sub-departments. This parameter is returned only when estimateStatus is Ready and the downstream system provides the value.
	//
	// example:
	//
	// 2
	ExpandedDepartmentCount *int64 `json:"expandedDepartmentCount,omitempty" xml:"expandedDepartmentCount,omitempty"`
	// The names and validity of the saved selected departments. This parameter can be an empty array when the mode is AllAuthorized. An empty array does not indicate that the organization has no departments.
	//
	// example:
	//
	// [{"id":"10001","name":"R&D Department","exists":true}]
	SelectedDepartments []*GetIdentityProviderSyncScopeResponseBodyDataSelectedDepartments `json:"selectedDepartments,omitempty" xml:"selectedDepartments,omitempty" type:"Repeated"`
	// The saved synchronization scope downstream. AllAuthorized indicates the address book scope authorized for the application.
	//
	// example:
	//
	// {"mode":"SelectedDepartments","departmentIds":["10001"]}
	SyncScope *GetIdentityProviderSyncScopeResponseBodyDataSyncScope `json:"syncScope,omitempty" xml:"syncScope,omitempty" type:"Struct"`
}

func (s GetIdentityProviderSyncScopeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSyncScopeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSyncScopeResponseBodyData) GetEstimateStatus() *string {
	return s.EstimateStatus
}

func (s *GetIdentityProviderSyncScopeResponseBodyData) GetEstimatedUserCount() *int64 {
	return s.EstimatedUserCount
}

func (s *GetIdentityProviderSyncScopeResponseBodyData) GetExpandedDepartmentCount() *int64 {
	return s.ExpandedDepartmentCount
}

func (s *GetIdentityProviderSyncScopeResponseBodyData) GetSelectedDepartments() []*GetIdentityProviderSyncScopeResponseBodyDataSelectedDepartments {
	return s.SelectedDepartments
}

func (s *GetIdentityProviderSyncScopeResponseBodyData) GetSyncScope() *GetIdentityProviderSyncScopeResponseBodyDataSyncScope {
	return s.SyncScope
}

func (s *GetIdentityProviderSyncScopeResponseBodyData) SetEstimateStatus(v string) *GetIdentityProviderSyncScopeResponseBodyData {
	s.EstimateStatus = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBodyData) SetEstimatedUserCount(v int64) *GetIdentityProviderSyncScopeResponseBodyData {
	s.EstimatedUserCount = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBodyData) SetExpandedDepartmentCount(v int64) *GetIdentityProviderSyncScopeResponseBodyData {
	s.ExpandedDepartmentCount = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBodyData) SetSelectedDepartments(v []*GetIdentityProviderSyncScopeResponseBodyDataSelectedDepartments) *GetIdentityProviderSyncScopeResponseBodyData {
	s.SelectedDepartments = v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBodyData) SetSyncScope(v *GetIdentityProviderSyncScopeResponseBodyDataSyncScope) *GetIdentityProviderSyncScopeResponseBodyData {
	s.SyncScope = v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBodyData) Validate() error {
	if s.SelectedDepartments != nil {
		for _, item := range s.SelectedDepartments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SyncScope != nil {
		if err := s.SyncScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderSyncScopeResponseBodyDataSelectedDepartments struct {
	// Indicates whether the selected department still exists. A value of false indicates that the selection is invalid and the synchronization scope must be reconfirmed.
	//
	// example:
	//
	// true
	Exists *bool `json:"exists,omitempty" xml:"exists,omitempty"`
	// The unique identifier of the department returned by the identity provider. When querying the sub-departments of this node, use this value as parentDepartmentId.
	//
	// example:
	//
	// 10001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the department returned by the identity provider.
	//
	// example:
	//
	// R&D Department
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetIdentityProviderSyncScopeResponseBodyDataSelectedDepartments) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSyncScopeResponseBodyDataSelectedDepartments) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSyncScopeResponseBodyDataSelectedDepartments) GetExists() *bool {
	return s.Exists
}

func (s *GetIdentityProviderSyncScopeResponseBodyDataSelectedDepartments) GetId() *string {
	return s.Id
}

func (s *GetIdentityProviderSyncScopeResponseBodyDataSelectedDepartments) GetName() *string {
	return s.Name
}

func (s *GetIdentityProviderSyncScopeResponseBodyDataSelectedDepartments) SetExists(v bool) *GetIdentityProviderSyncScopeResponseBodyDataSelectedDepartments {
	s.Exists = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBodyDataSelectedDepartments) SetId(v string) *GetIdentityProviderSyncScopeResponseBodyDataSelectedDepartments {
	s.Id = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBodyDataSelectedDepartments) SetName(v string) *GetIdentityProviderSyncScopeResponseBodyDataSelectedDepartments {
	s.Name = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBodyDataSelectedDepartments) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderSyncScopeResponseBodyDataSyncScope struct {
	// When the mode is SelectedDepartments, this parameter contains at least one non-empty department ID, which should be obtained from ListIdentityProviderDepartments. When the mode is AllAuthorized, pass an empty array or omit this parameter.
	//
	// example:
	//
	// ["10001"]
	DepartmentIds []*string `json:"departmentIds,omitempty" xml:"departmentIds,omitempty" type:"Repeated"`
	// The scope mode. Valid values: AllAuthorized (the entire authorized scope, limited by the application\\"s address book authorization) and SelectedDepartments (the selected departments and all their sub-departments). When explicitly submitting syncScope, a valid mode must be provided and verified by the backend.
	//
	// example:
	//
	// SelectedDepartments
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s GetIdentityProviderSyncScopeResponseBodyDataSyncScope) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSyncScopeResponseBodyDataSyncScope) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSyncScopeResponseBodyDataSyncScope) GetDepartmentIds() []*string {
	return s.DepartmentIds
}

func (s *GetIdentityProviderSyncScopeResponseBodyDataSyncScope) GetMode() *string {
	return s.Mode
}

func (s *GetIdentityProviderSyncScopeResponseBodyDataSyncScope) SetDepartmentIds(v []*string) *GetIdentityProviderSyncScopeResponseBodyDataSyncScope {
	s.DepartmentIds = v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBodyDataSyncScope) SetMode(v string) *GetIdentityProviderSyncScopeResponseBodyDataSyncScope {
	s.Mode = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBodyDataSyncScope) Validate() error {
	return dara.Validate(s)
}
