// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateConfigRequest
	GetCode() *string
	SetDescription(v string) *CreateConfigRequest
	GetDescription() *string
	SetFeatureType(v int32) *CreateConfigRequest
	GetFeatureType() *int32
	SetLang(v string) *CreateConfigRequest
	GetLang() *string
	SetSourceIp(v string) *CreateConfigRequest
	GetSourceIp() *string
	SetValue(v string) *CreateConfigRequest
	GetValue() *string
}

type CreateConfigRequest struct {
	// The code of the common configuration item. Valid values:
	//
	// 	- **access_failed_cnt**: the maximum number of access attempts allowed when Data Security Center (DSC) fails to access an unauthorized resource.
	//
	// 	- **access_permission_exprie_max_days**: the maximum idle period allowed for access permissions before an alert is triggered.
	//
	// 	- **log_datasize_avg_days**: the minimum percentage of the volume of logs of a specific type generated on the current day to the average volume of logs generated in the previous 10 days before an alert is triggered.
	//
	// example:
	//
	// access_failed_cnt
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the common configuration item.
	//
	// example:
	//
	// Maximum number of access attempts allowed when DSC fails to access an unauthorized resource: 10
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 39.170.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The value of the common configuration item. The meaning of this parameter varies with the value of the Code parameter.
	//
	// 	- If you set the Code parameter to **access_failed_cnt**, the Value parameter specifies the maximum number of access attempts allowed when DSC fails to access an unauthorized resource.
	//
	// 	- If you set the Code parameter to **access_permission_exprie_max_days**, the Value parameter specifies the maximum idle period allowed for access permissions before an alert is triggered.
	//
	// 	- If you set the Code parameter to **log_datasize_avg_days**, the Value parameter specifies the minimum percentage of the volume of logs of a specific type generated on the current day to the average amount of logs generated in the previous 10 days before an alert is triggered.
	//
	// example:
	//
	// 30
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigRequest) GetCode() *string {
	return s.Code
}

func (s *CreateConfigRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateConfigRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *CreateConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateConfigRequest) GetValue() *string {
	return s.Value
}

func (s *CreateConfigRequest) SetCode(v string) *CreateConfigRequest {
	s.Code = &v
	return s
}

func (s *CreateConfigRequest) SetDescription(v string) *CreateConfigRequest {
	s.Description = &v
	return s
}

func (s *CreateConfigRequest) SetFeatureType(v int32) *CreateConfigRequest {
	s.FeatureType = &v
	return s
}

func (s *CreateConfigRequest) SetLang(v string) *CreateConfigRequest {
	s.Lang = &v
	return s
}

func (s *CreateConfigRequest) SetSourceIp(v string) *CreateConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateConfigRequest) SetValue(v string) *CreateConfigRequest {
	s.Value = &v
	return s
}

func (s *CreateConfigRequest) Validate() error {
	return dara.Validate(s)
}
