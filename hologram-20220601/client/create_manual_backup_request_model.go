// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateManualBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateManualBackupRequest
	GetRegionId() *string
	SetInstanceId(v string) *CreateManualBackupRequest
	GetInstanceId() *string
}

type CreateManualBackupRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// hgpostcn-cn-721344a2z001
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s CreateManualBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateManualBackupRequest) GoString() string {
	return s.String()
}

func (s *CreateManualBackupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateManualBackupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateManualBackupRequest) SetRegionId(v string) *CreateManualBackupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateManualBackupRequest) SetInstanceId(v string) *CreateManualBackupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateManualBackupRequest) Validate() error {
	return dara.Validate(s)
}
