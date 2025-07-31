// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetBackupConfigRequest
	GetRegionId() *string
	SetResourceType(v string) *GetBackupConfigRequest
	GetResourceType() *string
	SetServiceCode(v string) *GetBackupConfigRequest
	GetServiceCode() *string
}

type GetBackupConfigRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s GetBackupConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBackupConfigRequest) GoString() string {
	return s.String()
}

func (s *GetBackupConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBackupConfigRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetBackupConfigRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetBackupConfigRequest) SetRegionId(v string) *GetBackupConfigRequest {
	s.RegionId = &v
	return s
}

func (s *GetBackupConfigRequest) SetResourceType(v string) *GetBackupConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *GetBackupConfigRequest) SetServiceCode(v string) *GetBackupConfigRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetBackupConfigRequest) Validate() error {
	return dara.Validate(s)
}
