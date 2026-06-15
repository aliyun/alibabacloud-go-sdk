// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCacheClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartCacheClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartCacheClusterResponse
	GetStatusCode() *int32
	SetBody(v *StartCacheClusterResponseBody) *StartCacheClusterResponse
	GetBody() *StartCacheClusterResponseBody
}

type StartCacheClusterResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartCacheClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartCacheClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s StartCacheClusterResponse) GoString() string {
	return s.String()
}

func (s *StartCacheClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartCacheClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartCacheClusterResponse) GetBody() *StartCacheClusterResponseBody {
	return s.Body
}

func (s *StartCacheClusterResponse) SetHeaders(v map[string]*string) *StartCacheClusterResponse {
	s.Headers = v
	return s
}

func (s *StartCacheClusterResponse) SetStatusCode(v int32) *StartCacheClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *StartCacheClusterResponse) SetBody(v *StartCacheClusterResponseBody) *StartCacheClusterResponse {
	s.Body = v
	return s
}

func (s *StartCacheClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
