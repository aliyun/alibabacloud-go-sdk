// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUsedPropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckUsedPropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckUsedPropertyResponse
	GetStatusCode() *int32
	SetBody(v *CheckUsedPropertyResponseBody) *CheckUsedPropertyResponse
	GetBody() *CheckUsedPropertyResponseBody
}

type CheckUsedPropertyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUsedPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUsedPropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckUsedPropertyResponse) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckUsedPropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckUsedPropertyResponse) GetBody() *CheckUsedPropertyResponseBody {
	return s.Body
}

func (s *CheckUsedPropertyResponse) SetHeaders(v map[string]*string) *CheckUsedPropertyResponse {
	s.Headers = v
	return s
}

func (s *CheckUsedPropertyResponse) SetStatusCode(v int32) *CheckUsedPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUsedPropertyResponse) SetBody(v *CheckUsedPropertyResponseBody) *CheckUsedPropertyResponse {
	s.Body = v
	return s
}

func (s *CheckUsedPropertyResponse) Validate() error {
	return dara.Validate(s)
}
