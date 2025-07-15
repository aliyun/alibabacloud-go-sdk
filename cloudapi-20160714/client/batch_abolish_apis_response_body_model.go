// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAbolishApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *BatchAbolishApisResponseBody
	GetOperationId() *string
	SetRequestId(v string) *BatchAbolishApisResponseBody
	GetRequestId() *string
}

type BatchAbolishApisResponseBody struct {
	// The ID of the operation.
	//
	// example:
	//
	// f7834d74be4e41aa8e607b0fafae9b33
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E7FE7172-AA75-5880-B6F7-C00893E9BC06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchAbolishApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchAbolishApisResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAbolishApisResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *BatchAbolishApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchAbolishApisResponseBody) SetOperationId(v string) *BatchAbolishApisResponseBody {
	s.OperationId = &v
	return s
}

func (s *BatchAbolishApisResponseBody) SetRequestId(v string) *BatchAbolishApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchAbolishApisResponseBody) Validate() error {
	return dara.Validate(s)
}
