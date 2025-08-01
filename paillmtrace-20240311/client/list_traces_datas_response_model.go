// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTracesDatasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTracesDatasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTracesDatasResponse
	GetStatusCode() *int32
	SetBody(v *ListTracesDatasResponseBody) *ListTracesDatasResponse
	GetBody() *ListTracesDatasResponseBody
}

type ListTracesDatasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTracesDatasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTracesDatasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTracesDatasResponse) GoString() string {
	return s.String()
}

func (s *ListTracesDatasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTracesDatasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTracesDatasResponse) GetBody() *ListTracesDatasResponseBody {
	return s.Body
}

func (s *ListTracesDatasResponse) SetHeaders(v map[string]*string) *ListTracesDatasResponse {
	s.Headers = v
	return s
}

func (s *ListTracesDatasResponse) SetStatusCode(v int32) *ListTracesDatasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTracesDatasResponse) SetBody(v *ListTracesDatasResponseBody) *ListTracesDatasResponse {
	s.Body = v
	return s
}

func (s *ListTracesDatasResponse) Validate() error {
	return dara.Validate(s)
}
