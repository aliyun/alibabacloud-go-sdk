// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTagResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTagResourcesResponse
	GetStatusCode() *int32
	SetBody(v *CreateTagResourcesResponseBody) *CreateTagResourcesResponse
	GetBody() *CreateTagResourcesResponseBody
}

type CreateTagResourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTagResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *CreateTagResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTagResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTagResourcesResponse) GetBody() *CreateTagResourcesResponseBody {
	return s.Body
}

func (s *CreateTagResourcesResponse) SetHeaders(v map[string]*string) *CreateTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *CreateTagResourcesResponse) SetStatusCode(v int32) *CreateTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTagResourcesResponse) SetBody(v *CreateTagResourcesResponseBody) *CreateTagResourcesResponse {
	s.Body = v
	return s
}

func (s *CreateTagResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
