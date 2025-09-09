// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSwitchAzoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeInstanceSwitchAzoneResponseBody
	GetRequestId() *string
	SetResult(v *DescribeInstanceSwitchAzoneResponseBodyResult) *DescribeInstanceSwitchAzoneResponseBody
	GetResult() *DescribeInstanceSwitchAzoneResponseBodyResult
	SetSuccess(v bool) *DescribeInstanceSwitchAzoneResponseBody
	GetSuccess() *bool
}

type DescribeInstanceSwitchAzoneResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DSSDF-SEWE-23ERW
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the operation.
	Result *DescribeInstanceSwitchAzoneResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstanceSwitchAzoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSwitchAzoneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchAzoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceSwitchAzoneResponseBody) GetResult() *DescribeInstanceSwitchAzoneResponseBodyResult {
	return s.Result
}

func (s *DescribeInstanceSwitchAzoneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceSwitchAzoneResponseBody) SetRequestId(v string) *DescribeInstanceSwitchAzoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBody) SetResult(v *DescribeInstanceSwitchAzoneResponseBodyResult) *DescribeInstanceSwitchAzoneResponseBody {
	s.Result = v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBody) SetSuccess(v bool) *DescribeInstanceSwitchAzoneResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceSwitchAzoneResponseBodyResult struct {
	// The ID of the source azoneId.
	//
	// example:
	//
	// cn-hangzhou-a
	OriginAzoneId *string `json:"OriginAzoneId,omitempty" xml:"OriginAzoneId,omitempty"`
	// regionId.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Indicates whether the job can be switched.
	//
	// example:
	//
	// true
	SwitchAble *bool `json:"SwitchAble,omitempty" xml:"SwitchAble,omitempty"`
	// Target azones.
	TargetAzones *DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones `json:"TargetAzones,omitempty" xml:"TargetAzones,omitempty" type:"Struct"`
}

func (s DescribeInstanceSwitchAzoneResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSwitchAzoneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) GetOriginAzoneId() *string {
	return s.OriginAzoneId
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) GetSwitchAble() *bool {
	return s.SwitchAble
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) GetTargetAzones() *DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones {
	return s.TargetAzones
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) SetOriginAzoneId(v string) *DescribeInstanceSwitchAzoneResponseBodyResult {
	s.OriginAzoneId = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) SetRegionId(v string) *DescribeInstanceSwitchAzoneResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) SetSwitchAble(v bool) *DescribeInstanceSwitchAzoneResponseBodyResult {
	s.SwitchAble = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) SetTargetAzones(v *DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones) *DescribeInstanceSwitchAzoneResponseBodyResult {
	s.TargetAzones = v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones struct {
	TargetAzone []*string `json:"TargetAzone,omitempty" xml:"TargetAzone,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones) GetTargetAzone() []*string {
	return s.TargetAzone
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones) SetTargetAzone(v []*string) *DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones {
	s.TargetAzone = v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponseBodyResultTargetAzones) Validate() error {
	return dara.Validate(s)
}
