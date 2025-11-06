// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMdsCubeTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMdsCubeTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMdsCubeTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateMdsCubeTemplateResponseBody) *CreateMdsCubeTemplateResponse
	GetBody() *CreateMdsCubeTemplateResponseBody
}

type CreateMdsCubeTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMdsCubeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMdsCubeTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsCubeTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateMdsCubeTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMdsCubeTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMdsCubeTemplateResponse) GetBody() *CreateMdsCubeTemplateResponseBody {
	return s.Body
}

func (s *CreateMdsCubeTemplateResponse) SetHeaders(v map[string]*string) *CreateMdsCubeTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateMdsCubeTemplateResponse) SetStatusCode(v int32) *CreateMdsCubeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMdsCubeTemplateResponse) SetBody(v *CreateMdsCubeTemplateResponseBody) *CreateMdsCubeTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateMdsCubeTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
