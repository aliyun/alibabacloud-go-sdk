// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDatabaseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RestartDatabaseInstanceRequest
	GetClientToken() *string
	SetDatabaseInstanceId(v string) *RestartDatabaseInstanceRequest
	GetDatabaseInstanceId() *string
	SetRegionId(v string) *RestartDatabaseInstanceRequest
	GetRegionId() *string
}

type RestartDatabaseInstanceRequest struct {
	// A client-generated token to ensure request idempotency. The token must be unique for each request. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the lightweight database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the lightweight database instance. Call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to get the latest list of supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RestartDatabaseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartDatabaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDatabaseInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestartDatabaseInstanceRequest) GetDatabaseInstanceId() *string {
	return s.DatabaseInstanceId
}

func (s *RestartDatabaseInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartDatabaseInstanceRequest) SetClientToken(v string) *RestartDatabaseInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartDatabaseInstanceRequest) SetDatabaseInstanceId(v string) *RestartDatabaseInstanceRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *RestartDatabaseInstanceRequest) SetRegionId(v string) *RestartDatabaseInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RestartDatabaseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
