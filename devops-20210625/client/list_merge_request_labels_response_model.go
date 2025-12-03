// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMergeRequestLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMergeRequestLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMergeRequestLabelsResponse
	GetStatusCode() *int32
	SetBody(v *ListMergeRequestLabelsResponseBody) *ListMergeRequestLabelsResponse
	GetBody() *ListMergeRequestLabelsResponseBody
}

type ListMergeRequestLabelsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMergeRequestLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMergeRequestLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestLabelsResponse) GoString() string {
	return s.String()
}

func (s *ListMergeRequestLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMergeRequestLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMergeRequestLabelsResponse) GetBody() *ListMergeRequestLabelsResponseBody {
	return s.Body
}

func (s *ListMergeRequestLabelsResponse) SetHeaders(v map[string]*string) *ListMergeRequestLabelsResponse {
	s.Headers = v
	return s
}

func (s *ListMergeRequestLabelsResponse) SetStatusCode(v int32) *ListMergeRequestLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMergeRequestLabelsResponse) SetBody(v *ListMergeRequestLabelsResponseBody) *ListMergeRequestLabelsResponse {
	s.Body = v
	return s
}

func (s *ListMergeRequestLabelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
