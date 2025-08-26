// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeTrainTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeTrainTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeTrainTicketResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeTrainTicketResponseBody) *RecognizeTrainTicketResponse
	GetBody() *RecognizeTrainTicketResponseBody
}

type RecognizeTrainTicketResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeTrainTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTrainTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTrainTicketResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTrainTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeTrainTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeTrainTicketResponse) GetBody() *RecognizeTrainTicketResponseBody {
	return s.Body
}

func (s *RecognizeTrainTicketResponse) SetHeaders(v map[string]*string) *RecognizeTrainTicketResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTrainTicketResponse) SetStatusCode(v int32) *RecognizeTrainTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTrainTicketResponse) SetBody(v *RecognizeTrainTicketResponseBody) *RecognizeTrainTicketResponse {
	s.Body = v
	return s
}

func (s *RecognizeTrainTicketResponse) Validate() error {
	return dara.Validate(s)
}
