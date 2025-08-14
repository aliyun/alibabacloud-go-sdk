// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySearchIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySearchIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySearchIndexResponse
	GetStatusCode() *int32
	SetBody(v *QuerySearchIndexResponseBody) *QuerySearchIndexResponse
	GetBody() *QuerySearchIndexResponseBody
}

type QuerySearchIndexResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySearchIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySearchIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySearchIndexResponse) GoString() string {
	return s.String()
}

func (s *QuerySearchIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySearchIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySearchIndexResponse) GetBody() *QuerySearchIndexResponseBody {
	return s.Body
}

func (s *QuerySearchIndexResponse) SetHeaders(v map[string]*string) *QuerySearchIndexResponse {
	s.Headers = v
	return s
}

func (s *QuerySearchIndexResponse) SetStatusCode(v int32) *QuerySearchIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySearchIndexResponse) SetBody(v *QuerySearchIndexResponseBody) *QuerySearchIndexResponse {
	s.Body = v
	return s
}

func (s *QuerySearchIndexResponse) Validate() error {
	return dara.Validate(s)
}
