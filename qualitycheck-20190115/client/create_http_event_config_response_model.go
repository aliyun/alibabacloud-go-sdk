// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpEventConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHttpEventConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHttpEventConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateHttpEventConfigResponseBody) *CreateHttpEventConfigResponse
	GetBody() *CreateHttpEventConfigResponseBody
}

type CreateHttpEventConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHttpEventConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHttpEventConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpEventConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateHttpEventConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHttpEventConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHttpEventConfigResponse) GetBody() *CreateHttpEventConfigResponseBody {
	return s.Body
}

func (s *CreateHttpEventConfigResponse) SetHeaders(v map[string]*string) *CreateHttpEventConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateHttpEventConfigResponse) SetStatusCode(v int32) *CreateHttpEventConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHttpEventConfigResponse) SetBody(v *CreateHttpEventConfigResponseBody) *CreateHttpEventConfigResponse {
	s.Body = v
	return s
}

func (s *CreateHttpEventConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
