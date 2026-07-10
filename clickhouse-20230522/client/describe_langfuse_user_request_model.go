// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeLangfuseUserRequest
	GetDBInstanceId() *string
	SetEmail(v string) *DescribeLangfuseUserRequest
	GetEmail() *string
	SetRegionId(v string) *DescribeLangfuseUserRequest
	GetRegionId() *string
}

type DescribeLangfuseUserRequest struct {
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

func (s DescribeLangfuseUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseUserRequest) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseUserRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeLangfuseUserRequest) GetEmail() *string {
	return s.Email
}

func (s *DescribeLangfuseUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLangfuseUserRequest) SetDBInstanceId(v string) *DescribeLangfuseUserRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLangfuseUserRequest) SetEmail(v string) *DescribeLangfuseUserRequest {
	s.Email = &v
	return s
}

func (s *DescribeLangfuseUserRequest) SetRegionId(v string) *DescribeLangfuseUserRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLangfuseUserRequest) Validate() error {
	return dara.Validate(s)
}
