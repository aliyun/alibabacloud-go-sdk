// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResourceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResourceAccountResponse
	GetStatusCode() *int32
	SetBody(v *CreateResourceAccountResponseBody) *CreateResourceAccountResponse
	GetBody() *CreateResourceAccountResponseBody
}

type CreateResourceAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResourceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResourceAccountResponse) GetBody() *CreateResourceAccountResponseBody {
	return s.Body
}

func (s *CreateResourceAccountResponse) SetHeaders(v map[string]*string) *CreateResourceAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceAccountResponse) SetStatusCode(v int32) *CreateResourceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceAccountResponse) SetBody(v *CreateResourceAccountResponseBody) *CreateResourceAccountResponse {
	s.Body = v
	return s
}

func (s *CreateResourceAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
