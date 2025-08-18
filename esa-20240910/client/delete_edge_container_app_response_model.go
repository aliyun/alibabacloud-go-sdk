// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeContainerAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEdgeContainerAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEdgeContainerAppResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEdgeContainerAppResponseBody) *DeleteEdgeContainerAppResponse
	GetBody() *DeleteEdgeContainerAppResponseBody
}

type DeleteEdgeContainerAppResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEdgeContainerAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEdgeContainerAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeContainerAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteEdgeContainerAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEdgeContainerAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEdgeContainerAppResponse) GetBody() *DeleteEdgeContainerAppResponseBody {
	return s.Body
}

func (s *DeleteEdgeContainerAppResponse) SetHeaders(v map[string]*string) *DeleteEdgeContainerAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteEdgeContainerAppResponse) SetStatusCode(v int32) *DeleteEdgeContainerAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEdgeContainerAppResponse) SetBody(v *DeleteEdgeContainerAppResponseBody) *DeleteEdgeContainerAppResponse {
	s.Body = v
	return s
}

func (s *DeleteEdgeContainerAppResponse) Validate() error {
	return dara.Validate(s)
}
