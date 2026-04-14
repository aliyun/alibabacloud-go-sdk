// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStackResponse
	GetStatusCode() *int32
	SetBody(v *CreateStackResponseBody) *CreateStackResponse
	GetBody() *CreateStackResponseBody
}

type CreateStackResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStackResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStackResponse) GoString() string {
	return s.String()
}

func (s *CreateStackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStackResponse) GetBody() *CreateStackResponseBody {
	return s.Body
}

func (s *CreateStackResponse) SetHeaders(v map[string]*string) *CreateStackResponse {
	s.Headers = v
	return s
}

func (s *CreateStackResponse) SetStatusCode(v int32) *CreateStackResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStackResponse) SetBody(v *CreateStackResponseBody) *CreateStackResponse {
	s.Body = v
	return s
}

func (s *CreateStackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
