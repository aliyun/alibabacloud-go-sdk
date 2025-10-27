// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAdviceByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ApplyAdviceByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *ApplyAdviceByIdResponseBody
	GetRequestId() *string
}

type ApplyAdviceByIdResponseBody struct {
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
	// 5DC10091-348D-12B1-906D-AB49D658012E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyAdviceByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyAdviceByIdResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyAdviceByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyAdviceByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyAdviceByIdResponseBody) SetMessage(v string) *ApplyAdviceByIdResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyAdviceByIdResponseBody) SetRequestId(v string) *ApplyAdviceByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyAdviceByIdResponseBody) Validate() error {
	return dara.Validate(s)
}
