// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateBatchResponseBody
	GetRequestId() *string
}

type UpdateBatchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CB4D73A3-BAF4-4A9D-A631-15F219AF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBatchResponseBody) SetRequestId(v string) *UpdateBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
