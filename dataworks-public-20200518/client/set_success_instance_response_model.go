// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSuccessInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSuccessInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSuccessInstanceResponse
	GetStatusCode() *int32
	SetBody(v *SetSuccessInstanceResponseBody) *SetSuccessInstanceResponse
	GetBody() *SetSuccessInstanceResponseBody
}

type SetSuccessInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSuccessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSuccessInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSuccessInstanceResponse) GoString() string {
	return s.String()
}

func (s *SetSuccessInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSuccessInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSuccessInstanceResponse) GetBody() *SetSuccessInstanceResponseBody {
	return s.Body
}

func (s *SetSuccessInstanceResponse) SetHeaders(v map[string]*string) *SetSuccessInstanceResponse {
	s.Headers = v
	return s
}

func (s *SetSuccessInstanceResponse) SetStatusCode(v int32) *SetSuccessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSuccessInstanceResponse) SetBody(v *SetSuccessInstanceResponseBody) *SetSuccessInstanceResponse {
	s.Body = v
	return s
}

func (s *SetSuccessInstanceResponse) Validate() error {
	return dara.Validate(s)
}
