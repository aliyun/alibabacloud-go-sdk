// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDropConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInstanceDropConnectionsRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeInstanceDropConnectionsRequest
	GetInstanceId() *string
	SetSbcName(v string) *DescribeInstanceDropConnectionsRequest
	GetSbcName() *string
	SetSecurityToken(v string) *DescribeInstanceDropConnectionsRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeInstanceDropConnectionsRequest
	GetStartTime() *string
}

type DescribeInstanceDropConnectionsRequest struct {
	// The end time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-31T07:00:09Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the dedicated instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// apigateway-cn-n6w1v1234501
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
	// 2023-01-31T06:00:09Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInstanceDropConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDropConnectionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDropConnectionsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstanceDropConnectionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceDropConnectionsRequest) GetSbcName() *string {
	return s.SbcName
}

func (s *DescribeInstanceDropConnectionsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstanceDropConnectionsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstanceDropConnectionsRequest) SetEndTime(v string) *DescribeInstanceDropConnectionsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceDropConnectionsRequest) SetInstanceId(v string) *DescribeInstanceDropConnectionsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceDropConnectionsRequest) SetSbcName(v string) *DescribeInstanceDropConnectionsRequest {
	s.SbcName = &v
	return s
}

func (s *DescribeInstanceDropConnectionsRequest) SetSecurityToken(v string) *DescribeInstanceDropConnectionsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceDropConnectionsRequest) SetStartTime(v string) *DescribeInstanceDropConnectionsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceDropConnectionsRequest) Validate() error {
	return dara.Validate(s)
}
