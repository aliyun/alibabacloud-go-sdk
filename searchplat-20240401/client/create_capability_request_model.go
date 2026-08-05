// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCapabilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetItemDesc(v string) *CreateCapabilityRequest
	GetItemDesc() *string
	SetItemName(v string) *CreateCapabilityRequest
	GetItemName() *string
	SetItemValue(v map[string]interface{}) *CreateCapabilityRequest
	GetItemValue() map[string]interface{}
	SetDryRun(v bool) *CreateCapabilityRequest
	GetDryRun() *bool
}

type CreateCapabilityRequest struct {
	// The configuration description.
	//
	// example:
	//
	// 描述
	ItemDesc *string `json:"itemDesc,omitempty" xml:"itemDesc,omitempty"`
	// The configuration name.
	//
	// example:
	//
	// es_knowledge_base
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// The configuration item.
	//
	// example:
	//
	// {
	//
	//         "name": "知识库名称",
	//
	//         "type": "aliyun-es",
	//
	//         "endpoint": {
	//
	//             "address": "http://xxxxx.es-serverless.aliyuncs.com:9200",
	//
	//             "authorization": "xxx:ABCDxxxx",
	//
	//             "network": {
	//
	//                 "type": "public"
	//
	//             },
	//
	//             "region": ""
	//
	//         },
	//
	//         "function": {
	//
	//             "indexName": "product_info",
	//
	//             "queryType": "hybrid-search",
	//
	//             "description": "金融理财产品知识库，理财产品名称列表及详细描述",
	//
	//             "parameters": {
	//
	//                 "properties": {
	//
	//                     "query": {
	//
	//                         "defaultValue": "",
	//
	//                         "description": "理财产品名称或详细描述",
	//
	//                         "type": "string"
	//
	//                     }
	//
	//                 },
	//
	//                 "required": [
	//
	//                     "query"
	//
	//                 ]
	//
	//             },
	//
	//             "template": "{\\"query\\":{\\"multi_match\\":{\\"query\\":\\"${parameters.query}\\",\\"fields\\": [\\"*\\"]}}}",
	//
	//             "type": "dsl",
	//
	//             "embedding": {
	//
	//                 "denseModel": "dense",
	//
	//                 "sparseModel": "sparse",
	//
	//                 "address": "address",
	//
	//                 "authorization": "authorization"
	//
	//             },
	//
	//             "filter": "es或os语句",
	//
	//             "size": 10
	//
	//         }
	//
	//     }
	ItemValue map[string]interface{} `json:"itemValue,omitempty" xml:"itemValue,omitempty"`
	// Specifies whether to check the validity of the request parameters without performing the actual operation. Default value: false.
	//
	// Valid values:
	//
	// - **true**
	//
	// - **false**.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateCapabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCapabilityRequest) GoString() string {
	return s.String()
}

func (s *CreateCapabilityRequest) GetItemDesc() *string {
	return s.ItemDesc
}

func (s *CreateCapabilityRequest) GetItemName() *string {
	return s.ItemName
}

func (s *CreateCapabilityRequest) GetItemValue() map[string]interface{} {
	return s.ItemValue
}

func (s *CreateCapabilityRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateCapabilityRequest) SetItemDesc(v string) *CreateCapabilityRequest {
	s.ItemDesc = &v
	return s
}

func (s *CreateCapabilityRequest) SetItemName(v string) *CreateCapabilityRequest {
	s.ItemName = &v
	return s
}

func (s *CreateCapabilityRequest) SetItemValue(v map[string]interface{}) *CreateCapabilityRequest {
	s.ItemValue = v
	return s
}

func (s *CreateCapabilityRequest) SetDryRun(v bool) *CreateCapabilityRequest {
	s.DryRun = &v
	return s
}

func (s *CreateCapabilityRequest) Validate() error {
	return dara.Validate(s)
}
