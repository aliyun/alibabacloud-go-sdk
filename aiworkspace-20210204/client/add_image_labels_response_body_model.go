// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddImageLabelsResponseBody
	GetRequestId() *string
}

type AddImageLabelsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddImageLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddImageLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddImageLabelsResponseBody) SetRequestId(v string) *AddImageLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImageLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
