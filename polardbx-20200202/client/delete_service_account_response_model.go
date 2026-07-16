// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceAccountResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceAccountResponseBody) *DeleteServiceAccountResponse
	GetBody() *DeleteServiceAccountResponseBody
}

type DeleteServiceAccountResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceAccountResponse) GetBody() *DeleteServiceAccountResponseBody {
	return s.Body
}

func (s *DeleteServiceAccountResponse) SetHeaders(v map[string]*string) *DeleteServiceAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceAccountResponse) SetStatusCode(v int32) *DeleteServiceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceAccountResponse) SetBody(v *DeleteServiceAccountResponseBody) *DeleteServiceAccountResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
