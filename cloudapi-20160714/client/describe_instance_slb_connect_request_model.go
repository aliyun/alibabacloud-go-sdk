// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSlbConnectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInstanceSlbConnectRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeInstanceSlbConnectRequest
	GetInstanceId() *string
	SetSbcName(v string) *DescribeInstanceSlbConnectRequest
	GetSbcName() *string
	SetSecurityToken(v string) *DescribeInstanceSlbConnectRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeInstanceSlbConnectRequest
	GetStartTime() *string
}

type DescribeInstanceSlbConnectRequest struct {
	// The end time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-15T15:07:06Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// apigateway-sz-1574cc7c5a31
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
	// 2022-09-01T02:09:33Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInstanceSlbConnectRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSlbConnectRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSlbConnectRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstanceSlbConnectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceSlbConnectRequest) GetSbcName() *string {
	return s.SbcName
}

func (s *DescribeInstanceSlbConnectRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstanceSlbConnectRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstanceSlbConnectRequest) SetEndTime(v string) *DescribeInstanceSlbConnectRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceSlbConnectRequest) SetInstanceId(v string) *DescribeInstanceSlbConnectRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSlbConnectRequest) SetSbcName(v string) *DescribeInstanceSlbConnectRequest {
	s.SbcName = &v
	return s
}

func (s *DescribeInstanceSlbConnectRequest) SetSecurityToken(v string) *DescribeInstanceSlbConnectRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceSlbConnectRequest) SetStartTime(v string) *DescribeInstanceSlbConnectRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceSlbConnectRequest) Validate() error {
	return dara.Validate(s)
}
