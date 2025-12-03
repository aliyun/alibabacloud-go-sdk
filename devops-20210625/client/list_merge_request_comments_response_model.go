// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMergeRequestCommentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMergeRequestCommentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMergeRequestCommentsResponse
	GetStatusCode() *int32
	SetBody(v *ListMergeRequestCommentsResponseBody) *ListMergeRequestCommentsResponse
	GetBody() *ListMergeRequestCommentsResponseBody
}

type ListMergeRequestCommentsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMergeRequestCommentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMergeRequestCommentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestCommentsResponse) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMergeRequestCommentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMergeRequestCommentsResponse) GetBody() *ListMergeRequestCommentsResponseBody {
	return s.Body
}

func (s *ListMergeRequestCommentsResponse) SetHeaders(v map[string]*string) *ListMergeRequestCommentsResponse {
	s.Headers = v
	return s
}

func (s *ListMergeRequestCommentsResponse) SetStatusCode(v int32) *ListMergeRequestCommentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMergeRequestCommentsResponse) SetBody(v *ListMergeRequestCommentsResponseBody) *ListMergeRequestCommentsResponse {
	s.Body = v
	return s
}

func (s *ListMergeRequestCommentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
