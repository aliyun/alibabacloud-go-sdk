// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancePacketsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInstancePacketsRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeInstancePacketsRequest
	GetInstanceId() *string
	SetSbcName(v string) *DescribeInstancePacketsRequest
	GetSbcName() *string
	SetSecurityToken(v string) *DescribeInstancePacketsRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeInstancePacketsRequest
	GetStartTime() *string
}

type DescribeInstancePacketsRequest struct {
	// The end time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-05-24T10:14:53Z
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
	// 2022-05-18T01:14:26Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInstancePacketsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancePacketsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancePacketsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstancePacketsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancePacketsRequest) GetSbcName() *string {
	return s.SbcName
}

func (s *DescribeInstancePacketsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstancePacketsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstancePacketsRequest) SetEndTime(v string) *DescribeInstancePacketsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstancePacketsRequest) SetInstanceId(v string) *DescribeInstancePacketsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancePacketsRequest) SetSbcName(v string) *DescribeInstancePacketsRequest {
	s.SbcName = &v
	return s
}

func (s *DescribeInstancePacketsRequest) SetSecurityToken(v string) *DescribeInstancePacketsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstancePacketsRequest) SetStartTime(v string) *DescribeInstancePacketsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstancePacketsRequest) Validate() error {
	return dara.Validate(s)
}
