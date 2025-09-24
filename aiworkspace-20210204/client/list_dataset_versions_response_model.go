// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatasetVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatasetVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListDatasetVersionsResponseBody) *ListDatasetVersionsResponse
	GetBody() *ListDatasetVersionsResponseBody
}

type ListDatasetVersionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatasetVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatasetVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListDatasetVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatasetVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatasetVersionsResponse) GetBody() *ListDatasetVersionsResponseBody {
	return s.Body
}

func (s *ListDatasetVersionsResponse) SetHeaders(v map[string]*string) *ListDatasetVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListDatasetVersionsResponse) SetStatusCode(v int32) *ListDatasetVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatasetVersionsResponse) SetBody(v *ListDatasetVersionsResponseBody) *ListDatasetVersionsResponse {
	s.Body = v
	return s
}

func (s *ListDatasetVersionsResponse) Validate() error {
	return dara.Validate(s)
}
