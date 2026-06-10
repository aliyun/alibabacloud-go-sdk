// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAlertItemsResponseBody
	GetCode() *string
	SetData(v interface{}) *ListAlertItemsResponseBody
	GetData() interface{}
	SetMessage(v string) *ListAlertItemsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAlertItemsResponseBody
	GetRequestId() *string
}

type ListAlertItemsResponseBody struct {
	// Status code
	//
	// - If `code == Success`, authorization succeeded.
	//
	// - Any other status code indicates authorization failed. When authorization fails, view the `message` field to obtain the detailed error message.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Returned data.
	//
	// example:
	//
	// {
	//
	//     "NODE": {
	//
	//       "饱和度": [
	//
	//         "节点CPU使用率检测",
	//
	//         "节点内核态CPU使用检测",
	//
	//         "节点软中断CPU使用检测",
	//
	//         "节点内存使用检测",
	//
	//         "节点内核内存使用检测",
	//
	//         "节点文件描述符使用检测",
	//
	//         "节点根文件系统使用检测",
	//
	//         "节点cgroup泄漏检测",
	//
	//         "节点Sockets使用检测",
	//
	//         "节点TCP内存使用检测"
	//
	//       ],
	//
	//       "延时": [
	//
	//         "节点调度延时检测",
	//
	//         "节点网络延时检测",
	//
	//         "节点磁盘写入延迟检测",
	//
	//         "节点磁盘读取延迟检测"
	//
	//       ],
	//
	//       "负载": [
	//
	//         "节点磁盘IO流量检测",
	//
	//         "节点load average检测"
	//
	//       ],
	//
	//       "错误": [
	//
	//         "节点网络丢包检测",
	//
	//         "节点OOM夯机预测及检测"
	//
	//       ]
	//
	//     },
	//
	//     "POD": {
	//
	//       "饱和度": [
	//
	//         "Pod内存使用检测",
	//
	//         "Pod CPU使用率检测"
	//
	//       ],
	//
	//       "错误": [
	//
	//         "Pod CPU限流检测",
	//
	//         "Pod OOM事件检测"
	//
	//       ]
	//
	//     }
	//
	// }
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// Error message
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAlertItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertItemsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAlertItemsResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ListAlertItemsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAlertItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertItemsResponseBody) SetCode(v string) *ListAlertItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlertItemsResponseBody) SetData(v interface{}) *ListAlertItemsResponseBody {
	s.Data = v
	return s
}

func (s *ListAlertItemsResponseBody) SetMessage(v string) *ListAlertItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlertItemsResponseBody) SetRequestId(v string) *ListAlertItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertItemsResponseBody) Validate() error {
	return dara.Validate(s)
}
