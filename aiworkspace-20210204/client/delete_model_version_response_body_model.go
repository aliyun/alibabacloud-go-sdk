// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteModelVersionResponseBody
	GetRequestId() *string
}

type DeleteModelVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteModelVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelVersionResponseBody) SetRequestId(v string) *DeleteModelVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
