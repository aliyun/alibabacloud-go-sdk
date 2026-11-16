// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetVideoCharacterReplaceJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetVideoCharacterReplaceJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetVideoCharacterReplaceJobsResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetVideoCharacterReplaceJobsResponseBody) *BatchGetVideoCharacterReplaceJobsResponse
	GetBody() *BatchGetVideoCharacterReplaceJobsResponseBody
}

type BatchGetVideoCharacterReplaceJobsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetVideoCharacterReplaceJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetVideoCharacterReplaceJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetVideoCharacterReplaceJobsResponse) GoString() string {
	return s.String()
}

func (s *BatchGetVideoCharacterReplaceJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetVideoCharacterReplaceJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetVideoCharacterReplaceJobsResponse) GetBody() *BatchGetVideoCharacterReplaceJobsResponseBody {
	return s.Body
}

func (s *BatchGetVideoCharacterReplaceJobsResponse) SetHeaders(v map[string]*string) *BatchGetVideoCharacterReplaceJobsResponse {
	s.Headers = v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponse) SetStatusCode(v int32) *BatchGetVideoCharacterReplaceJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponse) SetBody(v *BatchGetVideoCharacterReplaceJobsResponseBody) *BatchGetVideoCharacterReplaceJobsResponse {
	s.Body = v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
