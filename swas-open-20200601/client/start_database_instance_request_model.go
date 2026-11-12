// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDatabaseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartDatabaseInstanceRequest
	GetClientToken() *string
	SetDatabaseInstanceId(v string) *StartDatabaseInstanceRequest
	GetDatabaseInstanceId() *string
	SetRegionId(v string) *StartDatabaseInstanceRequest
	GetRegionId() *string
}

type StartDatabaseInstanceRequest struct {
	// A client-generated token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the database instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to obtain the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartDatabaseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDatabaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartDatabaseInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartDatabaseInstanceRequest) GetDatabaseInstanceId() *string {
	return s.DatabaseInstanceId
}

func (s *StartDatabaseInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartDatabaseInstanceRequest) SetClientToken(v string) *StartDatabaseInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDatabaseInstanceRequest) SetDatabaseInstanceId(v string) *StartDatabaseInstanceRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *StartDatabaseInstanceRequest) SetRegionId(v string) *StartDatabaseInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartDatabaseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
