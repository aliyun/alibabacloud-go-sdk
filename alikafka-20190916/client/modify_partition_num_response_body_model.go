// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPartitionNumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ModifyPartitionNumResponseBody
	GetCode() *int32
	SetMessage(v string) *ModifyPartitionNumResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyPartitionNumResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyPartitionNumResponseBody
	GetSuccess() *bool
}

type ModifyPartitionNumResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B7A39AE5-0B36-4442-A304-E088526****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyPartitionNumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPartitionNumResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPartitionNumResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ModifyPartitionNumResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyPartitionNumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPartitionNumResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyPartitionNumResponseBody) SetCode(v int32) *ModifyPartitionNumResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyPartitionNumResponseBody) SetMessage(v string) *ModifyPartitionNumResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyPartitionNumResponseBody) SetRequestId(v string) *ModifyPartitionNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPartitionNumResponseBody) SetSuccess(v bool) *ModifyPartitionNumResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyPartitionNumResponseBody) Validate() error {
	return dara.Validate(s)
}
