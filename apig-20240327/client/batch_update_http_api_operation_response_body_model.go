// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateHttpApiOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchUpdateHttpApiOperationResponseBody
	GetCode() *string
	SetMessage(v string) *BatchUpdateHttpApiOperationResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchUpdateHttpApiOperationResponseBody
	GetRequestId() *string
}

type BatchUpdateHttpApiOperationResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4CF2E0A5-xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchUpdateHttpApiOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateHttpApiOperationResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateHttpApiOperationResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchUpdateHttpApiOperationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchUpdateHttpApiOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUpdateHttpApiOperationResponseBody) SetCode(v string) *BatchUpdateHttpApiOperationResponseBody {
	s.Code = &v
	return s
}

func (s *BatchUpdateHttpApiOperationResponseBody) SetMessage(v string) *BatchUpdateHttpApiOperationResponseBody {
	s.Message = &v
	return s
}

func (s *BatchUpdateHttpApiOperationResponseBody) SetRequestId(v string) *BatchUpdateHttpApiOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUpdateHttpApiOperationResponseBody) Validate() error {
	return dara.Validate(s)
}
