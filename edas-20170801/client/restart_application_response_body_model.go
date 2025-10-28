// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *RestartApplicationResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *RestartApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *RestartApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *RestartApplicationResponseBody
	GetRequestId() *string
}

type RestartApplicationResponseBody struct {
	// The ID of the change process.
	//
	// You can call the GetChangeOrderInfo operation to query the progress of this restart. For more information, see [GetChangeOrderInfo](https://help.aliyun.com/document_detail/62072.html).
	//
	// example:
	//
	// a9557bac-ddd7-*********************
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-****************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RestartApplicationResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *RestartApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RestartApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RestartApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartApplicationResponseBody) SetChangeOrderId(v string) *RestartApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *RestartApplicationResponseBody) SetCode(v int32) *RestartApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *RestartApplicationResponseBody) SetMessage(v string) *RestartApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *RestartApplicationResponseBody) SetRequestId(v string) *RestartApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
