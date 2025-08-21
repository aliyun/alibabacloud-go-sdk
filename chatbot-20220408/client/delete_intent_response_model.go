// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIntentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIntentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIntentResponseBody) *DeleteIntentResponse
	GetBody() *DeleteIntentResponseBody
}

type DeleteIntentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIntentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntentResponse) GoString() string {
	return s.String()
}

func (s *DeleteIntentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIntentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIntentResponse) GetBody() *DeleteIntentResponseBody {
	return s.Body
}

func (s *DeleteIntentResponse) SetHeaders(v map[string]*string) *DeleteIntentResponse {
	s.Headers = v
	return s
}

func (s *DeleteIntentResponse) SetStatusCode(v int32) *DeleteIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIntentResponse) SetBody(v *DeleteIntentResponseBody) *DeleteIntentResponse {
	s.Body = v
	return s
}

func (s *DeleteIntentResponse) Validate() error {
	return dara.Validate(s)
}
