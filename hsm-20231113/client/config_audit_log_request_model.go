// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigAuditLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditAction(v string) *ConfigAuditLogRequest
	GetAuditAction() *string
	SetAuditOssBucket(v string) *ConfigAuditLogRequest
	GetAuditOssBucket() *string
	SetRegionId(v string) *ConfigAuditLogRequest
	GetRegionId() *string
}

type ConfigAuditLogRequest struct {
	// Specifies whether to enable the audit log feature. Valid values:
	//
	// 	- enable
	//
	// 	- disable
	//
	// This parameter is required.
	//
	// example:
	//
	// enable
	AuditAction *string `json:"AuditAction,omitempty" xml:"AuditAction,omitempty"`
	// The bucket to which audit logs are delivered.
	//
	// example:
	//
	// hsm-log
	AuditOssBucket *string `json:"AuditOssBucket,omitempty" xml:"AuditOssBucket,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfigAuditLogRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigAuditLogRequest) GoString() string {
	return s.String()
}

func (s *ConfigAuditLogRequest) GetAuditAction() *string {
	return s.AuditAction
}

func (s *ConfigAuditLogRequest) GetAuditOssBucket() *string {
	return s.AuditOssBucket
}

func (s *ConfigAuditLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigAuditLogRequest) SetAuditAction(v string) *ConfigAuditLogRequest {
	s.AuditAction = &v
	return s
}

func (s *ConfigAuditLogRequest) SetAuditOssBucket(v string) *ConfigAuditLogRequest {
	s.AuditOssBucket = &v
	return s
}

func (s *ConfigAuditLogRequest) SetRegionId(v string) *ConfigAuditLogRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigAuditLogRequest) Validate() error {
	return dara.Validate(s)
}
