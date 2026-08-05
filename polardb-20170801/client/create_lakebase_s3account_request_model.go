// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLakebaseS3AccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPfsInstanceId(v string) *CreateLakebaseS3AccountRequest
	GetPfsInstanceId() *string
	SetRegionId(v string) *CreateLakebaseS3AccountRequest
	GetRegionId() *string
	SetUserAccAk(v string) *CreateLakebaseS3AccountRequest
	GetUserAccAk() *string
	SetUserAccPolicy(v string) *CreateLakebaseS3AccountRequest
	GetUserAccPolicy() *string
	SetUserAccSk(v string) *CreateLakebaseS3AccountRequest
	GetUserAccSk() *string
}

type CreateLakebaseS3AccountRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Access Key of the S3 account.
	//
	// > The account name can contain only uppercase letters, lowercase letters, and digits, and cannot exceed 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// accname
	UserAccAk *string `json:"UserAccAk,omitempty" xml:"UserAccAk,omitempty"`
	// A policy document in JSON format that defines the permissions of the S3 account. If this parameter is not specified, the default policy is used.
	//
	// example:
	//
	// {"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":["s3:*"],"Resource":["*"]}]}
	UserAccPolicy *string `json:"UserAccPolicy,omitempty" xml:"UserAccPolicy,omitempty"`
	// The Secret Key of the S3 account (@sensitive, encryption in transit).
	//
	// > The key must contain uppercase letters, lowercase letters, and digits, and must be greater than 18 and no more than 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// password***
	UserAccSk *string `json:"UserAccSk,omitempty" xml:"UserAccSk,omitempty"`
}

func (s CreateLakebaseS3AccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLakebaseS3AccountRequest) GoString() string {
	return s.String()
}

func (s *CreateLakebaseS3AccountRequest) GetPfsInstanceId() *string {
	return s.PfsInstanceId
}

func (s *CreateLakebaseS3AccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLakebaseS3AccountRequest) GetUserAccAk() *string {
	return s.UserAccAk
}

func (s *CreateLakebaseS3AccountRequest) GetUserAccPolicy() *string {
	return s.UserAccPolicy
}

func (s *CreateLakebaseS3AccountRequest) GetUserAccSk() *string {
	return s.UserAccSk
}

func (s *CreateLakebaseS3AccountRequest) SetPfsInstanceId(v string) *CreateLakebaseS3AccountRequest {
	s.PfsInstanceId = &v
	return s
}

func (s *CreateLakebaseS3AccountRequest) SetRegionId(v string) *CreateLakebaseS3AccountRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLakebaseS3AccountRequest) SetUserAccAk(v string) *CreateLakebaseS3AccountRequest {
	s.UserAccAk = &v
	return s
}

func (s *CreateLakebaseS3AccountRequest) SetUserAccPolicy(v string) *CreateLakebaseS3AccountRequest {
	s.UserAccPolicy = &v
	return s
}

func (s *CreateLakebaseS3AccountRequest) SetUserAccSk(v string) *CreateLakebaseS3AccountRequest {
	s.UserAccSk = &v
	return s
}

func (s *CreateLakebaseS3AccountRequest) Validate() error {
	return dara.Validate(s)
}
