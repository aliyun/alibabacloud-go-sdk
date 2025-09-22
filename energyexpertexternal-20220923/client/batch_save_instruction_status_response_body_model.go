// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSaveInstructionStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *BatchSaveInstructionStatusResponseBody
	GetData() *bool
	SetRequestId(v string) *BatchSaveInstructionStatusResponseBody
	GetRequestId() *string
}

type BatchSaveInstructionStatusResponseBody struct {
	// true
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchSaveInstructionStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchSaveInstructionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSaveInstructionStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *BatchSaveInstructionStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchSaveInstructionStatusResponseBody) SetData(v bool) *BatchSaveInstructionStatusResponseBody {
	s.Data = &v
	return s
}

func (s *BatchSaveInstructionStatusResponseBody) SetRequestId(v string) *BatchSaveInstructionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSaveInstructionStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
