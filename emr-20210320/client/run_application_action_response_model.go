// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunApplicationActionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunApplicationActionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunApplicationActionResponse
	GetStatusCode() *int32
	SetBody(v *RunApplicationActionResponseBody) *RunApplicationActionResponse
	GetBody() *RunApplicationActionResponseBody
}

type RunApplicationActionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunApplicationActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunApplicationActionResponse) String() string {
	return dara.Prettify(s)
}

func (s RunApplicationActionResponse) GoString() string {
	return s.String()
}

func (s *RunApplicationActionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunApplicationActionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunApplicationActionResponse) GetBody() *RunApplicationActionResponseBody {
	return s.Body
}

func (s *RunApplicationActionResponse) SetHeaders(v map[string]*string) *RunApplicationActionResponse {
	s.Headers = v
	return s
}

func (s *RunApplicationActionResponse) SetStatusCode(v int32) *RunApplicationActionResponse {
	s.StatusCode = &v
	return s
}

func (s *RunApplicationActionResponse) SetBody(v *RunApplicationActionResponseBody) *RunApplicationActionResponse {
	s.Body = v
	return s
}

func (s *RunApplicationActionResponse) Validate() error {
	return dara.Validate(s)
}
