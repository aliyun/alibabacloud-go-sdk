// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecreaseNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *DecreaseNodesResponseBody
	GetOperationId() *string
	SetRequestId(v string) *DecreaseNodesResponseBody
	GetRequestId() *string
}

type DecreaseNodesResponseBody struct {
	// Operation ID.
	//
	// example:
	//
	// op-13c37a77c505****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 44430037-E59A-3E66-A2B0-97D155346F22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DecreaseNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DecreaseNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DecreaseNodesResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *DecreaseNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DecreaseNodesResponseBody) SetOperationId(v string) *DecreaseNodesResponseBody {
	s.OperationId = &v
	return s
}

func (s *DecreaseNodesResponseBody) SetRequestId(v string) *DecreaseNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DecreaseNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
