// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserSlsLogStorageTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyUserSlsLogStorageTimeRequest
	GetInstanceId() *string
	SetStorageTime(v int32) *ModifyUserSlsLogStorageTimeRequest
	GetStorageTime() *int32
}

type ModifyUserSlsLogStorageTimeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// vipcloudfw-cn-uqm3fz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 180
	StorageTime *int32 `json:"StorageTime,omitempty" xml:"StorageTime,omitempty"`
}

func (s ModifyUserSlsLogStorageTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserSlsLogStorageTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserSlsLogStorageTimeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyUserSlsLogStorageTimeRequest) GetStorageTime() *int32 {
	return s.StorageTime
}

func (s *ModifyUserSlsLogStorageTimeRequest) SetInstanceId(v string) *ModifyUserSlsLogStorageTimeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUserSlsLogStorageTimeRequest) SetStorageTime(v int32) *ModifyUserSlsLogStorageTimeRequest {
	s.StorageTime = &v
	return s
}

func (s *ModifyUserSlsLogStorageTimeRequest) Validate() error {
	return dara.Validate(s)
}
