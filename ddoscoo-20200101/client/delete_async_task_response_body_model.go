// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAsyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAsyncTaskResponseBody
	GetRequestId() *string
}

type DeleteAsyncTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAsyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAsyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAsyncTaskResponseBody) SetRequestId(v string) *DeleteAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAsyncTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
