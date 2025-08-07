// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadMessageResponse
	GetStatusCode() *int32
	SetBody(v *ReadMessageResponseBody) *ReadMessageResponse
	GetBody() *ReadMessageResponseBody
}

type ReadMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageResponse) GoString() string {
	return s.String()
}

func (s *ReadMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadMessageResponse) GetBody() *ReadMessageResponseBody {
	return s.Body
}

func (s *ReadMessageResponse) SetHeaders(v map[string]*string) *ReadMessageResponse {
	s.Headers = v
	return s
}

func (s *ReadMessageResponse) SetStatusCode(v int32) *ReadMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadMessageResponse) SetBody(v *ReadMessageResponseBody) *ReadMessageResponse {
	s.Body = v
	return s
}

func (s *ReadMessageResponse) Validate() error {
	return dara.Validate(s)
}
