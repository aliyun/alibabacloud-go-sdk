// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEaiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEaiResponse
	GetStatusCode() *int32
	SetBody(v *CreateEaiResponseBody) *CreateEaiResponse
	GetBody() *CreateEaiResponseBody
}

type CreateEaiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEaiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEaiResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiResponse) GoString() string {
	return s.String()
}

func (s *CreateEaiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEaiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEaiResponse) GetBody() *CreateEaiResponseBody {
	return s.Body
}

func (s *CreateEaiResponse) SetHeaders(v map[string]*string) *CreateEaiResponse {
	s.Headers = v
	return s
}

func (s *CreateEaiResponse) SetStatusCode(v int32) *CreateEaiResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEaiResponse) SetBody(v *CreateEaiResponseBody) *CreateEaiResponse {
	s.Body = v
	return s
}

func (s *CreateEaiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
