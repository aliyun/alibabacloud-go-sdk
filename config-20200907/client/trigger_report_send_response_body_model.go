// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerReportSendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TriggerReportSendResponseBody
	GetRequestId() *string
}

type TriggerReportSendResponseBody struct {
	// example:
	//
	// DE9FFFE5-FCAD-4B24-9546-BF49273C562B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TriggerReportSendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TriggerReportSendResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerReportSendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TriggerReportSendResponseBody) SetRequestId(v string) *TriggerReportSendResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerReportSendResponseBody) Validate() error {
	return dara.Validate(s)
}
