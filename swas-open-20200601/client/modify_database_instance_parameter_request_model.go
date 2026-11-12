// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseInstanceParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDatabaseInstanceParameterRequest
	GetClientToken() *string
	SetDatabaseInstanceId(v string) *ModifyDatabaseInstanceParameterRequest
	GetDatabaseInstanceId() *string
	SetForceRestart(v bool) *ModifyDatabaseInstanceParameterRequest
	GetForceRestart() *bool
	SetParameters(v string) *ModifyDatabaseInstanceParameterRequest
	GetParameters() *string
	SetRegionId(v string) *ModifyDatabaseInstanceParameterRequest
	GetRegionId() *string
}

type ModifyDatabaseInstanceParameterRequest struct {
	// A client-generated token to ensure request idempotence. You must ensure that the token is unique for each request. It can contain only ASCII characters and cannot be longer than 64 characters. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
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
	// Specifies whether to restart the database for the parameter changes to take effect. Valid values:
	//
	// - true: Forces a restart. Set this to true if any modified parameter requires a restart to take effect. Otherwise, the change will not take effect.
	//
	// - false: Do not force a restart.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceRestart *bool `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	// A JSON string that specifies the database parameters and their values. The value for each parameter must be a string. Format: `{"ParameterName1":"Value1","ParameterName2":"Value2",...}`.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"delayed_insert_timeout":"600","max_length_for_sort_data":"2048"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID of the database instance. Call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to view the supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDatabaseInstanceParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseInstanceParameterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceParameterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDatabaseInstanceParameterRequest) GetDatabaseInstanceId() *string {
	return s.DatabaseInstanceId
}

func (s *ModifyDatabaseInstanceParameterRequest) GetForceRestart() *bool {
	return s.ForceRestart
}

func (s *ModifyDatabaseInstanceParameterRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyDatabaseInstanceParameterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDatabaseInstanceParameterRequest) SetClientToken(v string) *ModifyDatabaseInstanceParameterRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterRequest) SetDatabaseInstanceId(v string) *ModifyDatabaseInstanceParameterRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterRequest) SetForceRestart(v bool) *ModifyDatabaseInstanceParameterRequest {
	s.ForceRestart = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterRequest) SetParameters(v string) *ModifyDatabaseInstanceParameterRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterRequest) SetRegionId(v string) *ModifyDatabaseInstanceParameterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterRequest) Validate() error {
	return dara.Validate(s)
}
