// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCsrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCsrResponseBody
	GetRequestId() *string
}

type UpdateCsrResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 082FAB35-6AB9-4FD5-8750-D36673548E76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCsrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCsrResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCsrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCsrResponseBody) SetRequestId(v string) *UpdateCsrResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCsrResponseBody) Validate() error {
	return dara.Validate(s)
}
