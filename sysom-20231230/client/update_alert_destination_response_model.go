// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertDestinationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlertDestinationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlertDestinationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlertDestinationResponseBody) *UpdateAlertDestinationResponse
	GetBody() *UpdateAlertDestinationResponseBody
}

type UpdateAlertDestinationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlertDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlertDestinationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertDestinationResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertDestinationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlertDestinationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlertDestinationResponse) GetBody() *UpdateAlertDestinationResponseBody {
	return s.Body
}

func (s *UpdateAlertDestinationResponse) SetHeaders(v map[string]*string) *UpdateAlertDestinationResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertDestinationResponse) SetStatusCode(v int32) *UpdateAlertDestinationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertDestinationResponse) SetBody(v *UpdateAlertDestinationResponseBody) *UpdateAlertDestinationResponse {
	s.Body = v
	return s
}

func (s *UpdateAlertDestinationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
