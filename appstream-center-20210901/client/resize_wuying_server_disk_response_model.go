// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeWuyingServerDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResizeWuyingServerDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResizeWuyingServerDiskResponse
	GetStatusCode() *int32
	SetBody(v *ResizeWuyingServerDiskResponseBody) *ResizeWuyingServerDiskResponse
	GetBody() *ResizeWuyingServerDiskResponseBody
}

type ResizeWuyingServerDiskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeWuyingServerDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResizeWuyingServerDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s ResizeWuyingServerDiskResponse) GoString() string {
	return s.String()
}

func (s *ResizeWuyingServerDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResizeWuyingServerDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResizeWuyingServerDiskResponse) GetBody() *ResizeWuyingServerDiskResponseBody {
	return s.Body
}

func (s *ResizeWuyingServerDiskResponse) SetHeaders(v map[string]*string) *ResizeWuyingServerDiskResponse {
	s.Headers = v
	return s
}

func (s *ResizeWuyingServerDiskResponse) SetStatusCode(v int32) *ResizeWuyingServerDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeWuyingServerDiskResponse) SetBody(v *ResizeWuyingServerDiskResponseBody) *ResizeWuyingServerDiskResponse {
	s.Body = v
	return s
}

func (s *ResizeWuyingServerDiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
