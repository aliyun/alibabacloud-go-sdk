// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServerIdeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServerIdeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServerIdeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateServerIdeInstanceResponseBody) *CreateServerIdeInstanceResponse
	GetBody() *CreateServerIdeInstanceResponseBody
}

type CreateServerIdeInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServerIdeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServerIdeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServerIdeInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateServerIdeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServerIdeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServerIdeInstanceResponse) GetBody() *CreateServerIdeInstanceResponseBody {
	return s.Body
}

func (s *CreateServerIdeInstanceResponse) SetHeaders(v map[string]*string) *CreateServerIdeInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateServerIdeInstanceResponse) SetStatusCode(v int32) *CreateServerIdeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServerIdeInstanceResponse) SetBody(v *CreateServerIdeInstanceResponseBody) *CreateServerIdeInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateServerIdeInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
