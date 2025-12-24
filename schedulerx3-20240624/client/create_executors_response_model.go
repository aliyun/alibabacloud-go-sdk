// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExecutorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExecutorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExecutorsResponse
	GetStatusCode() *int32
	SetBody(v *CreateExecutorsResponseBody) *CreateExecutorsResponse
	GetBody() *CreateExecutorsResponseBody
}

type CreateExecutorsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExecutorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExecutorsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExecutorsResponse) GoString() string {
	return s.String()
}

func (s *CreateExecutorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExecutorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExecutorsResponse) GetBody() *CreateExecutorsResponseBody {
	return s.Body
}

func (s *CreateExecutorsResponse) SetHeaders(v map[string]*string) *CreateExecutorsResponse {
	s.Headers = v
	return s
}

func (s *CreateExecutorsResponse) SetStatusCode(v int32) *CreateExecutorsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExecutorsResponse) SetBody(v *CreateExecutorsResponseBody) *CreateExecutorsResponse {
	s.Body = v
	return s
}

func (s *CreateExecutorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
