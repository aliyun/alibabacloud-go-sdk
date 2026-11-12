// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAccountResourcesForCloudPermissionTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountResources(v []*ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody
	GetCloudAccountResources() []*ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources
	SetMaxResults(v int64) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody
	GetTotalCount() *int64
}

type ListCloudAccountResourcesForCloudPermissionTemplateResponseBody struct {
	// The list of cloud account resources.
	CloudAccountResources []*ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources `json:"CloudAccountResources,omitempty" xml:"CloudAccountResources,omitempty" type:"Repeated"`
	// The number of entries per page in a paging query.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token returned by the current call. Use this token for the next paginated query.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries in the list.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudAccountResourcesForCloudPermissionTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountResourcesForCloudPermissionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody) GetCloudAccountResources() []*ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources {
	return s.CloudAccountResources
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody) SetCloudAccountResources(v []*ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody {
	s.CloudAccountResources = v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody) SetMaxResults(v int64) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody) SetNextToken(v string) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody) SetRequestId(v string) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody) SetTotalCount(v int64) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody) Validate() error {
	if s.CloudAccountResources != nil {
		for _, item := range s.CloudAccountResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources struct {
	// The ID of the cloud account to which the resource belongs.
	//
	// example:
	//
	// 123456789012xxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The resource entity identifier, which is a unique remote identifier.
	//
	// example:
	//
	// System:AliyunIntelligentcockpitReadOnlyAccessexample
	CloudAccountResourceEntityId *string `json:"CloudAccountResourceEntityId,omitempty" xml:"CloudAccountResourceEntityId,omitempty"`
	// The resource entity type. Valid values: ram_role.
	//
	// example:
	//
	// ram_role
	CloudAccountResourceEntityType *string `json:"CloudAccountResourceEntityType,omitempty" xml:"CloudAccountResourceEntityType,omitempty"`
	// The identifier of the cloud account resource.
	//
	// example:
	//
	// car_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
	CloudAccountResourceId *string `json:"CloudAccountResourceId,omitempty" xml:"CloudAccountResourceId,omitempty"`
	// The resource metadata.
	CloudAccountResourceMeta *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta `json:"CloudAccountResourceMeta,omitempty" xml:"CloudAccountResourceMeta,omitempty" type:"Struct"`
	// The name of the resource.
	//
	// example:
	//
	// eiam-sso-role
	CloudAccountResourceName *string `json:"CloudAccountResourceName,omitempty" xml:"CloudAccountResourceName,omitempty"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
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
	// The last update time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649916625000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) GoString() string {
	return s.String()
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) GetCloudAccountResourceEntityId() *string {
	return s.CloudAccountResourceEntityId
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) GetCloudAccountResourceEntityType() *string {
	return s.CloudAccountResourceEntityType
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) GetCloudAccountResourceId() *string {
	return s.CloudAccountResourceId
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) GetCloudAccountResourceMeta() *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta {
	return s.CloudAccountResourceMeta
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) GetCloudAccountResourceName() *string {
	return s.CloudAccountResourceName
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) GetDescription() *string {
	return s.Description
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) SetCloudAccountId(v string) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources {
	s.CloudAccountId = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) SetCloudAccountResourceEntityId(v string) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources {
	s.CloudAccountResourceEntityId = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) SetCloudAccountResourceEntityType(v string) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources {
	s.CloudAccountResourceEntityType = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) SetCloudAccountResourceId(v string) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources {
	s.CloudAccountResourceId = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) SetCloudAccountResourceMeta(v *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources {
	s.CloudAccountResourceMeta = v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) SetCloudAccountResourceName(v string) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources {
	s.CloudAccountResourceName = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) SetCreateTime(v int64) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources {
	s.CreateTime = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) SetDescription(v string) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources {
	s.Description = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) SetUpdateTime(v int64) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources {
	s.UpdateTime = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResources) Validate() error {
	if s.CloudAccountResourceMeta != nil {
		if err := s.CloudAccountResourceMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta struct {
	// The JSON string of the Tencent Cloud CAM policy document. This parameter is returned only for Tencent Cloud accounts.
	//
	// example:
	//
	// {"Version":"1","Statement":[]}
	CamPolicyDocument *string `json:"CamPolicyDocument,omitempty" xml:"CamPolicyDocument,omitempty"`
	// The Tencent Cloud CAM policy type. Valid values: qcs and local. This parameter is returned only for Tencent Cloud accounts.
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
	// The AWS IAM policy type. Valid values: aws and local. This parameter is returned only for AWS cloud accounts.
	//
	// example:
	//
	// aws
	IamPolicyType *string `json:"IamPolicyType,omitempty" xml:"IamPolicyType,omitempty"`
	// The JSON string of the RAM policy document. This parameter is returned only by the get operation.
	//
	// example:
	//
	// {"Version":"1","Statement":[]}
	RamPolicyDocument *string `json:"RamPolicyDocument,omitempty" xml:"RamPolicyDocument,omitempty"`
	// The RAM policy type. Valid values: system and custom.
	//
	// example:
	//
	// custom
	RamPolicyType *string `json:"RamPolicyType,omitempty" xml:"RamPolicyType,omitempty"`
}

func (s ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta) GoString() string {
	return s.String()
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta) GetCamPolicyDocument() *string {
	return s.CamPolicyDocument
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta) GetCamPolicyType() *string {
	return s.CamPolicyType
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta) GetIamPolicyDocument() *string {
	return s.IamPolicyDocument
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta) GetIamPolicyType() *string {
	return s.IamPolicyType
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta) GetRamPolicyDocument() *string {
	return s.RamPolicyDocument
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta) GetRamPolicyType() *string {
	return s.RamPolicyType
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta) SetCamPolicyDocument(v string) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta {
	s.CamPolicyDocument = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta) SetCamPolicyType(v string) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta {
	s.CamPolicyType = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta) SetIamPolicyDocument(v string) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta {
	s.IamPolicyDocument = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta) SetIamPolicyType(v string) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta {
	s.IamPolicyType = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta) SetRamPolicyDocument(v string) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta {
	s.RamPolicyDocument = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta) SetRamPolicyType(v string) *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta {
	s.RamPolicyType = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponseBodyCloudAccountResourcesCloudAccountResourceMeta) Validate() error {
	return dara.Validate(s)
}
