// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpenSourceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOpenSourceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOpenSourceAccountResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOpenSourceAccountResponseBody) *DeleteOpenSourceAccountResponse
	GetBody() *DeleteOpenSourceAccountResponseBody
}

type DeleteOpenSourceAccountResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOpenSourceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOpenSourceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSourceAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteOpenSourceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOpenSourceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOpenSourceAccountResponse) GetBody() *DeleteOpenSourceAccountResponseBody {
	return s.Body
}

func (s *DeleteOpenSourceAccountResponse) SetHeaders(v map[string]*string) *DeleteOpenSourceAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteOpenSourceAccountResponse) SetStatusCode(v int32) *DeleteOpenSourceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOpenSourceAccountResponse) SetBody(v *DeleteOpenSourceAccountResponseBody) *DeleteOpenSourceAccountResponse {
	s.Body = v
	return s
}

func (s *DeleteOpenSourceAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
