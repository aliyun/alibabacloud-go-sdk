// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TagResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TagResourceResponse
	GetStatusCode() *int32
	SetBody(v *TagResourceResponseBody) *TagResourceResponse
	GetBody() *TagResourceResponseBody
}

type TagResourceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s TagResourceResponse) GoString() string {
	return s.String()
}

func (s *TagResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TagResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TagResourceResponse) GetBody() *TagResourceResponseBody {
	return s.Body
}

func (s *TagResourceResponse) SetHeaders(v map[string]*string) *TagResourceResponse {
	s.Headers = v
	return s
}

func (s *TagResourceResponse) SetStatusCode(v int32) *TagResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourceResponse) SetBody(v *TagResourceResponseBody) *TagResourceResponse {
	s.Body = v
	return s
}

func (s *TagResourceResponse) Validate() error {
	return dara.Validate(s)
}
