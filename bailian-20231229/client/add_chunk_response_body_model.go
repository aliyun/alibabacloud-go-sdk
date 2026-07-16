// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddChunkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddChunkResponseBody
	GetCode() *string
	SetData(v bool) *AddChunkResponseBody
	GetData() *bool
	SetMessage(v string) *AddChunkResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddChunkResponseBody
	GetRequestId() *string
	SetStatus(v string) *AddChunkResponseBody
	GetStatus() *string
	SetSuccess(v bool) *AddChunkResponseBody
	GetSuccess() *bool
}

type AddChunkResponseBody struct {
	// The error status code.
	//
	// example:
	//
	// Index.InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned upon a successful request.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 35A267BF-xxxx-54DB-8394-AA3B0742D833
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code returned by the operation.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddChunkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddChunkResponseBody) GoString() string {
	return s.String()
}

func (s *AddChunkResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddChunkResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddChunkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddChunkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddChunkResponseBody) GetStatus() *string {
	return s.Status
}

func (s *AddChunkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddChunkResponseBody) SetCode(v string) *AddChunkResponseBody {
	s.Code = &v
	return s
}

func (s *AddChunkResponseBody) SetData(v bool) *AddChunkResponseBody {
	s.Data = &v
	return s
}

func (s *AddChunkResponseBody) SetMessage(v string) *AddChunkResponseBody {
	s.Message = &v
	return s
}

func (s *AddChunkResponseBody) SetRequestId(v string) *AddChunkResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddChunkResponseBody) SetStatus(v string) *AddChunkResponseBody {
	s.Status = &v
	return s
}

func (s *AddChunkResponseBody) SetSuccess(v bool) *AddChunkResponseBody {
	s.Success = &v
	return s
}

func (s *AddChunkResponseBody) Validate() error {
	return dara.Validate(s)
}
