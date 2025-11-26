// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponentInstances(v []*ListComponentInstancesResponseBodyComponentInstances) *ListComponentInstancesResponseBody
	GetComponentInstances() []*ListComponentInstancesResponseBodyComponentInstances
	SetMaxResults(v int32) *ListComponentInstancesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListComponentInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListComponentInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListComponentInstancesResponseBody
	GetTotalCount() *int32
}

type ListComponentInstancesResponseBody struct {
	// The list of instance component installation information.
	ComponentInstances []*ListComponentInstancesResponseBodyComponentInstances `json:"ComponentInstances,omitempty" xml:"ComponentInstances,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 2
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// “”
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7345241A-014C-17D2-A3AC-C72771188F46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListComponentInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComponentInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListComponentInstancesResponseBody) GetComponentInstances() []*ListComponentInstancesResponseBodyComponentInstances {
	return s.ComponentInstances
}

func (s *ListComponentInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListComponentInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListComponentInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComponentInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListComponentInstancesResponseBody) SetComponentInstances(v []*ListComponentInstancesResponseBodyComponentInstances) *ListComponentInstancesResponseBody {
	s.ComponentInstances = v
	return s
}

func (s *ListComponentInstancesResponseBody) SetMaxResults(v int32) *ListComponentInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListComponentInstancesResponseBody) SetNextToken(v string) *ListComponentInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListComponentInstancesResponseBody) SetRequestId(v string) *ListComponentInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComponentInstancesResponseBody) SetTotalCount(v int32) *ListComponentInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListComponentInstancesResponseBody) Validate() error {
	if s.ComponentInstances != nil {
		for _, item := range s.ComponentInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListComponentInstancesResponseBodyComponentInstances struct {
	// The application name.
	//
	// example:
	//
	// KNOX
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The status of the component service. Valid values:
	//
	// 	- active: the primary service.
	//
	// 	- standby: the standby service.
	//
	// example:
	//
	// active
	BizState *string `json:"BizState,omitempty" xml:"BizState,omitempty"`
	// The status of the Commission. Valid values:
	//
	// 	- COMMISSIONED
	//
	// 	- COMMISSIONING
	//
	// 	- DECOMMISSIONED
	//
	// 	- DECOMMISSIONINPROGRESS
	//
	// 	- DECOMMISSIONFAILED
	//
	// 	- INSERVICE
	//
	// 	- UNKNOWN
	//
	// example:
	//
	// INSERVICE
	CommissionState *string `json:"CommissionState,omitempty" xml:"CommissionState,omitempty"`
	// The status of the component. Valid values:
	//
	// 	- WAITING
	//
	// 	- INSTALLING
	//
	// 	- INSTALLED
	//
	// 	- INSTALL_FAILED
	//
	// 	- STARTING
	//
	// 	- STARTED
	//
	// 	- START_FAILED
	//
	// 	- STOPPING
	//
	// 	- STOPPED
	//
	// 	- STOP_FAILED
	//
	// example:
	//
	// STARTED
	ComponentInstanceState *string `json:"ComponentInstanceState,omitempty" xml:"ComponentInstanceState,omitempty"`
	// The component name.
	//
	// example:
	//
	// KNOX
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The timestamp of the installation.
	//
	// example:
	//
	// 1628248947000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Valid values:
	//
	// 	- WAITING
	//
	// 	- INSTALLING
	//
	// 	- INSTALLED
	//
	// 	- INSTALL_FAILED
	//
	// 	- STARTING
	//
	// 	- STARTED
	//
	// 	- START_FAILED
	//
	// 	- STOPPING
	//
	// 	- STOPPED
	//
	// 	- STOP_FAILED
	//
	// example:
	//
	// STARTED
	DesiredState *string `json:"DesiredState,omitempty" xml:"DesiredState,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp17yy050pxo01m2****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// emr-worker-1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListComponentInstancesResponseBodyComponentInstances) String() string {
	return dara.Prettify(s)
}

func (s ListComponentInstancesResponseBodyComponentInstances) GoString() string {
	return s.String()
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetBizState() *string {
	return s.BizState
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetCommissionState() *string {
	return s.CommissionState
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetComponentInstanceState() *string {
	return s.ComponentInstanceState
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetComponentName() *string {
	return s.ComponentName
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetDesiredState() *string {
	return s.DesiredState
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetNodeId() *string {
	return s.NodeId
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetNodeName() *string {
	return s.NodeName
}

func (s *ListComponentInstancesResponseBodyComponentInstances) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetApplicationName(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.ApplicationName = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetBizState(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.BizState = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetCommissionState(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.CommissionState = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetComponentInstanceState(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.ComponentInstanceState = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetComponentName(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.ComponentName = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetCreateTime(v int64) *ListComponentInstancesResponseBodyComponentInstances {
	s.CreateTime = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetDesiredState(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.DesiredState = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetNodeId(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.NodeId = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetNodeName(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.NodeName = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) SetZoneId(v string) *ListComponentInstancesResponseBodyComponentInstances {
	s.ZoneId = &v
	return s
}

func (s *ListComponentInstancesResponseBodyComponentInstances) Validate() error {
	return dara.Validate(s)
}
