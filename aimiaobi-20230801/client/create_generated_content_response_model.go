// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGeneratedContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGeneratedContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGeneratedContentResponse
	GetStatusCode() *int32
	SetBody(v *CreateGeneratedContentResponseBody) *CreateGeneratedContentResponse
	GetBody() *CreateGeneratedContentResponseBody
}

type CreateGeneratedContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGeneratedContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGeneratedContentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGeneratedContentResponse) GoString() string {
	return s.String()
}

func (s *CreateGeneratedContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGeneratedContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGeneratedContentResponse) GetBody() *CreateGeneratedContentResponseBody {
	return s.Body
}

func (s *CreateGeneratedContentResponse) SetHeaders(v map[string]*string) *CreateGeneratedContentResponse {
	s.Headers = v
	return s
}

func (s *CreateGeneratedContentResponse) SetStatusCode(v int32) *CreateGeneratedContentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGeneratedContentResponse) SetBody(v *CreateGeneratedContentResponseBody) *CreateGeneratedContentResponse {
	s.Body = v
	return s
}

func (s *CreateGeneratedContentResponse) Validate() error {
	return dara.Validate(s)
}
