// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSecondApplyPhysicalConnectionLOAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SecondApplyPhysicalConnectionLOAResponseBody
	GetRequestId() *string
}

type SecondApplyPhysicalConnectionLOAResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A47BD386-7FDE-42C4-8D22-C6223D18AA1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SecondApplyPhysicalConnectionLOAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SecondApplyPhysicalConnectionLOAResponseBody) GoString() string {
	return s.String()
}

func (s *SecondApplyPhysicalConnectionLOAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SecondApplyPhysicalConnectionLOAResponseBody) SetRequestId(v string) *SecondApplyPhysicalConnectionLOAResponseBody {
	s.RequestId = &v
	return s
}

func (s *SecondApplyPhysicalConnectionLOAResponseBody) Validate() error {
	return dara.Validate(s)
}
