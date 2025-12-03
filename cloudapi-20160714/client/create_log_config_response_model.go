// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateLogConfigResponseBody) *CreateLogConfigResponse
	GetBody() *CreateLogConfigResponseBody
}

type CreateLogConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLogConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLogConfigResponse) GetBody() *CreateLogConfigResponseBody {
	return s.Body
}

func (s *CreateLogConfigResponse) SetHeaders(v map[string]*string) *CreateLogConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateLogConfigResponse) SetStatusCode(v int32) *CreateLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLogConfigResponse) SetBody(v *CreateLogConfigResponseBody) *CreateLogConfigResponse {
	s.Body = v
	return s
}

func (s *CreateLogConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
