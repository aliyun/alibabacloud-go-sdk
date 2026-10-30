// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRouteMapsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRouteMapsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListRouteMapsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRouteMapsResponseBody
	GetRequestId() *string
	SetRouteMaps(v []*ListRouteMapsResponseBodyRouteMaps) *ListRouteMapsResponseBody
	GetRouteMaps() []*ListRouteMapsResponseBodyRouteMaps
	SetTotalCount(v int32) *ListRouteMapsResponseBody
	GetTotalCount() *int32
}

type ListRouteMapsResponseBody struct {
	// The number of entries returned on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If **NextToken*	- is empty, no next query is available.
	//
	// - If **NextToken*	- returns a value, the value is the token for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B2063B16-852B-5B66-B73D-4ED4D1A5E5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the route maps.
	RouteMaps []*ListRouteMapsResponseBodyRouteMaps `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRouteMapsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRouteMapsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRouteMapsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRouteMapsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRouteMapsResponseBody) GetRouteMaps() []*ListRouteMapsResponseBodyRouteMaps {
	return s.RouteMaps
}

func (s *ListRouteMapsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRouteMapsResponseBody) SetMaxResults(v int32) *ListRouteMapsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRouteMapsResponseBody) SetNextToken(v string) *ListRouteMapsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRouteMapsResponseBody) SetRequestId(v string) *ListRouteMapsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRouteMapsResponseBody) SetRouteMaps(v []*ListRouteMapsResponseBodyRouteMaps) *ListRouteMapsResponseBody {
	s.RouteMaps = v
	return s
}

func (s *ListRouteMapsResponseBody) SetTotalCount(v int32) *ListRouteMapsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRouteMapsResponseBody) Validate() error {
	if s.RouteMaps != nil {
		for _, item := range s.RouteMaps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRouteMapsResponseBodyRouteMaps struct {
	// The time when the route map was created.
	//
	// The time follows the ISO 8601 standard in the `YYYY-MM-DDThh:mmZ` format. The time is displayed in UTC.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-03-28T02:29:19Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the route map.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The direction of the route map. Valid values:
	//
	// - Inbound: inbound.
	//
	// - Outbound: outbound.
	//
	// example:
	//
	// Inbound
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The name of the route map.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority of the route map. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmydlzjd6dlhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the route map.
	//
	// example:
	//
	// rm-y40mxdvf7joc12****
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	// The status of the route map. Valid values:
	//
	// - **Creating**: The route map is being created.
	//
	// - **Updating**: The route map is being updated.
	//
	// - **Active**: The route map is available.
	//
	// - **Deleting**: The route map is being deleted.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tag []*ListRouteMapsResponseBodyRouteMapsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The version number of the route map configuration.
	//
	// The version number increases each time the route map and its associated route map rules are modified.
	//
	// example:
	//
	// 7
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListRouteMapsResponseBodyRouteMaps) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapsResponseBodyRouteMaps) GoString() string {
	return s.String()
}

func (s *ListRouteMapsResponseBodyRouteMaps) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListRouteMapsResponseBodyRouteMaps) GetDescription() *string {
	return s.Description
}

func (s *ListRouteMapsResponseBodyRouteMaps) GetDirection() *string {
	return s.Direction
}

func (s *ListRouteMapsResponseBodyRouteMaps) GetName() *string {
	return s.Name
}

func (s *ListRouteMapsResponseBodyRouteMaps) GetPriority() *int32 {
	return s.Priority
}

func (s *ListRouteMapsResponseBodyRouteMaps) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListRouteMapsResponseBodyRouteMaps) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *ListRouteMapsResponseBodyRouteMaps) GetStatus() *string {
	return s.Status
}

func (s *ListRouteMapsResponseBodyRouteMaps) GetTag() []*ListRouteMapsResponseBodyRouteMapsTag {
	return s.Tag
}

func (s *ListRouteMapsResponseBodyRouteMaps) GetVersion() *int32 {
	return s.Version
}

func (s *ListRouteMapsResponseBodyRouteMaps) SetCreationTime(v string) *ListRouteMapsResponseBodyRouteMaps {
	s.CreationTime = &v
	return s
}

func (s *ListRouteMapsResponseBodyRouteMaps) SetDescription(v string) *ListRouteMapsResponseBodyRouteMaps {
	s.Description = &v
	return s
}

func (s *ListRouteMapsResponseBodyRouteMaps) SetDirection(v string) *ListRouteMapsResponseBodyRouteMaps {
	s.Direction = &v
	return s
}

func (s *ListRouteMapsResponseBodyRouteMaps) SetName(v string) *ListRouteMapsResponseBodyRouteMaps {
	s.Name = &v
	return s
}

func (s *ListRouteMapsResponseBodyRouteMaps) SetPriority(v int32) *ListRouteMapsResponseBodyRouteMaps {
	s.Priority = &v
	return s
}

func (s *ListRouteMapsResponseBodyRouteMaps) SetResourceGroupId(v string) *ListRouteMapsResponseBodyRouteMaps {
	s.ResourceGroupId = &v
	return s
}

func (s *ListRouteMapsResponseBodyRouteMaps) SetRouteMapId(v string) *ListRouteMapsResponseBodyRouteMaps {
	s.RouteMapId = &v
	return s
}

func (s *ListRouteMapsResponseBodyRouteMaps) SetStatus(v string) *ListRouteMapsResponseBodyRouteMaps {
	s.Status = &v
	return s
}

func (s *ListRouteMapsResponseBodyRouteMaps) SetTag(v []*ListRouteMapsResponseBodyRouteMapsTag) *ListRouteMapsResponseBodyRouteMaps {
	s.Tag = v
	return s
}

func (s *ListRouteMapsResponseBodyRouteMaps) SetVersion(v int32) *ListRouteMapsResponseBodyRouteMaps {
	s.Version = &v
	return s
}

func (s *ListRouteMapsResponseBodyRouteMaps) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRouteMapsResponseBodyRouteMapsTag struct {
	// The tag key.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// jg
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListRouteMapsResponseBodyRouteMapsTag) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapsResponseBodyRouteMapsTag) GoString() string {
	return s.String()
}

func (s *ListRouteMapsResponseBodyRouteMapsTag) GetKey() *string {
	return s.Key
}

func (s *ListRouteMapsResponseBodyRouteMapsTag) GetValue() *string {
	return s.Value
}

func (s *ListRouteMapsResponseBodyRouteMapsTag) SetKey(v string) *ListRouteMapsResponseBodyRouteMapsTag {
	s.Key = &v
	return s
}

func (s *ListRouteMapsResponseBodyRouteMapsTag) SetValue(v string) *ListRouteMapsResponseBodyRouteMapsTag {
	s.Value = &v
	return s
}

func (s *ListRouteMapsResponseBodyRouteMapsTag) Validate() error {
	return dara.Validate(s)
}
