// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyPhysicalConnectionLOAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApplyPhysicalConnectionLOAResponseBody
	GetRequestId() *string
}

type ApplyPhysicalConnectionLOAResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A47BD386-7FDE-42C4-8D22-C6223D18AA1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyPhysicalConnectionLOAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyPhysicalConnectionLOAResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyPhysicalConnectionLOAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyPhysicalConnectionLOAResponseBody) SetRequestId(v string) *ApplyPhysicalConnectionLOAResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyPhysicalConnectionLOAResponseBody) Validate() error {
	return dara.Validate(s)
}
