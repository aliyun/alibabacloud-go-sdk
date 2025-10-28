// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSubAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSubAccountResponse
	GetStatusCode() *int32
	SetBody(v *ListSubAccountResponseBody) *ListSubAccountResponse
	GetBody() *ListSubAccountResponseBody
}

type ListSubAccountResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSubAccountResponse) GoString() string {
	return s.String()
}

func (s *ListSubAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSubAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSubAccountResponse) GetBody() *ListSubAccountResponseBody {
	return s.Body
}

func (s *ListSubAccountResponse) SetHeaders(v map[string]*string) *ListSubAccountResponse {
	s.Headers = v
	return s
}

func (s *ListSubAccountResponse) SetStatusCode(v int32) *ListSubAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubAccountResponse) SetBody(v *ListSubAccountResponseBody) *ListSubAccountResponse {
	s.Body = v
	return s
}

func (s *ListSubAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
