// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUdfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUdfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUdfResponse
	GetStatusCode() *int32
	SetBody(v *GetUdfResponseBody) *GetUdfResponse
	GetBody() *GetUdfResponseBody
}

type GetUdfResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUdfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUdfResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUdfResponse) GoString() string {
	return s.String()
}

func (s *GetUdfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUdfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUdfResponse) GetBody() *GetUdfResponseBody {
	return s.Body
}

func (s *GetUdfResponse) SetHeaders(v map[string]*string) *GetUdfResponse {
	s.Headers = v
	return s
}

func (s *GetUdfResponse) SetStatusCode(v int32) *GetUdfResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUdfResponse) SetBody(v *GetUdfResponseBody) *GetUdfResponse {
	s.Body = v
	return s
}

func (s *GetUdfResponse) Validate() error {
	return dara.Validate(s)
}
