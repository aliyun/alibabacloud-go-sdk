// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteVpdResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *DeleteVpdResponseBody
	GetCode() *int32
	SetContent(v interface{}) *DeleteVpdResponseBody
	GetContent() interface{}
	SetMessage(v string) *DeleteVpdResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteVpdResponseBody
	GetRequestId() *string
}

type DeleteVpdResponseBody struct {
	// The detailed reason why the access was denied.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters. (If a dependent resource exists, the existing dependent resource is returned.)
	//
	// example:
	//
	// {
	//
	//       "subnet": [
	//
	//             {
	//
	//                   "tenantId": "1620939556166277",
	//
	//                   "regionId": "cn-wulanchabu",
	//
	//                   "zoneId": "cn",
	//
	//                   "type": null,
	//
	//                   "subnetId": "subnet-zqebaxa0",
	//
	//                   "name": "lql_testVPD"
	//
	//             }
	//
	//       ],
	//
	//       "nc": [
	//
	//             {}
	//
	//       ]
	//
	// }
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The response message.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BDBCC783-84CA-5733-8EEA-645C88B9009C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpdResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteVpdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteVpdResponseBody) GetContent() interface{} {
	return s.Content
}

func (s *DeleteVpdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVpdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVpdResponseBody) SetAccessDeniedDetail(v string) *DeleteVpdResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteVpdResponseBody) SetCode(v int32) *DeleteVpdResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVpdResponseBody) SetContent(v interface{}) *DeleteVpdResponseBody {
	s.Content = v
	return s
}

func (s *DeleteVpdResponseBody) SetMessage(v string) *DeleteVpdResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVpdResponseBody) SetRequestId(v string) *DeleteVpdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpdResponseBody) Validate() error {
	return dara.Validate(s)
}
