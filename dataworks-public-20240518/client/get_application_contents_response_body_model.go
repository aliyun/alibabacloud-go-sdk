// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationContentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetApplicationContentsResponseBodyData) *GetApplicationContentsResponseBody
	GetData() *GetApplicationContentsResponseBodyData
	SetRequestId(v string) *GetApplicationContentsResponseBody
	GetRequestId() *string
}

type GetApplicationContentsResponseBody struct {
	// The process instance and its associated application contents.
	Data *GetApplicationContentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. Use this ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 34267E2E-0335-1A60-A1F0-ADA530890CBA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationContentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationContentsResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationContentsResponseBody) GetData() *GetApplicationContentsResponseBodyData {
	return s.Data
}

func (s *GetApplicationContentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationContentsResponseBody) SetData(v *GetApplicationContentsResponseBodyData) *GetApplicationContentsResponseBody {
	s.Data = v
	return s
}

func (s *GetApplicationContentsResponseBody) SetRequestId(v string) *GetApplicationContentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationContentsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationContentsResponseBodyData struct {
	// The time when the application was submitted. This value is a millisecond-precision timestamp.
	//
	// example:
	//
	// 1779675618000
	ApplicationTime *int64 `json:"ApplicationTime,omitempty" xml:"ApplicationTime,omitempty"`
	// A list of the application contents.
	Contents []*GetApplicationContentsResponseBodyDataContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	// The resource type.
	//
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// The process instance ID.
	//
	// example:
	//
	// 332066440109224007
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	// The reason for the application.
	//
	// example:
	//
	// 业务需要
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The approval status. Valid values:
	//
	// - `WaitApproval`: The application is pending approval.
	//
	// - `Confirmed`: The application is pending authorization.
	//
	// - `RejectApproval`: The application was rejected.
	//
	// - `AuthorizeSucceed`: Authorization was successful.
	//
	// - `AuthorizeFailed`: Authorization failed.
	//
	// - `Deleted`: The application was deleted.
	//
	// - `Canceled`: The application was canceled.
	//
	// example:
	//
	// WaitApproval
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetApplicationContentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationContentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApplicationContentsResponseBodyData) GetApplicationTime() *int64 {
	return s.ApplicationTime
}

func (s *GetApplicationContentsResponseBodyData) GetContents() []*GetApplicationContentsResponseBodyDataContents {
	return s.Contents
}

func (s *GetApplicationContentsResponseBodyData) GetDefSchema() *string {
	return s.DefSchema
}

func (s *GetApplicationContentsResponseBodyData) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetApplicationContentsResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *GetApplicationContentsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetApplicationContentsResponseBodyData) SetApplicationTime(v int64) *GetApplicationContentsResponseBodyData {
	s.ApplicationTime = &v
	return s
}

func (s *GetApplicationContentsResponseBodyData) SetContents(v []*GetApplicationContentsResponseBodyDataContents) *GetApplicationContentsResponseBodyData {
	s.Contents = v
	return s
}

func (s *GetApplicationContentsResponseBodyData) SetDefSchema(v string) *GetApplicationContentsResponseBodyData {
	s.DefSchema = &v
	return s
}

func (s *GetApplicationContentsResponseBodyData) SetProcessInstanceId(v string) *GetApplicationContentsResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetApplicationContentsResponseBodyData) SetReason(v string) *GetApplicationContentsResponseBodyData {
	s.Reason = &v
	return s
}

func (s *GetApplicationContentsResponseBodyData) SetStatus(v string) *GetApplicationContentsResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetApplicationContentsResponseBodyData) Validate() error {
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

type GetApplicationContentsResponseBodyDataContents struct {
	// A list of the permissions requested for the resource.
	AccessTypes []*string `json:"AccessTypes,omitempty" xml:"AccessTypes,omitempty" type:"Repeated"`
	// The authorization method.
	//
	// example:
	//
	// ranger
	AuthMethod *string `json:"AuthMethod,omitempty" xml:"AuthMethod,omitempty"`
	// The time when the content item was created. This value is a millisecond-precision timestamp.
	//
	// example:
	//
	// 1773972024000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The resource type.
	//
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// The time when the permissions expire. This value is a millisecond-precision timestamp.
	//
	// example:
	//
	// 1785835708000
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// A list of the permissions granted in the final approval.
	FinalAccessTypes []*string `json:"FinalAccessTypes,omitempty" xml:"FinalAccessTypes,omitempty" type:"Repeated"`
	// The grantee.
	Grantee *GetApplicationContentsResponseBodyDataContentsGrantee `json:"Grantee,omitempty" xml:"Grantee,omitempty" type:"Struct"`
	// The unique ID of the application content item.
	//
	// example:
	//
	// Y9H7AKFmjhWzLYdZNDZA5
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the approval process instance for the application.
	//
	// example:
	//
	// 777799223
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	// The resource declaration.
	Resource *GetApplicationContentsResponseBodyDataContentsResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// The specific type of the resource, such as a table.
	//
	// example:
	//
	// table
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The approval status. Valid values:
	//
	// - `WaitApproval`: The item is pending approval.
	//
	// - `Confirmed`: The item is pending authorization.
	//
	// - `RejectApproval`: The item was rejected.
	//
	// - `AuthorizeSucceed`: Authorization was successful.
	//
	// - `AuthorizeFailed`: Authorization failed.
	//
	// - `Deleted`: The item was deleted during the approval process.
	//
	// - `Canceled`: The item was canceled.
	//
	// example:
	//
	// Deleted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 69973837489
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The time when the content item was last updated. This value is a millisecond-precision timestamp.
	//
	// example:
	//
	// 1773972024000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetApplicationContentsResponseBodyDataContents) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationContentsResponseBodyDataContents) GoString() string {
	return s.String()
}

func (s *GetApplicationContentsResponseBodyDataContents) GetAccessTypes() []*string {
	return s.AccessTypes
}

func (s *GetApplicationContentsResponseBodyDataContents) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *GetApplicationContentsResponseBodyDataContents) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetApplicationContentsResponseBodyDataContents) GetDefSchema() *string {
	return s.DefSchema
}

func (s *GetApplicationContentsResponseBodyDataContents) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *GetApplicationContentsResponseBodyDataContents) GetFinalAccessTypes() []*string {
	return s.FinalAccessTypes
}

func (s *GetApplicationContentsResponseBodyDataContents) GetGrantee() *GetApplicationContentsResponseBodyDataContentsGrantee {
	return s.Grantee
}

func (s *GetApplicationContentsResponseBodyDataContents) GetId() *string {
	return s.Id
}

func (s *GetApplicationContentsResponseBodyDataContents) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetApplicationContentsResponseBodyDataContents) GetResource() *GetApplicationContentsResponseBodyDataContentsResource {
	return s.Resource
}

func (s *GetApplicationContentsResponseBodyDataContents) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetApplicationContentsResponseBodyDataContents) GetStatus() *string {
	return s.Status
}

func (s *GetApplicationContentsResponseBodyDataContents) GetTenantId() *string {
	return s.TenantId
}

func (s *GetApplicationContentsResponseBodyDataContents) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetApplicationContentsResponseBodyDataContents) SetAccessTypes(v []*string) *GetApplicationContentsResponseBodyDataContents {
	s.AccessTypes = v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContents) SetAuthMethod(v string) *GetApplicationContentsResponseBodyDataContents {
	s.AuthMethod = &v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContents) SetCreateTime(v int64) *GetApplicationContentsResponseBodyDataContents {
	s.CreateTime = &v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContents) SetDefSchema(v string) *GetApplicationContentsResponseBodyDataContents {
	s.DefSchema = &v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContents) SetExpirationTime(v int64) *GetApplicationContentsResponseBodyDataContents {
	s.ExpirationTime = &v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContents) SetFinalAccessTypes(v []*string) *GetApplicationContentsResponseBodyDataContents {
	s.FinalAccessTypes = v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContents) SetGrantee(v *GetApplicationContentsResponseBodyDataContentsGrantee) *GetApplicationContentsResponseBodyDataContents {
	s.Grantee = v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContents) SetId(v string) *GetApplicationContentsResponseBodyDataContents {
	s.Id = &v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContents) SetProcessInstanceId(v string) *GetApplicationContentsResponseBodyDataContents {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContents) SetResource(v *GetApplicationContentsResponseBodyDataContentsResource) *GetApplicationContentsResponseBodyDataContents {
	s.Resource = v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContents) SetResourceName(v string) *GetApplicationContentsResponseBodyDataContents {
	s.ResourceName = &v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContents) SetStatus(v string) *GetApplicationContentsResponseBodyDataContents {
	s.Status = &v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContents) SetTenantId(v string) *GetApplicationContentsResponseBodyDataContents {
	s.TenantId = &v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContents) SetUpdateTime(v int64) *GetApplicationContentsResponseBodyDataContents {
	s.UpdateTime = &v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContents) Validate() error {
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

type GetApplicationContentsResponseBodyDataContentsGrantee struct {
	// The ID of the principal. The format of the ID varies based on the `PrincipalType` value:
	//
	// - If `PrincipalType` is `RamUser`, this parameter specifies the ID of a DataWorks user.
	//
	// - If `PrincipalType` is `RamRole`, this parameter specifies the ID of a role in DataWorks. The ID must be prefixed with `ROLE_`.
	//
	// - If `PrincipalType` is `DlfRole`, this parameter specifies the name of a DlfNext role.
	//
	// example:
	//
	// ROLE_3133343434
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The principal type. Valid values:
	//
	// - `RamUser`
	//
	// - `RamRole`
	//
	// - `DlfRole`
	//
	// example:
	//
	// RamRole
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s GetApplicationContentsResponseBodyDataContentsGrantee) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationContentsResponseBodyDataContentsGrantee) GoString() string {
	return s.String()
}

func (s *GetApplicationContentsResponseBodyDataContentsGrantee) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *GetApplicationContentsResponseBodyDataContentsGrantee) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *GetApplicationContentsResponseBodyDataContentsGrantee) SetPrincipalId(v string) *GetApplicationContentsResponseBodyDataContentsGrantee {
	s.PrincipalId = &v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContentsGrantee) SetPrincipalType(v string) *GetApplicationContentsResponseBodyDataContentsGrantee {
	s.PrincipalType = &v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContentsGrantee) Validate() error {
	return dara.Validate(s)
}

type GetApplicationContentsResponseBodyDataContentsResource struct {
	// The name of the `ResourceSchema` that defines how to parse this resource.
	//
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// The version of the `ResourceSchema` that defines how to parse this resource.
	//
	// example:
	//
	// v1.0.0
	DefVersion *string `json:"DefVersion,omitempty" xml:"DefVersion,omitempty"`
	// The resource metadata. The structure of the metadata is defined by the `ResourceSchema`.
	//
	// example:
	//
	// "{\\"schema\\":\\"default\\",\\"threeTierModel\\":false,\\"workspace\\":\\"449656\\",\\"project\\":\\"sync_destination\\",\\"table\\":\\"order_table\\",\\"tenant\\":\\"524997424564736\\"}"
	MetaData *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
}

func (s GetApplicationContentsResponseBodyDataContentsResource) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationContentsResponseBodyDataContentsResource) GoString() string {
	return s.String()
}

func (s *GetApplicationContentsResponseBodyDataContentsResource) GetDefSchema() *string {
	return s.DefSchema
}

func (s *GetApplicationContentsResponseBodyDataContentsResource) GetDefVersion() *string {
	return s.DefVersion
}

func (s *GetApplicationContentsResponseBodyDataContentsResource) GetMetaData() *string {
	return s.MetaData
}

func (s *GetApplicationContentsResponseBodyDataContentsResource) SetDefSchema(v string) *GetApplicationContentsResponseBodyDataContentsResource {
	s.DefSchema = &v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContentsResource) SetDefVersion(v string) *GetApplicationContentsResponseBodyDataContentsResource {
	s.DefVersion = &v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContentsResource) SetMetaData(v string) *GetApplicationContentsResponseBodyDataContentsResource {
	s.MetaData = &v
	return s
}

func (s *GetApplicationContentsResponseBodyDataContentsResource) Validate() error {
	return dara.Validate(s)
}
