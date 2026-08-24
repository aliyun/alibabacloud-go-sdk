// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendDiagnosticInterruptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendDiagnosticInterruptResponseBody
	GetRequestId() *string
}

type SendDiagnosticInterruptResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendDiagnosticInterruptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendDiagnosticInterruptResponseBody) GoString() string {
	return s.String()
}

func (s *SendDiagnosticInterruptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendDiagnosticInterruptResponseBody) SetRequestId(v string) *SendDiagnosticInterruptResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendDiagnosticInterruptResponseBody) Validate() error {
	return dara.Validate(s)
}
