// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelVersionLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteModelVersionLabelsResponseBody
	GetRequestId() *string
}

type DeleteModelVersionLabelsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteModelVersionLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelVersionLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelVersionLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelVersionLabelsResponseBody) SetRequestId(v string) *DeleteModelVersionLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelVersionLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
