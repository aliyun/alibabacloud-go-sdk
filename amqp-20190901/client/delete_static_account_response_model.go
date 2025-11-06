// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStaticAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStaticAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStaticAccountResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStaticAccountResponseBody) *DeleteStaticAccountResponse
	GetBody() *DeleteStaticAccountResponseBody
}

type DeleteStaticAccountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStaticAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStaticAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStaticAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteStaticAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStaticAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStaticAccountResponse) GetBody() *DeleteStaticAccountResponseBody {
	return s.Body
}

func (s *DeleteStaticAccountResponse) SetHeaders(v map[string]*string) *DeleteStaticAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteStaticAccountResponse) SetStatusCode(v int32) *DeleteStaticAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStaticAccountResponse) SetBody(v *DeleteStaticAccountResponseBody) *DeleteStaticAccountResponse {
	s.Body = v
	return s
}

func (s *DeleteStaticAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
