// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceAccountResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceAccountResponseBody) *GetInstanceAccountResponse
	GetBody() *GetInstanceAccountResponseBody
}

type GetInstanceAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAccountResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceAccountResponse) GetBody() *GetInstanceAccountResponseBody {
	return s.Body
}

func (s *GetInstanceAccountResponse) SetHeaders(v map[string]*string) *GetInstanceAccountResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceAccountResponse) SetStatusCode(v int32) *GetInstanceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceAccountResponse) SetBody(v *GetInstanceAccountResponseBody) *GetInstanceAccountResponse {
	s.Body = v
	return s
}

func (s *GetInstanceAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
