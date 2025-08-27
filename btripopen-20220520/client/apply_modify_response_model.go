// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyModifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyModifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyModifyResponse
	GetStatusCode() *int32
	SetBody(v *ApplyModifyResponseBody) *ApplyModifyResponse
	GetBody() *ApplyModifyResponseBody
}

type ApplyModifyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyModifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyModifyResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyResponse) GoString() string {
	return s.String()
}

func (s *ApplyModifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyModifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyModifyResponse) GetBody() *ApplyModifyResponseBody {
	return s.Body
}

func (s *ApplyModifyResponse) SetHeaders(v map[string]*string) *ApplyModifyResponse {
	s.Headers = v
	return s
}

func (s *ApplyModifyResponse) SetStatusCode(v int32) *ApplyModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyModifyResponse) SetBody(v *ApplyModifyResponseBody) *ApplyModifyResponse {
	s.Body = v
	return s
}

func (s *ApplyModifyResponse) Validate() error {
	return dara.Validate(s)
}
