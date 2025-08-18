// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOriginProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOriginProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOriginProtectionResponse
	GetStatusCode() *int32
	SetBody(v *CreateOriginProtectionResponseBody) *CreateOriginProtectionResponse
	GetBody() *CreateOriginProtectionResponseBody
}

type CreateOriginProtectionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOriginProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOriginProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOriginProtectionResponse) GoString() string {
	return s.String()
}

func (s *CreateOriginProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOriginProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOriginProtectionResponse) GetBody() *CreateOriginProtectionResponseBody {
	return s.Body
}

func (s *CreateOriginProtectionResponse) SetHeaders(v map[string]*string) *CreateOriginProtectionResponse {
	s.Headers = v
	return s
}

func (s *CreateOriginProtectionResponse) SetStatusCode(v int32) *CreateOriginProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOriginProtectionResponse) SetBody(v *CreateOriginProtectionResponseBody) *CreateOriginProtectionResponse {
	s.Body = v
	return s
}

func (s *CreateOriginProtectionResponse) Validate() error {
	return dara.Validate(s)
}
