// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCycleTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCycleTaskResponseBody
	GetRequestId() *string
}

type DeleteCycleTaskResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7DBB3D54-AF29-5BF4-8B44-9CFA94F50****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCycleTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCycleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCycleTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCycleTaskResponseBody) SetRequestId(v string) *DeleteCycleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCycleTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
