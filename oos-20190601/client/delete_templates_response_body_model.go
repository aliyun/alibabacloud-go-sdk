// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTemplatesResponseBody
	GetRequestId() *string
}

type DeleteTemplatesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2075899A-585D-4A41-A9B2-28DA8534
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTemplatesResponseBody) SetRequestId(v string) *DeleteTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}
