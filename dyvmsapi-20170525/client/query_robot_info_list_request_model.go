// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotInfoListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditStatus(v string) *QueryRobotInfoListRequest
	GetAuditStatus() *string
	SetOwnerId(v int64) *QueryRobotInfoListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryRobotInfoListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryRobotInfoListRequest
	GetResourceOwnerId() *int64
}

type QueryRobotInfoListRequest struct {
	// The review state. Valid values:
	//
	// 	- **CONFIGURABLE**
	//
	// 	- **AUDITING**
	//
	// 	- **AUDITPASS**
	//
	// 	- **AUDITFAIL**
	//
	// example:
	//
	// AUDITING
	AuditStatus          *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryRobotInfoListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotInfoListRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotInfoListRequest) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *QueryRobotInfoListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryRobotInfoListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryRobotInfoListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryRobotInfoListRequest) SetAuditStatus(v string) *QueryRobotInfoListRequest {
	s.AuditStatus = &v
	return s
}

func (s *QueryRobotInfoListRequest) SetOwnerId(v int64) *QueryRobotInfoListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRobotInfoListRequest) SetResourceOwnerAccount(v string) *QueryRobotInfoListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRobotInfoListRequest) SetResourceOwnerId(v int64) *QueryRobotInfoListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryRobotInfoListRequest) Validate() error {
	return dara.Validate(s)
}
