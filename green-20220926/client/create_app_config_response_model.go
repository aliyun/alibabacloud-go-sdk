// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppConfigResponseBody) *CreateAppConfigResponse
	GetBody() *CreateAppConfigResponseBody
}

type CreateAppConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateAppConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppConfigResponse) GetBody() *CreateAppConfigResponseBody {
	return s.Body
}

func (s *CreateAppConfigResponse) SetHeaders(v map[string]*string) *CreateAppConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateAppConfigResponse) SetStatusCode(v int32) *CreateAppConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppConfigResponse) SetBody(v *CreateAppConfigResponseBody) *CreateAppConfigResponse {
	s.Body = v
	return s
}

func (s *CreateAppConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
