// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHoneyPotSuspStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeHoneyPotSuspStatisticsResponseBody
	GetRequestId() *string
	SetSuspHoneyPotStatisticsResponse(v []*DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) *DescribeHoneyPotSuspStatisticsResponseBody
	GetSuspHoneyPotStatisticsResponse() []*DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse
}

type DescribeHoneyPotSuspStatisticsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9E3969FA-5ACC-4256-9FDE-BB6918CD0410
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the top 5 VPCs or assets for which alerts are most frequently generated.
	SuspHoneyPotStatisticsResponse []*DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse `json:"SuspHoneyPotStatisticsResponse,omitempty" xml:"SuspHoneyPotStatisticsResponse,omitempty" type:"Repeated"`
}

func (s DescribeHoneyPotSuspStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHoneyPotSuspStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHoneyPotSuspStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHoneyPotSuspStatisticsResponseBody) GetSuspHoneyPotStatisticsResponse() []*DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse {
	return s.SuspHoneyPotStatisticsResponse
}

func (s *DescribeHoneyPotSuspStatisticsResponseBody) SetRequestId(v string) *DescribeHoneyPotSuspStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponseBody) SetSuspHoneyPotStatisticsResponse(v []*DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) *DescribeHoneyPotSuspStatisticsResponseBody {
	s.SuspHoneyPotStatisticsResponse = v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse struct {
	// The total number of alerts that are generated for the asset.
	//
	// example:
	//
	// 8793
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the server.
	//
	// > This parameter is returned only when **StatisticsKeyType*	- is set to **uuid**.
	//
	// example:
	//
	// i-p0whhoba24wd28p8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	//
	// > This parameter is returned only when **StatisticsKeyType*	- is set to **uuid**.
	//
	// example:
	//
	// abc-launch-advisor
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- **vpcInstanceId**: VPC
	//
	// 	- **uuid**: server
	//
	// example:
	//
	// vpcInstanceId
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the VPC.
	//
	// > This parameter is returned only when **StatisticsKeyType*	- is set to **vpcInstanceId**.
	//
	// example:
	//
	// vpc-p0wwdsuutdyu1ygkt****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// > This parameter is returned only when **StatisticsKeyType*	- is set to **vpcInstanceId**.
	//
	// example:
	//
	// abc01
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) GetCount() *int32 {
	return s.Count
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) GetType() *string {
	return s.Type
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) SetCount(v int32) *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse {
	s.Count = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) SetInstanceId(v string) *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse {
	s.InstanceId = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) SetInstanceName(v string) *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse {
	s.InstanceName = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) SetType(v string) *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse {
	s.Type = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) SetVpcId(v string) *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse {
	s.VpcId = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) SetVpcName(v string) *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse {
	s.VpcName = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) Validate() error {
	return dara.Validate(s)
}
