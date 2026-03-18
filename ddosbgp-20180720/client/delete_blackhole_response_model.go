// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBlackholeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBlackholeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBlackholeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBlackholeResponseBody) *DeleteBlackholeResponse
	GetBody() *DeleteBlackholeResponseBody
}

type DeleteBlackholeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBlackholeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBlackholeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBlackholeResponse) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBlackholeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBlackholeResponse) GetBody() *DeleteBlackholeResponseBody {
	return s.Body
}

func (s *DeleteBlackholeResponse) SetHeaders(v map[string]*string) *DeleteBlackholeResponse {
	s.Headers = v
	return s
}

func (s *DeleteBlackholeResponse) SetStatusCode(v int32) *DeleteBlackholeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBlackholeResponse) SetBody(v *DeleteBlackholeResponseBody) *DeleteBlackholeResponse {
	s.Body = v
	return s
}

func (s *DeleteBlackholeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
