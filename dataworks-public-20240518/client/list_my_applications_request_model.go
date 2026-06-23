// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefSchema(v string) *ListMyApplicationsRequest
	GetDefSchema() *string
	SetEndTime(v int64) *ListMyApplicationsRequest
	GetEndTime() *int64
	SetNextToken(v string) *ListMyApplicationsRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListMyApplicationsRequest
	GetPageSize() *int32
	SetResource(v *ListMyApplicationsRequestResource) *ListMyApplicationsRequest
	GetResource() *ListMyApplicationsRequestResource
	SetResourceType(v []*string) *ListMyApplicationsRequest
	GetResourceType() []*string
	SetStartTime(v int64) *ListMyApplicationsRequest
	GetStartTime() *int64
	SetStatuses(v []*string) *ListMyApplicationsRequest
	GetStatuses() []*string
}

type ListMyApplicationsRequest struct {
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// The end time of the application, specified as a Unix timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1779724799999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// A token that you can use in a subsequent request to retrieve the next page of results.
	//
	// example:
	//
	// eyJpZCI6MTIzfQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to return on each page. Default value: 10. Maximum value: 200.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search criteria for the resource.
	Resource *ListMyApplicationsRequestResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// The name of the leaf node that specifies the resource type. You can specify multiple resource types. Note that different leaf node names can map to the same business logic.
	//
	// This parameter is required.
	ResourceType []*string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" type:"Repeated"`
	// The start time of the application, specified as a Unix timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1771948800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The approval statuses for filtering. Valid values:
	//
	// - `WaitApproval`: Pending approval
	//
	// - `Confirmed`: Pending authorization
	//
	// - `RejectApproval`: Approval rejected
	//
	// - `AuthorizeSucceed`: Authorization succeeded
	//
	// - `AuthorizeFailed`: Authorization failed
	//
	// - `Deleted`: The application was deleted.
	//
	// - `Canceled`: The application was canceled.
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s ListMyApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMyApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListMyApplicationsRequest) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListMyApplicationsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListMyApplicationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMyApplicationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMyApplicationsRequest) GetResource() *ListMyApplicationsRequestResource {
	return s.Resource
}

func (s *ListMyApplicationsRequest) GetResourceType() []*string {
	return s.ResourceType
}

func (s *ListMyApplicationsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListMyApplicationsRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListMyApplicationsRequest) SetDefSchema(v string) *ListMyApplicationsRequest {
	s.DefSchema = &v
	return s
}

func (s *ListMyApplicationsRequest) SetEndTime(v int64) *ListMyApplicationsRequest {
	s.EndTime = &v
	return s
}

func (s *ListMyApplicationsRequest) SetNextToken(v string) *ListMyApplicationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMyApplicationsRequest) SetPageSize(v int32) *ListMyApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMyApplicationsRequest) SetResource(v *ListMyApplicationsRequestResource) *ListMyApplicationsRequest {
	s.Resource = v
	return s
}

func (s *ListMyApplicationsRequest) SetResourceType(v []*string) *ListMyApplicationsRequest {
	s.ResourceType = v
	return s
}

func (s *ListMyApplicationsRequest) SetStartTime(v int64) *ListMyApplicationsRequest {
	s.StartTime = &v
	return s
}

func (s *ListMyApplicationsRequest) SetStatuses(v []*string) *ListMyApplicationsRequest {
	s.Statuses = v
	return s
}

func (s *ListMyApplicationsRequest) Validate() error {
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMyApplicationsRequestResource struct {
	// The name of the resource schema (`ResourceSchema.name`) required for resource parsing.
	//
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// The version of the resource schema (`ResourceSchema.version`) required for resource parsing.
	//
	// example:
	//
	// v1.0.0
	DefVersion *string `json:"DefVersion,omitempty" xml:"DefVersion,omitempty"`
	// The resource metadata. The content is constrained by the `ResourceSchema`.
	MetaData map[string]interface{} `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
}

func (s ListMyApplicationsRequestResource) String() string {
	return dara.Prettify(s)
}

func (s ListMyApplicationsRequestResource) GoString() string {
	return s.String()
}

func (s *ListMyApplicationsRequestResource) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListMyApplicationsRequestResource) GetDefVersion() *string {
	return s.DefVersion
}

func (s *ListMyApplicationsRequestResource) GetMetaData() map[string]interface{} {
	return s.MetaData
}

func (s *ListMyApplicationsRequestResource) SetDefSchema(v string) *ListMyApplicationsRequestResource {
	s.DefSchema = &v
	return s
}

func (s *ListMyApplicationsRequestResource) SetDefVersion(v string) *ListMyApplicationsRequestResource {
	s.DefVersion = &v
	return s
}

func (s *ListMyApplicationsRequestResource) SetMetaData(v map[string]interface{}) *ListMyApplicationsRequestResource {
	s.MetaData = v
	return s
}

func (s *ListMyApplicationsRequestResource) Validate() error {
	return dara.Validate(s)
}
