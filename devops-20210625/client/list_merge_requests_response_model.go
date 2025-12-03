// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMergeRequestsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMergeRequestsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMergeRequestsResponse
	GetStatusCode() *int32
	SetBody(v *ListMergeRequestsResponseBody) *ListMergeRequestsResponse
	GetBody() *ListMergeRequestsResponseBody
}

type ListMergeRequestsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMergeRequestsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMergeRequestsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestsResponse) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMergeRequestsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMergeRequestsResponse) GetBody() *ListMergeRequestsResponseBody {
	return s.Body
}

func (s *ListMergeRequestsResponse) SetHeaders(v map[string]*string) *ListMergeRequestsResponse {
	s.Headers = v
	return s
}

func (s *ListMergeRequestsResponse) SetStatusCode(v int32) *ListMergeRequestsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMergeRequestsResponse) SetBody(v *ListMergeRequestsResponseBody) *ListMergeRequestsResponse {
	s.Body = v
	return s
}

func (s *ListMergeRequestsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
