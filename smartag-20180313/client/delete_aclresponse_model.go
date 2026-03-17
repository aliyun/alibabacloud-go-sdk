// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteACLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteACLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteACLResponse
	GetStatusCode() *int32
	SetBody(v *DeleteACLResponseBody) *DeleteACLResponse
	GetBody() *DeleteACLResponseBody
}

type DeleteACLResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteACLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteACLResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteACLResponse) GoString() string {
	return s.String()
}

func (s *DeleteACLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteACLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteACLResponse) GetBody() *DeleteACLResponseBody {
	return s.Body
}

func (s *DeleteACLResponse) SetHeaders(v map[string]*string) *DeleteACLResponse {
	s.Headers = v
	return s
}

func (s *DeleteACLResponse) SetStatusCode(v int32) *DeleteACLResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteACLResponse) SetBody(v *DeleteACLResponseBody) *DeleteACLResponse {
	s.Body = v
	return s
}

func (s *DeleteACLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
