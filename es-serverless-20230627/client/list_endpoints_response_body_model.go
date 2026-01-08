// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListEndpointsResponseBody
	GetRequestId() *string
	SetResult(v []*ListEndpointsResponseBodyResult) *ListEndpointsResponseBody
	GetResult() []*ListEndpointsResponseBodyResult
	SetTotalCount(v int32) *ListEndpointsResponseBody
	GetTotalCount() *int32
}

type ListEndpointsResponseBody struct {
	// example:
	//
	// D6030CE6-9FEB-5B2F-84AC-3ADE3CBA89E5
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListEndpointsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEndpointsResponseBody) GetResult() []*ListEndpointsResponseBodyResult {
	return s.Result
}

func (s *ListEndpointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEndpointsResponseBody) SetRequestId(v string) *ListEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEndpointsResponseBody) SetResult(v []*ListEndpointsResponseBodyResult) *ListEndpointsResponseBody {
	s.Result = v
	return s
}

func (s *ListEndpointsResponseBody) SetTotalCount(v int32) *ListEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEndpointsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEndpointsResponseBodyResult struct {
	// example:
	//
	// Pending
	ConnectionStatus *string `json:"connectionStatus,omitempty" xml:"connectionStatus,omitempty"`
	// example:
	//
	// 1701259721
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// ep-bp1i522d****a3.epsrv-bp1f****gei.cn-hangzhou.privatelink.aliyuncs.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// example:
	//
	// essep-2f46b743f60****
	EndpointId    *string                                         `json:"endpointId,omitempty" xml:"endpointId,omitempty"`
	EndpointZones []*ListEndpointsResponseBodyResultEndpointZones `json:"endpointZones,omitempty" xml:"endpointZones,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ep-bp1id41dd116e52e****
	ResourceId       *string   `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// Active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// VPC
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1701259721
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
	// example:
	//
	// vpc-uf6gykvwcirp886ef****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListEndpointsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponseBodyResult) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *ListEndpointsResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ListEndpointsResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *ListEndpointsResponseBodyResult) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListEndpointsResponseBodyResult) GetEndpointZones() []*ListEndpointsResponseBodyResultEndpointZones {
	return s.EndpointZones
}

func (s *ListEndpointsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListEndpointsResponseBodyResult) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListEndpointsResponseBodyResult) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *ListEndpointsResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListEndpointsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListEndpointsResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ListEndpointsResponseBodyResult) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEndpointsResponseBodyResult) SetConnectionStatus(v string) *ListEndpointsResponseBodyResult {
	s.ConnectionStatus = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetCreated(v int32) *ListEndpointsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetDomain(v string) *ListEndpointsResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetEndpointId(v string) *ListEndpointsResponseBodyResult {
	s.EndpointId = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetEndpointZones(v []*ListEndpointsResponseBodyResultEndpointZones) *ListEndpointsResponseBodyResult {
	s.EndpointZones = v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetName(v string) *ListEndpointsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetResourceId(v string) *ListEndpointsResponseBodyResult {
	s.ResourceId = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetSecurityGroupIds(v []*string) *ListEndpointsResponseBodyResult {
	s.SecurityGroupIds = v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetStatus(v string) *ListEndpointsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetType(v string) *ListEndpointsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetUpdated(v int32) *ListEndpointsResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetVpcId(v string) *ListEndpointsResponseBodyResult {
	s.VpcId = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) Validate() error {
	if s.EndpointZones != nil {
		for _, item := range s.EndpointZones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEndpointsResponseBodyResultEndpointZones struct {
	// example:
	//
	// vsw-bp194pz9iez****
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ListEndpointsResponseBodyResultEndpointZones) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointsResponseBodyResultEndpointZones) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponseBodyResultEndpointZones) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListEndpointsResponseBodyResultEndpointZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListEndpointsResponseBodyResultEndpointZones) SetVSwitchId(v string) *ListEndpointsResponseBodyResultEndpointZones {
	s.VSwitchId = &v
	return s
}

func (s *ListEndpointsResponseBodyResultEndpointZones) SetZoneId(v string) *ListEndpointsResponseBodyResultEndpointZones {
	s.ZoneId = &v
	return s
}

func (s *ListEndpointsResponseBodyResultEndpointZones) Validate() error {
	return dara.Validate(s)
}
