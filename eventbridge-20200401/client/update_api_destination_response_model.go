// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiDestinationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApiDestinationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApiDestinationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApiDestinationResponseBody) *UpdateApiDestinationResponse
	GetBody() *UpdateApiDestinationResponseBody
}

type UpdateApiDestinationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApiDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApiDestinationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiDestinationResponse) GoString() string {
	return s.String()
}

func (s *UpdateApiDestinationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApiDestinationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApiDestinationResponse) GetBody() *UpdateApiDestinationResponseBody {
	return s.Body
}

func (s *UpdateApiDestinationResponse) SetHeaders(v map[string]*string) *UpdateApiDestinationResponse {
	s.Headers = v
	return s
}

func (s *UpdateApiDestinationResponse) SetStatusCode(v int32) *UpdateApiDestinationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApiDestinationResponse) SetBody(v *UpdateApiDestinationResponseBody) *UpdateApiDestinationResponse {
	s.Body = v
	return s
}

func (s *UpdateApiDestinationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
