// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnPublishProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnPublishProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnPublishProjectResponse
	GetStatusCode() *int32
	SetBody(v *UnPublishProjectResponseBody) *UnPublishProjectResponse
	GetBody() *UnPublishProjectResponseBody
}

type UnPublishProjectResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnPublishProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnPublishProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s UnPublishProjectResponse) GoString() string {
	return s.String()
}

func (s *UnPublishProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnPublishProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnPublishProjectResponse) GetBody() *UnPublishProjectResponseBody {
	return s.Body
}

func (s *UnPublishProjectResponse) SetHeaders(v map[string]*string) *UnPublishProjectResponse {
	s.Headers = v
	return s
}

func (s *UnPublishProjectResponse) SetStatusCode(v int32) *UnPublishProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UnPublishProjectResponse) SetBody(v *UnPublishProjectResponseBody) *UnPublishProjectResponse {
	s.Body = v
	return s
}

func (s *UnPublishProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
