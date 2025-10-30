// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAxgGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAxgGroupResponseBody
	GetCode() *string
	SetGroupId(v int64) *CreateAxgGroupResponseBody
	GetGroupId() *int64
	SetMessage(v string) *CreateAxgGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAxgGroupResponseBody
	GetRequestId() *string
}

type CreateAxgGroupResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of number group G. The value of this parameter is required when the [BindAxg](https://help.aliyun.com/document_detail/110249.html) operation is called to add an AXG binding.
	//
	// example:
	//
	// 2000000000001
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 635C0FDA-9EBC-43D7-B368-9F583C08A126
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAxgGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAxgGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAxgGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAxgGroupResponseBody) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateAxgGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAxgGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAxgGroupResponseBody) SetCode(v string) *CreateAxgGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAxgGroupResponseBody) SetGroupId(v int64) *CreateAxgGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateAxgGroupResponseBody) SetMessage(v string) *CreateAxgGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAxgGroupResponseBody) SetRequestId(v string) *CreateAxgGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAxgGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
