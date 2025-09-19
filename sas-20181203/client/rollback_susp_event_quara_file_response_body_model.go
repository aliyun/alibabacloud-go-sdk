// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackSuspEventQuaraFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RollbackSuspEventQuaraFileResponseBody
	GetRequestId() *string
}

type RollbackSuspEventQuaraFileResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 52A3AEE6-114A-499D-8990-4BA9B27FE0AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackSuspEventQuaraFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackSuspEventQuaraFileResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackSuspEventQuaraFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackSuspEventQuaraFileResponseBody) SetRequestId(v string) *RollbackSuspEventQuaraFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackSuspEventQuaraFileResponseBody) Validate() error {
	return dara.Validate(s)
}
