// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseRenderingInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseRenderingInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseRenderingInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseRenderingInstanceResponseBody) *ReleaseRenderingInstanceResponse
	GetBody() *ReleaseRenderingInstanceResponseBody
}

type ReleaseRenderingInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseRenderingInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseRenderingInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseRenderingInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseRenderingInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseRenderingInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseRenderingInstanceResponse) GetBody() *ReleaseRenderingInstanceResponseBody {
	return s.Body
}

func (s *ReleaseRenderingInstanceResponse) SetHeaders(v map[string]*string) *ReleaseRenderingInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseRenderingInstanceResponse) SetStatusCode(v int32) *ReleaseRenderingInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseRenderingInstanceResponse) SetBody(v *ReleaseRenderingInstanceResponseBody) *ReleaseRenderingInstanceResponse {
	s.Body = v
	return s
}

func (s *ReleaseRenderingInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
