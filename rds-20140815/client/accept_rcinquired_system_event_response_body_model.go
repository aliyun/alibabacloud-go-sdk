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
	// The ID of the request.
	//
	// example:
	//
	// 08A3B71B-FE08-4B03-974F-CC7EA6DB1828
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
