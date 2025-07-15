// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDropPacketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInstanceDropPacketRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeInstanceDropPacketRequest
	GetInstanceId() *string
	SetSbcName(v string) *DescribeInstanceDropPacketRequest
	GetSbcName() *string
	SetSecurityToken(v string) *DescribeInstanceDropPacketRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeInstanceDropPacketRequest
	GetStartTime() *string
}

type DescribeInstanceDropPacketRequest struct {
	// The end time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-12-16T02:04:36Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// apigateway-cn-v641b9dxc00p
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The statistical metric. Valid values:
	//
	// 	- Maximum
	//
	// 	- Minimum
	//
	// 	- Average
	//
	// This parameter is required.
	//
	// example:
	//
	// Maximum
	SbcName       *string `json:"SbcName,omitempty" xml:"SbcName,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The start time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-06T04:00:36Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInstanceDropPacketRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDropPacketRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDropPacketRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstanceDropPacketRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceDropPacketRequest) GetSbcName() *string {
	return s.SbcName
}

func (s *DescribeInstanceDropPacketRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstanceDropPacketRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstanceDropPacketRequest) SetEndTime(v string) *DescribeInstanceDropPacketRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceDropPacketRequest) SetInstanceId(v string) *DescribeInstanceDropPacketRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceDropPacketRequest) SetSbcName(v string) *DescribeInstanceDropPacketRequest {
	s.SbcName = &v
	return s
}

func (s *DescribeInstanceDropPacketRequest) SetSecurityToken(v string) *DescribeInstanceDropPacketRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceDropPacketRequest) SetStartTime(v string) *DescribeInstanceDropPacketRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceDropPacketRequest) Validate() error {
	return dara.Validate(s)
}
