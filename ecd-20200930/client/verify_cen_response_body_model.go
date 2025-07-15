// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlocks(v []*string) *VerifyCenResponseBody
	GetCidrBlocks() []*string
	SetRequestId(v string) *VerifyCenResponseBody
	GetRequestId() *string
	SetRouteEntries(v []*VerifyCenResponseBodyRouteEntries) *VerifyCenResponseBody
	GetRouteEntries() []*VerifyCenResponseBodyRouteEntries
	SetStatus(v string) *VerifyCenResponseBody
	GetStatus() *string
}

type VerifyCenResponseBody struct {
	// The recommended IPv4 CIDR blocks. Three CIDR blocks are randomly recommended. This parameter is returned when the `Status` value is `Conflict`.
	CidrBlocks []*string `json:"CidrBlocks,omitempty" xml:"CidrBlocks,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0AF9D857-AE96-53D6-B317-5DD665EC4EC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The routes provided by the CEN instance.
	RouteEntries []*VerifyCenResponseBodyRouteEntries `json:"RouteEntries,omitempty" xml:"RouteEntries,omitempty" type:"Repeated"`
	// The check result of CIDR block conflict.
	//
	// Valid values:
	//
	// 	- InvalidCen.CenUidInvalid: The Alibaba Cloud account is invalid or the Alibaba Cloud account does not have the permission to access Elastic Desktop Service.
	//
	// 	- VerifyCode.InvalidTokenCode: The verification code is invalid.
	//
	// 	- VerifyCode.ReachTokenRetryTime: The maximum number of times for entering a verification code reaches the limit.
	//
	// 	- Conflict: A CIDR block conflict exists. If the verification result of at least one route is Conflict, Conflict is returned for this parameter.
	//
	// 	- Access: The verification is passed. If the verification result for all routes is Access, Access is returned for this parameter.
	//
	// 	- InvalidCen.ParameterCenInstanceId: The Alibaba Cloud account does not own the CEN instance.
	//
	// example:
	//
	// Access
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s VerifyCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyCenResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyCenResponseBody) GetCidrBlocks() []*string {
	return s.CidrBlocks
}

func (s *VerifyCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyCenResponseBody) GetRouteEntries() []*VerifyCenResponseBodyRouteEntries {
	return s.RouteEntries
}

func (s *VerifyCenResponseBody) GetStatus() *string {
	return s.Status
}

func (s *VerifyCenResponseBody) SetCidrBlocks(v []*string) *VerifyCenResponseBody {
	s.CidrBlocks = v
	return s
}

func (s *VerifyCenResponseBody) SetRequestId(v string) *VerifyCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyCenResponseBody) SetRouteEntries(v []*VerifyCenResponseBodyRouteEntries) *VerifyCenResponseBody {
	s.RouteEntries = v
	return s
}

func (s *VerifyCenResponseBody) SetStatus(v string) *VerifyCenResponseBody {
	s.Status = &v
	return s
}

func (s *VerifyCenResponseBody) Validate() error {
	return dara.Validate(s)
}

type VerifyCenResponseBodyRouteEntries struct {
	// The CIDR block of the route.
	//
	// example:
	//
	// 172.16.111.3****
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ID of the instance corresponding to the route.
	//
	// example:
	//
	// vpc-uf62bu0xs5j8md54p****
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
	// The region ID of the route.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The verification result of the route.
	//
	// Valid values:
	//
	// 	- Conflict: A CIDR block conflict exists.
	//
	// 	- Access: The verification is passed.
	//
	// example:
	//
	// Access
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s VerifyCenResponseBodyRouteEntries) String() string {
	return dara.Prettify(s)
}

func (s VerifyCenResponseBodyRouteEntries) GoString() string {
	return s.String()
}

func (s *VerifyCenResponseBodyRouteEntries) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *VerifyCenResponseBodyRouteEntries) GetNextHopInstanceId() *string {
	return s.NextHopInstanceId
}

func (s *VerifyCenResponseBodyRouteEntries) GetRegionId() *string {
	return s.RegionId
}

func (s *VerifyCenResponseBodyRouteEntries) GetStatus() *string {
	return s.Status
}

func (s *VerifyCenResponseBodyRouteEntries) SetDestinationCidrBlock(v string) *VerifyCenResponseBodyRouteEntries {
	s.DestinationCidrBlock = &v
	return s
}

func (s *VerifyCenResponseBodyRouteEntries) SetNextHopInstanceId(v string) *VerifyCenResponseBodyRouteEntries {
	s.NextHopInstanceId = &v
	return s
}

func (s *VerifyCenResponseBodyRouteEntries) SetRegionId(v string) *VerifyCenResponseBodyRouteEntries {
	s.RegionId = &v
	return s
}

func (s *VerifyCenResponseBodyRouteEntries) SetStatus(v string) *VerifyCenResponseBodyRouteEntries {
	s.Status = &v
	return s
}

func (s *VerifyCenResponseBodyRouteEntries) Validate() error {
	return dara.Validate(s)
}
