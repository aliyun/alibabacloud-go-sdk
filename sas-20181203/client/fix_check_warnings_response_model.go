// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFixCheckWarningsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FixCheckWarningsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FixCheckWarningsResponse
	GetStatusCode() *int32
	SetBody(v *FixCheckWarningsResponseBody) *FixCheckWarningsResponse
	GetBody() *FixCheckWarningsResponseBody
}

type FixCheckWarningsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FixCheckWarningsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FixCheckWarningsResponse) String() string {
	return dara.Prettify(s)
}

func (s FixCheckWarningsResponse) GoString() string {
	return s.String()
}

func (s *FixCheckWarningsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FixCheckWarningsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FixCheckWarningsResponse) GetBody() *FixCheckWarningsResponseBody {
	return s.Body
}

func (s *FixCheckWarningsResponse) SetHeaders(v map[string]*string) *FixCheckWarningsResponse {
	s.Headers = v
	return s
}

func (s *FixCheckWarningsResponse) SetStatusCode(v int32) *FixCheckWarningsResponse {
	s.StatusCode = &v
	return s
}

func (s *FixCheckWarningsResponse) SetBody(v *FixCheckWarningsResponseBody) *FixCheckWarningsResponse {
	s.Body = v
	return s
}

func (s *FixCheckWarningsResponse) Validate() error {
	return dara.Validate(s)
}
