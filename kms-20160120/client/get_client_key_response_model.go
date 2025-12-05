// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClientKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClientKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetClientKeyResponseBody) *GetClientKeyResponse
	GetBody() *GetClientKeyResponseBody
}

type GetClientKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClientKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClientKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClientKeyResponse) GoString() string {
	return s.String()
}

func (s *GetClientKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClientKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClientKeyResponse) GetBody() *GetClientKeyResponseBody {
	return s.Body
}

func (s *GetClientKeyResponse) SetHeaders(v map[string]*string) *GetClientKeyResponse {
	s.Headers = v
	return s
}

func (s *GetClientKeyResponse) SetStatusCode(v int32) *GetClientKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClientKeyResponse) SetBody(v *GetClientKeyResponseBody) *GetClientKeyResponse {
	s.Body = v
	return s
}

func (s *GetClientKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
