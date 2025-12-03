// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppCodeResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppCodeResponseBody) *CreateAppCodeResponse
	GetBody() *CreateAppCodeResponseBody
}

type CreateAppCodeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppCodeResponse) GoString() string {
	return s.String()
}

func (s *CreateAppCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppCodeResponse) GetBody() *CreateAppCodeResponseBody {
	return s.Body
}

func (s *CreateAppCodeResponse) SetHeaders(v map[string]*string) *CreateAppCodeResponse {
	s.Headers = v
	return s
}

func (s *CreateAppCodeResponse) SetStatusCode(v int32) *CreateAppCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppCodeResponse) SetBody(v *CreateAppCodeResponseBody) *CreateAppCodeResponse {
	s.Body = v
	return s
}

func (s *CreateAppCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
