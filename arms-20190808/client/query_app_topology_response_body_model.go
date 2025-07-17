// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppTopologyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *QueryAppTopologyResponseBody
	GetCode() *int64
	SetData(v interface{}) *QueryAppTopologyResponseBody
	GetData() interface{}
	SetMessage(v string) *QueryAppTopologyResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAppTopologyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAppTopologyResponseBody
	GetSuccess() *bool
}

type QueryAppTopologyResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful. If another status code is returned, the request failed.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	//
	// example:
	//
	// "Data": {
	//
	//     "nodes": [
	//
	//       {
	//
	//         "data": {
	//
	//           "duration": 0.2254335260115607,
	//
	//           "requests": 1211,
	//
	//           "type": "MYSQL",
	//
	//           "errors": 0
	//
	//         },
	//
	//         "id": "ggxw4lnjuz@c0507xxxx##MYSQL",
	//
	//         "label": "mysql-pod:3306(cart_db)"
	//
	//       }
	//
	//     ],
	//
	//     "edges": [
	//
	//       {
	//
	//         "data": {
	//
	//           "duration": 0.03333333333333333,
	//
	//           "requests": 600,
	//
	//           "type": "UNKNOWN",
	//
	//           "errors": 0
	//
	//         },
	//
	//         "id": "1974097372",
	//
	//         "source": "ggxw4lnjuz@456xxxxx",
	//
	//         "target": "ggxw4lnjuz@c0507xxxx"
	//
	//       }
	//
	//     ]
	//
	//   }
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// Internal error. Please try again. Contact the DingTalk service account if the issue                              persists after multiple retries.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 78901766-3806-4E96-8E47-CFEF59E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- true: The call was successful.
	//
	// 	- false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAppTopologyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAppTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAppTopologyResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *QueryAppTopologyResponseBody) GetData() interface{} {
	return s.Data
}

func (s *QueryAppTopologyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAppTopologyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAppTopologyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAppTopologyResponseBody) SetCode(v int64) *QueryAppTopologyResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAppTopologyResponseBody) SetData(v interface{}) *QueryAppTopologyResponseBody {
	s.Data = v
	return s
}

func (s *QueryAppTopologyResponseBody) SetMessage(v string) *QueryAppTopologyResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAppTopologyResponseBody) SetRequestId(v string) *QueryAppTopologyResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAppTopologyResponseBody) SetSuccess(v bool) *QueryAppTopologyResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAppTopologyResponseBody) Validate() error {
	return dara.Validate(s)
}
