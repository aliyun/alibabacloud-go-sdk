// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRenderingInstanceCommandsStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRenderingInstanceCommandsStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRenderingInstanceCommandsStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetRenderingInstanceCommandsStatusResponseBody) *GetRenderingInstanceCommandsStatusResponse
	GetBody() *GetRenderingInstanceCommandsStatusResponseBody
}

type GetRenderingInstanceCommandsStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRenderingInstanceCommandsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRenderingInstanceCommandsStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRenderingInstanceCommandsStatusResponse) GoString() string {
	return s.String()
}

func (s *GetRenderingInstanceCommandsStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRenderingInstanceCommandsStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRenderingInstanceCommandsStatusResponse) GetBody() *GetRenderingInstanceCommandsStatusResponseBody {
	return s.Body
}

func (s *GetRenderingInstanceCommandsStatusResponse) SetHeaders(v map[string]*string) *GetRenderingInstanceCommandsStatusResponse {
	s.Headers = v
	return s
}

func (s *GetRenderingInstanceCommandsStatusResponse) SetStatusCode(v int32) *GetRenderingInstanceCommandsStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRenderingInstanceCommandsStatusResponse) SetBody(v *GetRenderingInstanceCommandsStatusResponseBody) *GetRenderingInstanceCommandsStatusResponse {
	s.Body = v
	return s
}

func (s *GetRenderingInstanceCommandsStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
