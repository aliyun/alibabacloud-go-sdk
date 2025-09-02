// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSensitiveDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSensitiveDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSensitiveDataResponse
	GetStatusCode() *int32
	SetBody(v *GetSensitiveDataResponseBody) *GetSensitiveDataResponse
	GetBody() *GetSensitiveDataResponseBody
}

type GetSensitiveDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSensitiveDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSensitiveDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSensitiveDataResponse) GoString() string {
	return s.String()
}

func (s *GetSensitiveDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSensitiveDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSensitiveDataResponse) GetBody() *GetSensitiveDataResponseBody {
	return s.Body
}

func (s *GetSensitiveDataResponse) SetHeaders(v map[string]*string) *GetSensitiveDataResponse {
	s.Headers = v
	return s
}

func (s *GetSensitiveDataResponse) SetStatusCode(v int32) *GetSensitiveDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSensitiveDataResponse) SetBody(v *GetSensitiveDataResponseBody) *GetSensitiveDataResponse {
	s.Body = v
	return s
}

func (s *GetSensitiveDataResponse) Validate() error {
	return dara.Validate(s)
}
