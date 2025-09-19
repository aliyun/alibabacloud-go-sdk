// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceRecordConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceRecordConfigRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetInstanceRecordConfigRequest
	GetRegionId() *string
}

type GetInstanceRecordConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetInstanceRecordConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceRecordConfigRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRecordConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceRecordConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceRecordConfigRequest) SetInstanceId(v string) *GetInstanceRecordConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceRecordConfigRequest) SetRegionId(v string) *GetInstanceRecordConfigRequest {
	s.RegionId = &v
	return s
}

func (s *GetInstanceRecordConfigRequest) Validate() error {
	return dara.Validate(s)
}
