// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAICInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetAICInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetAICInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ResetAICInstanceResponseBody) *ResetAICInstanceResponse
	GetBody() *ResetAICInstanceResponseBody
}

type ResetAICInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAICInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAICInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetAICInstanceResponse) GoString() string {
	return s.String()
}

func (s *ResetAICInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetAICInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetAICInstanceResponse) GetBody() *ResetAICInstanceResponseBody {
	return s.Body
}

func (s *ResetAICInstanceResponse) SetHeaders(v map[string]*string) *ResetAICInstanceResponse {
	s.Headers = v
	return s
}

func (s *ResetAICInstanceResponse) SetStatusCode(v int32) *ResetAICInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAICInstanceResponse) SetBody(v *ResetAICInstanceResponseBody) *ResetAICInstanceResponse {
	s.Body = v
	return s
}

func (s *ResetAICInstanceResponse) Validate() error {
	return dara.Validate(s)
}
