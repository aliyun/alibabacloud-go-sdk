// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRunResponse
	GetStatusCode() *int32
	SetBody(v *CreateRunResponseBody) *CreateRunResponse
	GetBody() *CreateRunResponseBody
}

type CreateRunResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRunResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRunResponse) GoString() string {
	return s.String()
}

func (s *CreateRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRunResponse) GetBody() *CreateRunResponseBody {
	return s.Body
}

func (s *CreateRunResponse) SetHeaders(v map[string]*string) *CreateRunResponse {
	s.Headers = v
	return s
}

func (s *CreateRunResponse) SetStatusCode(v int32) *CreateRunResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRunResponse) SetBody(v *CreateRunResponseBody) *CreateRunResponse {
	s.Body = v
	return s
}

func (s *CreateRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
