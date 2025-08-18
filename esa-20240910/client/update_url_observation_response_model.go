// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUrlObservationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUrlObservationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUrlObservationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUrlObservationResponseBody) *UpdateUrlObservationResponse
	GetBody() *UpdateUrlObservationResponseBody
}

type UpdateUrlObservationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUrlObservationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUrlObservationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUrlObservationResponse) GoString() string {
	return s.String()
}

func (s *UpdateUrlObservationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUrlObservationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUrlObservationResponse) GetBody() *UpdateUrlObservationResponseBody {
	return s.Body
}

func (s *UpdateUrlObservationResponse) SetHeaders(v map[string]*string) *UpdateUrlObservationResponse {
	s.Headers = v
	return s
}

func (s *UpdateUrlObservationResponse) SetStatusCode(v int32) *UpdateUrlObservationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUrlObservationResponse) SetBody(v *UpdateUrlObservationResponseBody) *UpdateUrlObservationResponse {
	s.Body = v
	return s
}

func (s *UpdateUrlObservationResponse) Validate() error {
	return dara.Validate(s)
}
