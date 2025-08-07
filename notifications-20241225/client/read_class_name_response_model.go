// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadClassNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadClassNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadClassNameResponse
	GetStatusCode() *int32
	SetBody(v *ReadClassNameResponseBody) *ReadClassNameResponse
	GetBody() *ReadClassNameResponseBody
}

type ReadClassNameResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadClassNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadClassNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadClassNameResponse) GoString() string {
	return s.String()
}

func (s *ReadClassNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadClassNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadClassNameResponse) GetBody() *ReadClassNameResponseBody {
	return s.Body
}

func (s *ReadClassNameResponse) SetHeaders(v map[string]*string) *ReadClassNameResponse {
	s.Headers = v
	return s
}

func (s *ReadClassNameResponse) SetStatusCode(v int32) *ReadClassNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadClassNameResponse) SetBody(v *ReadClassNameResponseBody) *ReadClassNameResponse {
	s.Body = v
	return s
}

func (s *ReadClassNameResponse) Validate() error {
	return dara.Validate(s)
}
