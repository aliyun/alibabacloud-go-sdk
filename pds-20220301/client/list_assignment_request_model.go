// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssignmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListAssignmentRequest
	GetLimit() *int32
	SetManageResourceId(v string) *ListAssignmentRequest
	GetManageResourceId() *string
	SetManageResourceType(v string) *ListAssignmentRequest
	GetManageResourceType() *string
	SetMarker(v string) *ListAssignmentRequest
	GetMarker() *string
}

type ListAssignmentRequest struct {
	// The maximum number of results to return. Valid values: 1 to 100.
	//
	// The number of returned results must be less than or equal to the specified number.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The ID of the managed resource, such as a group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 105***b82
	ManageResourceId *string `json:"manage_resource_id,omitempty" xml:"manage_resource_id,omitempty"`
	// The type of the managed resource. Set the value to RT_Group, which specifies that the administrators of a group are queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// RT_Group
	ManageResourceType *string `json:"manage_resource_type,omitempty" xml:"manage_resource_type,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListAssignmentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAssignmentRequest) GoString() string {
	return s.String()
}

func (s *ListAssignmentRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListAssignmentRequest) GetManageResourceId() *string {
	return s.ManageResourceId
}

func (s *ListAssignmentRequest) GetManageResourceType() *string {
	return s.ManageResourceType
}

func (s *ListAssignmentRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListAssignmentRequest) SetLimit(v int32) *ListAssignmentRequest {
	s.Limit = &v
	return s
}

func (s *ListAssignmentRequest) SetManageResourceId(v string) *ListAssignmentRequest {
	s.ManageResourceId = &v
	return s
}

func (s *ListAssignmentRequest) SetManageResourceType(v string) *ListAssignmentRequest {
	s.ManageResourceType = &v
	return s
}

func (s *ListAssignmentRequest) SetMarker(v string) *ListAssignmentRequest {
	s.Marker = &v
	return s
}

func (s *ListAssignmentRequest) Validate() error {
	return dara.Validate(s)
}
