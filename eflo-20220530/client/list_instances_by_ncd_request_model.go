// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesByNcdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListInstancesByNcdRequest
	GetInstanceId() *string
	SetInstanceType(v string) *ListInstancesByNcdRequest
	GetInstanceType() *string
	SetMaxNcd(v int32) *ListInstancesByNcdRequest
	GetMaxNcd() *int32
	SetRegionId(v string) *ListInstancesByNcdRequest
	GetRegionId() *string
}

type ListInstancesByNcdRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lni-1234****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// Maximum network communication distance
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	MaxNcd *int32 `json:"MaxNcd,omitempty" xml:"MaxNcd,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancesByNcdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesByNcdRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesByNcdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesByNcdRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListInstancesByNcdRequest) GetMaxNcd() *int32 {
	return s.MaxNcd
}

func (s *ListInstancesByNcdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesByNcdRequest) SetInstanceId(v string) *ListInstancesByNcdRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesByNcdRequest) SetInstanceType(v string) *ListInstancesByNcdRequest {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesByNcdRequest) SetMaxNcd(v int32) *ListInstancesByNcdRequest {
	s.MaxNcd = &v
	return s
}

func (s *ListInstancesByNcdRequest) SetRegionId(v string) *ListInstancesByNcdRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancesByNcdRequest) Validate() error {
	return dara.Validate(s)
}
