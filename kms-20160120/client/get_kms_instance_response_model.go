// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKmsInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKmsInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKmsInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetKmsInstanceResponseBody) *GetKmsInstanceResponse
	GetBody() *GetKmsInstanceResponseBody
}

type GetKmsInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKmsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKmsInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKmsInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetKmsInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKmsInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKmsInstanceResponse) GetBody() *GetKmsInstanceResponseBody {
	return s.Body
}

func (s *GetKmsInstanceResponse) SetHeaders(v map[string]*string) *GetKmsInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetKmsInstanceResponse) SetStatusCode(v int32) *GetKmsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKmsInstanceResponse) SetBody(v *GetKmsInstanceResponseBody) *GetKmsInstanceResponse {
	s.Body = v
	return s
}

func (s *GetKmsInstanceResponse) Validate() error {
	return dara.Validate(s)
}
