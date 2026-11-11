// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAccountResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountResource(v *GetCloudAccountResourceResponseBodyCloudAccountResource) *GetCloudAccountResourceResponseBody
	GetCloudAccountResource() *GetCloudAccountResourceResponseBodyCloudAccountResource
	SetRequestId(v string) *GetCloudAccountResourceResponseBody
	GetRequestId() *string
}

type GetCloudAccountResourceResponseBody struct {
	// The cloud account resource information.
	CloudAccountResource *GetCloudAccountResourceResponseBodyCloudAccountResource `json:"CloudAccountResource,omitempty" xml:"CloudAccountResource,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCloudAccountResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudAccountResourceResponseBody) GetCloudAccountResource() *GetCloudAccountResourceResponseBodyCloudAccountResource {
	return s.CloudAccountResource
}

func (s *GetCloudAccountResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCloudAccountResourceResponseBody) SetCloudAccountResource(v *GetCloudAccountResourceResponseBodyCloudAccountResource) *GetCloudAccountResourceResponseBody {
	s.CloudAccountResource = v
	return s
}

func (s *GetCloudAccountResourceResponseBody) SetRequestId(v string) *GetCloudAccountResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudAccountResourceResponseBody) Validate() error {
	if s.CloudAccountResource != nil {
		if err := s.CloudAccountResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCloudAccountResourceResponseBodyCloudAccountResource struct {
	// The ID of the cloud account.
	//
	// example:
	//
	// 123456789012xxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The identifier of the resource entity, which is a unique remote identifier.
	//
	// example:
	//
	// System:AliyunIntelligentcockpitReadOnlyAccess
	CloudAccountResourceEntityId *string `json:"CloudAccountResourceEntityId,omitempty" xml:"CloudAccountResourceEntityId,omitempty"`
	// The type of the resource entity. Valid values: ram_policy.
	//
	// example:
	//
	// ram_policy
	CloudAccountResourceEntityType *string `json:"CloudAccountResourceEntityType,omitempty" xml:"CloudAccountResourceEntityType,omitempty"`
	// The ID of the cloud account resource.
	//
	// example:
	//
	// car_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
	CloudAccountResourceId *string `json:"CloudAccountResourceId,omitempty" xml:"CloudAccountResourceId,omitempty"`
	// The metadata information of the resource.
	CloudAccountResourceMeta *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta `json:"CloudAccountResourceMeta,omitempty" xml:"CloudAccountResourceMeta,omitempty" type:"Struct"`
	// The name of the resource.
	//
	// example:
	//
	// eiam-sso-role
	CloudAccountResourceName *string `json:"CloudAccountResourceName,omitempty" xml:"CloudAccountResourceName,omitempty"`
	// The time when the resource was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// Cloud account resource used for SSO
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the resource was updated.
	//
	// example:
	//
	// 1649916625000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetCloudAccountResourceResponseBodyCloudAccountResource) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountResourceResponseBodyCloudAccountResource) GoString() string {
	return s.String()
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) GetCloudAccountResourceEntityId() *string {
	return s.CloudAccountResourceEntityId
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) GetCloudAccountResourceEntityType() *string {
	return s.CloudAccountResourceEntityType
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) GetCloudAccountResourceId() *string {
	return s.CloudAccountResourceId
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) GetCloudAccountResourceMeta() *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta {
	return s.CloudAccountResourceMeta
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) GetCloudAccountResourceName() *string {
	return s.CloudAccountResourceName
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) GetDescription() *string {
	return s.Description
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) SetCloudAccountId(v string) *GetCloudAccountResourceResponseBodyCloudAccountResource {
	s.CloudAccountId = &v
	return s
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) SetCloudAccountResourceEntityId(v string) *GetCloudAccountResourceResponseBodyCloudAccountResource {
	s.CloudAccountResourceEntityId = &v
	return s
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) SetCloudAccountResourceEntityType(v string) *GetCloudAccountResourceResponseBodyCloudAccountResource {
	s.CloudAccountResourceEntityType = &v
	return s
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) SetCloudAccountResourceId(v string) *GetCloudAccountResourceResponseBodyCloudAccountResource {
	s.CloudAccountResourceId = &v
	return s
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) SetCloudAccountResourceMeta(v *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta) *GetCloudAccountResourceResponseBodyCloudAccountResource {
	s.CloudAccountResourceMeta = v
	return s
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) SetCloudAccountResourceName(v string) *GetCloudAccountResourceResponseBodyCloudAccountResource {
	s.CloudAccountResourceName = &v
	return s
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) SetCreateTime(v int64) *GetCloudAccountResourceResponseBodyCloudAccountResource {
	s.CreateTime = &v
	return s
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) SetDescription(v string) *GetCloudAccountResourceResponseBodyCloudAccountResource {
	s.Description = &v
	return s
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) SetUpdateTime(v int64) *GetCloudAccountResourceResponseBodyCloudAccountResource {
	s.UpdateTime = &v
	return s
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResource) Validate() error {
	if s.CloudAccountResourceMeta != nil {
		if err := s.CloudAccountResourceMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta struct {
	// The JSON string of the Tencent Cloud CAM policy document. This parameter is returned only for Tencent Cloud accounts.
	//
	// example:
	//
	// {"Version":"1","Statement":[]}
	CamPolicyDocument *string `json:"CamPolicyDocument,omitempty" xml:"CamPolicyDocument,omitempty"`
	// The type of the Tencent Cloud CAM policy. Valid values: qcs and local. This parameter is returned only for Tencent Cloud accounts.
	//
	// example:
	//
	// qcs
	CamPolicyType *string `json:"CamPolicyType,omitempty" xml:"CamPolicyType,omitempty"`
	// The JSON string of the AWS IAM policy document. This parameter is returned only for AWS cloud accounts.
	//
	// example:
	//
	// {"Version":"2012-10-17","Statement":[]}
	IamPolicyDocument *string `json:"IamPolicyDocument,omitempty" xml:"IamPolicyDocument,omitempty"`
	// The type of the AWS IAM policy. Valid values: aws and local. This parameter is returned only for AWS cloud accounts.
	//
	// example:
	//
	// aws
	IamPolicyType *string `json:"IamPolicyType,omitempty" xml:"IamPolicyType,omitempty"`
	// The JSON string of the RAM policy document. This parameter is returned only by the Get operation.
	//
	// example:
	//
	// {"Version":"1","Statement":[]}
	RamPolicyDocument *string `json:"RamPolicyDocument,omitempty" xml:"RamPolicyDocument,omitempty"`
	// The type of the RAM policy. Valid values: system and custom.
	//
	// example:
	//
	// custom
	RamPolicyType *string `json:"RamPolicyType,omitempty" xml:"RamPolicyType,omitempty"`
}

func (s GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta) GoString() string {
	return s.String()
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta) GetCamPolicyDocument() *string {
	return s.CamPolicyDocument
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta) GetCamPolicyType() *string {
	return s.CamPolicyType
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta) GetIamPolicyDocument() *string {
	return s.IamPolicyDocument
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta) GetIamPolicyType() *string {
	return s.IamPolicyType
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta) GetRamPolicyDocument() *string {
	return s.RamPolicyDocument
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta) GetRamPolicyType() *string {
	return s.RamPolicyType
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta) SetCamPolicyDocument(v string) *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta {
	s.CamPolicyDocument = &v
	return s
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta) SetCamPolicyType(v string) *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta {
	s.CamPolicyType = &v
	return s
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta) SetIamPolicyDocument(v string) *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta {
	s.IamPolicyDocument = &v
	return s
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta) SetIamPolicyType(v string) *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta {
	s.IamPolicyType = &v
	return s
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta) SetRamPolicyDocument(v string) *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta {
	s.RamPolicyDocument = &v
	return s
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta) SetRamPolicyType(v string) *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta {
	s.RamPolicyType = &v
	return s
}

func (s *GetCloudAccountResourceResponseBodyCloudAccountResourceCloudAccountResourceMeta) Validate() error {
	return dara.Validate(s)
}
