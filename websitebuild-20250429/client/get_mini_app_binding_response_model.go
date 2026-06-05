// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMiniAppBindingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMiniAppBindingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMiniAppBindingResponse
	GetStatusCode() *int32
	SetBody(v *GetMiniAppBindingResponseBody) *GetMiniAppBindingResponse
	GetBody() *GetMiniAppBindingResponseBody
}

type GetMiniAppBindingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMiniAppBindingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMiniAppBindingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMiniAppBindingResponse) GoString() string {
	return s.String()
}

func (s *GetMiniAppBindingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMiniAppBindingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMiniAppBindingResponse) GetBody() *GetMiniAppBindingResponseBody {
	return s.Body
}

func (s *GetMiniAppBindingResponse) SetHeaders(v map[string]*string) *GetMiniAppBindingResponse {
	s.Headers = v
	return s
}

func (s *GetMiniAppBindingResponse) SetStatusCode(v int32) *GetMiniAppBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMiniAppBindingResponse) SetBody(v *GetMiniAppBindingResponseBody) *GetMiniAppBindingResponse {
	s.Body = v
	return s
}

func (s *GetMiniAppBindingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
