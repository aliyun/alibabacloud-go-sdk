// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLangfuseUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteLangfuseUserRequest
	GetDBInstanceId() *string
	SetEmail(v string) *DeleteLangfuseUserRequest
	GetEmail() *string
	SetRegionId(v string) *DeleteLangfuseUserRequest
	GetRegionId() *string
}

type DeleteLangfuseUserRequest struct {
	// The Langfuse instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lfs-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The email address of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// john@company.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteLangfuseUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseUserRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteLangfuseUserRequest) GetEmail() *string {
	return s.Email
}

func (s *DeleteLangfuseUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLangfuseUserRequest) SetDBInstanceId(v string) *DeleteLangfuseUserRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteLangfuseUserRequest) SetEmail(v string) *DeleteLangfuseUserRequest {
	s.Email = &v
	return s
}

func (s *DeleteLangfuseUserRequest) SetRegionId(v string) *DeleteLangfuseUserRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLangfuseUserRequest) Validate() error {
	return dara.Validate(s)
}
