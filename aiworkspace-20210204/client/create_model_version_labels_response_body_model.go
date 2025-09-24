// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelVersionLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateModelVersionLabelsResponseBody
	GetRequestId() *string
}

type CreateModelVersionLabelsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateModelVersionLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelVersionLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelVersionLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelVersionLabelsResponseBody) SetRequestId(v string) *CreateModelVersionLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelVersionLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
