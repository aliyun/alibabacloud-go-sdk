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
	// The log code.
	//
	// example:
	//
	// cloud_siem_config_log
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated and can be ignored.
	//
	// example:
	//
	// audit log
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated and can be ignored.
	//
	// example:
	//
	// audit log
	LogNameEn *string `json:"LogNameEn,omitempty" xml:"LogNameEn,omitempty"`
	// The language code of the log, which is used to display multilingual names.
	//
	// example:
	//
	// ${sas.cloudsiem.prod.cloud_siem_aegis_crack_from_beaver}
	LogNameKey *string `json:"LogNameKey,omitempty" xml:"LogNameKey,omitempty"`
	// The log delivery status. Valid values:
	//
	// - true: Delivery is in progress.
	//
	// - false: Delivery is disabled.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// Specifies whether the log delivery switch can be operated. Only the delegated administrator of Threat Analysis can operate the log delivery switch. Valid values:
	//
	// - true: The switch can be operated.
	//
	// - false: The switch cannot be operated.
	//
	// example:
	//
	// true
	CanOperateOrNot *bool `json:"CanOperateOrNot,omitempty" xml:"CanOperateOrNot,omitempty"`
	// The topic of the log in the Logstore. It is an index field of the Logstore and can be used to distinguish different logs.
	//
	// example:
	//
	// sas_login_event
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The extended parameters.
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
	if s.ExtraParameters != nil {
		for _, item := range s.ExtraParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DataProductListLogMapValueExtraParameters struct {
	// The identifier of the extended parameter.
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
