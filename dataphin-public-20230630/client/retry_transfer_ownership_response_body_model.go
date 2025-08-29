// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryTransferOwnershipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RetryTransferOwnershipResponseBody
	GetCode() *string
	SetData(v int64) *RetryTransferOwnershipResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *RetryTransferOwnershipResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RetryTransferOwnershipResponseBody
	GetMessage() *string
	SetRequestId(v string) *RetryTransferOwnershipResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RetryTransferOwnershipResponseBody
	GetSuccess() *bool
}

type RetryTransferOwnershipResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RetryTransferOwnershipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryTransferOwnershipResponseBody) GoString() string {
	return s.String()
}

func (s *RetryTransferOwnershipResponseBody) GetCode() *string {
	return s.Code
}

func (s *RetryTransferOwnershipResponseBody) GetData() *int64 {
	return s.Data
}

func (s *RetryTransferOwnershipResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RetryTransferOwnershipResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RetryTransferOwnershipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryTransferOwnershipResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RetryTransferOwnershipResponseBody) SetCode(v string) *RetryTransferOwnershipResponseBody {
	s.Code = &v
	return s
}

func (s *RetryTransferOwnershipResponseBody) SetData(v int64) *RetryTransferOwnershipResponseBody {
	s.Data = &v
	return s
}

func (s *RetryTransferOwnershipResponseBody) SetHttpStatusCode(v int32) *RetryTransferOwnershipResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RetryTransferOwnershipResponseBody) SetMessage(v string) *RetryTransferOwnershipResponseBody {
	s.Message = &v
	return s
}

func (s *RetryTransferOwnershipResponseBody) SetRequestId(v string) *RetryTransferOwnershipResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryTransferOwnershipResponseBody) SetSuccess(v bool) *RetryTransferOwnershipResponseBody {
	s.Success = &v
	return s
}

func (s *RetryTransferOwnershipResponseBody) Validate() error {
	return dara.Validate(s)
}
