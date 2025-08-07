// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DelMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DelMessageResponse
	GetStatusCode() *int32
	SetBody(v *DelMessageResponseBody) *DelMessageResponse
	GetBody() *DelMessageResponseBody
}

type DelMessageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DelMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DelMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s DelMessageResponse) GoString() string {
	return s.String()
}

func (s *DelMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DelMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DelMessageResponse) GetBody() *DelMessageResponseBody {
	return s.Body
}

func (s *DelMessageResponse) SetHeaders(v map[string]*string) *DelMessageResponse {
	s.Headers = v
	return s
}

func (s *DelMessageResponse) SetStatusCode(v int32) *DelMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *DelMessageResponse) SetBody(v *DelMessageResponseBody) *DelMessageResponse {
	s.Body = v
	return s
}

func (s *DelMessageResponse) Validate() error {
	return dara.Validate(s)
}
