// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSearchStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSearchStrategyResponseBody
	GetRequestId() *string
}

type UpdateSearchStrategyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateSearchStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSearchStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSearchStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSearchStrategyResponseBody) SetRequestId(v string) *UpdateSearchStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSearchStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
