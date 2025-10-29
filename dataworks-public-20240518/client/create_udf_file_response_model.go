// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUdfFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUdfFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUdfFileResponse
	GetStatusCode() *int32
	SetBody(v *CreateUdfFileResponseBody) *CreateUdfFileResponse
	GetBody() *CreateUdfFileResponseBody
}

type CreateUdfFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUdfFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUdfFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUdfFileResponse) GoString() string {
	return s.String()
}

func (s *CreateUdfFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUdfFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUdfFileResponse) GetBody() *CreateUdfFileResponseBody {
	return s.Body
}

func (s *CreateUdfFileResponse) SetHeaders(v map[string]*string) *CreateUdfFileResponse {
	s.Headers = v
	return s
}

func (s *CreateUdfFileResponse) SetStatusCode(v int32) *CreateUdfFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUdfFileResponse) SetBody(v *CreateUdfFileResponseBody) *CreateUdfFileResponse {
	s.Body = v
	return s
}

func (s *CreateUdfFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
