// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *NotifyExecutionResponseBody
	GetRequestId() *string
}

type NotifyExecutionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 491DF8C2-34C9-4679-9DB3-4C0F49B129AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s NotifyExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s NotifyExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *NotifyExecutionResponseBody) SetRequestId(v string) *NotifyExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *NotifyExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
