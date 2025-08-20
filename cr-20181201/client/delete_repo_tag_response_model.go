// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepoTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRepoTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRepoTagResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRepoTagResponseBody) *DeleteRepoTagResponse
	GetBody() *DeleteRepoTagResponseBody
}

type DeleteRepoTagResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRepoTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRepoTagResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepoTagResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepoTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRepoTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRepoTagResponse) GetBody() *DeleteRepoTagResponseBody {
	return s.Body
}

func (s *DeleteRepoTagResponse) SetHeaders(v map[string]*string) *DeleteRepoTagResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepoTagResponse) SetStatusCode(v int32) *DeleteRepoTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepoTagResponse) SetBody(v *DeleteRepoTagResponseBody) *DeleteRepoTagResponse {
	s.Body = v
	return s
}

func (s *DeleteRepoTagResponse) Validate() error {
	return dara.Validate(s)
}
