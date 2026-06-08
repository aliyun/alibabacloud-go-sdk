// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListMyApplicationsResponseBodyData) *ListMyApplicationsResponseBody
	GetData() *ListMyApplicationsResponseBodyData
	SetRequestId(v string) *ListMyApplicationsResponseBody
	GetRequestId() *string
}

type ListMyApplicationsResponseBody struct {
	Data *ListMyApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0bc5df3a17****903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMyApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMyApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMyApplicationsResponseBody) GetData() *ListMyApplicationsResponseBodyData {
	return s.Data
}

func (s *ListMyApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMyApplicationsResponseBody) SetData(v *ListMyApplicationsResponseBodyData) *ListMyApplicationsResponseBody {
	s.Data = v
	return s
}

func (s *ListMyApplicationsResponseBody) SetRequestId(v string) *ListMyApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMyApplicationsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMyApplicationsResponseBodyData struct {
	Data []*ListMyApplicationsResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// false
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// example:
	//
	// eyJpZCI6NDU2fQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMyApplicationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMyApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMyApplicationsResponseBodyData) GetData() []*ListMyApplicationsResponseBodyDataData {
	return s.Data
}

func (s *ListMyApplicationsResponseBodyData) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMyApplicationsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMyApplicationsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMyApplicationsResponseBodyData) SetData(v []*ListMyApplicationsResponseBodyDataData) *ListMyApplicationsResponseBodyData {
	s.Data = v
	return s
}

func (s *ListMyApplicationsResponseBodyData) SetHasMore(v bool) *ListMyApplicationsResponseBodyData {
	s.HasMore = &v
	return s
}

func (s *ListMyApplicationsResponseBodyData) SetNextToken(v string) *ListMyApplicationsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListMyApplicationsResponseBodyData) SetPageSize(v int32) *ListMyApplicationsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMyApplicationsResponseBodyData) Validate() error {
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

type ListMyApplicationsResponseBodyDataData struct {
	// example:
	//
	// 1779695088000
	ApplicationTime *int64                                            `json:"ApplicationTime,omitempty" xml:"ApplicationTime,omitempty"`
	Contents        []*ListMyApplicationsResponseBodyDataDataContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
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

func (s ListMyApplicationsResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s ListMyApplicationsResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ListMyApplicationsResponseBodyDataData) GetApplicationTime() *int64 {
	return s.ApplicationTime
}

func (s *ListMyApplicationsResponseBodyDataData) GetContents() []*ListMyApplicationsResponseBodyDataDataContents {
	return s.Contents
}

func (s *ListMyApplicationsResponseBodyDataData) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListMyApplicationsResponseBodyDataData) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *ListMyApplicationsResponseBodyDataData) GetReason() *string {
	return s.Reason
}

func (s *ListMyApplicationsResponseBodyDataData) GetStatus() *string {
	return s.Status
}

func (s *ListMyApplicationsResponseBodyDataData) SetApplicationTime(v int64) *ListMyApplicationsResponseBodyDataData {
	s.ApplicationTime = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataData) SetContents(v []*ListMyApplicationsResponseBodyDataDataContents) *ListMyApplicationsResponseBodyDataData {
	s.Contents = v
	return s
}

func (s *ListMyApplicationsResponseBodyDataData) SetDefSchema(v string) *ListMyApplicationsResponseBodyDataData {
	s.DefSchema = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataData) SetProcessInstanceId(v string) *ListMyApplicationsResponseBodyDataData {
	s.ProcessInstanceId = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataData) SetReason(v string) *ListMyApplicationsResponseBodyDataData {
	s.Reason = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataData) SetStatus(v string) *ListMyApplicationsResponseBodyDataData {
	s.Status = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataData) Validate() error {
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

type ListMyApplicationsResponseBodyDataDataContents struct {
	AccessTypes []*string `json:"AccessTypes,omitempty" xml:"AccessTypes,omitempty" type:"Repeated"`
	// example:
	//
	// default
	AuthMethod *string `json:"AuthMethod,omitempty" xml:"AuthMethod,omitempty"`
	// example:
	//
	// 2022-11-29 15:04:52
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// MAXCOMPUTE
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// example:
	//
	// 1782354014507
	ExpirationTime   *int64                                                 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	FinalAccessTypes []*string                                              `json:"FinalAccessTypes,omitempty" xml:"FinalAccessTypes,omitempty" type:"Repeated"`
	Grantee          *ListMyApplicationsResponseBodyDataDataContentsGrantee `json:"Grantee,omitempty" xml:"Grantee,omitempty" type:"Struct"`
	// example:
	//
	// a8aa620037bb410ea13837f9b4d053d8
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 777799223
	ProcessInstanceId *string                                                 `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	Resource          *ListMyApplicationsResponseBodyDataDataContentsResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
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
	// 2022-07-08 23:58:59
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListMyApplicationsResponseBodyDataDataContents) String() string {
	return dara.Prettify(s)
}

func (s ListMyApplicationsResponseBodyDataDataContents) GoString() string {
	return s.String()
}

func (s *ListMyApplicationsResponseBodyDataDataContents) GetAccessTypes() []*string {
	return s.AccessTypes
}

func (s *ListMyApplicationsResponseBodyDataDataContents) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *ListMyApplicationsResponseBodyDataDataContents) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListMyApplicationsResponseBodyDataDataContents) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListMyApplicationsResponseBodyDataDataContents) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *ListMyApplicationsResponseBodyDataDataContents) GetFinalAccessTypes() []*string {
	return s.FinalAccessTypes
}

func (s *ListMyApplicationsResponseBodyDataDataContents) GetGrantee() *ListMyApplicationsResponseBodyDataDataContentsGrantee {
	return s.Grantee
}

func (s *ListMyApplicationsResponseBodyDataDataContents) GetId() *string {
	return s.Id
}

func (s *ListMyApplicationsResponseBodyDataDataContents) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *ListMyApplicationsResponseBodyDataDataContents) GetResource() *ListMyApplicationsResponseBodyDataDataContentsResource {
	return s.Resource
}

func (s *ListMyApplicationsResponseBodyDataDataContents) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListMyApplicationsResponseBodyDataDataContents) GetStatus() *string {
	return s.Status
}

func (s *ListMyApplicationsResponseBodyDataDataContents) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMyApplicationsResponseBodyDataDataContents) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListMyApplicationsResponseBodyDataDataContents) SetAccessTypes(v []*string) *ListMyApplicationsResponseBodyDataDataContents {
	s.AccessTypes = v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContents) SetAuthMethod(v string) *ListMyApplicationsResponseBodyDataDataContents {
	s.AuthMethod = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContents) SetCreateTime(v int64) *ListMyApplicationsResponseBodyDataDataContents {
	s.CreateTime = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContents) SetDefSchema(v string) *ListMyApplicationsResponseBodyDataDataContents {
	s.DefSchema = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContents) SetExpirationTime(v int64) *ListMyApplicationsResponseBodyDataDataContents {
	s.ExpirationTime = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContents) SetFinalAccessTypes(v []*string) *ListMyApplicationsResponseBodyDataDataContents {
	s.FinalAccessTypes = v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContents) SetGrantee(v *ListMyApplicationsResponseBodyDataDataContentsGrantee) *ListMyApplicationsResponseBodyDataDataContents {
	s.Grantee = v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContents) SetId(v string) *ListMyApplicationsResponseBodyDataDataContents {
	s.Id = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContents) SetProcessInstanceId(v string) *ListMyApplicationsResponseBodyDataDataContents {
	s.ProcessInstanceId = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContents) SetResource(v *ListMyApplicationsResponseBodyDataDataContentsResource) *ListMyApplicationsResponseBodyDataDataContents {
	s.Resource = v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContents) SetResourceName(v string) *ListMyApplicationsResponseBodyDataDataContents {
	s.ResourceName = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContents) SetStatus(v string) *ListMyApplicationsResponseBodyDataDataContents {
	s.Status = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContents) SetTenantId(v string) *ListMyApplicationsResponseBodyDataDataContents {
	s.TenantId = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContents) SetUpdateTime(v int64) *ListMyApplicationsResponseBodyDataDataContents {
	s.UpdateTime = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContents) Validate() error {
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

type ListMyApplicationsResponseBodyDataDataContentsGrantee struct {
	// example:
	//
	// ROLE_3133343434
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// example:
	//
	// RamRole
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s ListMyApplicationsResponseBodyDataDataContentsGrantee) String() string {
	return dara.Prettify(s)
}

func (s ListMyApplicationsResponseBodyDataDataContentsGrantee) GoString() string {
	return s.String()
}

func (s *ListMyApplicationsResponseBodyDataDataContentsGrantee) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListMyApplicationsResponseBodyDataDataContentsGrantee) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListMyApplicationsResponseBodyDataDataContentsGrantee) SetPrincipalId(v string) *ListMyApplicationsResponseBodyDataDataContentsGrantee {
	s.PrincipalId = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContentsGrantee) SetPrincipalType(v string) *ListMyApplicationsResponseBodyDataDataContentsGrantee {
	s.PrincipalType = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContentsGrantee) Validate() error {
	return dara.Validate(s)
}

type ListMyApplicationsResponseBodyDataDataContentsResource struct {
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

func (s ListMyApplicationsResponseBodyDataDataContentsResource) String() string {
	return dara.Prettify(s)
}

func (s ListMyApplicationsResponseBodyDataDataContentsResource) GoString() string {
	return s.String()
}

func (s *ListMyApplicationsResponseBodyDataDataContentsResource) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListMyApplicationsResponseBodyDataDataContentsResource) GetDefVersion() *string {
	return s.DefVersion
}

func (s *ListMyApplicationsResponseBodyDataDataContentsResource) GetMetaData() map[string]interface{} {
	return s.MetaData
}

func (s *ListMyApplicationsResponseBodyDataDataContentsResource) SetDefSchema(v string) *ListMyApplicationsResponseBodyDataDataContentsResource {
	s.DefSchema = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContentsResource) SetDefVersion(v string) *ListMyApplicationsResponseBodyDataDataContentsResource {
	s.DefVersion = &v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContentsResource) SetMetaData(v map[string]interface{}) *ListMyApplicationsResponseBodyDataDataContentsResource {
	s.MetaData = v
	return s
}

func (s *ListMyApplicationsResponseBodyDataDataContentsResource) Validate() error {
	return dara.Validate(s)
}
