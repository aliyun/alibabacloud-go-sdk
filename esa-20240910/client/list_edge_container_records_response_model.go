// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeContainerRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEdgeContainerRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEdgeContainerRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListEdgeContainerRecordsResponseBody) *ListEdgeContainerRecordsResponse
	GetBody() *ListEdgeContainerRecordsResponseBody
}

type ListEdgeContainerRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEdgeContainerRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEdgeContainerRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEdgeContainerRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEdgeContainerRecordsResponse) GetBody() *ListEdgeContainerRecordsResponseBody {
	return s.Body
}

func (s *ListEdgeContainerRecordsResponse) SetHeaders(v map[string]*string) *ListEdgeContainerRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListEdgeContainerRecordsResponse) SetStatusCode(v int32) *ListEdgeContainerRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEdgeContainerRecordsResponse) SetBody(v *ListEdgeContainerRecordsResponseBody) *ListEdgeContainerRecordsResponse {
	s.Body = v
	return s
}

func (s *ListEdgeContainerRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
