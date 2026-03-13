// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesSystemTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UntagResourcesSystemTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UntagResourcesSystemTagsResponse
	GetStatusCode() *int32
	SetBody(v *UntagResourcesSystemTagsResponseBody) *UntagResourcesSystemTagsResponse
	GetBody() *UntagResourcesSystemTagsResponseBody
}

type UntagResourcesSystemTagsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesSystemTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesSystemTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesSystemTagsResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesSystemTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UntagResourcesSystemTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UntagResourcesSystemTagsResponse) GetBody() *UntagResourcesSystemTagsResponseBody {
	return s.Body
}

func (s *UntagResourcesSystemTagsResponse) SetHeaders(v map[string]*string) *UntagResourcesSystemTagsResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesSystemTagsResponse) SetStatusCode(v int32) *UntagResourcesSystemTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesSystemTagsResponse) SetBody(v *UntagResourcesSystemTagsResponseBody) *UntagResourcesSystemTagsResponse {
	s.Body = v
	return s
}

func (s *UntagResourcesSystemTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
