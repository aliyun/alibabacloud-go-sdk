// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetInstanceRecordConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *SetInstanceRecordConfigRequest
	GetEnabled() *bool
	SetExpirationDays(v int32) *SetInstanceRecordConfigRequest
	GetExpirationDays() *int32
	SetInstanceId(v string) *SetInstanceRecordConfigRequest
	GetInstanceId() *string
	SetRecordStorageTarget(v string) *SetInstanceRecordConfigRequest
	GetRecordStorageTarget() *string
	SetRegionId(v string) *SetInstanceRecordConfigRequest
	GetRegionId() *string
}

type SetInstanceRecordConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 7
	ExpirationDays *int32 `json:"ExpirationDays,omitempty" xml:"ExpirationDays,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// acs:oss:cn-shanghai:123:workbench-record-123-1/record
	RecordStorageTarget *string `json:"RecordStorageTarget,omitempty" xml:"RecordStorageTarget,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetInstanceRecordConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetInstanceRecordConfigRequest) GoString() string {
	return s.String()
}

func (s *SetInstanceRecordConfigRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *SetInstanceRecordConfigRequest) GetExpirationDays() *int32 {
	return s.ExpirationDays
}

func (s *SetInstanceRecordConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetInstanceRecordConfigRequest) GetRecordStorageTarget() *string {
	return s.RecordStorageTarget
}

func (s *SetInstanceRecordConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetInstanceRecordConfigRequest) SetEnabled(v bool) *SetInstanceRecordConfigRequest {
	s.Enabled = &v
	return s
}

func (s *SetInstanceRecordConfigRequest) SetExpirationDays(v int32) *SetInstanceRecordConfigRequest {
	s.ExpirationDays = &v
	return s
}

func (s *SetInstanceRecordConfigRequest) SetInstanceId(v string) *SetInstanceRecordConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *SetInstanceRecordConfigRequest) SetRecordStorageTarget(v string) *SetInstanceRecordConfigRequest {
	s.RecordStorageTarget = &v
	return s
}

func (s *SetInstanceRecordConfigRequest) SetRegionId(v string) *SetInstanceRecordConfigRequest {
	s.RegionId = &v
	return s
}

func (s *SetInstanceRecordConfigRequest) Validate() error {
	return dara.Validate(s)
}
