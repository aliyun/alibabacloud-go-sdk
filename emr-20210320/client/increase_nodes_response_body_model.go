// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncreaseNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *IncreaseNodesResponseBody
	GetOperationId() *string
	SetRequestId(v string) *IncreaseNodesResponseBody
	GetRequestId() *string
}

type IncreaseNodesResponseBody struct {
	// The ID of the operation.
	//
	// example:
	//
	// op-13c37a77c505****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IncreaseNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IncreaseNodesResponseBody) GoString() string {
	return s.String()
}

func (s *IncreaseNodesResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *IncreaseNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IncreaseNodesResponseBody) SetOperationId(v string) *IncreaseNodesResponseBody {
	s.OperationId = &v
	return s
}

func (s *IncreaseNodesResponseBody) SetRequestId(v string) *IncreaseNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *IncreaseNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
