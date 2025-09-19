// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSoarStrategyTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSoarStrategyTaskResponseBody
	GetRequestId() *string
}

type DeleteSoarStrategyTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C699E4E4-F2F4-58FC-A949-457FFE59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSoarStrategyTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSoarStrategyTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSoarStrategyTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSoarStrategyTaskResponseBody) SetRequestId(v string) *DeleteSoarStrategyTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSoarStrategyTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
