// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConsumerProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConsumerProgressResponse
	GetStatusCode() *int32
	SetBody(v *GetConsumerProgressResponseBody) *GetConsumerProgressResponse
	GetBody() *GetConsumerProgressResponseBody
}

type GetConsumerProgressResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumerProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumerProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerProgressResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConsumerProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConsumerProgressResponse) GetBody() *GetConsumerProgressResponseBody {
	return s.Body
}

func (s *GetConsumerProgressResponse) SetHeaders(v map[string]*string) *GetConsumerProgressResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerProgressResponse) SetStatusCode(v int32) *GetConsumerProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumerProgressResponse) SetBody(v *GetConsumerProgressResponseBody) *GetConsumerProgressResponse {
	s.Body = v
	return s
}

func (s *GetConsumerProgressResponse) Validate() error {
	return dara.Validate(s)
}
