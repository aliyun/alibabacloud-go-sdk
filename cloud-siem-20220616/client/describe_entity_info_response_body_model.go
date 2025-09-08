// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEntityInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeEntityInfoResponseBody
	GetCode() *int32
	SetData(v *DescribeEntityInfoResponseBodyData) *DescribeEntityInfoResponseBody
	GetData() *DescribeEntityInfoResponseBodyData
	SetMessage(v string) *DescribeEntityInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEntityInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeEntityInfoResponseBody
	GetSuccess() *bool
}

type DescribeEntityInfoResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeEntityInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEntityInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEntityInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEntityInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeEntityInfoResponseBody) GetData() *DescribeEntityInfoResponseBodyData {
	return s.Data
}

func (s *DescribeEntityInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEntityInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEntityInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEntityInfoResponseBody) SetCode(v int32) *DescribeEntityInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEntityInfoResponseBody) SetData(v *DescribeEntityInfoResponseBodyData) *DescribeEntityInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEntityInfoResponseBody) SetMessage(v string) *DescribeEntityInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEntityInfoResponseBody) SetRequestId(v string) *DescribeEntityInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEntityInfoResponseBody) SetSuccess(v bool) *DescribeEntityInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEntityInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEntityInfoResponseBodyData struct {
	// The logical ID of the entity.
	//
	// example:
	//
	// 12345
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The information about the entry.
	//
	// example:
	//
	// { location: "xian", net_connect_dir: "in", malware_type: "${aliyun.siem.sas.alert_tag.login_unusual_account}" }
	EntityInfo map[string]interface{} `json:"EntityInfo,omitempty" xml:"EntityInfo,omitempty"`
	// The type of the entity. Valid values:
	//
	// 	- ip
	//
	// 	- domain
	//
	// 	- url
	//
	// 	- process
	//
	// 	- file
	//
	// 	- host
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The information about the risk Intelligence.
	//
	// example:
	//
	// {
	//
	//       "Ip": {
	//
	//             "queryHot": "0",
	//
	//             "country": "China",
	//
	//             "province": "shanxi",
	//
	//             "ip": "221.11.XX.XXX",
	//
	//             "asn": "4837",
	//
	//             "asn_label": "CHINAXXX-Backbone - CHINA UNICOM ChinaXXX Backbone, CN"
	//
	//       }
	//
	// }
	TipInfo map[string]interface{} `json:"TipInfo,omitempty" xml:"TipInfo,omitempty"`
}

func (s DescribeEntityInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEntityInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEntityInfoResponseBodyData) GetEntityId() *int64 {
	return s.EntityId
}

func (s *DescribeEntityInfoResponseBodyData) GetEntityInfo() map[string]interface{} {
	return s.EntityInfo
}

func (s *DescribeEntityInfoResponseBodyData) GetEntityType() *string {
	return s.EntityType
}

func (s *DescribeEntityInfoResponseBodyData) GetTipInfo() map[string]interface{} {
	return s.TipInfo
}

func (s *DescribeEntityInfoResponseBodyData) SetEntityId(v int64) *DescribeEntityInfoResponseBodyData {
	s.EntityId = &v
	return s
}

func (s *DescribeEntityInfoResponseBodyData) SetEntityInfo(v map[string]interface{}) *DescribeEntityInfoResponseBodyData {
	s.EntityInfo = v
	return s
}

func (s *DescribeEntityInfoResponseBodyData) SetEntityType(v string) *DescribeEntityInfoResponseBodyData {
	s.EntityType = &v
	return s
}

func (s *DescribeEntityInfoResponseBodyData) SetTipInfo(v map[string]interface{}) *DescribeEntityInfoResponseBodyData {
	s.TipInfo = v
	return s
}

func (s *DescribeEntityInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
