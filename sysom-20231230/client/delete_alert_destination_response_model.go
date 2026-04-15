// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertDestinationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlertDestinationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlertDestinationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlertDestinationResponseBody) *DeleteAlertDestinationResponse
	GetBody() *DeleteAlertDestinationResponseBody
}

type DeleteAlertDestinationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlertDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlertDestinationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertDestinationResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertDestinationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlertDestinationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlertDestinationResponse) GetBody() *DeleteAlertDestinationResponseBody {
	return s.Body
}

func (s *DeleteAlertDestinationResponse) SetHeaders(v map[string]*string) *DeleteAlertDestinationResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertDestinationResponse) SetStatusCode(v int32) *DeleteAlertDestinationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertDestinationResponse) SetBody(v *DeleteAlertDestinationResponseBody) *DeleteAlertDestinationResponse {
	s.Body = v
	return s
}

func (s *DeleteAlertDestinationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
