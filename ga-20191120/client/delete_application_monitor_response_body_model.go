// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApplicationMonitorResponseBody
	GetRequestId() *string
}

type DeleteApplicationMonitorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApplicationMonitorResponseBody) SetRequestId(v string) *DeleteApplicationMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
