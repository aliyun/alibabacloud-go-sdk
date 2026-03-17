// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSmartAGDpiMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableSmartAGDpiMonitorResponseBody
	GetRequestId() *string
}

type DisableSmartAGDpiMonitorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 64966488-B3E3-41E2-9570-4596117EC12E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableSmartAGDpiMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSmartAGDpiMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSmartAGDpiMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSmartAGDpiMonitorResponseBody) SetRequestId(v string) *DisableSmartAGDpiMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSmartAGDpiMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
