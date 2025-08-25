// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushSimpleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushSimpleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushSimpleResponse
	GetStatusCode() *int32
	SetBody(v *PushSimpleResponseBody) *PushSimpleResponse
	GetBody() *PushSimpleResponseBody
}

type PushSimpleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushSimpleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushSimpleResponse) String() string {
	return dara.Prettify(s)
}

func (s PushSimpleResponse) GoString() string {
	return s.String()
}

func (s *PushSimpleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushSimpleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushSimpleResponse) GetBody() *PushSimpleResponseBody {
	return s.Body
}

func (s *PushSimpleResponse) SetHeaders(v map[string]*string) *PushSimpleResponse {
	s.Headers = v
	return s
}

func (s *PushSimpleResponse) SetStatusCode(v int32) *PushSimpleResponse {
	s.StatusCode = &v
	return s
}

func (s *PushSimpleResponse) SetBody(v *PushSimpleResponseBody) *PushSimpleResponse {
	s.Body = v
	return s
}

func (s *PushSimpleResponse) Validate() error {
	return dara.Validate(s)
}
