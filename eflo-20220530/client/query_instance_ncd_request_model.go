// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceNcdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId1(v string) *QueryInstanceNcdRequest
	GetInstanceId1() *string
	SetInstanceId2(v string) *QueryInstanceNcdRequest
	GetInstanceId2() *string
	SetInstanceType(v string) *QueryInstanceNcdRequest
	GetInstanceType() *string
	SetRegionId(v string) *QueryInstanceNcdRequest
	GetRegionId() *string
}

type QueryInstanceNcdRequest struct {
	// Instance 1ID
	//
	// This parameter is required.
	//
	// example:
	//
	// lni-1235****
	InstanceId1 *string `json:"InstanceId1,omitempty" xml:"InstanceId1,omitempty"`
	// Instance 2ID
	//
	// This parameter is required.
	//
	// example:
	//
	// lni-1234****
	InstanceId2 *string `json:"InstanceId2,omitempty" xml:"InstanceId2,omitempty"`
	// The parameter that specifies the instance type.
	//
	// Valid value:
	//
	// 	- node: Lingjun node.
	//
	// 	- lni: lingjun network interface controller.
	//
	// This parameter is required.
	//
	// example:
	//
	// lni
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryInstanceNcdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceNcdRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceNcdRequest) GetInstanceId1() *string {
	return s.InstanceId1
}

func (s *QueryInstanceNcdRequest) GetInstanceId2() *string {
	return s.InstanceId2
}

func (s *QueryInstanceNcdRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *QueryInstanceNcdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryInstanceNcdRequest) SetInstanceId1(v string) *QueryInstanceNcdRequest {
	s.InstanceId1 = &v
	return s
}

func (s *QueryInstanceNcdRequest) SetInstanceId2(v string) *QueryInstanceNcdRequest {
	s.InstanceId2 = &v
	return s
}

func (s *QueryInstanceNcdRequest) SetInstanceType(v string) *QueryInstanceNcdRequest {
	s.InstanceType = &v
	return s
}

func (s *QueryInstanceNcdRequest) SetRegionId(v string) *QueryInstanceNcdRequest {
	s.RegionId = &v
	return s
}

func (s *QueryInstanceNcdRequest) Validate() error {
	return dara.Validate(s)
}
