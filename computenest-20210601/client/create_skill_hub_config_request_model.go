// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillHubConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSkillHubConfigRequest
	GetClientToken() *string
	SetOssBucketName(v string) *CreateSkillHubConfigRequest
	GetOssBucketName() *string
	SetOssRegionId(v string) *CreateSkillHubConfigRequest
	GetOssRegionId() *string
}

type CreateSkillHubConfigRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tidb-test-a
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	OssRegionId *string `json:"OssRegionId,omitempty" xml:"OssRegionId,omitempty"`
}

func (s CreateSkillHubConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillHubConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillHubConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSkillHubConfigRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *CreateSkillHubConfigRequest) GetOssRegionId() *string {
	return s.OssRegionId
}

func (s *CreateSkillHubConfigRequest) SetClientToken(v string) *CreateSkillHubConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSkillHubConfigRequest) SetOssBucketName(v string) *CreateSkillHubConfigRequest {
	s.OssBucketName = &v
	return s
}

func (s *CreateSkillHubConfigRequest) SetOssRegionId(v string) *CreateSkillHubConfigRequest {
	s.OssRegionId = &v
	return s
}

func (s *CreateSkillHubConfigRequest) Validate() error {
	return dara.Validate(s)
}
