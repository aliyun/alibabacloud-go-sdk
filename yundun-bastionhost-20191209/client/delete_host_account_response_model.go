// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHostAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHostAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHostAccountResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHostAccountResponseBody) *DeleteHostAccountResponse
	GetBody() *DeleteHostAccountResponseBody
}

type DeleteHostAccountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHostAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHostAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteHostAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHostAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHostAccountResponse) GetBody() *DeleteHostAccountResponseBody {
	return s.Body
}

func (s *DeleteHostAccountResponse) SetHeaders(v map[string]*string) *DeleteHostAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteHostAccountResponse) SetStatusCode(v int32) *DeleteHostAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHostAccountResponse) SetBody(v *DeleteHostAccountResponseBody) *DeleteHostAccountResponse {
	s.Body = v
	return s
}

func (s *DeleteHostAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
