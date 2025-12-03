// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *OpenBackupRequest
	GetClusterId() *string
	SetRegionId(v string) *OpenBackupRequest
	GetRegionId() *string
	SetZoneId(v string) *OpenBackupRequest
	GetZoneId() *string
}

type OpenBackupRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s OpenBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenBackupRequest) GoString() string {
	return s.String()
}

func (s *OpenBackupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OpenBackupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenBackupRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *OpenBackupRequest) SetClusterId(v string) *OpenBackupRequest {
	s.ClusterId = &v
	return s
}

func (s *OpenBackupRequest) SetRegionId(v string) *OpenBackupRequest {
	s.RegionId = &v
	return s
}

func (s *OpenBackupRequest) SetZoneId(v string) *OpenBackupRequest {
	s.ZoneId = &v
	return s
}

func (s *OpenBackupRequest) Validate() error {
	return dara.Validate(s)
}
