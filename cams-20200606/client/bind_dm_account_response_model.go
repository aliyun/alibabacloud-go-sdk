// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDmAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindDmAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindDmAccountResponse
	GetStatusCode() *int32
	SetBody(v *BindDmAccountResponseBody) *BindDmAccountResponse
	GetBody() *BindDmAccountResponseBody
}

type BindDmAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindDmAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindDmAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s BindDmAccountResponse) GoString() string {
	return s.String()
}

func (s *BindDmAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindDmAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindDmAccountResponse) GetBody() *BindDmAccountResponseBody {
	return s.Body
}

func (s *BindDmAccountResponse) SetHeaders(v map[string]*string) *BindDmAccountResponse {
	s.Headers = v
	return s
}

func (s *BindDmAccountResponse) SetStatusCode(v int32) *BindDmAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *BindDmAccountResponse) SetBody(v *BindDmAccountResponseBody) *BindDmAccountResponse {
	s.Body = v
	return s
}

func (s *BindDmAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
