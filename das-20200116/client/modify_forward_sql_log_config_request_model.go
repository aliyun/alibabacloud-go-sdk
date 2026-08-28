// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyForwardSqlLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *ModifyForwardSqlLogConfigRequest
	GetEnable() *bool
	SetInstanceId(v string) *ModifyForwardSqlLogConfigRequest
	GetInstanceId() *string
	SetService(v string) *ModifyForwardSqlLogConfigRequest
	GetService() *string
	SetSource(v string) *ModifyForwardSqlLogConfigRequest
	GetSource() *string
}

type ModifyForwardSqlLogConfigRequest struct {
	// Specifies whether to enable the feature. Valid values:
	//
	// - **true**: Enable.
	//
	// - **false**: Disable.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The database instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The service type. Valid values:
	//
	// DAS_OPS: enables TOP KEY delivery.
	//
	// This parameter is required.
	//
	// example:
	//
	// DAS_OPS
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The task source. Valid values:
	//
	// - TOP_KEY: enables TOP KEY delivery.
	//
	// This parameter is required.
	//
	// example:
	//
	// TOP_KEY
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ModifyForwardSqlLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardSqlLogConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyForwardSqlLogConfigRequest) GetEnable() *bool {
	return s.Enable
}

func (s *ModifyForwardSqlLogConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyForwardSqlLogConfigRequest) GetService() *string {
	return s.Service
}

func (s *ModifyForwardSqlLogConfigRequest) GetSource() *string {
	return s.Source
}

func (s *ModifyForwardSqlLogConfigRequest) SetEnable(v bool) *ModifyForwardSqlLogConfigRequest {
	s.Enable = &v
	return s
}

func (s *ModifyForwardSqlLogConfigRequest) SetInstanceId(v string) *ModifyForwardSqlLogConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyForwardSqlLogConfigRequest) SetService(v string) *ModifyForwardSqlLogConfigRequest {
	s.Service = &v
	return s
}

func (s *ModifyForwardSqlLogConfigRequest) SetSource(v string) *ModifyForwardSqlLogConfigRequest {
	s.Source = &v
	return s
}

func (s *ModifyForwardSqlLogConfigRequest) Validate() error {
	return dara.Validate(s)
}
