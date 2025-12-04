// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteErAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteErAttachmentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *DeleteErAttachmentResponseBody
	GetCode() *int32
	SetContent(v interface{}) *DeleteErAttachmentResponseBody
	GetContent() interface{}
	SetMessage(v string) *DeleteErAttachmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteErAttachmentResponseBody
	GetRequestId() *string
}

type DeleteErAttachmentResponseBody struct {
	// 访问被拒绝的详细原因。
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
	// The response content. If a resource has dependent resources, the existing dependent resources are returned.
	//
	// example:
	//
	// {
	//
	//     "ER_RMAP": [
	//
	//         {
	//
	//             "erId": "er-opy1wrfv",
	//
	//             "destinationCidrBlock": "0.0.0.0/0",
	//
	//             "regionId": "cn-wulanchabu",
	//
	//             "routeMapNum": 3000,
	//
	//             "erRouteMapId": "er-rmap-v5lfhmvm",
	//
	//             "action": "permit",
	//
	//             "status": "Available"
	//
	//         },
	//
	//         {
	//
	//             "erId": "er-opy1wrfv",
	//
	//             "destinationCidrBlock": "0.0.0.0/0",
	//
	//             "regionId": "cn-wulanchabu",
	//
	//             "routeMapNum": 3000,
	//
	//             "erRouteMapId": "er-rmap-of3r0ndh",
	//
	//             "action": "permit",
	//
	//             "status": "Available"
	//
	//         }
	//
	//     ]
	//
	// }
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// response message, if the success request is
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteErAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteErAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteErAttachmentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteErAttachmentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteErAttachmentResponseBody) GetContent() interface{} {
	return s.Content
}

func (s *DeleteErAttachmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteErAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteErAttachmentResponseBody) SetAccessDeniedDetail(v string) *DeleteErAttachmentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteErAttachmentResponseBody) SetCode(v int32) *DeleteErAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteErAttachmentResponseBody) SetContent(v interface{}) *DeleteErAttachmentResponseBody {
	s.Content = v
	return s
}

func (s *DeleteErAttachmentResponseBody) SetMessage(v string) *DeleteErAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteErAttachmentResponseBody) SetRequestId(v string) *DeleteErAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteErAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
