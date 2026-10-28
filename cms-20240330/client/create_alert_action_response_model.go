// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertActionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAlertActionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAlertActionResponse
	GetStatusCode() *int32
	SetBody(v *CreateAlertActionResponseBody) *CreateAlertActionResponse
	GetBody() *CreateAlertActionResponseBody
}

type CreateAlertActionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlertActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlertActionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertActionResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertActionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAlertActionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAlertActionResponse) GetBody() *CreateAlertActionResponseBody {
	return s.Body
}

func (s *CreateAlertActionResponse) SetHeaders(v map[string]*string) *CreateAlertActionResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertActionResponse) SetStatusCode(v int32) *CreateAlertActionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertActionResponse) SetBody(v *CreateAlertActionResponseBody) *CreateAlertActionResponse {
	s.Body = v
	return s
}

func (s *CreateAlertActionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
