// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLocationDateClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLocationDateClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLocationDateClusterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLocationDateClusterResponseBody) *DeleteLocationDateClusterResponse
	GetBody() *DeleteLocationDateClusterResponseBody
}

type DeleteLocationDateClusterResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLocationDateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLocationDateClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLocationDateClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteLocationDateClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLocationDateClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLocationDateClusterResponse) GetBody() *DeleteLocationDateClusterResponseBody {
	return s.Body
}

func (s *DeleteLocationDateClusterResponse) SetHeaders(v map[string]*string) *DeleteLocationDateClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteLocationDateClusterResponse) SetStatusCode(v int32) *DeleteLocationDateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLocationDateClusterResponse) SetBody(v *DeleteLocationDateClusterResponseBody) *DeleteLocationDateClusterResponse {
	s.Body = v
	return s
}

func (s *DeleteLocationDateClusterResponse) Validate() error {
	return dara.Validate(s)
}
