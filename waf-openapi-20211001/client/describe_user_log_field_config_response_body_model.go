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
	// example:
	//
	// acl_action,acl_rule_id
	AddList *string `json:"AddList,omitempty" xml:"AddList,omitempty"`
	// example:
	//
	// success_finished
	ConfigStatus *string `json:"ConfigStatus,omitempty" xml:"ConfigStatus,omitempty"`
	// example:
	//
	// waf_rule_id,waf_rule_type
	DelList *string `json:"DelList,omitempty" xml:"DelList,omitempty"`
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// example:
	//
	// {\\"request_header\\":\\"Ali-Cdn-Real-Ip\\"}
	ExtendConfig *string `json:"ExtendConfig,omitempty" xml:"ExtendConfig,omitempty"`
	// example:
	//
	// account,acl_action,acl_rule_id,acl_rule_type
	FieldList *string `json:"FieldList,omitempty" xml:"FieldList,omitempty"`
	// example:
	//
	// [{\\"logType\\":\\"blockLog\\",\\"rate\\":100},{\\"logType\\":\\"normalRequestLog\\",\\"rate\\":100},{\\"logType\\":\\"checkLog\\",\\"rate\\":100}]
	LogDeliveryStrategy *string `json:"LogDeliveryStrategy,omitempty" xml:"LogDeliveryStrategy,omitempty"`
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
