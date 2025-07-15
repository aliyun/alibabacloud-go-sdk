// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeployApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *BatchDeployApisResponseBody
	GetOperationId() *string
	SetRequestId(v string) *BatchDeployApisResponseBody
	GetRequestId() *string
}

type BatchDeployApisResponseBody struct {
	// The ID of the operation.
	//
	// example:
	//
	// 2a322599-8e38-428a-a306-9b21ea2129bf
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E7FE7172-AA75-5880-B6F7-C00893E9BC06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeployApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeployApisResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeployApisResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *BatchDeployApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeployApisResponseBody) SetOperationId(v string) *BatchDeployApisResponseBody {
	s.OperationId = &v
	return s
}

func (s *BatchDeployApisResponseBody) SetRequestId(v string) *BatchDeployApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeployApisResponseBody) Validate() error {
	return dara.Validate(s)
}
