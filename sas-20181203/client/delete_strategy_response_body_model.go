// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteStrategyResponseBody
	GetRequestId() *string
}

type DeleteStrategyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB3936FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStrategyResponseBody) SetRequestId(v string) *DeleteStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
