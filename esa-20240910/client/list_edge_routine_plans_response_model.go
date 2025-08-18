// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeRoutinePlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEdgeRoutinePlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEdgeRoutinePlansResponse
	GetStatusCode() *int32
	SetBody(v *ListEdgeRoutinePlansResponseBody) *ListEdgeRoutinePlansResponse
	GetBody() *ListEdgeRoutinePlansResponseBody
}

type ListEdgeRoutinePlansResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEdgeRoutinePlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEdgeRoutinePlansResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeRoutinePlansResponse) GoString() string {
	return s.String()
}

func (s *ListEdgeRoutinePlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEdgeRoutinePlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEdgeRoutinePlansResponse) GetBody() *ListEdgeRoutinePlansResponseBody {
	return s.Body
}

func (s *ListEdgeRoutinePlansResponse) SetHeaders(v map[string]*string) *ListEdgeRoutinePlansResponse {
	s.Headers = v
	return s
}

func (s *ListEdgeRoutinePlansResponse) SetStatusCode(v int32) *ListEdgeRoutinePlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEdgeRoutinePlansResponse) SetBody(v *ListEdgeRoutinePlansResponseBody) *ListEdgeRoutinePlansResponse {
	s.Body = v
	return s
}

func (s *ListEdgeRoutinePlansResponse) Validate() error {
	return dara.Validate(s)
}
