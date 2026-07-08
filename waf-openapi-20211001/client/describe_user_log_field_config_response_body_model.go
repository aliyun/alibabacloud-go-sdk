// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserLogFieldConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddList(v string) *DescribeUserLogFieldConfigResponseBody
	GetAddList() *string
	SetConfigStatus(v string) *DescribeUserLogFieldConfigResponseBody
	GetConfigStatus() *string
	SetDelList(v string) *DescribeUserLogFieldConfigResponseBody
	GetDelList() *string
	SetDeliveryType(v string) *DescribeUserLogFieldConfigResponseBody
	GetDeliveryType() *string
	SetExtendConfig(v string) *DescribeUserLogFieldConfigResponseBody
	GetExtendConfig() *string
	SetFieldList(v string) *DescribeUserLogFieldConfigResponseBody
	GetFieldList() *string
	SetLogDeliveryStrategy(v string) *DescribeUserLogFieldConfigResponseBody
	GetLogDeliveryStrategy() *string
	SetRequestId(v string) *DescribeUserLogFieldConfigResponseBody
	GetRequestId() *string
}

type DescribeUserLogFieldConfigResponseBody struct {
	// The additional log fields that are added to the default configuration. Multiple fields are separated by commas (,) in the `a,b,c,...` format.
	//
	// example:
	//
	// acl_action,acl_rule_id
	AddList *string `json:"AddList,omitempty" xml:"AddList,omitempty"`
	// The status of the log field configuration. Valid values:
	//
	// - **initial**: The log field configuration is being initialized.
	//
	// - **updating**: The log field configuration is being updated.
	//
	// - **failed_finished**: The log field configuration update failed.
	//
	// - **success_finished**: The log field configuration update succeeded.
	//
	// example:
	//
	// success_finished
	ConfigStatus *string `json:"ConfigStatus,omitempty" xml:"ConfigStatus,omitempty"`
	// The default log fields that are excluded from the log delivery configuration. Multiple fields are separated by commas (,) in the `a,b,c,...` format.
	//
	// example:
	//
	// waf_rule_id,waf_rule_type
	DelList *string `json:"DelList,omitempty" xml:"DelList,omitempty"`
	// The log delivery type. Valid values:
	//
	// - **sls**: Simple Log Service.
	//
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// The extended configuration for log delivery. The value is a JSON-formatted string that contains configuration key-value pairs, such as custom request headers.
	//
	// > For more information, see the **ExtendConfig*	- parameter description in [ModifyUserLogFieldConfig](~~ModifyUserLogFieldConfig~~).
	//
	// example:
	//
	// {\\"request_header\\":\\"Ali-Cdn-Real-Ip\\"}
	ExtendConfig *string `json:"ExtendConfig,omitempty" xml:"ExtendConfig,omitempty"`
	// The complete list of log fields that are delivered. Multiple fields are separated by commas (,) in the `a,b,c,...` format.
	//
	// example:
	//
	// account,acl_action,acl_rule_id,acl_rule_type
	FieldList *string `json:"FieldList,omitempty" xml:"FieldList,omitempty"`
	// The log delivery policies. Multiple policies are supported. The value is a JSON-formatted string that contains an array of policy objects.
	//
	// > For more information, see the **LogDeliveryStrategy*	- parameter description in [ModifyUserLogFieldConfig](~~ModifyUserLogFieldConfig~~).
	//
	// example:
	//
	// [{\\"logType\\":\\"blockLog\\",\\"rate\\":100},{\\"logType\\":\\"normalRequestLog\\",\\"rate\\":100},{\\"logType\\":\\"checkLog\\",\\"rate\\":100}]
	LogDeliveryStrategy *string `json:"LogDeliveryStrategy,omitempty" xml:"LogDeliveryStrategy,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 653778B4-4D47-5223-855B-4E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserLogFieldConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserLogFieldConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserLogFieldConfigResponseBody) GetAddList() *string {
	return s.AddList
}

func (s *DescribeUserLogFieldConfigResponseBody) GetConfigStatus() *string {
	return s.ConfigStatus
}

func (s *DescribeUserLogFieldConfigResponseBody) GetDelList() *string {
	return s.DelList
}

func (s *DescribeUserLogFieldConfigResponseBody) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *DescribeUserLogFieldConfigResponseBody) GetExtendConfig() *string {
	return s.ExtendConfig
}

func (s *DescribeUserLogFieldConfigResponseBody) GetFieldList() *string {
	return s.FieldList
}

func (s *DescribeUserLogFieldConfigResponseBody) GetLogDeliveryStrategy() *string {
	return s.LogDeliveryStrategy
}

func (s *DescribeUserLogFieldConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserLogFieldConfigResponseBody) SetAddList(v string) *DescribeUserLogFieldConfigResponseBody {
	s.AddList = &v
	return s
}

func (s *DescribeUserLogFieldConfigResponseBody) SetConfigStatus(v string) *DescribeUserLogFieldConfigResponseBody {
	s.ConfigStatus = &v
	return s
}

func (s *DescribeUserLogFieldConfigResponseBody) SetDelList(v string) *DescribeUserLogFieldConfigResponseBody {
	s.DelList = &v
	return s
}

func (s *DescribeUserLogFieldConfigResponseBody) SetDeliveryType(v string) *DescribeUserLogFieldConfigResponseBody {
	s.DeliveryType = &v
	return s
}

func (s *DescribeUserLogFieldConfigResponseBody) SetExtendConfig(v string) *DescribeUserLogFieldConfigResponseBody {
	s.ExtendConfig = &v
	return s
}

func (s *DescribeUserLogFieldConfigResponseBody) SetFieldList(v string) *DescribeUserLogFieldConfigResponseBody {
	s.FieldList = &v
	return s
}

func (s *DescribeUserLogFieldConfigResponseBody) SetLogDeliveryStrategy(v string) *DescribeUserLogFieldConfigResponseBody {
	s.LogDeliveryStrategy = &v
	return s
}

func (s *DescribeUserLogFieldConfigResponseBody) SetRequestId(v string) *DescribeUserLogFieldConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserLogFieldConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
