// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceLogFieldConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddList(v string) *DescribeResourceLogFieldConfigResponseBody
	GetAddList() *string
	SetDelList(v string) *DescribeResourceLogFieldConfigResponseBody
	GetDelList() *string
	SetDeliveryType(v string) *DescribeResourceLogFieldConfigResponseBody
	GetDeliveryType() *string
	SetExtendConfig(v string) *DescribeResourceLogFieldConfigResponseBody
	GetExtendConfig() *string
	SetFieldList(v string) *DescribeResourceLogFieldConfigResponseBody
	GetFieldList() *string
	SetLogDeliveryStrategy(v string) *DescribeResourceLogFieldConfigResponseBody
	GetLogDeliveryStrategy() *string
	SetRequestId(v string) *DescribeResourceLogFieldConfigResponseBody
	GetRequestId() *string
}

type DescribeResourceLogFieldConfigResponseBody struct {
	// example:
	//
	// acl_test,acl_action,acl_rule_id,waf_test,acl_rule_type
	AddList *string `json:"AddList,omitempty" xml:"AddList,omitempty"`
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
	// account,acl_action,acl_rule_id,acl_rule_type,acl_test,antiscan_action,antiscan_rule_id,antiscan_rule_type,antiscan_test,body_bytes_sent,bypass_matched_ids
	FieldList *string `json:"FieldList,omitempty" xml:"FieldList,omitempty"`
	// example:
	//
	// [{\\"logType\\":\\"blockLog\\",\\"rate\\":100},{\\"logType\\":\\"normalRequestLog\\",\\"rate\\":100},{\\"logType\\":\\"checkLog\\",\\"rate\\":100}]
	LogDeliveryStrategy *string `json:"LogDeliveryStrategy,omitempty" xml:"LogDeliveryStrategy,omitempty"`
	// example:
	//
	// 7B53B47C-D368-5344-BB5E-79******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeResourceLogFieldConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogFieldConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogFieldConfigResponseBody) GetAddList() *string {
	return s.AddList
}

func (s *DescribeResourceLogFieldConfigResponseBody) GetDelList() *string {
	return s.DelList
}

func (s *DescribeResourceLogFieldConfigResponseBody) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *DescribeResourceLogFieldConfigResponseBody) GetExtendConfig() *string {
	return s.ExtendConfig
}

func (s *DescribeResourceLogFieldConfigResponseBody) GetFieldList() *string {
	return s.FieldList
}

func (s *DescribeResourceLogFieldConfigResponseBody) GetLogDeliveryStrategy() *string {
	return s.LogDeliveryStrategy
}

func (s *DescribeResourceLogFieldConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceLogFieldConfigResponseBody) SetAddList(v string) *DescribeResourceLogFieldConfigResponseBody {
	s.AddList = &v
	return s
}

func (s *DescribeResourceLogFieldConfigResponseBody) SetDelList(v string) *DescribeResourceLogFieldConfigResponseBody {
	s.DelList = &v
	return s
}

func (s *DescribeResourceLogFieldConfigResponseBody) SetDeliveryType(v string) *DescribeResourceLogFieldConfigResponseBody {
	s.DeliveryType = &v
	return s
}

func (s *DescribeResourceLogFieldConfigResponseBody) SetExtendConfig(v string) *DescribeResourceLogFieldConfigResponseBody {
	s.ExtendConfig = &v
	return s
}

func (s *DescribeResourceLogFieldConfigResponseBody) SetFieldList(v string) *DescribeResourceLogFieldConfigResponseBody {
	s.FieldList = &v
	return s
}

func (s *DescribeResourceLogFieldConfigResponseBody) SetLogDeliveryStrategy(v string) *DescribeResourceLogFieldConfigResponseBody {
	s.LogDeliveryStrategy = &v
	return s
}

func (s *DescribeResourceLogFieldConfigResponseBody) SetRequestId(v string) *DescribeResourceLogFieldConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceLogFieldConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
