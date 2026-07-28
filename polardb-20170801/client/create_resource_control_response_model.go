// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResourceControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResourceControlResponse
	GetStatusCode() *int32
	SetBody(v *CreateResourceControlResponseBody) *CreateResourceControlResponse
	GetBody() *CreateResourceControlResponseBody
}

type CreateResourceControlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceControlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceControlResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResourceControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResourceControlResponse) GetBody() *CreateResourceControlResponseBody {
	return s.Body
}

func (s *CreateResourceControlResponse) SetHeaders(v map[string]*string) *CreateResourceControlResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceControlResponse) SetStatusCode(v int32) *CreateResourceControlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceControlResponse) SetBody(v *CreateResourceControlResponseBody) *CreateResourceControlResponse {
	s.Body = v
	return s
}

func (s *CreateResourceControlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
