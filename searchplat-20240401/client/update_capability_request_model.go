// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCapabilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetItemDesc(v string) *UpdateCapabilityRequest
	GetItemDesc() *string
	SetItemValue(v map[string]interface{}) *UpdateCapabilityRequest
	GetItemValue() map[string]interface{}
	SetDryRun(v bool) *UpdateCapabilityRequest
	GetDryRun() *bool
}

type UpdateCapabilityRequest struct {
	// The configuration description.
	//
	// example:
	//
	// 描述
	ItemDesc *string `json:"itemDesc,omitempty" xml:"itemDesc,omitempty"`
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
	// Specifies whether to validate the request parameters without applying the changes. Default value: false.
	//
	// Valid values:
	//
	// - **true**
	//
	// - **false**.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateCapabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCapabilityRequest) GoString() string {
	return s.String()
}

func (s *UpdateCapabilityRequest) GetItemDesc() *string {
	return s.ItemDesc
}

func (s *UpdateCapabilityRequest) GetItemValue() map[string]interface{} {
	return s.ItemValue
}

func (s *UpdateCapabilityRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateCapabilityRequest) SetItemDesc(v string) *UpdateCapabilityRequest {
	s.ItemDesc = &v
	return s
}

func (s *UpdateCapabilityRequest) SetItemValue(v map[string]interface{}) *UpdateCapabilityRequest {
	s.ItemValue = v
	return s
}

func (s *UpdateCapabilityRequest) SetDryRun(v bool) *UpdateCapabilityRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateCapabilityRequest) Validate() error {
	return dara.Validate(s)
}
