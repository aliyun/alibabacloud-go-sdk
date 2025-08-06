// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DelAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DelAppResponse
	GetStatusCode() *int32
	SetBody(v *DelAppResponseBody) *DelAppResponse
	GetBody() *DelAppResponseBody
}

type DelAppResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DelAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DelAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DelAppResponse) GoString() string {
	return s.String()
}

func (s *DelAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DelAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DelAppResponse) GetBody() *DelAppResponseBody {
	return s.Body
}

func (s *DelAppResponse) SetHeaders(v map[string]*string) *DelAppResponse {
	s.Headers = v
	return s
}

func (s *DelAppResponse) SetStatusCode(v int32) *DelAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DelAppResponse) SetBody(v *DelAppResponseBody) *DelAppResponse {
	s.Body = v
	return s
}

func (s *DelAppResponse) Validate() error {
	return dara.Validate(s)
}
