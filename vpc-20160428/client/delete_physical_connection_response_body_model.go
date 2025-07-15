// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePhysicalConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePhysicalConnectionResponseBody
	GetRequestId() *string
}

type DeletePhysicalConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePhysicalConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePhysicalConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePhysicalConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePhysicalConnectionResponseBody) SetRequestId(v string) *DeletePhysicalConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePhysicalConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
