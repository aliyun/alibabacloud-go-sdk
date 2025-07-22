// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMPUTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartMPUTaskResponseBody
	GetRequestId() *string
}

type StartMPUTaskResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartMPUTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartMPUTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartMPUTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartMPUTaskResponseBody) SetRequestId(v string) *StartMPUTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartMPUTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
