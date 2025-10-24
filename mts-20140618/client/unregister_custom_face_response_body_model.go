// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnregisterCustomFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnregisterCustomFaceResponseBody
	GetRequestId() *string
}

type UnregisterCustomFaceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1A3347BF-7BCE-40A6-B33E-43C2B8A9A278
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnregisterCustomFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnregisterCustomFaceResponseBody) GoString() string {
	return s.String()
}

func (s *UnregisterCustomFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnregisterCustomFaceResponseBody) SetRequestId(v string) *UnregisterCustomFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnregisterCustomFaceResponseBody) Validate() error {
	return dara.Validate(s)
}
