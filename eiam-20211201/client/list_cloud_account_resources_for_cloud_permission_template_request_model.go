// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAccountResourcesForCloudPermissionTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountId(v string) *ListCloudAccountResourcesForCloudPermissionTemplateRequest
	GetCloudAccountId() *string
	SetCloudAccountResourceEntityType(v string) *ListCloudAccountResourcesForCloudPermissionTemplateRequest
	GetCloudAccountResourceEntityType() *string
	SetCloudPermissionTemplateId(v string) *ListCloudAccountResourcesForCloudPermissionTemplateRequest
	GetCloudPermissionTemplateId() *string
	SetInstanceId(v string) *ListCloudAccountResourcesForCloudPermissionTemplateRequest
	GetInstanceId() *string
	SetMaxResults(v int64) *ListCloudAccountResourcesForCloudPermissionTemplateRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListCloudAccountResourcesForCloudPermissionTemplateRequest
	GetNextToken() *string
}

type ListCloudAccountResourcesForCloudPermissionTemplateRequest struct {
	// The ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789012xxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The resource entity type used for filtering. Valid values: ram_policy.
	//
	// example:
	//
	// ram_policy
	CloudAccountResourceEntityType *string `json:"CloudAccountResourceEntityType,omitempty" xml:"CloudAccountResourceEntityType,omitempty"`
	// The ID of the permission template.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpt_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
	CloudPermissionTemplateId *string `json:"CloudPermissionTemplateId,omitempty" xml:"CloudPermissionTemplateId,omitempty"`
	// The ID of the IDaaS EIAM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page in a paging query. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set the value to the NextToken value returned from the previous API call.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListCloudAccountResourcesForCloudPermissionTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountResourcesForCloudPermissionTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateRequest) GetCloudAccountResourceEntityType() *string {
	return s.CloudAccountResourceEntityType
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateRequest) GetCloudPermissionTemplateId() *string {
	return s.CloudPermissionTemplateId
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateRequest) SetCloudAccountId(v string) *ListCloudAccountResourcesForCloudPermissionTemplateRequest {
	s.CloudAccountId = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateRequest) SetCloudAccountResourceEntityType(v string) *ListCloudAccountResourcesForCloudPermissionTemplateRequest {
	s.CloudAccountResourceEntityType = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateRequest) SetCloudPermissionTemplateId(v string) *ListCloudAccountResourcesForCloudPermissionTemplateRequest {
	s.CloudPermissionTemplateId = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateRequest) SetInstanceId(v string) *ListCloudAccountResourcesForCloudPermissionTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateRequest) SetMaxResults(v int64) *ListCloudAccountResourcesForCloudPermissionTemplateRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateRequest) SetNextToken(v string) *ListCloudAccountResourcesForCloudPermissionTemplateRequest {
	s.NextToken = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateRequest) Validate() error {
	return dara.Validate(s)
}
