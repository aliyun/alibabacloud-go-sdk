// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCacheClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopCacheClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopCacheClusterResponse
	GetStatusCode() *int32
	SetBody(v *StopCacheClusterResponseBody) *StopCacheClusterResponse
	GetBody() *StopCacheClusterResponseBody
}

type StopCacheClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopCacheClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopCacheClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s StopCacheClusterResponse) GoString() string {
	return s.String()
}

func (s *StopCacheClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopCacheClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopCacheClusterResponse) GetBody() *StopCacheClusterResponseBody {
	return s.Body
}

func (s *StopCacheClusterResponse) SetHeaders(v map[string]*string) *StopCacheClusterResponse {
	s.Headers = v
	return s
}

func (s *StopCacheClusterResponse) SetStatusCode(v int32) *StopCacheClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *StopCacheClusterResponse) SetBody(v *StopCacheClusterResponseBody) *StopCacheClusterResponse {
	s.Body = v
	return s
}

func (s *StopCacheClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
