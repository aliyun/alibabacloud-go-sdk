// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySlsDispatchConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetailConfig(v string) *ModifySlsDispatchConfigRequest
	GetDetailConfig() *string
	SetLogVersion(v int32) *ModifySlsDispatchConfigRequest
	GetLogVersion() *int32
	SetModifyType(v string) *ModifySlsDispatchConfigRequest
	GetModifyType() *string
}

type ModifySlsDispatchConfigRequest struct {
	// The detailed configuration to modify.
	//
	// <details>
	//
	// <summary>Format for version 1</summary>
	//
	// {"global":{"slsRegionId":"ap-southeast-1","logTime":180,"logStorage":1000}}
	//
	// </details>
	//
	// <details>
	//
	// <summary>Format for version 2</summary>
	//
	// {"cn":{"slsRegionId":"ap-southeast-1","logTime":180,"logStorage":3000},"intl":{"slsRegionId":"ap-southeast-1","logTime":180,"logStorage":2000}}
	//
	// </details>
	//
	// The fields are described as follows:
	//
	// - slsRegionId: The region ID to which logs are delivered.
	//
	// - logTime: The storage duration of logs. Unit: days.
	//
	// - logStorage: The log storage capacity. Unit: GB. The total capacity specified must not exceed the total capacity purchased by the user.
	//
	// example:
	//
	// {"global":{"slsRegionId":"cn-hangzhou","logTime":180,"logStorage":1000}}
	DetailConfig *string `json:"DetailConfig,omitempty" xml:"DetailConfig,omitempty"`
	// The log version. A value of 1 indicates one Logstore. A value of 2 indicates two Logstores.
	//
	//
	// 	Notice: If ModifyType is set to version, set LogVersion to the target version. If ModifyType is set to config, set LogVersion to the current version of the user.
	//
	// example:
	//
	// 1
	LogVersion *int32 `json:"LogVersion,omitempty" xml:"LogVersion,omitempty"`
	// The modification type. Valid values:
	//
	// - version: The version is changed. For example, the version is changed from 1 (logs are delivered to one Logstore) to 2 (logs are delivered to two Logstores).
	//
	// - config: The configuration is changed. For example, the log delivery region or the storage duration of logs is modified.
	//
	// example:
	//
	// version
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
}

func (s ModifySlsDispatchConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySlsDispatchConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifySlsDispatchConfigRequest) GetDetailConfig() *string {
	return s.DetailConfig
}

func (s *ModifySlsDispatchConfigRequest) GetLogVersion() *int32 {
	return s.LogVersion
}

func (s *ModifySlsDispatchConfigRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *ModifySlsDispatchConfigRequest) SetDetailConfig(v string) *ModifySlsDispatchConfigRequest {
	s.DetailConfig = &v
	return s
}

func (s *ModifySlsDispatchConfigRequest) SetLogVersion(v int32) *ModifySlsDispatchConfigRequest {
	s.LogVersion = &v
	return s
}

func (s *ModifySlsDispatchConfigRequest) SetModifyType(v string) *ModifySlsDispatchConfigRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifySlsDispatchConfigRequest) Validate() error {
	return dara.Validate(s)
}
