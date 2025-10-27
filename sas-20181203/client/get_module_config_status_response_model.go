// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModuleConfigStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModuleConfigStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModuleConfigStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetModuleConfigStatusResponseBody) *GetModuleConfigStatusResponse
	GetBody() *GetModuleConfigStatusResponseBody
}

type GetModuleConfigStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModuleConfigStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModuleConfigStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModuleConfigStatusResponse) GoString() string {
	return s.String()
}

func (s *GetModuleConfigStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModuleConfigStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModuleConfigStatusResponse) GetBody() *GetModuleConfigStatusResponseBody {
	return s.Body
}

func (s *GetModuleConfigStatusResponse) SetHeaders(v map[string]*string) *GetModuleConfigStatusResponse {
	s.Headers = v
	return s
}

func (s *GetModuleConfigStatusResponse) SetStatusCode(v int32) *GetModuleConfigStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModuleConfigStatusResponse) SetBody(v *GetModuleConfigStatusResponseBody) *GetModuleConfigStatusResponse {
	s.Body = v
	return s
}

func (s *GetModuleConfigStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
