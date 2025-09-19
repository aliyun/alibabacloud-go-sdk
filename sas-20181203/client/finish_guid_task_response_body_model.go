// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishGuidTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *FinishGuidTaskResponseBody
	GetRequestId() *string
}

type FinishGuidTaskResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 08DCAABC-82E7-5EF5-A9E7-A82DC07C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FinishGuidTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FinishGuidTaskResponseBody) GoString() string {
	return s.String()
}

func (s *FinishGuidTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FinishGuidTaskResponseBody) SetRequestId(v string) *FinishGuidTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *FinishGuidTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
