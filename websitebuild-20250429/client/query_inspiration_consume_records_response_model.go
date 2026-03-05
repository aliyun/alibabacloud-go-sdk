// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInspirationConsumeRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInspirationConsumeRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInspirationConsumeRecordsResponse
	GetStatusCode() *int32
	SetBody(v *QueryInspirationConsumeRecordsResponseBody) *QueryInspirationConsumeRecordsResponse
	GetBody() *QueryInspirationConsumeRecordsResponseBody
}

type QueryInspirationConsumeRecordsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInspirationConsumeRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInspirationConsumeRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationConsumeRecordsResponse) GoString() string {
	return s.String()
}

func (s *QueryInspirationConsumeRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInspirationConsumeRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInspirationConsumeRecordsResponse) GetBody() *QueryInspirationConsumeRecordsResponseBody {
	return s.Body
}

func (s *QueryInspirationConsumeRecordsResponse) SetHeaders(v map[string]*string) *QueryInspirationConsumeRecordsResponse {
	s.Headers = v
	return s
}

func (s *QueryInspirationConsumeRecordsResponse) SetStatusCode(v int32) *QueryInspirationConsumeRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponse) SetBody(v *QueryInspirationConsumeRecordsResponseBody) *QueryInspirationConsumeRecordsResponse {
	s.Body = v
	return s
}

func (s *QueryInspirationConsumeRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
