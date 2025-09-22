// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionByActivityIdPopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySessionByActivityIdPopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySessionByActivityIdPopResponse
	GetStatusCode() *int32
	SetBody(v *QuerySessionByActivityIdPopResponseBody) *QuerySessionByActivityIdPopResponse
	GetBody() *QuerySessionByActivityIdPopResponseBody
}

type QuerySessionByActivityIdPopResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySessionByActivityIdPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySessionByActivityIdPopResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionByActivityIdPopResponse) GoString() string {
	return s.String()
}

func (s *QuerySessionByActivityIdPopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySessionByActivityIdPopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySessionByActivityIdPopResponse) GetBody() *QuerySessionByActivityIdPopResponseBody {
	return s.Body
}

func (s *QuerySessionByActivityIdPopResponse) SetHeaders(v map[string]*string) *QuerySessionByActivityIdPopResponse {
	s.Headers = v
	return s
}

func (s *QuerySessionByActivityIdPopResponse) SetStatusCode(v int32) *QuerySessionByActivityIdPopResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySessionByActivityIdPopResponse) SetBody(v *QuerySessionByActivityIdPopResponseBody) *QuerySessionByActivityIdPopResponse {
	s.Body = v
	return s
}

func (s *QuerySessionByActivityIdPopResponse) Validate() error {
	return dara.Validate(s)
}
