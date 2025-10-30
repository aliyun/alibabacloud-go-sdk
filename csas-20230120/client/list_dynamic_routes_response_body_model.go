// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicRoutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicRoutes(v []*ListDynamicRoutesResponseBodyDynamicRoutes) *ListDynamicRoutesResponseBody
	GetDynamicRoutes() []*ListDynamicRoutesResponseBodyDynamicRoutes
	SetRequestId(v string) *ListDynamicRoutesResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *ListDynamicRoutesResponseBody
	GetTotalNum() *int32
}

type ListDynamicRoutesResponseBody struct {
	DynamicRoutes []*ListDynamicRoutesResponseBodyDynamicRoutes `json:"DynamicRoutes,omitempty" xml:"DynamicRoutes,omitempty" type:"Repeated"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListDynamicRoutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDynamicRoutesResponseBody) GetDynamicRoutes() []*ListDynamicRoutesResponseBodyDynamicRoutes {
	return s.DynamicRoutes
}

func (s *ListDynamicRoutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDynamicRoutesResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListDynamicRoutesResponseBody) SetDynamicRoutes(v []*ListDynamicRoutesResponseBodyDynamicRoutes) *ListDynamicRoutesResponseBody {
	s.DynamicRoutes = v
	return s
}

func (s *ListDynamicRoutesResponseBody) SetRequestId(v string) *ListDynamicRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDynamicRoutesResponseBody) SetTotalNum(v int32) *ListDynamicRoutesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListDynamicRoutesResponseBody) Validate() error {
	if s.DynamicRoutes != nil {
		for _, item := range s.DynamicRoutes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDynamicRoutesResponseBodyDynamicRoutes struct {
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// example:
	//
	// 2023-03-21 11:50:03
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// dr-a0ca843f53cf****
	DynamicRouteId *string `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
	// example:
	//
	// connector
	DynamicRouteType *string `json:"DynamicRouteType,omitempty" xml:"DynamicRouteType,omitempty"`
	// example:
	//
	// dynamic_route_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// connector-8ccb13b6f52c****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// example:
	//
	// 1
	Priority  *int32    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// example:
	//
	// Enabled
	Status *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s ListDynamicRoutesResponseBodyDynamicRoutes) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicRoutesResponseBodyDynamicRoutes) GoString() string {
	return s.String()
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) GetDescription() *string {
	return s.Description
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) GetDynamicRouteId() *string {
	return s.DynamicRouteId
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) GetDynamicRouteType() *string {
	return s.DynamicRouteType
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) GetName() *string {
	return s.Name
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) GetNextHop() *string {
	return s.NextHop
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) GetPriority() *int32 {
	return s.Priority
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) GetStatus() *string {
	return s.Status
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) GetTagIds() []*string {
	return s.TagIds
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetApplicationIds(v []*string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.ApplicationIds = v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetApplicationType(v string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.ApplicationType = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetCreateTime(v string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.CreateTime = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetDescription(v string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.Description = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetDynamicRouteId(v string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.DynamicRouteId = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetDynamicRouteType(v string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.DynamicRouteType = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetName(v string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.Name = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetNextHop(v string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.NextHop = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetPriority(v int32) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.Priority = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetRegionIds(v []*string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.RegionIds = v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetStatus(v string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.Status = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetTagIds(v []*string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.TagIds = v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) Validate() error {
	return dara.Validate(s)
}
