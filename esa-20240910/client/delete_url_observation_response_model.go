// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUrlObservationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUrlObservationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUrlObservationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUrlObservationResponseBody) *DeleteUrlObservationResponse
	GetBody() *DeleteUrlObservationResponseBody
}

type DeleteUrlObservationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUrlObservationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUrlObservationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUrlObservationResponse) GoString() string {
	return s.String()
}

func (s *DeleteUrlObservationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUrlObservationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUrlObservationResponse) GetBody() *DeleteUrlObservationResponseBody {
	return s.Body
}

func (s *DeleteUrlObservationResponse) SetHeaders(v map[string]*string) *DeleteUrlObservationResponse {
	s.Headers = v
	return s
}

func (s *DeleteUrlObservationResponse) SetStatusCode(v int32) *DeleteUrlObservationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUrlObservationResponse) SetBody(v *DeleteUrlObservationResponseBody) *DeleteUrlObservationResponse {
	s.Body = v
	return s
}

func (s *DeleteUrlObservationResponse) Validate() error {
	return dara.Validate(s)
}
