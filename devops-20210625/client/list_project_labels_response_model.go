// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectLabelsResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectLabelsResponseBody) *ListProjectLabelsResponse
	GetBody() *ListProjectLabelsResponseBody
}

type ListProjectLabelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectLabelsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectLabelsResponse) GetBody() *ListProjectLabelsResponseBody {
	return s.Body
}

func (s *ListProjectLabelsResponse) SetHeaders(v map[string]*string) *ListProjectLabelsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectLabelsResponse) SetStatusCode(v int32) *ListProjectLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectLabelsResponse) SetBody(v *ListProjectLabelsResponseBody) *ListProjectLabelsResponse {
	s.Body = v
	return s
}

func (s *ListProjectLabelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
