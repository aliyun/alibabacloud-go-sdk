// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDoInsightsActionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DoInsightsActionResponseBody
	GetCode() *int32
	SetData(v string) *DoInsightsActionResponseBody
	GetData() *string
	SetMessage(v string) *DoInsightsActionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DoInsightsActionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DoInsightsActionResponseBody
	GetSuccess() *bool
}

type DoInsightsActionResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters vary with the value of module.
	//
	// 	- QueryTopo
	//
	//         {
	//
	//         "nodes": [Object] # The nodes. For more information, see node details in the supplementary notes of response parameters.
	//
	//         "edges": [Object] # The edges. For more information, see edge details in the supplementary notes of response parameters.
	//
	//         }
	//
	// 	- QueryTopoRed
	//
	//         {
	//
	//           "nodeRed": {
	//
	//           	"nodeId": {
	//
	//           		"count": double, # The total number of requests in the specified time range.
	//
	//           		"error": double, # The total number of errors in the specified time range.
	//
	//           		"rt": double, # The average response time in the specified time range. Unit: milliseconds.
	//
	//           	}
	//
	//           },
	//
	//           "edgeRed": {
	//
	//           	"edgeId": {
	//
	//           	    "count": double, # The total number of requests in the specified time range.
	//
	//           		"error": double, # The total number of errors in the specified time range.
	//
	//           		"rt": double, # The average response time in the specified time range. Unit: milliseconds.
	//
	//           	}
	//
	//           }
	//
	// }
	//
	// ```
	//
	// ```
	//
	// example:
	//
	// - QueryTopo
	//
	//
	// 	{
	//
	// 		"nodes": [
	//
	// 			{
	//
	// 				"nodeId": "3bfe1a747389273388182760406c079d",
	//
	// 				"entity": {
	//
	// 					"regionId": "cn-hangzhou",
	//
	// 					"appType": "TRACE",
	//
	// 					"appId": "xxxxxxxxxxxxxxxx",
	//
	// 					"name": "prometheus-pop-cn-hangzhou",
	//
	// 					"entityId": "3bfe1a747389273388182760406c079d",
	//
	// 					"firstSeenTms": 1721733226981,
	//
	// 					"lastSeenTms": 1721789171614,
	//
	// 					"type": "APPLICATION"
	//
	// 				},
	//
	// 				"attrs": {
	//
	// 					"RED": {
	//
	// 						"count": 643848.0,
	//
	// 						"error": 0.0,
	//
	// 						"rt": 172.31701892372112
	//
	// 					}
	//
	// 				}
	//
	// 			}
	//
	// 		],
	//
	// 		"edges": [
	//
	// 			{
	//
	// 				"from": "98b4184b22e588cf86e9a29aa4179606",
	//
	// 				"to": "98b4184b22e588cf86e9a29aa4179606",
	//
	// 				"type": "CALLS",
	//
	// 				"attrs": {
	//
	// 					"RED": {
	//
	// 						"count": 4.0,
	//
	// 						"error": 0.0,
	//
	// 						"rt": 37.0
	//
	// 					}
	//
	// 				},
	//
	// 				"edgeId": "5d611597e4b0013d0947615c9eca4de6",
	//
	// 				"firstSeenTms": 1721783795125,
	//
	// 				"lastSeenTms": 1721787371614
	//
	// 			}
	//
	// 		]
	//
	// 	}
	//
	//
	// - QueryTopoRed
	//
	// 	{
	//
	// 		"nodeRed": {
	//
	// 			"361d9f32e58cef316bf2355f3ff05575": {
	//
	// 				"count": 3258110.0,
	//
	// 				"error": 74.0,
	//
	// 				"rt": 167.39844355494878
	//
	// 			}
	//
	// 		},
	//
	// 		"edgeRed": {}
	//
	// 	}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 626037F5-FDEB-45B0-804C-B3C92797A64E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DoInsightsActionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DoInsightsActionResponseBody) GoString() string {
	return s.String()
}

func (s *DoInsightsActionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DoInsightsActionResponseBody) GetData() *string {
	return s.Data
}

func (s *DoInsightsActionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DoInsightsActionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DoInsightsActionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DoInsightsActionResponseBody) SetCode(v int32) *DoInsightsActionResponseBody {
	s.Code = &v
	return s
}

func (s *DoInsightsActionResponseBody) SetData(v string) *DoInsightsActionResponseBody {
	s.Data = &v
	return s
}

func (s *DoInsightsActionResponseBody) SetMessage(v string) *DoInsightsActionResponseBody {
	s.Message = &v
	return s
}

func (s *DoInsightsActionResponseBody) SetRequestId(v string) *DoInsightsActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DoInsightsActionResponseBody) SetSuccess(v bool) *DoInsightsActionResponseBody {
	s.Success = &v
	return s
}

func (s *DoInsightsActionResponseBody) Validate() error {
	return dara.Validate(s)
}
