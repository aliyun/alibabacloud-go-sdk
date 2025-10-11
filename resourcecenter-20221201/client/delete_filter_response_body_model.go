// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFilterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFilterResponseBody
	GetRequestId() *string
}

type DeleteFilterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A4A63E3C-89EC-51F9-9934-C9AF1BCBAAA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFilterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFilterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFilterResponseBody) SetRequestId(v string) *DeleteFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFilterResponseBody) Validate() error {
	return dara.Validate(s)
}
