// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecursionRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v string) *AddRecursionRecordResponseBody
	GetRecordId() *string
	SetRequestId(v string) *AddRecursionRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddRecursionRecordResponseBody
	GetSuccess() *bool
}

type AddRecursionRecordResponseBody struct {
	// The ID of the DNS record.
	//
	// example:
	//
	// 173671468000010
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRecursionRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRecursionRecordResponseBody) GoString() string {
	return s.String()
}

func (s *AddRecursionRecordResponseBody) GetRecordId() *string {
	return s.RecordId
}

func (s *AddRecursionRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRecursionRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddRecursionRecordResponseBody) SetRecordId(v string) *AddRecursionRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *AddRecursionRecordResponseBody) SetRequestId(v string) *AddRecursionRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRecursionRecordResponseBody) SetSuccess(v bool) *AddRecursionRecordResponseBody {
	s.Success = &v
	return s
}

func (s *AddRecursionRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
