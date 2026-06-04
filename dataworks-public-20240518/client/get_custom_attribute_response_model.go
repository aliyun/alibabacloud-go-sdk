// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomAttributeResponseBody) *GetCustomAttributeResponse
	GetBody() *GetCustomAttributeResponseBody
}

type GetCustomAttributeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetCustomAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomAttributeResponse) GetBody() *GetCustomAttributeResponseBody {
	return s.Body
}

func (s *GetCustomAttributeResponse) SetHeaders(v map[string]*string) *GetCustomAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetCustomAttributeResponse) SetStatusCode(v int32) *GetCustomAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomAttributeResponse) SetBody(v *GetCustomAttributeResponseBody) *GetCustomAttributeResponse {
	s.Body = v
	return s
}

func (s *GetCustomAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
