// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhiteListStrategyStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWhiteListStrategyStatusResponseBody
	GetRequestId() *string
}

type UpdateWhiteListStrategyStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 379a9b8f-107b-4630-9e95-2299a1ea****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWhiteListStrategyStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhiteListStrategyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWhiteListStrategyStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWhiteListStrategyStatusResponseBody) SetRequestId(v string) *UpdateWhiteListStrategyStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWhiteListStrategyStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
