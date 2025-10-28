// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySqlLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *ModifySqlLogConfigRequest
	GetEnable() *bool
	SetEnableAudit(v bool) *ModifySqlLogConfigRequest
	GetEnableAudit() *bool
	SetFilters(v []*ModifySqlLogConfigRequestFilters) *ModifySqlLogConfigRequest
	GetFilters() []*ModifySqlLogConfigRequestFilters
	SetHotRetention(v int32) *ModifySqlLogConfigRequest
	GetHotRetention() *int32
	SetInstanceId(v string) *ModifySqlLogConfigRequest
	GetInstanceId() *string
	SetRequestEnable(v bool) *ModifySqlLogConfigRequest
	GetRequestEnable() *bool
	SetRetention(v int32) *ModifySqlLogConfigRequest
	GetRetention() *int32
}

type ModifySqlLogConfigRequest struct {
	// Specifies whether to enable DAS Enterprise Edition. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is required if you want to enable DAS Enterprise Edition. By default, the latest version of DAS Enterprise Edition that supports the database instance is enabled.
	//
	// example:
	//
	// true
	Enable      *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	EnableAudit *bool `json:"EnableAudit,omitempty" xml:"EnableAudit,omitempty"`
	// A reserved parameter.
	Filters []*ModifySqlLogConfigRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of days for which the SQL Explorer and Audit data is stored in hot storage. Valid values: 1 to 7.
	//
	// >  This parameter is required if only DAS Enterprise Edition V3 can be enabled for the database instance.
	//
	// example:
	//
	// 1
	HotRetention *int32 `json:"HotRetention,omitempty" xml:"HotRetention,omitempty"`
	// The ID of the database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rr-2ze770smbq3tpr2o9
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to enable the SQL Explorer feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter is required if only DAS Enterprise Edition V3 can be enabled for the database instance.
	//
	// example:
	//
	// true
	RequestEnable *bool `json:"RequestEnable,omitempty" xml:"RequestEnable,omitempty"`
	// The total storage duration of the SQL Explorer and Audit data. Unit: day. Valid values:
	//
	// 	- 7
	//
	// 	- 30
	//
	// 	- 180
	//
	// 	- 365
	//
	// >  If you want to enable DAS Enterprise Edition V3, the value of this parameter must be greater than or equal to 30.
	//
	// example:
	//
	// 30
	Retention *int32 `json:"Retention,omitempty" xml:"Retention,omitempty"`
}

func (s ModifySqlLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySqlLogConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifySqlLogConfigRequest) GetEnable() *bool {
	return s.Enable
}

func (s *ModifySqlLogConfigRequest) GetEnableAudit() *bool {
	return s.EnableAudit
}

func (s *ModifySqlLogConfigRequest) GetFilters() []*ModifySqlLogConfigRequestFilters {
	return s.Filters
}

func (s *ModifySqlLogConfigRequest) GetHotRetention() *int32 {
	return s.HotRetention
}

func (s *ModifySqlLogConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifySqlLogConfigRequest) GetRequestEnable() *bool {
	return s.RequestEnable
}

func (s *ModifySqlLogConfigRequest) GetRetention() *int32 {
	return s.Retention
}

func (s *ModifySqlLogConfigRequest) SetEnable(v bool) *ModifySqlLogConfigRequest {
	s.Enable = &v
	return s
}

func (s *ModifySqlLogConfigRequest) SetEnableAudit(v bool) *ModifySqlLogConfigRequest {
	s.EnableAudit = &v
	return s
}

func (s *ModifySqlLogConfigRequest) SetFilters(v []*ModifySqlLogConfigRequestFilters) *ModifySqlLogConfigRequest {
	s.Filters = v
	return s
}

func (s *ModifySqlLogConfigRequest) SetHotRetention(v int32) *ModifySqlLogConfigRequest {
	s.HotRetention = &v
	return s
}

func (s *ModifySqlLogConfigRequest) SetInstanceId(v string) *ModifySqlLogConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySqlLogConfigRequest) SetRequestEnable(v bool) *ModifySqlLogConfigRequest {
	s.RequestEnable = &v
	return s
}

func (s *ModifySqlLogConfigRequest) SetRetention(v int32) *ModifySqlLogConfigRequest {
	s.Retention = &v
	return s
}

func (s *ModifySqlLogConfigRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifySqlLogConfigRequestFilters struct {
	// A reserved parameter.
	//
	// example:
	//
	// None
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifySqlLogConfigRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ModifySqlLogConfigRequestFilters) GoString() string {
	return s.String()
}

func (s *ModifySqlLogConfigRequestFilters) GetKey() *string {
	return s.Key
}

func (s *ModifySqlLogConfigRequestFilters) GetValue() *string {
	return s.Value
}

func (s *ModifySqlLogConfigRequestFilters) SetKey(v string) *ModifySqlLogConfigRequestFilters {
	s.Key = &v
	return s
}

func (s *ModifySqlLogConfigRequestFilters) SetValue(v string) *ModifySqlLogConfigRequestFilters {
	s.Value = &v
	return s
}

func (s *ModifySqlLogConfigRequestFilters) Validate() error {
	return dara.Validate(s)
}
