// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBatchResponseBody
	GetRequestId() *string
}

type DeleteBatchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 91AC8C98-0F36-49D2-8290-742E24******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBatchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBatchResponseBody) SetRequestId(v string) *DeleteBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
