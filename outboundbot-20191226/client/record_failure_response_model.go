// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecordFailureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecordFailureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecordFailureResponse
	GetStatusCode() *int32
	SetBody(v *RecordFailureResponseBody) *RecordFailureResponse
	GetBody() *RecordFailureResponseBody
}

type RecordFailureResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecordFailureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecordFailureResponse) String() string {
	return dara.Prettify(s)
}

func (s RecordFailureResponse) GoString() string {
	return s.String()
}

func (s *RecordFailureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecordFailureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecordFailureResponse) GetBody() *RecordFailureResponseBody {
	return s.Body
}

func (s *RecordFailureResponse) SetHeaders(v map[string]*string) *RecordFailureResponse {
	s.Headers = v
	return s
}

func (s *RecordFailureResponse) SetStatusCode(v int32) *RecordFailureResponse {
	s.StatusCode = &v
	return s
}

func (s *RecordFailureResponse) SetBody(v *RecordFailureResponseBody) *RecordFailureResponse {
	s.Body = v
	return s
}

func (s *RecordFailureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
