// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFoTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFoTaskStatusResponseBody
	GetCode() *string
	SetData(v string) *GetFoTaskStatusResponseBody
	GetData() *string
	SetMessage(v string) *GetFoTaskStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFoTaskStatusResponseBody
	GetRequestId() *string
}

type GetFoTaskStatusResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The status of the switchover task.
	//
	// example:
	//
	// Running
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message. If the request was successful, a success message is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// OKITHEVRQCN6ULQG
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 923692F0-A15B-58B4-BAF4-2AFA4AF46240
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFoTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFoTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetFoTaskStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFoTaskStatusResponseBody) GetData() *string {
	return s.Data
}

func (s *GetFoTaskStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFoTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFoTaskStatusResponseBody) SetCode(v string) *GetFoTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetFoTaskStatusResponseBody) SetData(v string) *GetFoTaskStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetFoTaskStatusResponseBody) SetMessage(v string) *GetFoTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetFoTaskStatusResponseBody) SetRequestId(v string) *GetFoTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFoTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
