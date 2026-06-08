// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPendingApprovalsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListPendingApprovalsResponseBodyData) *ListPendingApprovalsResponseBody
	GetData() *ListPendingApprovalsResponseBodyData
	SetRequestId(v string) *ListPendingApprovalsResponseBody
	GetRequestId() *string
}

type ListPendingApprovalsResponseBody struct {
	Data *ListPendingApprovalsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0bc5df3a17****903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPendingApprovalsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPendingApprovalsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPendingApprovalsResponseBody) GetData() *ListPendingApprovalsResponseBodyData {
	return s.Data
}

func (s *ListPendingApprovalsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPendingApprovalsResponseBody) SetData(v *ListPendingApprovalsResponseBodyData) *ListPendingApprovalsResponseBody {
	s.Data = v
	return s
}

func (s *ListPendingApprovalsResponseBody) SetRequestId(v string) *ListPendingApprovalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPendingApprovalsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPendingApprovalsResponseBodyData struct {
	Data []*ListPendingApprovalsResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// false
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// example:
	//
	// eyJpZCI6MTIzfQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPendingApprovalsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPendingApprovalsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPendingApprovalsResponseBodyData) GetData() []*ListPendingApprovalsResponseBodyDataData {
	return s.Data
}

func (s *ListPendingApprovalsResponseBodyData) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListPendingApprovalsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPendingApprovalsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPendingApprovalsResponseBodyData) SetData(v []*ListPendingApprovalsResponseBodyDataData) *ListPendingApprovalsResponseBodyData {
	s.Data = v
	return s
}

func (s *ListPendingApprovalsResponseBodyData) SetHasMore(v bool) *ListPendingApprovalsResponseBodyData {
	s.HasMore = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyData) SetNextToken(v string) *ListPendingApprovalsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyData) SetPageSize(v int32) *ListPendingApprovalsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyData) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPendingApprovalsResponseBodyDataData struct {
	ApplicationTime *int64                                              `json:"ApplicationTime,omitempty" xml:"ApplicationTime,omitempty"`
	Contents        []*ListPendingApprovalsResponseBodyDataDataContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// example:
	//
	// 176906667488145
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	Reason            *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// WAIT_APPROVAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListPendingApprovalsResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s ListPendingApprovalsResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ListPendingApprovalsResponseBodyDataData) GetApplicationTime() *int64 {
	return s.ApplicationTime
}

func (s *ListPendingApprovalsResponseBodyDataData) GetContents() []*ListPendingApprovalsResponseBodyDataDataContents {
	return s.Contents
}

func (s *ListPendingApprovalsResponseBodyDataData) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListPendingApprovalsResponseBodyDataData) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *ListPendingApprovalsResponseBodyDataData) GetReason() *string {
	return s.Reason
}

func (s *ListPendingApprovalsResponseBodyDataData) GetStatus() *string {
	return s.Status
}

func (s *ListPendingApprovalsResponseBodyDataData) SetApplicationTime(v int64) *ListPendingApprovalsResponseBodyDataData {
	s.ApplicationTime = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataData) SetContents(v []*ListPendingApprovalsResponseBodyDataDataContents) *ListPendingApprovalsResponseBodyDataData {
	s.Contents = v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataData) SetDefSchema(v string) *ListPendingApprovalsResponseBodyDataData {
	s.DefSchema = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataData) SetProcessInstanceId(v string) *ListPendingApprovalsResponseBodyDataData {
	s.ProcessInstanceId = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataData) SetReason(v string) *ListPendingApprovalsResponseBodyDataData {
	s.Reason = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataData) SetStatus(v string) *ListPendingApprovalsResponseBodyDataData {
	s.Status = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataData) Validate() error {
	if s.Contents != nil {
		for _, item := range s.Contents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPendingApprovalsResponseBodyDataDataContents struct {
	AccessTypes []*string `json:"AccessTypes,omitempty" xml:"AccessTypes,omitempty" type:"Repeated"`
	// example:
	//
	// default
	AuthMethod *string `json:"AuthMethod,omitempty" xml:"AuthMethod,omitempty"`
	// example:
	//
	// 2025-09-11 10:13:21
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// example:
	//
	// 1782354014507
	ExpirationTime   *int64                                                   `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	FinalAccessTypes []*string                                                `json:"FinalAccessTypes,omitempty" xml:"FinalAccessTypes,omitempty" type:"Repeated"`
	Grantee          *ListPendingApprovalsResponseBodyDataDataContentsGrantee `json:"Grantee,omitempty" xml:"Grantee,omitempty" type:"Struct"`
	// example:
	//
	// 210001918826
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 176906667488145
	ProcessInstanceId *string                                                   `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	Resource          *ListPendingApprovalsResponseBodyDataDataContentsResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// example:
	//
	// table
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// WAIT_APPROVAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 69973837489
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 2022-07-06 19:13:05
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListPendingApprovalsResponseBodyDataDataContents) String() string {
	return dara.Prettify(s)
}

func (s ListPendingApprovalsResponseBodyDataDataContents) GoString() string {
	return s.String()
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) GetAccessTypes() []*string {
	return s.AccessTypes
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) GetFinalAccessTypes() []*string {
	return s.FinalAccessTypes
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) GetGrantee() *ListPendingApprovalsResponseBodyDataDataContentsGrantee {
	return s.Grantee
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) GetId() *string {
	return s.Id
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) GetResource() *ListPendingApprovalsResponseBodyDataDataContentsResource {
	return s.Resource
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) GetStatus() *string {
	return s.Status
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) GetTenantId() *string {
	return s.TenantId
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) SetAccessTypes(v []*string) *ListPendingApprovalsResponseBodyDataDataContents {
	s.AccessTypes = v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) SetAuthMethod(v string) *ListPendingApprovalsResponseBodyDataDataContents {
	s.AuthMethod = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) SetCreateTime(v int64) *ListPendingApprovalsResponseBodyDataDataContents {
	s.CreateTime = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) SetDefSchema(v string) *ListPendingApprovalsResponseBodyDataDataContents {
	s.DefSchema = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) SetExpirationTime(v int64) *ListPendingApprovalsResponseBodyDataDataContents {
	s.ExpirationTime = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) SetFinalAccessTypes(v []*string) *ListPendingApprovalsResponseBodyDataDataContents {
	s.FinalAccessTypes = v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) SetGrantee(v *ListPendingApprovalsResponseBodyDataDataContentsGrantee) *ListPendingApprovalsResponseBodyDataDataContents {
	s.Grantee = v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) SetId(v string) *ListPendingApprovalsResponseBodyDataDataContents {
	s.Id = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) SetProcessInstanceId(v string) *ListPendingApprovalsResponseBodyDataDataContents {
	s.ProcessInstanceId = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) SetResource(v *ListPendingApprovalsResponseBodyDataDataContentsResource) *ListPendingApprovalsResponseBodyDataDataContents {
	s.Resource = v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) SetResourceName(v string) *ListPendingApprovalsResponseBodyDataDataContents {
	s.ResourceName = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) SetStatus(v string) *ListPendingApprovalsResponseBodyDataDataContents {
	s.Status = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) SetTenantId(v string) *ListPendingApprovalsResponseBodyDataDataContents {
	s.TenantId = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) SetUpdateTime(v int64) *ListPendingApprovalsResponseBodyDataDataContents {
	s.UpdateTime = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContents) Validate() error {
	if s.Grantee != nil {
		if err := s.Grantee.Validate(); err != nil {
			return err
		}
	}
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPendingApprovalsResponseBodyDataDataContentsGrantee struct {
	// example:
	//
	// 213463068144525171
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// example:
	//
	// User
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s ListPendingApprovalsResponseBodyDataDataContentsGrantee) String() string {
	return dara.Prettify(s)
}

func (s ListPendingApprovalsResponseBodyDataDataContentsGrantee) GoString() string {
	return s.String()
}

func (s *ListPendingApprovalsResponseBodyDataDataContentsGrantee) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListPendingApprovalsResponseBodyDataDataContentsGrantee) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListPendingApprovalsResponseBodyDataDataContentsGrantee) SetPrincipalId(v string) *ListPendingApprovalsResponseBodyDataDataContentsGrantee {
	s.PrincipalId = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContentsGrantee) SetPrincipalType(v string) *ListPendingApprovalsResponseBodyDataDataContentsGrantee {
	s.PrincipalType = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContentsGrantee) Validate() error {
	return dara.Validate(s)
}

type ListPendingApprovalsResponseBodyDataDataContentsResource struct {
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// example:
	//
	// v1.0.0
	DefVersion *string                `json:"DefVersion,omitempty" xml:"DefVersion,omitempty"`
	MetaData   map[string]interface{} `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
}

func (s ListPendingApprovalsResponseBodyDataDataContentsResource) String() string {
	return dara.Prettify(s)
}

func (s ListPendingApprovalsResponseBodyDataDataContentsResource) GoString() string {
	return s.String()
}

func (s *ListPendingApprovalsResponseBodyDataDataContentsResource) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListPendingApprovalsResponseBodyDataDataContentsResource) GetDefVersion() *string {
	return s.DefVersion
}

func (s *ListPendingApprovalsResponseBodyDataDataContentsResource) GetMetaData() map[string]interface{} {
	return s.MetaData
}

func (s *ListPendingApprovalsResponseBodyDataDataContentsResource) SetDefSchema(v string) *ListPendingApprovalsResponseBodyDataDataContentsResource {
	s.DefSchema = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContentsResource) SetDefVersion(v string) *ListPendingApprovalsResponseBodyDataDataContentsResource {
	s.DefVersion = &v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContentsResource) SetMetaData(v map[string]interface{}) *ListPendingApprovalsResponseBodyDataDataContentsResource {
	s.MetaData = v
	return s
}

func (s *ListPendingApprovalsResponseBodyDataDataContentsResource) Validate() error {
	return dara.Validate(s)
}
