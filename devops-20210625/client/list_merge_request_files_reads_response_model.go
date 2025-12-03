// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMergeRequestFilesReadsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMergeRequestFilesReadsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMergeRequestFilesReadsResponse
	GetStatusCode() *int32
	SetBody(v *ListMergeRequestFilesReadsResponseBody) *ListMergeRequestFilesReadsResponse
	GetBody() *ListMergeRequestFilesReadsResponseBody
}

type ListMergeRequestFilesReadsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMergeRequestFilesReadsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMergeRequestFilesReadsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestFilesReadsResponse) GoString() string {
	return s.String()
}

func (s *ListMergeRequestFilesReadsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMergeRequestFilesReadsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMergeRequestFilesReadsResponse) GetBody() *ListMergeRequestFilesReadsResponseBody {
	return s.Body
}

func (s *ListMergeRequestFilesReadsResponse) SetHeaders(v map[string]*string) *ListMergeRequestFilesReadsResponse {
	s.Headers = v
	return s
}

func (s *ListMergeRequestFilesReadsResponse) SetStatusCode(v int32) *ListMergeRequestFilesReadsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMergeRequestFilesReadsResponse) SetBody(v *ListMergeRequestFilesReadsResponseBody) *ListMergeRequestFilesReadsResponse {
	s.Body = v
	return s
}

func (s *ListMergeRequestFilesReadsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
