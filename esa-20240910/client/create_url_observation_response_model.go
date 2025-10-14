// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUrlObservationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUrlObservationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUrlObservationResponse
	GetStatusCode() *int32
	SetBody(v *CreateUrlObservationResponseBody) *CreateUrlObservationResponse
	GetBody() *CreateUrlObservationResponseBody
}

type CreateUrlObservationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUrlObservationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUrlObservationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUrlObservationResponse) GoString() string {
	return s.String()
}

func (s *CreateUrlObservationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUrlObservationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUrlObservationResponse) GetBody() *CreateUrlObservationResponseBody {
	return s.Body
}

func (s *CreateUrlObservationResponse) SetHeaders(v map[string]*string) *CreateUrlObservationResponse {
	s.Headers = v
	return s
}

func (s *CreateUrlObservationResponse) SetStatusCode(v int32) *CreateUrlObservationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUrlObservationResponse) SetBody(v *CreateUrlObservationResponseBody) *CreateUrlObservationResponse {
	s.Body = v
	return s
}

func (s *CreateUrlObservationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
