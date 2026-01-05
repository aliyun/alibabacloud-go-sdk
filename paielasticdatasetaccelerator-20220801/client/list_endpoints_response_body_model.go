// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoints(v []*ListEndpointsResponseBodyEndpoints) *ListEndpointsResponseBody
	GetEndpoints() []*ListEndpointsResponseBodyEndpoints
	SetRequestId(v string) *ListEndpointsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEndpointsResponseBody
	GetTotalCount() *int32
}

type ListEndpointsResponseBody struct {
	Endpoints []*ListEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponseBody) GetEndpoints() []*ListEndpointsResponseBodyEndpoints {
	return s.Endpoints
}

func (s *ListEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEndpointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEndpointsResponseBody) SetEndpoints(v []*ListEndpointsResponseBodyEndpoints) *ListEndpointsResponseBody {
	s.Endpoints = v
	return s
}

func (s *ListEndpointsResponseBody) SetRequestId(v string) *ListEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEndpointsResponseBody) SetTotalCount(v int32) *ListEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEndpointsResponseBody) Validate() error {
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEndpointsResponseBodyEndpoints struct {
	// example:
	//
	// 2014-10-02T15:01:23Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2014-10-02T15:01:23Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 所属加速实例的ID。
	//
	// example:
	//
	// inst-ivrq92qhbyrg4jctih
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// endpoint-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1557702098194904
	OwnerId *string         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Status  *EndpointStatus `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 276065346797410278
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// end-ivrq92qhbyrg4jctih
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// example:
	//
	// vpc-j6co2fcdsl1q0gnuc3ym3
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-j6cmr00qjyrft6jo2mg7g
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s ListEndpointsResponseBodyEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointsResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponseBodyEndpoints) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListEndpointsResponseBodyEndpoints) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListEndpointsResponseBodyEndpoints) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEndpointsResponseBodyEndpoints) GetName() *string {
	return s.Name
}

func (s *ListEndpointsResponseBodyEndpoints) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListEndpointsResponseBodyEndpoints) GetStatus() *EndpointStatus {
	return s.Status
}

func (s *ListEndpointsResponseBodyEndpoints) GetType() *string {
	return s.Type
}

func (s *ListEndpointsResponseBodyEndpoints) GetUserId() *string {
	return s.UserId
}

func (s *ListEndpointsResponseBodyEndpoints) GetUuid() *string {
	return s.Uuid
}

func (s *ListEndpointsResponseBodyEndpoints) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEndpointsResponseBodyEndpoints) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ListEndpointsResponseBodyEndpoints) SetGmtCreateTime(v string) *ListEndpointsResponseBodyEndpoints {
	s.GmtCreateTime = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetGmtModifiedTime(v string) *ListEndpointsResponseBodyEndpoints {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetInstanceId(v string) *ListEndpointsResponseBodyEndpoints {
	s.InstanceId = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetName(v string) *ListEndpointsResponseBodyEndpoints {
	s.Name = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetOwnerId(v string) *ListEndpointsResponseBodyEndpoints {
	s.OwnerId = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetStatus(v *EndpointStatus) *ListEndpointsResponseBodyEndpoints {
	s.Status = v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetType(v string) *ListEndpointsResponseBodyEndpoints {
	s.Type = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetUserId(v string) *ListEndpointsResponseBodyEndpoints {
	s.UserId = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetUuid(v string) *ListEndpointsResponseBodyEndpoints {
	s.Uuid = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetVpcId(v string) *ListEndpointsResponseBodyEndpoints {
	s.VpcId = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetVswitchId(v string) *ListEndpointsResponseBodyEndpoints {
	s.VswitchId = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) Validate() error {
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}
