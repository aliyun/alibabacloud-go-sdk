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
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
