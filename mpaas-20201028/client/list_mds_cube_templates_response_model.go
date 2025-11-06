// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMdsCubeTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMdsCubeTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMdsCubeTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListMdsCubeTemplatesResponseBody) *ListMdsCubeTemplatesResponse
	GetBody() *ListMdsCubeTemplatesResponseBody
}

type ListMdsCubeTemplatesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMdsCubeTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMdsCubeTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListMdsCubeTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMdsCubeTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMdsCubeTemplatesResponse) GetBody() *ListMdsCubeTemplatesResponseBody {
	return s.Body
}

func (s *ListMdsCubeTemplatesResponse) SetHeaders(v map[string]*string) *ListMdsCubeTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListMdsCubeTemplatesResponse) SetStatusCode(v int32) *ListMdsCubeTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMdsCubeTemplatesResponse) SetBody(v *ListMdsCubeTemplatesResponseBody) *ListMdsCubeTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListMdsCubeTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
