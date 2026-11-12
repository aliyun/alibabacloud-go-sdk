// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDatabaseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StopDatabaseInstanceRequest
	GetClientToken() *string
	SetDatabaseInstanceId(v string) *StopDatabaseInstanceRequest
	GetDatabaseInstanceId() *string
	SetRegionId(v string) *StopDatabaseInstanceRequest
	GetRegionId() *string
}

type StopDatabaseInstanceRequest struct {
	// A client-generated token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
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
	// The region ID of the lightweight database instance. You can call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to view the supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopDatabaseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDatabaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopDatabaseInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StopDatabaseInstanceRequest) GetDatabaseInstanceId() *string {
	return s.DatabaseInstanceId
}

func (s *StopDatabaseInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopDatabaseInstanceRequest) SetClientToken(v string) *StopDatabaseInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StopDatabaseInstanceRequest) SetDatabaseInstanceId(v string) *StopDatabaseInstanceRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *StopDatabaseInstanceRequest) SetRegionId(v string) *StopDatabaseInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StopDatabaseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
