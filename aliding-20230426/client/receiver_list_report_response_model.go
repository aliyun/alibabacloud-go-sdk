// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceiverListReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReceiverListReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReceiverListReportResponse
	GetStatusCode() *int32
	SetBody(v *ReceiverListReportResponseBody) *ReceiverListReportResponse
	GetBody() *ReceiverListReportResponseBody
}

type ReceiverListReportResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReceiverListReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReceiverListReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ReceiverListReportResponse) GoString() string {
	return s.String()
}

func (s *ReceiverListReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReceiverListReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReceiverListReportResponse) GetBody() *ReceiverListReportResponseBody {
	return s.Body
}

func (s *ReceiverListReportResponse) SetHeaders(v map[string]*string) *ReceiverListReportResponse {
	s.Headers = v
	return s
}

func (s *ReceiverListReportResponse) SetStatusCode(v int32) *ReceiverListReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ReceiverListReportResponse) SetBody(v *ReceiverListReportResponseBody) *ReceiverListReportResponse {
	s.Body = v
	return s
}

func (s *ReceiverListReportResponse) Validate() error {
	return dara.Validate(s)
}
