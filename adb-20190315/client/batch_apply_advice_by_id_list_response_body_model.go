// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchApplyAdviceByIdListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *BatchApplyAdviceByIdListResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchApplyAdviceByIdListResponseBody
	GetRequestId() *string
}

type BatchApplyAdviceByIdListResponseBody struct {
	// The message returned for the operation. Valid values:
	//
	// 	- **SUCCESS*	- is returned if the operation is successful.
	//
	// 	- An error message is returned if the operation fails.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 86F92D26-B774-5FA1-8E53-82CBEEEBB012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchApplyAdviceByIdListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchApplyAdviceByIdListResponseBody) GoString() string {
	return s.String()
}

func (s *BatchApplyAdviceByIdListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchApplyAdviceByIdListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchApplyAdviceByIdListResponseBody) SetMessage(v string) *BatchApplyAdviceByIdListResponseBody {
	s.Message = &v
	return s
}

func (s *BatchApplyAdviceByIdListResponseBody) SetRequestId(v string) *BatchApplyAdviceByIdListResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchApplyAdviceByIdListResponseBody) Validate() error {
	return dara.Validate(s)
}
