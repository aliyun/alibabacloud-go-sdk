// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDestroyInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DestroyInstanceResponseBody
	GetRequestId() *string
}

type DestroyInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 870422DE-B64C-5A15-9C48-4A9485D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DestroyInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DestroyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DestroyInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DestroyInstanceResponseBody) SetRequestId(v string) *DestroyInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DestroyInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
