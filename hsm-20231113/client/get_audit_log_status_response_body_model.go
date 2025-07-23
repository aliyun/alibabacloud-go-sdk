// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditLogStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuditLogStatus(v string) *GetAuditLogStatusResponseBody
	GetAuditLogStatus() *string
	SetAuditOssBucket(v string) *GetAuditLogStatusResponseBody
	GetAuditOssBucket() *string
	SetGrantedServiceAccess(v bool) *GetAuditLogStatusResponseBody
	GetGrantedServiceAccess() *bool
	SetOssBuckets(v []*string) *GetAuditLogStatusResponseBody
	GetOssBuckets() []*string
	SetRegionId(v string) *GetAuditLogStatusResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetAuditLogStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAuditLogStatusResponseBody
	GetSuccess() *bool
}

type GetAuditLogStatusResponseBody struct {
	// Indicates whether the audit log feature is enabled. Valid values:
	//
	// 	- enable
	//
	// 	- disable
	//
	// example:
	//
	// enable
	AuditLogStatus *string `json:"AuditLogStatus,omitempty" xml:"AuditLogStatus,omitempty"`
	// The bucket that stores audit logs.
	//
	// example:
	//
	// bucket-test
	AuditOssBucket *string `json:"AuditOssBucket,omitempty" xml:"AuditOssBucket,omitempty"`
	// Indicates whether Cloud Hardware Security Module is authorized to deliver logs. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	GrantedServiceAccess *bool `json:"GrantedServiceAccess,omitempty" xml:"GrantedServiceAccess,omitempty"`
	// A list of buckets that can be used to store audit logs.
	OssBuckets []*string `json:"OssBuckets,omitempty" xml:"OssBuckets,omitempty" type:"Repeated"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAuditLogStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuditLogStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuditLogStatusResponseBody) GetAuditLogStatus() *string {
	return s.AuditLogStatus
}

func (s *GetAuditLogStatusResponseBody) GetAuditOssBucket() *string {
	return s.AuditOssBucket
}

func (s *GetAuditLogStatusResponseBody) GetGrantedServiceAccess() *bool {
	return s.GrantedServiceAccess
}

func (s *GetAuditLogStatusResponseBody) GetOssBuckets() []*string {
	return s.OssBuckets
}

func (s *GetAuditLogStatusResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAuditLogStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuditLogStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAuditLogStatusResponseBody) SetAuditLogStatus(v string) *GetAuditLogStatusResponseBody {
	s.AuditLogStatus = &v
	return s
}

func (s *GetAuditLogStatusResponseBody) SetAuditOssBucket(v string) *GetAuditLogStatusResponseBody {
	s.AuditOssBucket = &v
	return s
}

func (s *GetAuditLogStatusResponseBody) SetGrantedServiceAccess(v bool) *GetAuditLogStatusResponseBody {
	s.GrantedServiceAccess = &v
	return s
}

func (s *GetAuditLogStatusResponseBody) SetOssBuckets(v []*string) *GetAuditLogStatusResponseBody {
	s.OssBuckets = v
	return s
}

func (s *GetAuditLogStatusResponseBody) SetRegionId(v string) *GetAuditLogStatusResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetAuditLogStatusResponseBody) SetRequestId(v string) *GetAuditLogStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuditLogStatusResponseBody) SetSuccess(v bool) *GetAuditLogStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetAuditLogStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
