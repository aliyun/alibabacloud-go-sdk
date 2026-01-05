// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreateTime(v string) *DescribeEndpointResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *DescribeEndpointResponseBody
	GetGmtModifiedTime() *string
	SetName(v string) *DescribeEndpointResponseBody
	GetName() *string
	SetOwnerId(v string) *DescribeEndpointResponseBody
	GetOwnerId() *string
	SetRequestId(v string) *DescribeEndpointResponseBody
	GetRequestId() *string
	SetStatus(v *EndpointStatus) *DescribeEndpointResponseBody
	GetStatus() *EndpointStatus
	SetType(v string) *DescribeEndpointResponseBody
	GetType() *string
	SetUserId(v string) *DescribeEndpointResponseBody
	GetUserId() *string
	SetUuid(v string) *DescribeEndpointResponseBody
	GetUuid() *string
	SetVpcId(v string) *DescribeEndpointResponseBody
	GetVpcId() *string
	SetVswitchId(v string) *DescribeEndpointResponseBody
	GetVswitchId() *string
}

type DescribeEndpointResponseBody struct {
	// example:
	//
	// 2014-10-02T15:01:23Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2014-10-02T15:01:23Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// endpoint-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1557702098194904
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *EndpointStatus `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s DescribeEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndpointResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *DescribeEndpointResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *DescribeEndpointResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeEndpointResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEndpointResponseBody) GetStatus() *EndpointStatus {
	return s.Status
}

func (s *DescribeEndpointResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeEndpointResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *DescribeEndpointResponseBody) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeEndpointResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeEndpointResponseBody) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeEndpointResponseBody) SetGmtCreateTime(v string) *DescribeEndpointResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetGmtModifiedTime(v string) *DescribeEndpointResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetName(v string) *DescribeEndpointResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetOwnerId(v string) *DescribeEndpointResponseBody {
	s.OwnerId = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetRequestId(v string) *DescribeEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetStatus(v *EndpointStatus) *DescribeEndpointResponseBody {
	s.Status = v
	return s
}

func (s *DescribeEndpointResponseBody) SetType(v string) *DescribeEndpointResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetUserId(v string) *DescribeEndpointResponseBody {
	s.UserId = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetUuid(v string) *DescribeEndpointResponseBody {
	s.Uuid = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetVpcId(v string) *DescribeEndpointResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetVswitchId(v string) *DescribeEndpointResponseBody {
	s.VswitchId = &v
	return s
}

func (s *DescribeEndpointResponseBody) Validate() error {
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}
