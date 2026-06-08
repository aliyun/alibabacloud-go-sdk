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
	Data *GetApplicationContentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// example:
	//
	// 1779675618000
	ApplicationTime *int64                                            `json:"ApplicationTime,omitempty" xml:"ApplicationTime,omitempty"`
	Contents        []*GetApplicationContentsResponseBodyDataContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// example:
	//
	// 332066440109224007
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	Reason            *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
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
	AccessTypes []*string `json:"AccessTypes,omitempty" xml:"AccessTypes,omitempty" type:"Repeated"`
	// example:
	//
	// ranger
	AuthMethod *string `json:"AuthMethod,omitempty" xml:"AuthMethod,omitempty"`
	// example:
	//
	// 1773972024000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// example:
	//
	// 1785835708000
	ExpirationTime   *int64                                                 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	FinalAccessTypes []*string                                              `json:"FinalAccessTypes,omitempty" xml:"FinalAccessTypes,omitempty" type:"Repeated"`
	Grantee          *GetApplicationContentsResponseBodyDataContentsGrantee `json:"Grantee,omitempty" xml:"Grantee,omitempty" type:"Struct"`
	// example:
	//
	// Y9H7AKFmjhWzLYdZNDZA5
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 777799223
	ProcessInstanceId *string                                                 `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	Resource          *GetApplicationContentsResponseBodyDataContentsResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
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
	// example:
	//
	// ROLE_3133343434
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
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
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// example:
	//
	// v1.0.0
	DefVersion *string `json:"DefVersion,omitempty" xml:"DefVersion,omitempty"`
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
