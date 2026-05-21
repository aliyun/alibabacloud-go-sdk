// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupDataDescRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UpdateBackupDataDescRequest
	GetRegionId() *string
	SetDesc(v string) *UpdateBackupDataDescRequest
	GetDesc() *string
	SetInstanceId(v string) *UpdateBackupDataDescRequest
	GetInstanceId() *string
}

type UpdateBackupDataDescRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// back up test
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// hgpostcn-cn-721344a2z001
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s UpdateBackupDataDescRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupDataDescRequest) GoString() string {
	return s.String()
}

func (s *UpdateBackupDataDescRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateBackupDataDescRequest) GetDesc() *string {
	return s.Desc
}

func (s *UpdateBackupDataDescRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateBackupDataDescRequest) SetRegionId(v string) *UpdateBackupDataDescRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBackupDataDescRequest) SetDesc(v string) *UpdateBackupDataDescRequest {
	s.Desc = &v
	return s
}

func (s *UpdateBackupDataDescRequest) SetInstanceId(v string) *UpdateBackupDataDescRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateBackupDataDescRequest) Validate() error {
	return dara.Validate(s)
}
