// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteBackupDataRequest
	GetRegionId() *string
}

type DeleteBackupDataRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBackupDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupDataRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBackupDataRequest) SetRegionId(v string) *DeleteBackupDataRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBackupDataRequest) Validate() error {
	return dara.Validate(s)
}
