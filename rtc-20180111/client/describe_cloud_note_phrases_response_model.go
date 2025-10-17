// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudNotePhrasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudNotePhrasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudNotePhrasesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudNotePhrasesResponseBody) *DescribeCloudNotePhrasesResponse
	GetBody() *DescribeCloudNotePhrasesResponseBody
}

type DescribeCloudNotePhrasesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudNotePhrasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudNotePhrasesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudNotePhrasesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudNotePhrasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudNotePhrasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudNotePhrasesResponse) GetBody() *DescribeCloudNotePhrasesResponseBody {
	return s.Body
}

func (s *DescribeCloudNotePhrasesResponse) SetHeaders(v map[string]*string) *DescribeCloudNotePhrasesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudNotePhrasesResponse) SetStatusCode(v int32) *DescribeCloudNotePhrasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudNotePhrasesResponse) SetBody(v *DescribeCloudNotePhrasesResponseBody) *DescribeCloudNotePhrasesResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudNotePhrasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
