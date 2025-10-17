// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSparkTemplateFileIdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSparkTemplateFileIdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSparkTemplateFileIdsResponse
	GetStatusCode() *int32
	SetBody(v *ListSparkTemplateFileIdsResponseBody) *ListSparkTemplateFileIdsResponse
	GetBody() *ListSparkTemplateFileIdsResponseBody
}

type ListSparkTemplateFileIdsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSparkTemplateFileIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSparkTemplateFileIdsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSparkTemplateFileIdsResponse) GoString() string {
	return s.String()
}

func (s *ListSparkTemplateFileIdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSparkTemplateFileIdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSparkTemplateFileIdsResponse) GetBody() *ListSparkTemplateFileIdsResponseBody {
	return s.Body
}

func (s *ListSparkTemplateFileIdsResponse) SetHeaders(v map[string]*string) *ListSparkTemplateFileIdsResponse {
	s.Headers = v
	return s
}

func (s *ListSparkTemplateFileIdsResponse) SetStatusCode(v int32) *ListSparkTemplateFileIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSparkTemplateFileIdsResponse) SetBody(v *ListSparkTemplateFileIdsResponseBody) *ListSparkTemplateFileIdsResponse {
	s.Body = v
	return s
}

func (s *ListSparkTemplateFileIdsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
