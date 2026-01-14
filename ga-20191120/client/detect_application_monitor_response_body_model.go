// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectApplicationMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetectApplicationMonitorResponseBody
	GetRequestId() *string
}

type DetectApplicationMonitorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectApplicationMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectApplicationMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DetectApplicationMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectApplicationMonitorResponseBody) SetRequestId(v string) *DetectApplicationMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectApplicationMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
