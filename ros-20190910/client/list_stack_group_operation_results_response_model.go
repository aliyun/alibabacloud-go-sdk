// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackGroupOperationResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStackGroupOperationResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStackGroupOperationResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListStackGroupOperationResultsResponseBody) *ListStackGroupOperationResultsResponse
	GetBody() *ListStackGroupOperationResultsResponseBody
}

type ListStackGroupOperationResultsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStackGroupOperationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStackGroupOperationResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStackGroupOperationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStackGroupOperationResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStackGroupOperationResultsResponse) GetBody() *ListStackGroupOperationResultsResponseBody {
	return s.Body
}

func (s *ListStackGroupOperationResultsResponse) SetHeaders(v map[string]*string) *ListStackGroupOperationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListStackGroupOperationResultsResponse) SetStatusCode(v int32) *ListStackGroupOperationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackGroupOperationResultsResponse) SetBody(v *ListStackGroupOperationResultsResponseBody) *ListStackGroupOperationResultsResponse {
	s.Body = v
	return s
}

func (s *ListStackGroupOperationResultsResponse) Validate() error {
	return dara.Validate(s)
}
