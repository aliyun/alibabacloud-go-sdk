// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UntagResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UntagResourceResponse
	GetStatusCode() *int32
	SetBody(v *UntagResourceResponseBody) *UntagResourceResponse
	GetBody() *UntagResourceResponseBody
}

type UntagResourceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UntagResourceResponse) GoString() string {
	return s.String()
}

func (s *UntagResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UntagResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UntagResourceResponse) GetBody() *UntagResourceResponseBody {
	return s.Body
}

func (s *UntagResourceResponse) SetHeaders(v map[string]*string) *UntagResourceResponse {
	s.Headers = v
	return s
}

func (s *UntagResourceResponse) SetStatusCode(v int32) *UntagResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourceResponse) SetBody(v *UntagResourceResponseBody) *UntagResourceResponse {
	s.Body = v
	return s
}

func (s *UntagResourceResponse) Validate() error {
	return dara.Validate(s)
}
