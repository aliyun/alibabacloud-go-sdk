// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetJobConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatasetJobConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatasetJobConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListDatasetJobConfigsResponseBody) *ListDatasetJobConfigsResponse
	GetBody() *ListDatasetJobConfigsResponseBody
}

type ListDatasetJobConfigsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatasetJobConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatasetJobConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetJobConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListDatasetJobConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatasetJobConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatasetJobConfigsResponse) GetBody() *ListDatasetJobConfigsResponseBody {
	return s.Body
}

func (s *ListDatasetJobConfigsResponse) SetHeaders(v map[string]*string) *ListDatasetJobConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListDatasetJobConfigsResponse) SetStatusCode(v int32) *ListDatasetJobConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatasetJobConfigsResponse) SetBody(v *ListDatasetJobConfigsResponseBody) *ListDatasetJobConfigsResponse {
	s.Body = v
	return s
}

func (s *ListDatasetJobConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
