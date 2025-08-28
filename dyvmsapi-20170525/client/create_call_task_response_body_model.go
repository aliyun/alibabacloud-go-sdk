// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCallTaskResponseBody
	GetCode() *string
	SetData(v int64) *CreateCallTaskResponseBody
	GetData() *int64
	SetRequestId(v string) *CreateCallTaskResponseBody
	GetRequestId() *string
}

type CreateCallTaskResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 123456
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ED815433-724A-4357-9991-A54AD2FF09FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCallTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCallTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCallTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCallTaskResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateCallTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCallTaskResponseBody) SetCode(v string) *CreateCallTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCallTaskResponseBody) SetData(v int64) *CreateCallTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCallTaskResponseBody) SetRequestId(v string) *CreateCallTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCallTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
