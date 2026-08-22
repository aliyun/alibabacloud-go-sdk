// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpenSearchAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOpenSearchAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOpenSearchAccountResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOpenSearchAccountResponseBody) *DeleteOpenSearchAccountResponse
	GetBody() *DeleteOpenSearchAccountResponseBody
}

type DeleteOpenSearchAccountResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOpenSearchAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOpenSearchAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSearchAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteOpenSearchAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOpenSearchAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOpenSearchAccountResponse) GetBody() *DeleteOpenSearchAccountResponseBody {
	return s.Body
}

func (s *DeleteOpenSearchAccountResponse) SetHeaders(v map[string]*string) *DeleteOpenSearchAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteOpenSearchAccountResponse) SetStatusCode(v int32) *DeleteOpenSearchAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOpenSearchAccountResponse) SetBody(v *DeleteOpenSearchAccountResponseBody) *DeleteOpenSearchAccountResponse {
	s.Body = v
	return s
}

func (s *DeleteOpenSearchAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
