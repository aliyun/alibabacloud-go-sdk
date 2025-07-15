// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceNewConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInstanceNewConnectionsRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeInstanceNewConnectionsRequest
	GetInstanceId() *string
	SetSbcName(v string) *DescribeInstanceNewConnectionsRequest
	GetSbcName() *string
	SetSecurityToken(v string) *DescribeInstanceNewConnectionsRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeInstanceNewConnectionsRequest
	GetStartTime() *string
}

type DescribeInstanceNewConnectionsRequest struct {
	// The end time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-10-08T02:08:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// apigateway-cn-2r426lavr001
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
	// 2022-10-01T02:08:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInstanceNewConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceNewConnectionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceNewConnectionsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstanceNewConnectionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceNewConnectionsRequest) GetSbcName() *string {
	return s.SbcName
}

func (s *DescribeInstanceNewConnectionsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstanceNewConnectionsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstanceNewConnectionsRequest) SetEndTime(v string) *DescribeInstanceNewConnectionsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceNewConnectionsRequest) SetInstanceId(v string) *DescribeInstanceNewConnectionsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceNewConnectionsRequest) SetSbcName(v string) *DescribeInstanceNewConnectionsRequest {
	s.SbcName = &v
	return s
}

func (s *DescribeInstanceNewConnectionsRequest) SetSecurityToken(v string) *DescribeInstanceNewConnectionsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceNewConnectionsRequest) SetStartTime(v string) *DescribeInstanceNewConnectionsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceNewConnectionsRequest) Validate() error {
	return dara.Validate(s)
}
