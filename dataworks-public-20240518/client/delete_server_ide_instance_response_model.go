// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServerIdeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServerIdeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServerIdeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServerIdeInstanceResponseBody) *DeleteServerIdeInstanceResponse
	GetBody() *DeleteServerIdeInstanceResponseBody
}

type DeleteServerIdeInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServerIdeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServerIdeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServerIdeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServerIdeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServerIdeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServerIdeInstanceResponse) GetBody() *DeleteServerIdeInstanceResponseBody {
	return s.Body
}

func (s *DeleteServerIdeInstanceResponse) SetHeaders(v map[string]*string) *DeleteServerIdeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServerIdeInstanceResponse) SetStatusCode(v int32) *DeleteServerIdeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServerIdeInstanceResponse) SetBody(v *DeleteServerIdeInstanceResponseBody) *DeleteServerIdeInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteServerIdeInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
