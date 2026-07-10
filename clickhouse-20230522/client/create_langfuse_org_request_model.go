// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLangfuseOrgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateLangfuseOrgRequest
	GetDBInstanceId() *string
	SetName(v string) *CreateLangfuseOrgRequest
	GetName() *string
	SetOwnerEmail(v string) *CreateLangfuseOrgRequest
	GetOwnerEmail() *string
	SetRegionId(v string) *CreateLangfuseOrgRequest
	GetRegionId() *string
}

type CreateLangfuseOrgRequest struct {
	// The Langfuse instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lfs-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the Langfuse organization.
	//
	// This parameter is required.
	//
	// example:
	//
	// org_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The email address of the Langfuse organization owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// john@company.com
	OwnerEmail *string `json:"OwnerEmail,omitempty" xml:"OwnerEmail,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateLangfuseOrgRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseOrgRequest) GoString() string {
	return s.String()
}

func (s *CreateLangfuseOrgRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateLangfuseOrgRequest) GetName() *string {
	return s.Name
}

func (s *CreateLangfuseOrgRequest) GetOwnerEmail() *string {
	return s.OwnerEmail
}

func (s *CreateLangfuseOrgRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLangfuseOrgRequest) SetDBInstanceId(v string) *CreateLangfuseOrgRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateLangfuseOrgRequest) SetName(v string) *CreateLangfuseOrgRequest {
	s.Name = &v
	return s
}

func (s *CreateLangfuseOrgRequest) SetOwnerEmail(v string) *CreateLangfuseOrgRequest {
	s.OwnerEmail = &v
	return s
}

func (s *CreateLangfuseOrgRequest) SetRegionId(v string) *CreateLangfuseOrgRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLangfuseOrgRequest) Validate() error {
	return dara.Validate(s)
}
