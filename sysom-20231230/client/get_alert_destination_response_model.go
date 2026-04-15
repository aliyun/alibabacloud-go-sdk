// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertDestinationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlertDestinationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlertDestinationResponse
	GetStatusCode() *int32
	SetBody(v *GetAlertDestinationResponseBody) *GetAlertDestinationResponse
	GetBody() *GetAlertDestinationResponseBody
}

type GetAlertDestinationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlertDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlertDestinationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlertDestinationResponse) GoString() string {
	return s.String()
}

func (s *GetAlertDestinationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlertDestinationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlertDestinationResponse) GetBody() *GetAlertDestinationResponseBody {
	return s.Body
}

func (s *GetAlertDestinationResponse) SetHeaders(v map[string]*string) *GetAlertDestinationResponse {
	s.Headers = v
	return s
}

func (s *GetAlertDestinationResponse) SetStatusCode(v int32) *GetAlertDestinationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlertDestinationResponse) SetBody(v *GetAlertDestinationResponseBody) *GetAlertDestinationResponse {
	s.Body = v
	return s
}

func (s *GetAlertDestinationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
