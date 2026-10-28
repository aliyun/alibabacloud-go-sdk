// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertActionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlertActionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlertActionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlertActionResponseBody) *UpdateAlertActionResponse
	GetBody() *UpdateAlertActionResponseBody
}

type UpdateAlertActionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlertActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlertActionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertActionResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertActionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlertActionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlertActionResponse) GetBody() *UpdateAlertActionResponseBody {
	return s.Body
}

func (s *UpdateAlertActionResponse) SetHeaders(v map[string]*string) *UpdateAlertActionResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertActionResponse) SetStatusCode(v int32) *UpdateAlertActionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertActionResponse) SetBody(v *UpdateAlertActionResponseBody) *UpdateAlertActionResponse {
	s.Body = v
	return s
}

func (s *UpdateAlertActionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
