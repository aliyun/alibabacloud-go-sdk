// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResizeClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResizeClusterResponse
	GetStatusCode() *int32
	SetBody(v *ResizeClusterResponseBody) *ResizeClusterResponse
	GetBody() *ResizeClusterResponseBody
}

type ResizeClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResizeClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ResizeClusterResponse) GoString() string {
	return s.String()
}

func (s *ResizeClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResizeClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResizeClusterResponse) GetBody() *ResizeClusterResponseBody {
	return s.Body
}

func (s *ResizeClusterResponse) SetHeaders(v map[string]*string) *ResizeClusterResponse {
	s.Headers = v
	return s
}

func (s *ResizeClusterResponse) SetStatusCode(v int32) *ResizeClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeClusterResponse) SetBody(v *ResizeClusterResponseBody) *ResizeClusterResponse {
	s.Body = v
	return s
}

func (s *ResizeClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
