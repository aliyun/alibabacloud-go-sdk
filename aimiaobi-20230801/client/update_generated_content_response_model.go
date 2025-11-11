// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGeneratedContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGeneratedContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGeneratedContentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGeneratedContentResponseBody) *UpdateGeneratedContentResponse
	GetBody() *UpdateGeneratedContentResponseBody
}

type UpdateGeneratedContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGeneratedContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGeneratedContentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGeneratedContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateGeneratedContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGeneratedContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGeneratedContentResponse) GetBody() *UpdateGeneratedContentResponseBody {
	return s.Body
}

func (s *UpdateGeneratedContentResponse) SetHeaders(v map[string]*string) *UpdateGeneratedContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateGeneratedContentResponse) SetStatusCode(v int32) *UpdateGeneratedContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGeneratedContentResponse) SetBody(v *UpdateGeneratedContentResponseBody) *UpdateGeneratedContentResponse {
	s.Body = v
	return s
}

func (s *UpdateGeneratedContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
