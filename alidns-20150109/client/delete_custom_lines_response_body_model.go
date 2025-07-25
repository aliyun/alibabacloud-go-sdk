// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomLinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomLinesResponseBody
	GetRequestId() *string
}

type DeleteCustomLinesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomLinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomLinesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomLinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomLinesResponseBody) SetRequestId(v string) *DeleteCustomLinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomLinesResponseBody) Validate() error {
	return dara.Validate(s)
}
