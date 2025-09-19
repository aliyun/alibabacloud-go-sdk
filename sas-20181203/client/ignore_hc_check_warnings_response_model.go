// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreHcCheckWarningsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IgnoreHcCheckWarningsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IgnoreHcCheckWarningsResponse
	GetStatusCode() *int32
	SetBody(v *IgnoreHcCheckWarningsResponseBody) *IgnoreHcCheckWarningsResponse
	GetBody() *IgnoreHcCheckWarningsResponseBody
}

type IgnoreHcCheckWarningsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IgnoreHcCheckWarningsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IgnoreHcCheckWarningsResponse) String() string {
	return dara.Prettify(s)
}

func (s IgnoreHcCheckWarningsResponse) GoString() string {
	return s.String()
}

func (s *IgnoreHcCheckWarningsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IgnoreHcCheckWarningsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IgnoreHcCheckWarningsResponse) GetBody() *IgnoreHcCheckWarningsResponseBody {
	return s.Body
}

func (s *IgnoreHcCheckWarningsResponse) SetHeaders(v map[string]*string) *IgnoreHcCheckWarningsResponse {
	s.Headers = v
	return s
}

func (s *IgnoreHcCheckWarningsResponse) SetStatusCode(v int32) *IgnoreHcCheckWarningsResponse {
	s.StatusCode = &v
	return s
}

func (s *IgnoreHcCheckWarningsResponse) SetBody(v *IgnoreHcCheckWarningsResponseBody) *IgnoreHcCheckWarningsResponse {
	s.Body = v
	return s
}

func (s *IgnoreHcCheckWarningsResponse) Validate() error {
	return dara.Validate(s)
}
