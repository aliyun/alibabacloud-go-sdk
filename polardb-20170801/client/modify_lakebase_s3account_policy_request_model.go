// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLakebaseS3AccountPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPfsInstanceId(v string) *ModifyLakebaseS3AccountPolicyRequest
	GetPfsInstanceId() *string
	SetRegionId(v string) *ModifyLakebaseS3AccountPolicyRequest
	GetRegionId() *string
	SetUserAccAk(v string) *ModifyLakebaseS3AccountPolicyRequest
	GetUserAccAk() *string
	SetUserAccPolicy(v string) *ModifyLakebaseS3AccountPolicyRequest
	GetUserAccPolicy() *string
}

type ModifyLakebaseS3AccountPolicyRequest struct {
	// The PolarFS instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-xxx
	PfsInstanceId *string `json:"PfsInstanceId,omitempty" xml:"PfsInstanceId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the regions of all clusters under the target account.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The access key of the S3 account whose policy you want to modify.
	//
	// > You cannot modify the default account.
	//
	// This parameter is required.
	//
	// example:
	//
	// accname
	UserAccAk *string `json:"UserAccAk,omitempty" xml:"UserAccAk,omitempty"`
	// The user account policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":["s3:GetObject","s3:PutObject"],"Resource":["arn:aws:s3:::my-bucket/*"]}]}
	UserAccPolicy *string `json:"UserAccPolicy,omitempty" xml:"UserAccPolicy,omitempty"`
}

func (s ModifyLakebaseS3AccountPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLakebaseS3AccountPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyLakebaseS3AccountPolicyRequest) GetPfsInstanceId() *string {
	return s.PfsInstanceId
}

func (s *ModifyLakebaseS3AccountPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLakebaseS3AccountPolicyRequest) GetUserAccAk() *string {
	return s.UserAccAk
}

func (s *ModifyLakebaseS3AccountPolicyRequest) GetUserAccPolicy() *string {
	return s.UserAccPolicy
}

func (s *ModifyLakebaseS3AccountPolicyRequest) SetPfsInstanceId(v string) *ModifyLakebaseS3AccountPolicyRequest {
	s.PfsInstanceId = &v
	return s
}

func (s *ModifyLakebaseS3AccountPolicyRequest) SetRegionId(v string) *ModifyLakebaseS3AccountPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLakebaseS3AccountPolicyRequest) SetUserAccAk(v string) *ModifyLakebaseS3AccountPolicyRequest {
	s.UserAccAk = &v
	return s
}

func (s *ModifyLakebaseS3AccountPolicyRequest) SetUserAccPolicy(v string) *ModifyLakebaseS3AccountPolicyRequest {
	s.UserAccPolicy = &v
	return s
}

func (s *ModifyLakebaseS3AccountPolicyRequest) Validate() error {
	return dara.Validate(s)
}
