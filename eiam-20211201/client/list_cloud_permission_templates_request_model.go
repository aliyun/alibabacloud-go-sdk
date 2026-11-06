// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudPermissionTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountId(v string) *ListCloudPermissionTemplatesRequest
	GetCloudAccountId() *string
	SetCloudPermissionTemplateNamePrefix(v string) *ListCloudPermissionTemplatesRequest
	GetCloudPermissionTemplateNamePrefix() *string
	SetCloudPermissionTemplateSharingMode(v string) *ListCloudPermissionTemplatesRequest
	GetCloudPermissionTemplateSharingMode() *string
	SetCloudPermissionTemplateUsageType(v string) *ListCloudPermissionTemplatesRequest
	GetCloudPermissionTemplateUsageType() *string
	SetCreatorType(v string) *ListCloudPermissionTemplatesRequest
	GetCreatorType() *string
	SetInstanceId(v string) *ListCloudPermissionTemplatesRequest
	GetInstanceId() *string
	SetMaxResults(v int64) *ListCloudPermissionTemplatesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListCloudPermissionTemplatesRequest
	GetNextToken() *string
}

type ListCloudPermissionTemplatesRequest struct {
	// The ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// Filters the results by the prefix of the template name.
	//
	// example:
	//
	// AdminAccess
	CloudPermissionTemplateNamePrefix *string `json:"CloudPermissionTemplateNamePrefix,omitempty" xml:"CloudPermissionTemplateNamePrefix,omitempty"`
	// Filters the results by permission template sharing mode. Valid values: shared and exclusive. If you do not specify this parameter, no filtering is applied.
	//
	// example:
	//
	// shared
	CloudPermissionTemplateSharingMode *string `json:"CloudPermissionTemplateSharingMode,omitempty" xml:"CloudPermissionTemplateSharingMode,omitempty"`
	// Filters the results by permission template usage type. Valid values: user (privilege escalation for human users) and application (privilege escalation for machine-to-machine applications).
	//
	// example:
	//
	// user
	CloudPermissionTemplateUsageType *string `json:"CloudPermissionTemplateUsageType,omitempty" xml:"CloudPermissionTemplateUsageType,omitempty"`
	// Filters the results by creator type. Valid values: idaas_admin and idaas_user.
	//
	// example:
	//
	// idaas_user
	CreatorType *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	// The ID of the IDaaS EIAM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return per page in a paging query. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set the value to the NextToken value returned from the previous call.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListCloudPermissionTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudPermissionTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudPermissionTemplatesRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *ListCloudPermissionTemplatesRequest) GetCloudPermissionTemplateNamePrefix() *string {
	return s.CloudPermissionTemplateNamePrefix
}

func (s *ListCloudPermissionTemplatesRequest) GetCloudPermissionTemplateSharingMode() *string {
	return s.CloudPermissionTemplateSharingMode
}

func (s *ListCloudPermissionTemplatesRequest) GetCloudPermissionTemplateUsageType() *string {
	return s.CloudPermissionTemplateUsageType
}

func (s *ListCloudPermissionTemplatesRequest) GetCreatorType() *string {
	return s.CreatorType
}

func (s *ListCloudPermissionTemplatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloudPermissionTemplatesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListCloudPermissionTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCloudPermissionTemplatesRequest) SetCloudAccountId(v string) *ListCloudPermissionTemplatesRequest {
	s.CloudAccountId = &v
	return s
}

func (s *ListCloudPermissionTemplatesRequest) SetCloudPermissionTemplateNamePrefix(v string) *ListCloudPermissionTemplatesRequest {
	s.CloudPermissionTemplateNamePrefix = &v
	return s
}

func (s *ListCloudPermissionTemplatesRequest) SetCloudPermissionTemplateSharingMode(v string) *ListCloudPermissionTemplatesRequest {
	s.CloudPermissionTemplateSharingMode = &v
	return s
}

func (s *ListCloudPermissionTemplatesRequest) SetCloudPermissionTemplateUsageType(v string) *ListCloudPermissionTemplatesRequest {
	s.CloudPermissionTemplateUsageType = &v
	return s
}

func (s *ListCloudPermissionTemplatesRequest) SetCreatorType(v string) *ListCloudPermissionTemplatesRequest {
	s.CreatorType = &v
	return s
}

func (s *ListCloudPermissionTemplatesRequest) SetInstanceId(v string) *ListCloudPermissionTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCloudPermissionTemplatesRequest) SetMaxResults(v int64) *ListCloudPermissionTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCloudPermissionTemplatesRequest) SetNextToken(v string) *ListCloudPermissionTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListCloudPermissionTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
