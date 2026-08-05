// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfflineTaskLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyOfflineTaskLogResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyOfflineTaskLogResponseBody
	GetResult() map[string]interface{}
}

type ModifyOfflineTaskLogResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1-2-3-4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	//
	// example:
	//
	// {
	//
	//   "network": {
	//
	//     "publicEs": {
	//
	//       "enabled": true,
	//
	//       "whiteIpGroup": [
	//
	//         {
	//
	//           "groupName": "kevintest",
	//
	//           "ips": [
	//
	//             "1.2.3.4",
	//
	//             "1.2.4.3"
	//
	//           ]
	//
	//         }
	//
	//       ]
	//
	//     }
	//
	//   }
	//
	// }
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyOfflineTaskLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskLogResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyOfflineTaskLogResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyOfflineTaskLogResponseBody) SetRequestId(v string) *ModifyOfflineTaskLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOfflineTaskLogResponseBody) SetResult(v map[string]interface{}) *ModifyOfflineTaskLogResponseBody {
	s.Result = v
	return s
}

func (s *ModifyOfflineTaskLogResponseBody) Validate() error {
	return dara.Validate(s)
}
