// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePortResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePortResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePortResponse
	GetStatusCode() *int32
	SetBody(v *CreatePortResponseBody) *CreatePortResponse
	GetBody() *CreatePortResponseBody
}

type CreatePortResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePortResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePortResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePortResponse) GoString() string {
	return s.String()
}

func (s *CreatePortResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePortResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePortResponse) GetBody() *CreatePortResponseBody {
	return s.Body
}

func (s *CreatePortResponse) SetHeaders(v map[string]*string) *CreatePortResponse {
	s.Headers = v
	return s
}

func (s *CreatePortResponse) SetStatusCode(v int32) *CreatePortResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePortResponse) SetBody(v *CreatePortResponseBody) *CreatePortResponse {
	s.Body = v
	return s
}

func (s *CreatePortResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
