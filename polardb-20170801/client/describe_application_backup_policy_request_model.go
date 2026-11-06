// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationBackupPolicyRequest
	GetApplicationId() *string
	SetApplicationType(v string) *DescribeApplicationBackupPolicyRequest
	GetApplicationType() *string
	SetRegionId(v string) *DescribeApplicationBackupPolicyRequest
	GetRegionId() *string
}

type DescribeApplicationBackupPolicyRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type, such as polarclaw.
	//
	// example:
	//
	// polarclaw
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The region ID of the application.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeApplicationBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationBackupPolicyRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationBackupPolicyRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *DescribeApplicationBackupPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationBackupPolicyRequest) SetApplicationId(v string) *DescribeApplicationBackupPolicyRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationBackupPolicyRequest) SetApplicationType(v string) *DescribeApplicationBackupPolicyRequest {
	s.ApplicationType = &v
	return s
}

func (s *DescribeApplicationBackupPolicyRequest) SetRegionId(v string) *DescribeApplicationBackupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
