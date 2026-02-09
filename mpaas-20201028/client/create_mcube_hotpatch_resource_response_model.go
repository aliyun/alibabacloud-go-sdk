// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeHotpatchResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcubeHotpatchResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcubeHotpatchResourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcubeHotpatchResourceResponseBody) *CreateMcubeHotpatchResourceResponse
	GetBody() *CreateMcubeHotpatchResourceResponseBody
}

type CreateMcubeHotpatchResourceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcubeHotpatchResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcubeHotpatchResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeHotpatchResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateMcubeHotpatchResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcubeHotpatchResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcubeHotpatchResourceResponse) GetBody() *CreateMcubeHotpatchResourceResponseBody {
	return s.Body
}

func (s *CreateMcubeHotpatchResourceResponse) SetHeaders(v map[string]*string) *CreateMcubeHotpatchResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateMcubeHotpatchResourceResponse) SetStatusCode(v int32) *CreateMcubeHotpatchResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcubeHotpatchResourceResponse) SetBody(v *CreateMcubeHotpatchResourceResponseBody) *CreateMcubeHotpatchResourceResponse {
	s.Body = v
	return s
}

func (s *CreateMcubeHotpatchResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
