// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptRCInquiredSystemEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AcceptRCInquiredSystemEventResponseBody
	GetRequestId() *string
}

type AcceptRCInquiredSystemEventResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AcceptRCInquiredSystemEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AcceptRCInquiredSystemEventResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptRCInquiredSystemEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AcceptRCInquiredSystemEventResponseBody) SetRequestId(v string) *AcceptRCInquiredSystemEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcceptRCInquiredSystemEventResponseBody) Validate() error {
	return dara.Validate(s)
}
