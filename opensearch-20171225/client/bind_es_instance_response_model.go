// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindEsInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindEsInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindEsInstanceResponse
	GetStatusCode() *int32
	SetBody(v *BindEsInstanceResponseBody) *BindEsInstanceResponse
	GetBody() *BindEsInstanceResponseBody
}

type BindEsInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindEsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindEsInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s BindEsInstanceResponse) GoString() string {
	return s.String()
}

func (s *BindEsInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindEsInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindEsInstanceResponse) GetBody() *BindEsInstanceResponseBody {
	return s.Body
}

func (s *BindEsInstanceResponse) SetHeaders(v map[string]*string) *BindEsInstanceResponse {
	s.Headers = v
	return s
}

func (s *BindEsInstanceResponse) SetStatusCode(v int32) *BindEsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *BindEsInstanceResponse) SetBody(v *BindEsInstanceResponseBody) *BindEsInstanceResponse {
	s.Body = v
	return s
}

func (s *BindEsInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
