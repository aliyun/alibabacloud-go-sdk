// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiDestinationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApiDestinationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApiDestinationResponse
	GetStatusCode() *int32
	SetBody(v *CreateApiDestinationResponseBody) *CreateApiDestinationResponse
	GetBody() *CreateApiDestinationResponseBody
}

type CreateApiDestinationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApiDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApiDestinationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApiDestinationResponse) GoString() string {
	return s.String()
}

func (s *CreateApiDestinationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApiDestinationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApiDestinationResponse) GetBody() *CreateApiDestinationResponseBody {
	return s.Body
}

func (s *CreateApiDestinationResponse) SetHeaders(v map[string]*string) *CreateApiDestinationResponse {
	s.Headers = v
	return s
}

func (s *CreateApiDestinationResponse) SetStatusCode(v int32) *CreateApiDestinationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApiDestinationResponse) SetBody(v *CreateApiDestinationResponseBody) *CreateApiDestinationResponse {
	s.Body = v
	return s
}

func (s *CreateApiDestinationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
