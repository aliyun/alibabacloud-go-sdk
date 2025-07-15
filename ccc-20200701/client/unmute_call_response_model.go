// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnmuteCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnmuteCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnmuteCallResponse
	GetStatusCode() *int32
	SetBody(v *UnmuteCallResponseBody) *UnmuteCallResponse
	GetBody() *UnmuteCallResponseBody
}

type UnmuteCallResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnmuteCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnmuteCallResponse) String() string {
	return dara.Prettify(s)
}

func (s UnmuteCallResponse) GoString() string {
	return s.String()
}

func (s *UnmuteCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnmuteCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnmuteCallResponse) GetBody() *UnmuteCallResponseBody {
	return s.Body
}

func (s *UnmuteCallResponse) SetHeaders(v map[string]*string) *UnmuteCallResponse {
	s.Headers = v
	return s
}

func (s *UnmuteCallResponse) SetStatusCode(v int32) *UnmuteCallResponse {
	s.StatusCode = &v
	return s
}

func (s *UnmuteCallResponse) SetBody(v *UnmuteCallResponseBody) *UnmuteCallResponse {
	s.Body = v
	return s
}

func (s *UnmuteCallResponse) Validate() error {
	return dara.Validate(s)
}
