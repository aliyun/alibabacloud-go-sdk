// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKibanaPvlNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListKibanaPvlNetworkResponseBody
	GetRequestId() *string
	SetResult(v []*ListKibanaPvlNetworkResponseBodyResult) *ListKibanaPvlNetworkResponseBody
	GetResult() []*ListKibanaPvlNetworkResponseBodyResult
}

type ListKibanaPvlNetworkResponseBody struct {
	// request id
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListKibanaPvlNetworkResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListKibanaPvlNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKibanaPvlNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *ListKibanaPvlNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKibanaPvlNetworkResponseBody) GetResult() []*ListKibanaPvlNetworkResponseBodyResult {
	return s.Result
}

func (s *ListKibanaPvlNetworkResponseBody) SetRequestId(v string) *ListKibanaPvlNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKibanaPvlNetworkResponseBody) SetResult(v []*ListKibanaPvlNetworkResponseBodyResult) *ListKibanaPvlNetworkResponseBody {
	s.Result = v
	return s
}

func (s *ListKibanaPvlNetworkResponseBody) Validate() error {
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

type ListKibanaPvlNetworkResponseBodyResult struct {
	// example:
	//
	// 2024-03-07T06:26:28Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// ep-bp1tah7zbrwmkjef****
	EndpointId *string `json:"endpointId,omitempty" xml:"endpointId,omitempty"`
	// example:
	//
	// es-cn-xxdjfia****-kibana
	EndpointName *string `json:"endpointName,omitempty" xml:"endpointName,omitempty"`
	// example:
	//
	// Connected
	EndpointStatus *string `json:"endpointStatus,omitempty" xml:"endpointStatus,omitempty"`
	// example:
	//
	// es-cn-27a3mul6l***-kibana-internal
	PvlId          *string                                                 `json:"pvlId,omitempty" xml:"pvlId,omitempty"`
	SecurityGroups []*string                                               `json:"securityGroups,omitempty" xml:"securityGroups,omitempty" type:"Repeated"`
	VSwitchIdsZone []*ListKibanaPvlNetworkResponseBodyResultVSwitchIdsZone `json:"vSwitchIdsZone,omitempty" xml:"vSwitchIdsZone,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-bp16k1dvzxtma*****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListKibanaPvlNetworkResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListKibanaPvlNetworkResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListKibanaPvlNetworkResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListKibanaPvlNetworkResponseBodyResult) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListKibanaPvlNetworkResponseBodyResult) GetEndpointName() *string {
	return s.EndpointName
}

func (s *ListKibanaPvlNetworkResponseBodyResult) GetEndpointStatus() *string {
	return s.EndpointStatus
}

func (s *ListKibanaPvlNetworkResponseBodyResult) GetPvlId() *string {
	return s.PvlId
}

func (s *ListKibanaPvlNetworkResponseBodyResult) GetSecurityGroups() []*string {
	return s.SecurityGroups
}

func (s *ListKibanaPvlNetworkResponseBodyResult) GetVSwitchIdsZone() []*ListKibanaPvlNetworkResponseBodyResultVSwitchIdsZone {
	return s.VSwitchIdsZone
}

func (s *ListKibanaPvlNetworkResponseBodyResult) GetVpcId() *string {
	return s.VpcId
}

func (s *ListKibanaPvlNetworkResponseBodyResult) SetCreateTime(v string) *ListKibanaPvlNetworkResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListKibanaPvlNetworkResponseBodyResult) SetEndpointId(v string) *ListKibanaPvlNetworkResponseBodyResult {
	s.EndpointId = &v
	return s
}

func (s *ListKibanaPvlNetworkResponseBodyResult) SetEndpointName(v string) *ListKibanaPvlNetworkResponseBodyResult {
	s.EndpointName = &v
	return s
}

func (s *ListKibanaPvlNetworkResponseBodyResult) SetEndpointStatus(v string) *ListKibanaPvlNetworkResponseBodyResult {
	s.EndpointStatus = &v
	return s
}

func (s *ListKibanaPvlNetworkResponseBodyResult) SetPvlId(v string) *ListKibanaPvlNetworkResponseBodyResult {
	s.PvlId = &v
	return s
}

func (s *ListKibanaPvlNetworkResponseBodyResult) SetSecurityGroups(v []*string) *ListKibanaPvlNetworkResponseBodyResult {
	s.SecurityGroups = v
	return s
}

func (s *ListKibanaPvlNetworkResponseBodyResult) SetVSwitchIdsZone(v []*ListKibanaPvlNetworkResponseBodyResultVSwitchIdsZone) *ListKibanaPvlNetworkResponseBodyResult {
	s.VSwitchIdsZone = v
	return s
}

func (s *ListKibanaPvlNetworkResponseBodyResult) SetVpcId(v string) *ListKibanaPvlNetworkResponseBodyResult {
	s.VpcId = &v
	return s
}

func (s *ListKibanaPvlNetworkResponseBodyResult) Validate() error {
	if s.VSwitchIdsZone != nil {
		for _, item := range s.VSwitchIdsZone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKibanaPvlNetworkResponseBodyResultVSwitchIdsZone struct {
	// example:
	//
	// vsw-xdefafns***
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ListKibanaPvlNetworkResponseBodyResultVSwitchIdsZone) String() string {
	return dara.Prettify(s)
}

func (s ListKibanaPvlNetworkResponseBodyResultVSwitchIdsZone) GoString() string {
	return s.String()
}

func (s *ListKibanaPvlNetworkResponseBodyResultVSwitchIdsZone) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ListKibanaPvlNetworkResponseBodyResultVSwitchIdsZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListKibanaPvlNetworkResponseBodyResultVSwitchIdsZone) SetVswitchId(v string) *ListKibanaPvlNetworkResponseBodyResultVSwitchIdsZone {
	s.VswitchId = &v
	return s
}

func (s *ListKibanaPvlNetworkResponseBodyResultVSwitchIdsZone) SetZoneId(v string) *ListKibanaPvlNetworkResponseBodyResultVSwitchIdsZone {
	s.ZoneId = &v
	return s
}

func (s *ListKibanaPvlNetworkResponseBodyResultVSwitchIdsZone) Validate() error {
	return dara.Validate(s)
}
