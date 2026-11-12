// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseInstanceDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDatabaseInstanceDescriptionRequest
	GetClientToken() *string
	SetDatabaseInstanceDescription(v string) *ModifyDatabaseInstanceDescriptionRequest
	GetDatabaseInstanceDescription() *string
	SetDatabaseInstanceId(v string) *ModifyDatabaseInstanceDescriptionRequest
	GetDatabaseInstanceId() *string
	SetRegionId(v string) *ModifyDatabaseInstanceDescriptionRequest
	GetRegionId() *string
}

type ModifyDatabaseInstanceDescriptionRequest struct {
	// Ensures the idempotence of the request. You can generate the value from your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters and can be up to 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the lightweight database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// This is a lightweight database instance.
	DatabaseInstanceDescription *string `json:"DatabaseInstanceDescription,omitempty" xml:"DatabaseInstanceDescription,omitempty"`
	// The instance ID of the lightweight database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the lightweight database instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the list of supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDatabaseInstanceDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceDescriptionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDatabaseInstanceDescriptionRequest) GetDatabaseInstanceDescription() *string {
	return s.DatabaseInstanceDescription
}

func (s *ModifyDatabaseInstanceDescriptionRequest) GetDatabaseInstanceId() *string {
	return s.DatabaseInstanceId
}

func (s *ModifyDatabaseInstanceDescriptionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDatabaseInstanceDescriptionRequest) SetClientToken(v string) *ModifyDatabaseInstanceDescriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionRequest) SetDatabaseInstanceDescription(v string) *ModifyDatabaseInstanceDescriptionRequest {
	s.DatabaseInstanceDescription = &v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionRequest) SetDatabaseInstanceId(v string) *ModifyDatabaseInstanceDescriptionRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionRequest) SetRegionId(v string) *ModifyDatabaseInstanceDescriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
