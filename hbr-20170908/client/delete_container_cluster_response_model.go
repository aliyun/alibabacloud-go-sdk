// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContainerClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContainerClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContainerClusterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContainerClusterResponseBody) *DeleteContainerClusterResponse
	GetBody() *DeleteContainerClusterResponseBody
}

type DeleteContainerClusterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContainerClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContainerClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContainerClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteContainerClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContainerClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContainerClusterResponse) GetBody() *DeleteContainerClusterResponseBody {
	return s.Body
}

func (s *DeleteContainerClusterResponse) SetHeaders(v map[string]*string) *DeleteContainerClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteContainerClusterResponse) SetStatusCode(v int32) *DeleteContainerClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContainerClusterResponse) SetBody(v *DeleteContainerClusterResponseBody) *DeleteContainerClusterResponse {
	s.Body = v
	return s
}

func (s *DeleteContainerClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
