// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateACLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateACLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateACLResponse
	GetStatusCode() *int32
	SetBody(v *CreateACLResponseBody) *CreateACLResponse
	GetBody() *CreateACLResponseBody
}

type CreateACLResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateACLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateACLResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateACLResponse) GoString() string {
	return s.String()
}

func (s *CreateACLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateACLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateACLResponse) GetBody() *CreateACLResponseBody {
	return s.Body
}

func (s *CreateACLResponse) SetHeaders(v map[string]*string) *CreateACLResponse {
	s.Headers = v
	return s
}

func (s *CreateACLResponse) SetStatusCode(v int32) *CreateACLResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateACLResponse) SetBody(v *CreateACLResponseBody) *CreateACLResponse {
	s.Body = v
	return s
}

func (s *CreateACLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
