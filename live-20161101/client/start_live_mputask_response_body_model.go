// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLiveMPUTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartLiveMPUTaskResponseBody
	GetRequestId() *string
}

type StartLiveMPUTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0F72851F-5DC1-1979-9B2C-450040316C3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartLiveMPUTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartLiveMPUTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartLiveMPUTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartLiveMPUTaskResponseBody) SetRequestId(v string) *StartLiveMPUTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartLiveMPUTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
