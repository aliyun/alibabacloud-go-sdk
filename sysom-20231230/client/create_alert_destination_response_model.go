// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertDestinationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAlertDestinationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAlertDestinationResponse
	GetStatusCode() *int32
	SetBody(v *CreateAlertDestinationResponseBody) *CreateAlertDestinationResponse
	GetBody() *CreateAlertDestinationResponseBody
}

type CreateAlertDestinationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlertDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlertDestinationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertDestinationResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertDestinationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAlertDestinationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAlertDestinationResponse) GetBody() *CreateAlertDestinationResponseBody {
	return s.Body
}

func (s *CreateAlertDestinationResponse) SetHeaders(v map[string]*string) *CreateAlertDestinationResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertDestinationResponse) SetStatusCode(v int32) *CreateAlertDestinationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertDestinationResponse) SetBody(v *CreateAlertDestinationResponseBody) *CreateAlertDestinationResponse {
	s.Body = v
	return s
}

func (s *CreateAlertDestinationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
