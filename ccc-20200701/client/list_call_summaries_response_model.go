// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallSummariesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCallSummariesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCallSummariesResponse
	GetStatusCode() *int32
	SetBody(v *ListCallSummariesResponseBody) *ListCallSummariesResponse
	GetBody() *ListCallSummariesResponseBody
}

type ListCallSummariesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCallSummariesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCallSummariesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCallSummariesResponse) GoString() string {
	return s.String()
}

func (s *ListCallSummariesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCallSummariesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCallSummariesResponse) GetBody() *ListCallSummariesResponseBody {
	return s.Body
}

func (s *ListCallSummariesResponse) SetHeaders(v map[string]*string) *ListCallSummariesResponse {
	s.Headers = v
	return s
}

func (s *ListCallSummariesResponse) SetStatusCode(v int32) *ListCallSummariesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCallSummariesResponse) SetBody(v *ListCallSummariesResponseBody) *ListCallSummariesResponse {
	s.Body = v
	return s
}

func (s *ListCallSummariesResponse) Validate() error {
	return dara.Validate(s)
}
