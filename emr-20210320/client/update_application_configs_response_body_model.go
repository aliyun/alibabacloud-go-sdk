// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *UpdateApplicationConfigsResponseBody
	GetOperationId() *string
	SetRequestId(v string) *UpdateApplicationConfigsResponseBody
	GetRequestId() *string
}

type UpdateApplicationConfigsResponseBody struct {
	// The operation ID.
	//
	// example:
	//
	// op-13c37a77c505****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9E3A7161-EB7B-172B-8D18-FFB06BA3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationConfigsResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *UpdateApplicationConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationConfigsResponseBody) SetOperationId(v string) *UpdateApplicationConfigsResponseBody {
	s.OperationId = &v
	return s
}

func (s *UpdateApplicationConfigsResponseBody) SetRequestId(v string) *UpdateApplicationConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
