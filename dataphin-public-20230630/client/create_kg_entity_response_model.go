// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKgEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKgEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKgEntityResponse
	GetStatusCode() *int32
	SetBody(v *CreateKgEntityResponseBody) *CreateKgEntityResponse
	GetBody() *CreateKgEntityResponseBody
}

type CreateKgEntityResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKgEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKgEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKgEntityResponse) GoString() string {
	return s.String()
}

func (s *CreateKgEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKgEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKgEntityResponse) GetBody() *CreateKgEntityResponseBody {
	return s.Body
}

func (s *CreateKgEntityResponse) SetHeaders(v map[string]*string) *CreateKgEntityResponse {
	s.Headers = v
	return s
}

func (s *CreateKgEntityResponse) SetStatusCode(v int32) *CreateKgEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKgEntityResponse) SetBody(v *CreateKgEntityResponseBody) *CreateKgEntityResponse {
	s.Body = v
	return s
}

func (s *CreateKgEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
