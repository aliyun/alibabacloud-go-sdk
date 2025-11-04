// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSourceLocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSourceLocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSourceLocationResponse
	GetStatusCode() *int32
	SetBody(v *CreateSourceLocationResponseBody) *CreateSourceLocationResponse
	GetBody() *CreateSourceLocationResponseBody
}

type CreateSourceLocationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSourceLocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSourceLocationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSourceLocationResponse) GoString() string {
	return s.String()
}

func (s *CreateSourceLocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSourceLocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSourceLocationResponse) GetBody() *CreateSourceLocationResponseBody {
	return s.Body
}

func (s *CreateSourceLocationResponse) SetHeaders(v map[string]*string) *CreateSourceLocationResponse {
	s.Headers = v
	return s
}

func (s *CreateSourceLocationResponse) SetStatusCode(v int32) *CreateSourceLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSourceLocationResponse) SetBody(v *CreateSourceLocationResponseBody) *CreateSourceLocationResponse {
	s.Body = v
	return s
}

func (s *CreateSourceLocationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
