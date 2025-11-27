// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParentPlatformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteParentPlatformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteParentPlatformResponse
	GetStatusCode() *int32
	SetBody(v *DeleteParentPlatformResponseBody) *DeleteParentPlatformResponse
	GetBody() *DeleteParentPlatformResponseBody
}

type DeleteParentPlatformResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteParentPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteParentPlatformResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteParentPlatformResponse) GoString() string {
	return s.String()
}

func (s *DeleteParentPlatformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteParentPlatformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteParentPlatformResponse) GetBody() *DeleteParentPlatformResponseBody {
	return s.Body
}

func (s *DeleteParentPlatformResponse) SetHeaders(v map[string]*string) *DeleteParentPlatformResponse {
	s.Headers = v
	return s
}

func (s *DeleteParentPlatformResponse) SetStatusCode(v int32) *DeleteParentPlatformResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteParentPlatformResponse) SetBody(v *DeleteParentPlatformResponseBody) *DeleteParentPlatformResponse {
	s.Body = v
	return s
}

func (s *DeleteParentPlatformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
