// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCVClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRCVClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRCVClusterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRCVClusterResponseBody) *DeleteRCVClusterResponse
	GetBody() *DeleteRCVClusterResponseBody
}

type DeleteRCVClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRCVClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRCVClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCVClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteRCVClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRCVClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRCVClusterResponse) GetBody() *DeleteRCVClusterResponseBody {
	return s.Body
}

func (s *DeleteRCVClusterResponse) SetHeaders(v map[string]*string) *DeleteRCVClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteRCVClusterResponse) SetStatusCode(v int32) *DeleteRCVClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRCVClusterResponse) SetBody(v *DeleteRCVClusterResponseBody) *DeleteRCVClusterResponse {
	s.Body = v
	return s
}

func (s *DeleteRCVClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
