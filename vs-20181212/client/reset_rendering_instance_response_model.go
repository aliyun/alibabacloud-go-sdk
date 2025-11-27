// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetRenderingInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetRenderingInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetRenderingInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ResetRenderingInstanceResponseBody) *ResetRenderingInstanceResponse
	GetBody() *ResetRenderingInstanceResponseBody
}

type ResetRenderingInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetRenderingInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetRenderingInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetRenderingInstanceResponse) GoString() string {
	return s.String()
}

func (s *ResetRenderingInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetRenderingInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetRenderingInstanceResponse) GetBody() *ResetRenderingInstanceResponseBody {
	return s.Body
}

func (s *ResetRenderingInstanceResponse) SetHeaders(v map[string]*string) *ResetRenderingInstanceResponse {
	s.Headers = v
	return s
}

func (s *ResetRenderingInstanceResponse) SetStatusCode(v int32) *ResetRenderingInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetRenderingInstanceResponse) SetBody(v *ResetRenderingInstanceResponseBody) *ResetRenderingInstanceResponse {
	s.Body = v
	return s
}

func (s *ResetRenderingInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
