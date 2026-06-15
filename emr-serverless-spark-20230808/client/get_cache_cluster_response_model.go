// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCacheClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCacheClusterResponse
	GetStatusCode() *int32
	SetBody(v *GetCacheClusterResponseBody) *GetCacheClusterResponse
	GetBody() *GetCacheClusterResponseBody
}

type GetCacheClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCacheClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCacheClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCacheClusterResponse) GoString() string {
	return s.String()
}

func (s *GetCacheClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCacheClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCacheClusterResponse) GetBody() *GetCacheClusterResponseBody {
	return s.Body
}

func (s *GetCacheClusterResponse) SetHeaders(v map[string]*string) *GetCacheClusterResponse {
	s.Headers = v
	return s
}

func (s *GetCacheClusterResponse) SetStatusCode(v int32) *GetCacheClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCacheClusterResponse) SetBody(v *GetCacheClusterResponseBody) *GetCacheClusterResponse {
	s.Body = v
	return s
}

func (s *GetCacheClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
