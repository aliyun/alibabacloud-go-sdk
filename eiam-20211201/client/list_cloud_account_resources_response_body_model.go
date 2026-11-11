// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAccountResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountResources(v []*ListCloudAccountResourcesResponseBodyCloudAccountResources) *ListCloudAccountResourcesResponseBody
	GetCloudAccountResources() []*ListCloudAccountResourcesResponseBodyCloudAccountResources
	SetMaxResults(v int64) *ListCloudAccountResourcesResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListCloudAccountResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCloudAccountResourcesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListCloudAccountResourcesResponseBody
	GetTotalCount() *int64
}

type ListCloudAccountResourcesResponseBody struct {
	// The list of cloud account resources.
	CloudAccountResources []*ListCloudAccountResourcesResponseBodyCloudAccountResources `json:"CloudAccountResources,omitempty" xml:"CloudAccountResources,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page for paging in a paged query.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token returned by the current call, which is used for the next paginated query.
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

func (s ListCloudAccountResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudAccountResourcesResponseBody) GetCloudAccountResources() []*ListCloudAccountResourcesResponseBodyCloudAccountResources {
	return s.CloudAccountResources
}

func (s *ListCloudAccountResourcesResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListCloudAccountResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCloudAccountResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudAccountResourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCloudAccountResourcesResponseBody) SetCloudAccountResources(v []*ListCloudAccountResourcesResponseBodyCloudAccountResources) *ListCloudAccountResourcesResponseBody {
	s.CloudAccountResources = v
	return s
}

func (s *ListCloudAccountResourcesResponseBody) SetMaxResults(v int64) *ListCloudAccountResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBody) SetNextToken(v string) *ListCloudAccountResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBody) SetRequestId(v string) *ListCloudAccountResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBody) SetTotalCount(v int64) *ListCloudAccountResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBody) Validate() error {
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

type ListCloudAccountResourcesResponseBodyCloudAccountResources struct {
	// The ID of the cloud account to which the resource belongs.
	//
	// example:
	//
	// 123456789012xxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The identifier of the resource entity, which is a unique identifier on the remote end.
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
	// The metadata of the resource.
	CloudAccountResourceMeta *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta `json:"CloudAccountResourceMeta,omitempty" xml:"CloudAccountResourceMeta,omitempty" type:"Struct"`
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
	// The time when the resource was last updated. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649916625000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListCloudAccountResourcesResponseBodyCloudAccountResources) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountResourcesResponseBodyCloudAccountResources) GoString() string {
	return s.String()
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) GetCloudAccountResourceEntityId() *string {
	return s.CloudAccountResourceEntityId
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) GetCloudAccountResourceEntityType() *string {
	return s.CloudAccountResourceEntityType
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) GetCloudAccountResourceId() *string {
	return s.CloudAccountResourceId
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) GetCloudAccountResourceMeta() *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta {
	return s.CloudAccountResourceMeta
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) GetCloudAccountResourceName() *string {
	return s.CloudAccountResourceName
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) GetDescription() *string {
	return s.Description
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) SetCloudAccountId(v string) *ListCloudAccountResourcesResponseBodyCloudAccountResources {
	s.CloudAccountId = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) SetCloudAccountResourceEntityId(v string) *ListCloudAccountResourcesResponseBodyCloudAccountResources {
	s.CloudAccountResourceEntityId = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) SetCloudAccountResourceEntityType(v string) *ListCloudAccountResourcesResponseBodyCloudAccountResources {
	s.CloudAccountResourceEntityType = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) SetCloudAccountResourceId(v string) *ListCloudAccountResourcesResponseBodyCloudAccountResources {
	s.CloudAccountResourceId = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) SetCloudAccountResourceMeta(v *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta) *ListCloudAccountResourcesResponseBodyCloudAccountResources {
	s.CloudAccountResourceMeta = v
	return s
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) SetCloudAccountResourceName(v string) *ListCloudAccountResourcesResponseBodyCloudAccountResources {
	s.CloudAccountResourceName = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) SetCreateTime(v int64) *ListCloudAccountResourcesResponseBodyCloudAccountResources {
	s.CreateTime = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) SetDescription(v string) *ListCloudAccountResourcesResponseBodyCloudAccountResources {
	s.Description = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) SetUpdateTime(v int64) *ListCloudAccountResourcesResponseBodyCloudAccountResources {
	s.UpdateTime = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResources) Validate() error {
	if s.CloudAccountResourceMeta != nil {
		if err := s.CloudAccountResourceMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta struct {
	// The JSON document of the Tencent Cloud CAM policy. This parameter is returned only for Tencent Cloud accounts.
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
	// The JSON document of the AWS IAM policy. This parameter is returned only for AWS cloud accounts.
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
	// The JSON document of the RAM policy. This parameter is returned only by the Get operation.
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

func (s ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta) GoString() string {
	return s.String()
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta) GetCamPolicyDocument() *string {
	return s.CamPolicyDocument
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta) GetCamPolicyType() *string {
	return s.CamPolicyType
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta) GetIamPolicyDocument() *string {
	return s.IamPolicyDocument
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta) GetIamPolicyType() *string {
	return s.IamPolicyType
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta) GetRamPolicyDocument() *string {
	return s.RamPolicyDocument
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta) GetRamPolicyType() *string {
	return s.RamPolicyType
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta) SetCamPolicyDocument(v string) *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta {
	s.CamPolicyDocument = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta) SetCamPolicyType(v string) *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta {
	s.CamPolicyType = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta) SetIamPolicyDocument(v string) *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta {
	s.IamPolicyDocument = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta) SetIamPolicyType(v string) *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta {
	s.IamPolicyType = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta) SetRamPolicyDocument(v string) *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta {
	s.RamPolicyDocument = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta) SetRamPolicyType(v string) *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta {
	s.RamPolicyType = &v
	return s
}

func (s *ListCloudAccountResourcesResponseBodyCloudAccountResourcesCloudAccountResourceMeta) Validate() error {
	return dara.Validate(s)
}
