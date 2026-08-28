// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpApiVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHttpApiVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHttpApiVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateHttpApiVersionResponseBody) *CreateHttpApiVersionResponse
	GetBody() *CreateHttpApiVersionResponseBody
}

type CreateHttpApiVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHttpApiVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHttpApiVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateHttpApiVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHttpApiVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHttpApiVersionResponse) GetBody() *CreateHttpApiVersionResponseBody {
	return s.Body
}

func (s *CreateHttpApiVersionResponse) SetHeaders(v map[string]*string) *CreateHttpApiVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateHttpApiVersionResponse) SetStatusCode(v int32) *CreateHttpApiVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHttpApiVersionResponse) SetBody(v *CreateHttpApiVersionResponseBody) *CreateHttpApiVersionResponse {
	s.Body = v
	return s
}

func (s *CreateHttpApiVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
