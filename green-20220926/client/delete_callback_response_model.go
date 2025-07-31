// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCallbackResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCallbackResponseBody) *DeleteCallbackResponse
	GetBody() *DeleteCallbackResponseBody
}

type DeleteCallbackResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCallbackResponse) GoString() string {
	return s.String()
}

func (s *DeleteCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCallbackResponse) GetBody() *DeleteCallbackResponseBody {
	return s.Body
}

func (s *DeleteCallbackResponse) SetHeaders(v map[string]*string) *DeleteCallbackResponse {
	s.Headers = v
	return s
}

func (s *DeleteCallbackResponse) SetStatusCode(v int32) *DeleteCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCallbackResponse) SetBody(v *DeleteCallbackResponseBody) *DeleteCallbackResponse {
	s.Body = v
	return s
}

func (s *DeleteCallbackResponse) Validate() error {
	return dara.Validate(s)
}
