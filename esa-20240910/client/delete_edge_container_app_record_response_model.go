// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeContainerAppRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEdgeContainerAppRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEdgeContainerAppRecordResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEdgeContainerAppRecordResponseBody) *DeleteEdgeContainerAppRecordResponse
	GetBody() *DeleteEdgeContainerAppRecordResponseBody
}

type DeleteEdgeContainerAppRecordResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEdgeContainerAppRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEdgeContainerAppRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeContainerAppRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteEdgeContainerAppRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEdgeContainerAppRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEdgeContainerAppRecordResponse) GetBody() *DeleteEdgeContainerAppRecordResponseBody {
	return s.Body
}

func (s *DeleteEdgeContainerAppRecordResponse) SetHeaders(v map[string]*string) *DeleteEdgeContainerAppRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteEdgeContainerAppRecordResponse) SetStatusCode(v int32) *DeleteEdgeContainerAppRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEdgeContainerAppRecordResponse) SetBody(v *DeleteEdgeContainerAppRecordResponseBody) *DeleteEdgeContainerAppRecordResponse {
	s.Body = v
	return s
}

func (s *DeleteEdgeContainerAppRecordResponse) Validate() error {
	return dara.Validate(s)
}
