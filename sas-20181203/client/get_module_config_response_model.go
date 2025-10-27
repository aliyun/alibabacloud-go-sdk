// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModuleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModuleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModuleConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetModuleConfigResponseBody) *GetModuleConfigResponse
	GetBody() *GetModuleConfigResponseBody
}

type GetModuleConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModuleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModuleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModuleConfigResponse) GoString() string {
	return s.String()
}

func (s *GetModuleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModuleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModuleConfigResponse) GetBody() *GetModuleConfigResponseBody {
	return s.Body
}

func (s *GetModuleConfigResponse) SetHeaders(v map[string]*string) *GetModuleConfigResponse {
	s.Headers = v
	return s
}

func (s *GetModuleConfigResponse) SetStatusCode(v int32) *GetModuleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModuleConfigResponse) SetBody(v *GetModuleConfigResponseBody) *GetModuleConfigResponse {
	s.Body = v
	return s
}

func (s *GetModuleConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
