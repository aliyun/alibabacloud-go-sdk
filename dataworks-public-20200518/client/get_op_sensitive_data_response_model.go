// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpSensitiveDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpSensitiveDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpSensitiveDataResponse
	GetStatusCode() *int32
	SetBody(v *GetOpSensitiveDataResponseBody) *GetOpSensitiveDataResponse
	GetBody() *GetOpSensitiveDataResponseBody
}

type GetOpSensitiveDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpSensitiveDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpSensitiveDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpSensitiveDataResponse) GoString() string {
	return s.String()
}

func (s *GetOpSensitiveDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpSensitiveDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpSensitiveDataResponse) GetBody() *GetOpSensitiveDataResponseBody {
	return s.Body
}

func (s *GetOpSensitiveDataResponse) SetHeaders(v map[string]*string) *GetOpSensitiveDataResponse {
	s.Headers = v
	return s
}

func (s *GetOpSensitiveDataResponse) SetStatusCode(v int32) *GetOpSensitiveDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpSensitiveDataResponse) SetBody(v *GetOpSensitiveDataResponseBody) *GetOpSensitiveDataResponse {
	s.Body = v
	return s
}

func (s *GetOpSensitiveDataResponse) Validate() error {
	return dara.Validate(s)
}
