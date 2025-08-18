// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduledPreloadExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteScheduledPreloadExecutionResponseBody
	GetRequestId() *string
}

type DeleteScheduledPreloadExecutionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScheduledPreloadExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledPreloadExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPreloadExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScheduledPreloadExecutionResponseBody) SetRequestId(v string) *DeleteScheduledPreloadExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScheduledPreloadExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
