// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcubeHotpatchResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMcubeHotpatchResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMcubeHotpatchResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMcubeHotpatchResourceResponseBody) *DeleteMcubeHotpatchResourceResponse
	GetBody() *DeleteMcubeHotpatchResourceResponseBody
}

type DeleteMcubeHotpatchResourceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMcubeHotpatchResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMcubeHotpatchResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeHotpatchResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteMcubeHotpatchResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMcubeHotpatchResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMcubeHotpatchResourceResponse) GetBody() *DeleteMcubeHotpatchResourceResponseBody {
	return s.Body
}

func (s *DeleteMcubeHotpatchResourceResponse) SetHeaders(v map[string]*string) *DeleteMcubeHotpatchResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteMcubeHotpatchResourceResponse) SetStatusCode(v int32) *DeleteMcubeHotpatchResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMcubeHotpatchResourceResponse) SetBody(v *DeleteMcubeHotpatchResourceResponseBody) *DeleteMcubeHotpatchResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteMcubeHotpatchResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
