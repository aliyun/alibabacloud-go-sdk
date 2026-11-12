// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ReleasePublicConnectionRequest
	GetClientToken() *string
	SetDatabaseInstanceId(v string) *ReleasePublicConnectionRequest
	GetDatabaseInstanceId() *string
	SetRegionId(v string) *ReleasePublicConnectionRequest
	GetRegionId() *string
}

type ReleasePublicConnectionRequest struct {
	// A client-generated token to ensure request idempotence. You must ensure that this token is unique for each request. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
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
	// The region ID of the lightweight database instance. You can call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to query for supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleasePublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleasePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleasePublicConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReleasePublicConnectionRequest) GetDatabaseInstanceId() *string {
	return s.DatabaseInstanceId
}

func (s *ReleasePublicConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleasePublicConnectionRequest) SetClientToken(v string) *ReleasePublicConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *ReleasePublicConnectionRequest) SetDatabaseInstanceId(v string) *ReleasePublicConnectionRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *ReleasePublicConnectionRequest) SetRegionId(v string) *ReleasePublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *ReleasePublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
