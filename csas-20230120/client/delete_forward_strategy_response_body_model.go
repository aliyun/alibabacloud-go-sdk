// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteForwardStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteForwardStrategyResponseBody
	GetRequestId() *string
}

type DeleteForwardStrategyResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5A232DC6-0A2E-5754-B9E6-C9A8E9EF784A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteForwardStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteForwardStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteForwardStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteForwardStrategyResponseBody) SetRequestId(v string) *DeleteForwardStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteForwardStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
