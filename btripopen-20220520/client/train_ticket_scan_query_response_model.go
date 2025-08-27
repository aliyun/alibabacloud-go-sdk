// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainTicketScanQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainTicketScanQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainTicketScanQueryResponse
	GetStatusCode() *int32
	SetBody(v *TrainTicketScanQueryResponseBody) *TrainTicketScanQueryResponse
	GetBody() *TrainTicketScanQueryResponseBody
}

type TrainTicketScanQueryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainTicketScanQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainTicketScanQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainTicketScanQueryResponse) GoString() string {
	return s.String()
}

func (s *TrainTicketScanQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainTicketScanQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainTicketScanQueryResponse) GetBody() *TrainTicketScanQueryResponseBody {
	return s.Body
}

func (s *TrainTicketScanQueryResponse) SetHeaders(v map[string]*string) *TrainTicketScanQueryResponse {
	s.Headers = v
	return s
}

func (s *TrainTicketScanQueryResponse) SetStatusCode(v int32) *TrainTicketScanQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainTicketScanQueryResponse) SetBody(v *TrainTicketScanQueryResponseBody) *TrainTicketScanQueryResponse {
	s.Body = v
	return s
}

func (s *TrainTicketScanQueryResponse) Validate() error {
	return dara.Validate(s)
}
