// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverAICInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecoverAICInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecoverAICInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RecoverAICInstanceResponseBody) *RecoverAICInstanceResponse
	GetBody() *RecoverAICInstanceResponseBody
}

type RecoverAICInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoverAICInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoverAICInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RecoverAICInstanceResponse) GoString() string {
	return s.String()
}

func (s *RecoverAICInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecoverAICInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecoverAICInstanceResponse) GetBody() *RecoverAICInstanceResponseBody {
	return s.Body
}

func (s *RecoverAICInstanceResponse) SetHeaders(v map[string]*string) *RecoverAICInstanceResponse {
	s.Headers = v
	return s
}

func (s *RecoverAICInstanceResponse) SetStatusCode(v int32) *RecoverAICInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverAICInstanceResponse) SetBody(v *RecoverAICInstanceResponseBody) *RecoverAICInstanceResponse {
	s.Body = v
	return s
}

func (s *RecoverAICInstanceResponse) Validate() error {
	return dara.Validate(s)
}
