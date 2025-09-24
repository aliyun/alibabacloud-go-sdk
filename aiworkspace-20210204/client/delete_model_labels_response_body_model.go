// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteModelLabelsResponseBody
	GetRequestId() *string
}

type DeleteModelLabelsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteModelLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelLabelsResponseBody) SetRequestId(v string) *DeleteModelLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
