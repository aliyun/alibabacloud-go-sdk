// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeContainerAppRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEdgeContainerAppRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEdgeContainerAppRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListEdgeContainerAppRecordsResponseBody) *ListEdgeContainerAppRecordsResponse
	GetBody() *ListEdgeContainerAppRecordsResponseBody
}

type ListEdgeContainerAppRecordsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEdgeContainerAppRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEdgeContainerAppRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEdgeContainerAppRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEdgeContainerAppRecordsResponse) GetBody() *ListEdgeContainerAppRecordsResponseBody {
	return s.Body
}

func (s *ListEdgeContainerAppRecordsResponse) SetHeaders(v map[string]*string) *ListEdgeContainerAppRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListEdgeContainerAppRecordsResponse) SetStatusCode(v int32) *ListEdgeContainerAppRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEdgeContainerAppRecordsResponse) SetBody(v *ListEdgeContainerAppRecordsResponseBody) *ListEdgeContainerAppRecordsResponse {
	s.Body = v
	return s
}

func (s *ListEdgeContainerAppRecordsResponse) Validate() error {
	return dara.Validate(s)
}
