// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataProductListLogMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetLogCode(v string) *DataProductListLogMapValue
	GetLogCode() *string
	SetLogName(v string) *DataProductListLogMapValue
	GetLogName() *string
	SetLogNameEn(v string) *DataProductListLogMapValue
	GetLogNameEn() *string
	SetLogNameKey(v string) *DataProductListLogMapValue
	GetLogNameKey() *string
	SetStatus(v bool) *DataProductListLogMapValue
	GetStatus() *bool
	SetCanOperateOrNot(v bool) *DataProductListLogMapValue
	GetCanOperateOrNot() *bool
	SetTopic(v string) *DataProductListLogMapValue
	GetTopic() *string
	SetExtraParameters(v []*DataProductListLogMapValueExtraParameters) *DataProductListLogMapValue
	GetExtraParameters() []*DataProductListLogMapValueExtraParameters
}

type DataProductListLogMapValue struct {
	// The code of the log.
	//
	// example:
	//
	// cloud_siem_config_log
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// audit log
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// audit log
	LogNameEn *string `json:"LogNameEn,omitempty" xml:"LogNameEn,omitempty"`
	// The language code of the log that is used to indicate the language in which the log is displayed.
	//
	// example:
	//
	// ${sas.cloudsiem.prod.cloud_siem_aegis_crack_from_beaver}
	LogNameKey *string `json:"LogNameKey,omitempty" xml:"LogNameKey,omitempty"`
	// The status of the log delivery. Valid values:
	//
	// 	- true: The logs are being delivered.
	//
	// 	- false: The log delivery feature is disabled.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the log delivery feature can be enabled or disabled. The feature can be enabled or disabled only by the administrator of the threat analysis feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	CanOperateOrNot *bool `json:"CanOperateOrNot,omitempty" xml:"CanOperateOrNot,omitempty"`
	// The topic of the log in the Logstore. The value is an index field in the Logstore that can be used to distinguish different logs.
	//
	// example:
	//
	// sas_login_event
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The extended parameter.
	ExtraParameters []*DataProductListLogMapValueExtraParameters `json:"ExtraParameters,omitempty" xml:"ExtraParameters,omitempty" type:"Repeated"`
}

func (s DataProductListLogMapValue) String() string {
	return dara.Prettify(s)
}

func (s DataProductListLogMapValue) GoString() string {
	return s.String()
}

func (s *DataProductListLogMapValue) GetLogCode() *string {
	return s.LogCode
}

func (s *DataProductListLogMapValue) GetLogName() *string {
	return s.LogName
}

func (s *DataProductListLogMapValue) GetLogNameEn() *string {
	return s.LogNameEn
}

func (s *DataProductListLogMapValue) GetLogNameKey() *string {
	return s.LogNameKey
}

func (s *DataProductListLogMapValue) GetStatus() *bool {
	return s.Status
}

func (s *DataProductListLogMapValue) GetCanOperateOrNot() *bool {
	return s.CanOperateOrNot
}

func (s *DataProductListLogMapValue) GetTopic() *string {
	return s.Topic
}

func (s *DataProductListLogMapValue) GetExtraParameters() []*DataProductListLogMapValueExtraParameters {
	return s.ExtraParameters
}

func (s *DataProductListLogMapValue) SetLogCode(v string) *DataProductListLogMapValue {
	s.LogCode = &v
	return s
}

func (s *DataProductListLogMapValue) SetLogName(v string) *DataProductListLogMapValue {
	s.LogName = &v
	return s
}

func (s *DataProductListLogMapValue) SetLogNameEn(v string) *DataProductListLogMapValue {
	s.LogNameEn = &v
	return s
}

func (s *DataProductListLogMapValue) SetLogNameKey(v string) *DataProductListLogMapValue {
	s.LogNameKey = &v
	return s
}

func (s *DataProductListLogMapValue) SetStatus(v bool) *DataProductListLogMapValue {
	s.Status = &v
	return s
}

func (s *DataProductListLogMapValue) SetCanOperateOrNot(v bool) *DataProductListLogMapValue {
	s.CanOperateOrNot = &v
	return s
}

func (s *DataProductListLogMapValue) SetTopic(v string) *DataProductListLogMapValue {
	s.Topic = &v
	return s
}

func (s *DataProductListLogMapValue) SetExtraParameters(v []*DataProductListLogMapValueExtraParameters) *DataProductListLogMapValue {
	s.ExtraParameters = v
	return s
}

func (s *DataProductListLogMapValue) Validate() error {
	return dara.Validate(s)
}

type DataProductListLogMapValueExtraParameters struct {
	// The ID of the extended parameter.
	//
	// example:
	//
	// flag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the extended parameter.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DataProductListLogMapValueExtraParameters) String() string {
	return dara.Prettify(s)
}

func (s DataProductListLogMapValueExtraParameters) GoString() string {
	return s.String()
}

func (s *DataProductListLogMapValueExtraParameters) GetKey() *string {
	return s.Key
}

func (s *DataProductListLogMapValueExtraParameters) GetValue() *string {
	return s.Value
}

func (s *DataProductListLogMapValueExtraParameters) SetKey(v string) *DataProductListLogMapValueExtraParameters {
	s.Key = &v
	return s
}

func (s *DataProductListLogMapValueExtraParameters) SetValue(v string) *DataProductListLogMapValueExtraParameters {
	s.Value = &v
	return s
}

func (s *DataProductListLogMapValueExtraParameters) Validate() error {
	return dara.Validate(s)
}
