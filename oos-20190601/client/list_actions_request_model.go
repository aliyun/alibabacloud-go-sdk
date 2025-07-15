// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListActionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListActionsRequest
	GetNextToken() *string
	SetOOSActionName(v string) *ListActionsRequest
	GetOOSActionName() *string
	SetRegionId(v string) *ListActionsRequest
	GetRegionId() *string
}

type ListActionsRequest struct {
	// The number of entries to return on each page. Valid values: 20 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the action. All actions whose names contain the specified action name are returned.
	//
	// example:
	//
	// MyTemplate
	OOSActionName *string `json:"OOSActionName,omitempty" xml:"OOSActionName,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListActionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListActionsRequest) GoString() string {
	return s.String()
}

func (s *ListActionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListActionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListActionsRequest) GetOOSActionName() *string {
	return s.OOSActionName
}

func (s *ListActionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListActionsRequest) SetMaxResults(v int32) *ListActionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListActionsRequest) SetNextToken(v string) *ListActionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListActionsRequest) SetOOSActionName(v string) *ListActionsRequest {
	s.OOSActionName = &v
	return s
}

func (s *ListActionsRequest) SetRegionId(v string) *ListActionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListActionsRequest) Validate() error {
	return dara.Validate(s)
}
