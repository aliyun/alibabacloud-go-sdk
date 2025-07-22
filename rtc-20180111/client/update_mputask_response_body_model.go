// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMPUTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMPUTaskResponseBody
	GetRequestId() *string
}

type UpdateMPUTaskResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMPUTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMPUTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMPUTaskResponseBody) SetRequestId(v string) *UpdateMPUTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMPUTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
