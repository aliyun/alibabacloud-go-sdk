// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTodoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TransferTodoResponseBody
	GetCode() *string
	SetMessage(v string) *TransferTodoResponseBody
	GetMessage() *string
	SetRequestId(v string) *TransferTodoResponseBody
	GetRequestId() *string
	SetTargets(v []*TransferTodoResponseBodyTargets) *TransferTodoResponseBody
	GetTargets() []*TransferTodoResponseBodyTargets
	SetTodoId(v string) *TransferTodoResponseBody
	GetTodoId() *string
}

type TransferTodoResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The processing results for each target.
	Targets []*TransferTodoResponseBodyTargets `json:"targets,omitempty" xml:"targets,omitempty" type:"Repeated"`
	// The ID of the to-do task.
	//
	// example:
	//
	// todo_abc123
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
}

func (s TransferTodoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferTodoResponseBody) GoString() string {
	return s.String()
}

func (s *TransferTodoResponseBody) GetCode() *string {
	return s.Code
}

func (s *TransferTodoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TransferTodoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferTodoResponseBody) GetTargets() []*TransferTodoResponseBodyTargets {
	return s.Targets
}

func (s *TransferTodoResponseBody) GetTodoId() *string {
	return s.TodoId
}

func (s *TransferTodoResponseBody) SetCode(v string) *TransferTodoResponseBody {
	s.Code = &v
	return s
}

func (s *TransferTodoResponseBody) SetMessage(v string) *TransferTodoResponseBody {
	s.Message = &v
	return s
}

func (s *TransferTodoResponseBody) SetRequestId(v string) *TransferTodoResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferTodoResponseBody) SetTargets(v []*TransferTodoResponseBodyTargets) *TransferTodoResponseBody {
	s.Targets = v
	return s
}

func (s *TransferTodoResponseBody) SetTodoId(v string) *TransferTodoResponseBody {
	s.TodoId = &v
	return s
}

func (s *TransferTodoResponseBody) Validate() error {
	if s.Targets != nil {
		for _, item := range s.Targets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TransferTodoResponseBodyTargets struct {
	// The processing result. Valid values: CREATED, REVIVED, and ALREADY_PENDING.
	//
	// example:
	//
	// CREATED
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// The ID of the target user.
	//
	// example:
	//
	// 43
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s TransferTodoResponseBodyTargets) String() string {
	return dara.Prettify(s)
}

func (s TransferTodoResponseBodyTargets) GoString() string {
	return s.String()
}

func (s *TransferTodoResponseBodyTargets) GetResult() *string {
	return s.Result
}

func (s *TransferTodoResponseBodyTargets) GetUserId() *string {
	return s.UserId
}

func (s *TransferTodoResponseBodyTargets) SetResult(v string) *TransferTodoResponseBodyTargets {
	s.Result = &v
	return s
}

func (s *TransferTodoResponseBodyTargets) SetUserId(v string) *TransferTodoResponseBodyTargets {
	s.UserId = &v
	return s
}

func (s *TransferTodoResponseBodyTargets) Validate() error {
	return dara.Validate(s)
}
