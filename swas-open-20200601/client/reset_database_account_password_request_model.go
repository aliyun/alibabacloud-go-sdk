// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDatabaseAccountPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountPassword(v string) *ResetDatabaseAccountPasswordRequest
	GetAccountPassword() *string
	SetClientToken(v string) *ResetDatabaseAccountPasswordRequest
	GetClientToken() *string
	SetDatabaseInstanceId(v string) *ResetDatabaseAccountPasswordRequest
	GetDatabaseInstanceId() *string
	SetRegionId(v string) *ResetDatabaseAccountPasswordRequest
	GetRegionId() *string
}

type ResetDatabaseAccountPasswordRequest struct {
	// The password of the administrator account for the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// Password****
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, and you must ensure that the value is unique among different requests. The client token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID of the simple application server database.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the simple application server database instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetDatabaseAccountPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetDatabaseAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetDatabaseAccountPasswordRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *ResetDatabaseAccountPasswordRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ResetDatabaseAccountPasswordRequest) GetDatabaseInstanceId() *string {
	return s.DatabaseInstanceId
}

func (s *ResetDatabaseAccountPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetDatabaseAccountPasswordRequest) SetAccountPassword(v string) *ResetDatabaseAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetDatabaseAccountPasswordRequest) SetClientToken(v string) *ResetDatabaseAccountPasswordRequest {
	s.ClientToken = &v
	return s
}

func (s *ResetDatabaseAccountPasswordRequest) SetDatabaseInstanceId(v string) *ResetDatabaseAccountPasswordRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *ResetDatabaseAccountPasswordRequest) SetRegionId(v string) *ResetDatabaseAccountPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetDatabaseAccountPasswordRequest) Validate() error {
	return dara.Validate(s)
}
