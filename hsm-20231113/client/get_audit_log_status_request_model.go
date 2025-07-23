// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditLogStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGetOssBucket(v bool) *GetAuditLogStatusRequest
	GetGetOssBucket() *bool
	SetRegionId(v string) *GetAuditLogStatusRequest
	GetRegionId() *string
}

type GetAuditLogStatusRequest struct {
	// Specifies whether to obtain the list of OSS buckets that can be used to store audit logs. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	GetOssBucket *bool `json:"GetOssBucket,omitempty" xml:"GetOssBucket,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAuditLogStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuditLogStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAuditLogStatusRequest) GetGetOssBucket() *bool {
	return s.GetOssBucket
}

func (s *GetAuditLogStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAuditLogStatusRequest) SetGetOssBucket(v bool) *GetAuditLogStatusRequest {
	s.GetOssBucket = &v
	return s
}

func (s *GetAuditLogStatusRequest) SetRegionId(v string) *GetAuditLogStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetAuditLogStatusRequest) Validate() error {
	return dara.Validate(s)
}
