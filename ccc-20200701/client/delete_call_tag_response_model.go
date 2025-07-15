// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCallTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCallTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCallTagResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCallTagResponseBody) *DeleteCallTagResponse
	GetBody() *DeleteCallTagResponseBody
}

type DeleteCallTagResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCallTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCallTagResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCallTagResponse) GoString() string {
	return s.String()
}

func (s *DeleteCallTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCallTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCallTagResponse) GetBody() *DeleteCallTagResponseBody {
	return s.Body
}

func (s *DeleteCallTagResponse) SetHeaders(v map[string]*string) *DeleteCallTagResponse {
	s.Headers = v
	return s
}

func (s *DeleteCallTagResponse) SetStatusCode(v int32) *DeleteCallTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCallTagResponse) SetBody(v *DeleteCallTagResponseBody) *DeleteCallTagResponse {
	s.Body = v
	return s
}

func (s *DeleteCallTagResponse) Validate() error {
	return dara.Validate(s)
}
