// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *RollbackApplicationResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *RollbackApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *RollbackApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *RollbackApplicationResponseBody
	GetRequestId() *string
}

type RollbackApplicationResponseBody struct {
	// The change process ID. You can call the GetChangeOrderInfo operation to query the progress of this rollback. For more information, see [GetChangeOrderInfo](https://help.aliyun.com/document_detail/423155.html).
	//
	// example:
	//
	// 921026b8-d1be-************
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D16979DC-4D42-*********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackApplicationResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *RollbackApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RollbackApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RollbackApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackApplicationResponseBody) SetChangeOrderId(v string) *RollbackApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *RollbackApplicationResponseBody) SetCode(v int32) *RollbackApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *RollbackApplicationResponseBody) SetMessage(v string) *RollbackApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *RollbackApplicationResponseBody) SetRequestId(v string) *RollbackApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
