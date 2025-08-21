// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaqResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFaqResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFaqResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFaqResponseBody) *DeleteFaqResponse
	GetBody() *DeleteFaqResponseBody
}

type DeleteFaqResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaqResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaqResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaqResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaqResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFaqResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFaqResponse) GetBody() *DeleteFaqResponseBody {
	return s.Body
}

func (s *DeleteFaqResponse) SetHeaders(v map[string]*string) *DeleteFaqResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaqResponse) SetStatusCode(v int32) *DeleteFaqResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaqResponse) SetBody(v *DeleteFaqResponseBody) *DeleteFaqResponse {
	s.Body = v
	return s
}

func (s *DeleteFaqResponse) Validate() error {
	return dara.Validate(s)
}
