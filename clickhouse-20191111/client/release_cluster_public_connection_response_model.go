// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseClusterPublicConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseClusterPublicConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseClusterPublicConnectionResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseClusterPublicConnectionResponseBody) *ReleaseClusterPublicConnectionResponse
	GetBody() *ReleaseClusterPublicConnectionResponseBody
}

type ReleaseClusterPublicConnectionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseClusterPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseClusterPublicConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseClusterPublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseClusterPublicConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseClusterPublicConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseClusterPublicConnectionResponse) GetBody() *ReleaseClusterPublicConnectionResponseBody {
	return s.Body
}

func (s *ReleaseClusterPublicConnectionResponse) SetHeaders(v map[string]*string) *ReleaseClusterPublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseClusterPublicConnectionResponse) SetStatusCode(v int32) *ReleaseClusterPublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseClusterPublicConnectionResponse) SetBody(v *ReleaseClusterPublicConnectionResponseBody) *ReleaseClusterPublicConnectionResponse {
	s.Body = v
	return s
}

func (s *ReleaseClusterPublicConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
