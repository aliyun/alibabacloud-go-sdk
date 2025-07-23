// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecuteOperationResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetExecuteOperationResultResponseBody
	GetCode() *int32
	SetData(v *GetExecuteOperationResultResponseBodyData) *GetExecuteOperationResultResponseBody
	GetData() *GetExecuteOperationResultResponseBodyData
	SetMessage(v string) *GetExecuteOperationResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetExecuteOperationResultResponseBody
	GetRequestId() *string
}

type GetExecuteOperationResultResponseBody struct {
	// The HTTP status code. A value of 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed result of the queried operation.
	Data *GetExecuteOperationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// " "
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExecuteOperationResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExecuteOperationResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetExecuteOperationResultResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetExecuteOperationResultResponseBody) GetData() *GetExecuteOperationResultResponseBodyData {
	return s.Data
}

func (s *GetExecuteOperationResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetExecuteOperationResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExecuteOperationResultResponseBody) SetCode(v int32) *GetExecuteOperationResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetExecuteOperationResultResponseBody) SetData(v *GetExecuteOperationResultResponseBodyData) *GetExecuteOperationResultResponseBody {
	s.Data = v
	return s
}

func (s *GetExecuteOperationResultResponseBody) SetMessage(v string) *GetExecuteOperationResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetExecuteOperationResultResponseBody) SetRequestId(v string) *GetExecuteOperationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExecuteOperationResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetExecuteOperationResultResponseBodyData struct {
	// The output of the operation.
	//
	// example:
	//
	// " "
	Arguments *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	// The returned message.
	//
	// example:
	//
	// “ ”
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the operation.
	//
	// example:
	//
	// op_xxxxxxxx
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The status of the operation.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetExecuteOperationResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetExecuteOperationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExecuteOperationResultResponseBodyData) GetArguments() *string {
	return s.Arguments
}

func (s *GetExecuteOperationResultResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GetExecuteOperationResultResponseBodyData) GetOperationId() *string {
	return s.OperationId
}

func (s *GetExecuteOperationResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetExecuteOperationResultResponseBodyData) SetArguments(v string) *GetExecuteOperationResultResponseBodyData {
	s.Arguments = &v
	return s
}

func (s *GetExecuteOperationResultResponseBodyData) SetMessage(v string) *GetExecuteOperationResultResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetExecuteOperationResultResponseBodyData) SetOperationId(v string) *GetExecuteOperationResultResponseBodyData {
	s.OperationId = &v
	return s
}

func (s *GetExecuteOperationResultResponseBodyData) SetStatus(v string) *GetExecuteOperationResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetExecuteOperationResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
