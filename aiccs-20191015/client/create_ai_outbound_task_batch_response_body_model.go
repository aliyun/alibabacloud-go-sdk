// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiOutboundTaskBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAiOutboundTaskBatchResponseBody
	GetCode() *string
	SetData(v int32) *CreateAiOutboundTaskBatchResponseBody
	GetData() *int32
	SetMessage(v string) *CreateAiOutboundTaskBatchResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAiOutboundTaskBatchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAiOutboundTaskBatchResponseBody
	GetSuccess() *bool
}

type CreateAiOutboundTaskBatchResponseBody struct {
	// The request status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The version ID of the newly created job batch.
	//
	// example:
	//
	// 123456
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAiOutboundTaskBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAiOutboundTaskBatchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAiOutboundTaskBatchResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAiOutboundTaskBatchResponseBody) GetData() *int32 {
	return s.Data
}

func (s *CreateAiOutboundTaskBatchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAiOutboundTaskBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAiOutboundTaskBatchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAiOutboundTaskBatchResponseBody) SetCode(v string) *CreateAiOutboundTaskBatchResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAiOutboundTaskBatchResponseBody) SetData(v int32) *CreateAiOutboundTaskBatchResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAiOutboundTaskBatchResponseBody) SetMessage(v string) *CreateAiOutboundTaskBatchResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAiOutboundTaskBatchResponseBody) SetRequestId(v string) *CreateAiOutboundTaskBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAiOutboundTaskBatchResponseBody) SetSuccess(v bool) *CreateAiOutboundTaskBatchResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAiOutboundTaskBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
