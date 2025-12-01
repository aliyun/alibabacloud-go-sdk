// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateConfigResponseBody) *CreateConfigResponse
	GetBody() *CreateConfigResponseBody
}

type CreateConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConfigResponse) GetBody() *CreateConfigResponseBody {
	return s.Body
}

func (s *CreateConfigResponse) SetHeaders(v map[string]*string) *CreateConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigResponse) SetStatusCode(v int32) *CreateConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConfigResponse) SetBody(v *CreateConfigResponseBody) *CreateConfigResponse {
	s.Body = v
	return s
}

func (s *CreateConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
