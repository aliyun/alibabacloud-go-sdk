// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSuperAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSuperAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSuperAccountResponse
	GetStatusCode() *int32
	SetBody(v *CreateSuperAccountResponseBody) *CreateSuperAccountResponse
	GetBody() *CreateSuperAccountResponseBody
}

type CreateSuperAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSuperAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSuperAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSuperAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateSuperAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSuperAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSuperAccountResponse) GetBody() *CreateSuperAccountResponseBody {
	return s.Body
}

func (s *CreateSuperAccountResponse) SetHeaders(v map[string]*string) *CreateSuperAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateSuperAccountResponse) SetStatusCode(v int32) *CreateSuperAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSuperAccountResponse) SetBody(v *CreateSuperAccountResponseBody) *CreateSuperAccountResponse {
	s.Body = v
	return s
}

func (s *CreateSuperAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
