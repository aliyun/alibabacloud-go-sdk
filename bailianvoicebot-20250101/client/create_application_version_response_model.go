// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApplicationVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApplicationVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateApplicationVersionResponseBody) *CreateApplicationVersionResponse
	GetBody() *CreateApplicationVersionResponseBody
}

type CreateApplicationVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApplicationVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApplicationVersionResponse) GetBody() *CreateApplicationVersionResponseBody {
	return s.Body
}

func (s *CreateApplicationVersionResponse) SetHeaders(v map[string]*string) *CreateApplicationVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationVersionResponse) SetStatusCode(v int32) *CreateApplicationVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationVersionResponse) SetBody(v *CreateApplicationVersionResponseBody) *CreateApplicationVersionResponse {
	s.Body = v
	return s
}

func (s *CreateApplicationVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
