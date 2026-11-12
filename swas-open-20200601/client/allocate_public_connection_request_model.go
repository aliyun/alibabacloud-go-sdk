// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocatePublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AllocatePublicConnectionRequest
	GetClientToken() *string
	SetDatabaseInstanceId(v string) *AllocatePublicConnectionRequest
	GetDatabaseInstanceId() *string
	SetRegionId(v string) *AllocatePublicConnectionRequest
	GetRegionId() *string
}

type AllocatePublicConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The client token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID of the lightweight database.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the specified lightweight database instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/2361076.html) operation to query the list of supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AllocatePublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocatePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocatePublicConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AllocatePublicConnectionRequest) GetDatabaseInstanceId() *string {
	return s.DatabaseInstanceId
}

func (s *AllocatePublicConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocatePublicConnectionRequest) SetClientToken(v string) *AllocatePublicConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *AllocatePublicConnectionRequest) SetDatabaseInstanceId(v string) *AllocatePublicConnectionRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *AllocatePublicConnectionRequest) SetRegionId(v string) *AllocatePublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *AllocatePublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
