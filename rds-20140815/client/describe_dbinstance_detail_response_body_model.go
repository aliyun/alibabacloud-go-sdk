// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActivationState(v string) *DescribeDBInstanceDetailResponseBody
	GetActivationState() *string
	SetDBInstanceId(v string) *DescribeDBInstanceDetailResponseBody
	GetDBInstanceId() *string
	SetLicenseType(v string) *DescribeDBInstanceDetailResponseBody
	GetLicenseType() *string
	SetRegionId(v string) *DescribeDBInstanceDetailResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeDBInstanceDetailResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceDetailResponseBody struct {
	// Indicates whether the instance is in the active state.
	//
	// example:
	//
	// Invalid
	ActivationState *string `json:"ActivationState,omitempty" xml:"ActivationState,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// rm-bp6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The type of the license.
	//
	// example:
	//
	// Normal
	LicenseType *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06B220E2-EAC5-4DBE-A1FC-1B62DB6A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDetailResponseBody) GetActivationState() *string {
	return s.ActivationState
}

func (s *DescribeDBInstanceDetailResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceDetailResponseBody) GetLicenseType() *string {
	return s.LicenseType
}

func (s *DescribeDBInstanceDetailResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceDetailResponseBody) SetActivationState(v string) *DescribeDBInstanceDetailResponseBody {
	s.ActivationState = &v
	return s
}

func (s *DescribeDBInstanceDetailResponseBody) SetDBInstanceId(v string) *DescribeDBInstanceDetailResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceDetailResponseBody) SetLicenseType(v string) *DescribeDBInstanceDetailResponseBody {
	s.LicenseType = &v
	return s
}

func (s *DescribeDBInstanceDetailResponseBody) SetRegionId(v string) *DescribeDBInstanceDetailResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceDetailResponseBody) SetRequestId(v string) *DescribeDBInstanceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
