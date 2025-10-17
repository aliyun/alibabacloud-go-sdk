// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRecordContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopRecordContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopRecordContentResponse
	GetStatusCode() *int32
	SetBody(v *StopRecordContentResponseBody) *StopRecordContentResponse
	GetBody() *StopRecordContentResponseBody
}

type StopRecordContentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRecordContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopRecordContentResponse) String() string {
	return dara.Prettify(s)
}

func (s StopRecordContentResponse) GoString() string {
	return s.String()
}

func (s *StopRecordContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopRecordContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopRecordContentResponse) GetBody() *StopRecordContentResponseBody {
	return s.Body
}

func (s *StopRecordContentResponse) SetHeaders(v map[string]*string) *StopRecordContentResponse {
	s.Headers = v
	return s
}

func (s *StopRecordContentResponse) SetStatusCode(v int32) *StopRecordContentResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRecordContentResponse) SetBody(v *StopRecordContentResponseBody) *StopRecordContentResponse {
	s.Body = v
	return s
}

func (s *StopRecordContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
