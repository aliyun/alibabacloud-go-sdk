// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceModuleInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceModuleInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceModuleInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceModuleInfoResponseBody) *GetInstanceModuleInfoResponse
	GetBody() *GetInstanceModuleInfoResponseBody
}

type GetInstanceModuleInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceModuleInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceModuleInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceModuleInfoResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceModuleInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceModuleInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceModuleInfoResponse) GetBody() *GetInstanceModuleInfoResponseBody {
	return s.Body
}

func (s *GetInstanceModuleInfoResponse) SetHeaders(v map[string]*string) *GetInstanceModuleInfoResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceModuleInfoResponse) SetStatusCode(v int32) *GetInstanceModuleInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceModuleInfoResponse) SetBody(v *GetInstanceModuleInfoResponseBody) *GetInstanceModuleInfoResponse {
	s.Body = v
	return s
}

func (s *GetInstanceModuleInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
