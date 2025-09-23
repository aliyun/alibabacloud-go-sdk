// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceShareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteResourceShareResponseBody
	GetRequestId() *string
}

type DeleteResourceShareResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A627EE2A-223D-4E1F-A954-394686AEA916
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceShareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceShareResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceShareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceShareResponseBody) SetRequestId(v string) *DeleteResourceShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceShareResponseBody) Validate() error {
	return dara.Validate(s)
}
