// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackendModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *DeleteBackendModelResponseBody
	GetOperationId() *string
	SetRequestId(v string) *DeleteBackendModelResponseBody
	GetRequestId() *string
}

type DeleteBackendModelResponseBody struct {
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
	// D1B18FFE-4A81-59D8-AA02-1817098977CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBackendModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackendModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackendModelResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *DeleteBackendModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBackendModelResponseBody) SetOperationId(v string) *DeleteBackendModelResponseBody {
	s.OperationId = &v
	return s
}

func (s *DeleteBackendModelResponseBody) SetRequestId(v string) *DeleteBackendModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackendModelResponseBody) Validate() error {
	return dara.Validate(s)
}
