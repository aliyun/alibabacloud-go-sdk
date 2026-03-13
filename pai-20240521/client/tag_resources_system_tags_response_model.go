// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesSystemTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TagResourcesSystemTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TagResourcesSystemTagsResponse
	GetStatusCode() *int32
	SetBody(v *TagResourcesSystemTagsResponseBody) *TagResourcesSystemTagsResponse
	GetBody() *TagResourcesSystemTagsResponseBody
}

type TagResourcesSystemTagsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesSystemTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesSystemTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesSystemTagsResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesSystemTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TagResourcesSystemTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TagResourcesSystemTagsResponse) GetBody() *TagResourcesSystemTagsResponseBody {
	return s.Body
}

func (s *TagResourcesSystemTagsResponse) SetHeaders(v map[string]*string) *TagResourcesSystemTagsResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesSystemTagsResponse) SetStatusCode(v int32) *TagResourcesSystemTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesSystemTagsResponse) SetBody(v *TagResourcesSystemTagsResponseBody) *TagResourcesSystemTagsResponse {
	s.Body = v
	return s
}

func (s *TagResourcesSystemTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
