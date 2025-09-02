// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMetaPartitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CheckMetaPartitionResponseBody
	GetData() *bool
	SetErrorCode(v string) *CheckMetaPartitionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CheckMetaPartitionResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CheckMetaPartitionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CheckMetaPartitionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckMetaPartitionResponseBody
	GetSuccess() *bool
}

type CheckMetaPartitionResponseBody struct {
	// Indicates whether the partition in the MaxCompute metatable exists. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckMetaPartitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckMetaPartitionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckMetaPartitionResponseBody) GetData() *bool {
	return s.Data
}

func (s *CheckMetaPartitionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CheckMetaPartitionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CheckMetaPartitionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckMetaPartitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckMetaPartitionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckMetaPartitionResponseBody) SetData(v bool) *CheckMetaPartitionResponseBody {
	s.Data = &v
	return s
}

func (s *CheckMetaPartitionResponseBody) SetErrorCode(v string) *CheckMetaPartitionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckMetaPartitionResponseBody) SetErrorMessage(v string) *CheckMetaPartitionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CheckMetaPartitionResponseBody) SetHttpStatusCode(v int32) *CheckMetaPartitionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckMetaPartitionResponseBody) SetRequestId(v string) *CheckMetaPartitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckMetaPartitionResponseBody) SetSuccess(v bool) *CheckMetaPartitionResponseBody {
	s.Success = &v
	return s
}

func (s *CheckMetaPartitionResponseBody) Validate() error {
	return dara.Validate(s)
}
