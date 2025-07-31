// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetBackupStatusRequest
	GetRegionId() *string
}

type GetBackupStatusRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBackupStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBackupStatusRequest) GoString() string {
	return s.String()
}

func (s *GetBackupStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBackupStatusRequest) SetRegionId(v string) *GetBackupStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetBackupStatusRequest) Validate() error {
	return dara.Validate(s)
}
