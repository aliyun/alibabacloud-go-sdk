// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyParameterRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyParameterRequest
	GetDBInstanceId() *string
	SetForcerestart(v bool) *ModifyParameterRequest
	GetForcerestart() *bool
	SetOwnerAccount(v string) *ModifyParameterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyParameterRequest
	GetOwnerId() *int64
	SetParameterGroupId(v string) *ModifyParameterRequest
	GetParameterGroupId() *string
	SetParameters(v string) *ModifyParameterRequest
	GetParameters() *string
	SetResourceOwnerAccount(v string) *ModifyParameterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyParameterRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *ModifyParameterRequest
	GetSwitchTime() *string
	SetSwitchTimeMode(v string) *ModifyParameterRequest
	GetSwitchTimeMode() *string
}

type ModifyParameterRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to restart the instance for a new parameter value to take effect. Valid values:
	//
	// 	- **true**: The system forcefully restarts the instance. If a new parameter value takes effect only after the instance restarts, you must set this parameter to true. Otherwise, the new parameter value cannot take effect.
	//
	// 	- **false**: The system does not forcefully restart the instance.
	//
	// Default value: **false**.
	//
	// example:
	//
	// false
	Forcerestart *bool   `json:"Forcerestart,omitempty" xml:"Forcerestart,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The parameter template ID.
	//
	// > 	- If you specify this parameter, you do not need to specify **Parameters**.
	//
	// > 	- If the parameter template can be applied only after the instance is restarted, you must specify **Forcerestart**.
	//
	// example:
	//
	// rpg-xxxxxxxxx
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The JSON strings of parameters and their values. All the parameter values are of the string type. Format: {"Parameter name 1":"Parameter value 1","Parameter name 2":"Parameter value 2"...}. You can call the DescribeParameterTemplates operation to query parameter names and values.
	//
	// >  If you specify this parameter, you do not need to specify **ParameterGroupId**.
	//
	// example:
	//
	// {"delayed_insert_timeout":"600","max_length_for_sort_data":"2048"}
	Parameters           *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The time at which the modification takes effect. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// > This time must be later than the time at which you call this operation.
	//
	// example:
	//
	// 2022-05-06T09:24:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The time at which the modification takes effect. Valid values:
	//
	// - **Immediate**: immediately modifies the parameter. This is the default value.
	//
	// - **MaintainTime**: modifies the parameter during the maintenance window of the instance. You can call the ModifyDBInstanceMaintainTime operation to change the maintenance window.
	//
	// - **ScheduleTime**: modifies the parameter at the point in time that you specify. If you specify this value, you must also specify **SwitchTime**.
	//
	// example:
	//
	// ScheduleTime
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s ModifyParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyParameterRequest) GoString() string {
	return s.String()
}

func (s *ModifyParameterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyParameterRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyParameterRequest) GetForcerestart() *bool {
	return s.Forcerestart
}

func (s *ModifyParameterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyParameterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyParameterRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *ModifyParameterRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyParameterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyParameterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyParameterRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyParameterRequest) GetSwitchTimeMode() *string {
	return s.SwitchTimeMode
}

func (s *ModifyParameterRequest) SetClientToken(v string) *ModifyParameterRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyParameterRequest) SetDBInstanceId(v string) *ModifyParameterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyParameterRequest) SetForcerestart(v bool) *ModifyParameterRequest {
	s.Forcerestart = &v
	return s
}

func (s *ModifyParameterRequest) SetOwnerAccount(v string) *ModifyParameterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyParameterRequest) SetOwnerId(v int64) *ModifyParameterRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyParameterRequest) SetParameterGroupId(v string) *ModifyParameterRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *ModifyParameterRequest) SetParameters(v string) *ModifyParameterRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyParameterRequest) SetResourceOwnerAccount(v string) *ModifyParameterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyParameterRequest) SetResourceOwnerId(v int64) *ModifyParameterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyParameterRequest) SetSwitchTime(v string) *ModifyParameterRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyParameterRequest) SetSwitchTimeMode(v string) *ModifyParameterRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *ModifyParameterRequest) Validate() error {
	return dara.Validate(s)
}
