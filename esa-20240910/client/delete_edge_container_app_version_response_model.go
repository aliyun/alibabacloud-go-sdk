// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeContainerAppVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEdgeContainerAppVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEdgeContainerAppVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEdgeContainerAppVersionResponseBody) *DeleteEdgeContainerAppVersionResponse
	GetBody() *DeleteEdgeContainerAppVersionResponseBody
}

type DeleteEdgeContainerAppVersionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEdgeContainerAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEdgeContainerAppVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeContainerAppVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteEdgeContainerAppVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEdgeContainerAppVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEdgeContainerAppVersionResponse) GetBody() *DeleteEdgeContainerAppVersionResponseBody {
	return s.Body
}

func (s *DeleteEdgeContainerAppVersionResponse) SetHeaders(v map[string]*string) *DeleteEdgeContainerAppVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteEdgeContainerAppVersionResponse) SetStatusCode(v int32) *DeleteEdgeContainerAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEdgeContainerAppVersionResponse) SetBody(v *DeleteEdgeContainerAppVersionResponseBody) *DeleteEdgeContainerAppVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteEdgeContainerAppVersionResponse) Validate() error {
	return dara.Validate(s)
}
